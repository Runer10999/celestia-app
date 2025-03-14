package app_test

import (
	"bytes"
	"testing"

	"github.com/celestiaorg/nmt/namespace"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	core "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/celestiaorg/celestia-app/app"
	"github.com/celestiaorg/celestia-app/app/encoding"
	"github.com/celestiaorg/celestia-app/pkg/appconsts"
	"github.com/celestiaorg/celestia-app/testutil"
	"github.com/celestiaorg/celestia-app/x/blob/types"
)

func TestPrepareProposal(t *testing.T) {
	signer := types.GenerateKeyringSigner(t, types.TestAccName)

	encCfg := encoding.MakeConfig(app.ModuleEncodingRegisters...)

	testApp := testutil.SetupTestAppWithGenesisValSet(t)

	type test struct {
		input         abci.RequestPrepareProposal
		expectedBlobs []core.Blob
		expectedTxs   int
	}

	firstNamespace := []byte{2, 2, 2, 2, 2, 2, 2, 2}
	firstBlob := bytes.Repeat([]byte{4}, 512)
	firstRawTx := app.GenerateRawWirePFB(t, encCfg.TxConfig, firstNamespace, firstBlob, signer)

	secondNamespace := []byte{1, 1, 1, 1, 1, 1, 1, 1}
	secondBlob := []byte{2}
	secondRawTx := app.GenerateRawWirePFB(t, encCfg.TxConfig, secondNamespace, secondBlob, signer)

	thirdNamespace := []byte{3, 3, 3, 3, 3, 3, 3, 3}
	thirdBlob := []byte{1}
	thirdRawTx := app.GenerateRawWirePFB(t, encCfg.TxConfig, thirdNamespace, thirdBlob, signer)

	tests := []test{
		{
			input: abci.RequestPrepareProposal{
				BlockData: &core.Data{
					Txs: [][]byte{firstRawTx, secondRawTx, thirdRawTx},
				},
			},
			expectedBlobs: []core.Blob{
				{
					NamespaceId: secondNamespace, // the second blob should be first
					Data:        []byte{2},
				},
				{
					NamespaceId: firstNamespace,
					Data:        firstBlob,
				},
				{
					NamespaceId: thirdNamespace,
					Data:        []byte{1},
				},
			},
			expectedTxs: 3,
		},
	}

	for _, tt := range tests {
		res := testApp.PrepareProposal(tt.input)
		assert.Equal(t, tt.expectedBlobs, res.BlockData.Blobs)
		assert.Equal(t, tt.expectedTxs, len(res.BlockData.Txs))

		// verify the signatures of the prepared txs
		sdata, err := signer.GetSignerData()
		if err != nil {
			require.NoError(t, err)
		}
		dec := encoding.MalleatedTxDecoder(encCfg.TxConfig.TxDecoder())
		for _, tx := range res.BlockData.Txs {
			sTx, err := dec(tx)
			require.NoError(t, err)

			sigTx, ok := sTx.(authsigning.SigVerifiableTx)
			require.True(t, ok)

			sigs, err := sigTx.GetSignaturesV2()
			require.NoError(t, err)
			require.Equal(t, 1, len(sigs))
			sig := sigs[0]

			err = authsigning.VerifySignature(
				sdata.PubKey,
				sdata,
				sig.Data,
				encCfg.TxConfig.SignModeHandler(),
				sTx,
			)
			assert.NoError(t, err)
		}
	}
}

func TestPrepareProposalWithReservedNamespaces(t *testing.T) {
	testApp := testutil.SetupTestAppWithGenesisValSet(t)
	encCfg := encoding.MakeConfig(app.ModuleEncodingRegisters...)

	signer := types.GenerateKeyringSigner(t, types.TestAccName)

	type test struct {
		name          string
		namespace     namespace.ID
		expectedBlobs int
	}

	tests := []test{
		{"transaction namespace", appconsts.TxNamespaceID, 0},
		{"evidence namespace", appconsts.EvidenceNamespaceID, 0},
		{"tail padding namespace", appconsts.TailPaddingNamespaceID, 0},
		{"parity shares namespace", appconsts.ParitySharesNamespaceID, 0},
		{"other reserved namespace", namespace.ID{0, 0, 0, 0, 0, 0, 0, 200}, 0},
		{"valid namespace", namespace.ID{3, 3, 2, 2, 2, 1, 1, 1}, 1},
	}

	for _, tt := range tests {
		blob := []byte{1}
		tx := app.GenerateRawWirePFB(t, encCfg.TxConfig, tt.namespace, blob, signer)
		input := abci.RequestPrepareProposal{
			BlockData: &core.Data{
				Txs: [][]byte{tx},
			},
		}
		res := testApp.PrepareProposal(input)
		assert.Equal(t, tt.expectedBlobs, len(res.BlockData.Blobs))
	}
}
