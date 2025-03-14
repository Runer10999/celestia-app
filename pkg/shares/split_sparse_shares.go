package shares

import (
	"encoding/binary"
	"fmt"

	"github.com/celestiaorg/celestia-app/pkg/appconsts"
	"github.com/celestiaorg/nmt/namespace"
	coretypes "github.com/tendermint/tendermint/types"
	"golang.org/x/exp/slices"
)

// SparseShareSplitter lazily splits blobs into shares that will eventually be
// included in a data square. It also has methods to help progressively count
// how many shares the blobs written take up.
type SparseShareSplitter struct {
	shares []Share
}

func NewSparseShareSplitter() *SparseShareSplitter {
	return &SparseShareSplitter{}
}

// Write writes the provided blob to this sparse share splitter. It returns an
// error or nil if no error is encountered.
func (sss *SparseShareSplitter) Write(blob coretypes.Blob) error {
	if !slices.Contains(appconsts.SupportedShareVersions, blob.ShareVersion) {
		return fmt.Errorf("unsupported share version: %d", blob.ShareVersion)
	}
	rawBlob, err := MarshalDelimitedBlob(blob)
	if err != nil {
		panic(fmt.Sprintf("app accepted a Blob that can not be encoded %#v", blob))
	}
	newShares := make([]Share, 0)
	newShares = AppendToShares(newShares, blob.NamespaceID, rawBlob, blob.ShareVersion)
	sss.shares = append(sss.shares, newShares...)
	return nil
}

// RemoveBlob will remove a blob from the underlying blob state. If
// there is namespaced padding after the blob, then that is also removed.
func (sss *SparseShareSplitter) RemoveBlob(i int) (int, error) {
	j := 1
	initialCount := len(sss.shares)
	if len(sss.shares) > i+1 {
		blobLen, err := sss.shares[i+1].SequenceLength()
		if err != nil {
			return 0, err
		}
		// 0 means that there is padding after the share that we are about to
		// remove. to remove this padding, we increase j by 1
		// with the blob
		if blobLen == 0 {
			j++
		}
	}
	copy(sss.shares[i:], sss.shares[i+j:])
	sss.shares = sss.shares[:len(sss.shares)-j]
	newCount := len(sss.shares)
	return initialCount - newCount, nil
}

// WriteNamespacedPaddedShares adds empty shares using the namespace of the last written share.
// This is useful to follow the message layout rules. It assumes that at least
// one share has already been written, if not it panics.
func (sss *SparseShareSplitter) WriteNamespacedPaddedShares(count int) {
	if len(sss.shares) == 0 {
		panic("cannot write empty namespaced shares on an empty SparseShareSplitter")
	}
	if count == 0 {
		return
	}
	lastBlob := sss.shares[len(sss.shares)-1]
	sss.shares = append(sss.shares, namespacedPaddedShares(lastBlob.NamespaceID(), count)...)
}

// Export finalizes and returns the underlying shares.
func (sss *SparseShareSplitter) Export() []Share {
	return sss.shares
}

// Count returns the current number of shares that will be made if exporting.
func (sss *SparseShareSplitter) Count() int {
	return len(sss.shares)
}

// AppendToShares appends raw data as shares.
// Used for blobs.
func AppendToShares(shares []Share, nid namespace.ID, rawData []byte, shareVersion uint8) []Share {
	if len(rawData) <= appconsts.SparseShareContentSize {
		infoByte, err := NewInfoByte(shareVersion, true)
		if err != nil {
			panic(err)
		}
		rawShare := append(append(append(
			make([]byte, 0, appconsts.ShareSize),
			nid...),
			byte(infoByte)),
			rawData...,
		)
		paddedShare, _ := zeroPadIfNecessary(rawShare, appconsts.ShareSize)
		shares = append(shares, paddedShare)
	} else { // len(rawData) > BlobShareSize
		shares = append(shares, splitBlob(rawData, nid, shareVersion)...)
	}
	return shares
}

// MarshalDelimitedBlob marshals the raw share data (excluding the namespace)
// of this blob and prefixes it with the length of the blob encoded as a
// varint.
func MarshalDelimitedBlob(blob coretypes.Blob) ([]byte, error) {
	lenBuf := make([]byte, binary.MaxVarintLen64)
	length := uint64(len(blob.Data))
	n := binary.PutUvarint(lenBuf, length)
	return append(lenBuf[:n], blob.Data...), nil
}

// splitBlob breaks the data in a blob into the minimum number of
// namespaced shares
func splitBlob(rawData []byte, nid namespace.ID, shareVersion uint8) (shares []Share) {
	infoByte, err := NewInfoByte(shareVersion, true)
	if err != nil {
		panic(err)
	}
	firstRawShare := append(append(append(
		make([]byte, 0, appconsts.ShareSize),
		nid...),
		byte(infoByte)),
		rawData[:appconsts.SparseShareContentSize]...,
	)
	shares = append(shares, firstRawShare)
	rawData = rawData[appconsts.SparseShareContentSize:]
	for len(rawData) > 0 {
		shareSizeOrLen := min(appconsts.SparseShareContentSize, len(rawData))
		infoByte, err := NewInfoByte(appconsts.ShareVersionZero, false)
		if err != nil {
			panic(err)
		}
		rawShare := append(append(append(
			make([]byte, 0, appconsts.ShareSize),
			nid...),
			byte(infoByte)),
			rawData[:shareSizeOrLen]...,
		)
		paddedShare, _ := zeroPadIfNecessary(rawShare, appconsts.ShareSize)
		shares = append(shares, paddedShare)
		rawData = rawData[shareSizeOrLen:]
	}
	return shares
}

func namespacedPaddedShares(ns namespace.ID, count int) []Share {
	shares := make([]Share, count)
	for i := 0; i < count; i++ {
		shares[i] = namespacedPaddedShare(ns)
	}
	return shares
}

func namespacedPaddedShare(ns namespace.ID) Share {
	infoByte, err := NewInfoByte(appconsts.ShareVersionZero, true)
	if err != nil {
		panic(err)
	}
	share := make([]byte, 0, appconsts.ShareSize)
	share = append(share, ns...)
	share = append(share, byte(infoByte))
	share = append(share, appconsts.NameSpacedPaddedShareBytes...)
	return share
}
