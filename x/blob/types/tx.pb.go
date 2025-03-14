// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: blob/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgWirePayForBlob describes the format of data that is sent over the wire
// for each PayForBlob
type MsgWirePayForBlob struct {
	Signer      string `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	NamespaceId []byte `protobuf:"bytes,2,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id,omitempty"`
	BlobSize    uint64 `protobuf:"varint,3,opt,name=blob_size,json=blobSize,proto3" json:"blob_size,omitempty"`
	Blob        []byte `protobuf:"bytes,4,opt,name=blob,proto3" json:"blob,omitempty"`
	// field number 6 is obsolete and was used for `repeated
	// ShareCommitAndSignature` when a MsgWirePayForBlob included multiple
	// blob share commitments (one per square size).
	ShareCommitment *ShareCommitAndSignature `protobuf:"bytes,7,opt,name=share_commitment,json=shareCommitment,proto3" json:"share_commitment,omitempty"`
	// share_version is the version of the share format that the blob in this
	// message should use when included in a block. The share_version specified
	// must match the share_version used to generate the share_commitment in this
	// message.
	ShareVersion uint32 `protobuf:"varint,8,opt,name=share_version,json=shareVersion,proto3" json:"share_version,omitempty"`
}

func (m *MsgWirePayForBlob) Reset()         { *m = MsgWirePayForBlob{} }
func (m *MsgWirePayForBlob) String() string { return proto.CompactTextString(m) }
func (*MsgWirePayForBlob) ProtoMessage()    {}
func (*MsgWirePayForBlob) Descriptor() ([]byte, []int) {
	return fileDescriptor_f945cb94fe124aae, []int{0}
}
func (m *MsgWirePayForBlob) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWirePayForBlob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWirePayForBlob.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWirePayForBlob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWirePayForBlob.Merge(m, src)
}
func (m *MsgWirePayForBlob) XXX_Size() int {
	return m.Size()
}
func (m *MsgWirePayForBlob) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWirePayForBlob.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWirePayForBlob proto.InternalMessageInfo

func (m *MsgWirePayForBlob) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *MsgWirePayForBlob) GetNamespaceId() []byte {
	if m != nil {
		return m.NamespaceId
	}
	return nil
}

func (m *MsgWirePayForBlob) GetBlobSize() uint64 {
	if m != nil {
		return m.BlobSize
	}
	return 0
}

func (m *MsgWirePayForBlob) GetBlob() []byte {
	if m != nil {
		return m.Blob
	}
	return nil
}

func (m *MsgWirePayForBlob) GetShareCommitment() *ShareCommitAndSignature {
	if m != nil {
		return m.ShareCommitment
	}
	return nil
}

func (m *MsgWirePayForBlob) GetShareVersion() uint32 {
	if m != nil {
		return m.ShareVersion
	}
	return 0
}

// MsgWirePayForBlobResponse describes the response returned after the
// submission of a WirePayForBlob
type MsgWirePayForBlobResponse struct {
}

func (m *MsgWirePayForBlobResponse) Reset()         { *m = MsgWirePayForBlobResponse{} }
func (m *MsgWirePayForBlobResponse) String() string { return proto.CompactTextString(m) }
func (*MsgWirePayForBlobResponse) ProtoMessage()    {}
func (*MsgWirePayForBlobResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f945cb94fe124aae, []int{1}
}
func (m *MsgWirePayForBlobResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWirePayForBlobResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWirePayForBlobResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWirePayForBlobResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWirePayForBlobResponse.Merge(m, src)
}
func (m *MsgWirePayForBlobResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgWirePayForBlobResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWirePayForBlobResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWirePayForBlobResponse proto.InternalMessageInfo

// ShareCommitAndSignature defines the
type ShareCommitAndSignature struct {
	// share_commitment is the root of a binary Merkle tree that has leaves which
	// are subtree roots of the relevant blob shares in the original data
	// square.
	ShareCommitment []byte `protobuf:"bytes,2,opt,name=share_commitment,json=shareCommitment,proto3" json:"share_commitment,omitempty"`
	Signature       []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *ShareCommitAndSignature) Reset()         { *m = ShareCommitAndSignature{} }
func (m *ShareCommitAndSignature) String() string { return proto.CompactTextString(m) }
func (*ShareCommitAndSignature) ProtoMessage()    {}
func (*ShareCommitAndSignature) Descriptor() ([]byte, []int) {
	return fileDescriptor_f945cb94fe124aae, []int{2}
}
func (m *ShareCommitAndSignature) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ShareCommitAndSignature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ShareCommitAndSignature.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ShareCommitAndSignature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShareCommitAndSignature.Merge(m, src)
}
func (m *ShareCommitAndSignature) XXX_Size() int {
	return m.Size()
}
func (m *ShareCommitAndSignature) XXX_DiscardUnknown() {
	xxx_messageInfo_ShareCommitAndSignature.DiscardUnknown(m)
}

var xxx_messageInfo_ShareCommitAndSignature proto.InternalMessageInfo

func (m *ShareCommitAndSignature) GetShareCommitment() []byte {
	if m != nil {
		return m.ShareCommitment
	}
	return nil
}

func (m *ShareCommitAndSignature) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

// MsgPayForBlob is what gets signed by users when creating
// ShareCommitSignatures. It is a subset of MsgWirePayForBlob that does not contain the blob.
type MsgPayForBlob struct {
	Signer      string `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	NamespaceId []byte `protobuf:"bytes,2,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id,omitempty"`
	BlobSize    uint64 `protobuf:"varint,3,opt,name=blob_size,json=blobSize,proto3" json:"blob_size,omitempty"`
	// share_commitment is the share_commitment from
	// ShareCommitAndSignature that will be included in a block
	ShareCommitment []byte `protobuf:"bytes,4,opt,name=share_commitment,json=shareCommitment,proto3" json:"share_commitment,omitempty"`
	// share_version is the version of the share format that the blob associated
	// with this message should use when included in a block. The share_version
	// specified must match the share_version used to generate the
	// share_commitment in this message.
	ShareVersion uint32 `protobuf:"varint,8,opt,name=share_version,json=shareVersion,proto3" json:"share_version,omitempty"`
}

func (m *MsgPayForBlob) Reset()         { *m = MsgPayForBlob{} }
func (m *MsgPayForBlob) String() string { return proto.CompactTextString(m) }
func (*MsgPayForBlob) ProtoMessage()    {}
func (*MsgPayForBlob) Descriptor() ([]byte, []int) {
	return fileDescriptor_f945cb94fe124aae, []int{3}
}
func (m *MsgPayForBlob) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPayForBlob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPayForBlob.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPayForBlob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPayForBlob.Merge(m, src)
}
func (m *MsgPayForBlob) XXX_Size() int {
	return m.Size()
}
func (m *MsgPayForBlob) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPayForBlob.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPayForBlob proto.InternalMessageInfo

func (m *MsgPayForBlob) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *MsgPayForBlob) GetNamespaceId() []byte {
	if m != nil {
		return m.NamespaceId
	}
	return nil
}

func (m *MsgPayForBlob) GetBlobSize() uint64 {
	if m != nil {
		return m.BlobSize
	}
	return 0
}

func (m *MsgPayForBlob) GetShareCommitment() []byte {
	if m != nil {
		return m.ShareCommitment
	}
	return nil
}

func (m *MsgPayForBlob) GetShareVersion() uint32 {
	if m != nil {
		return m.ShareVersion
	}
	return 0
}

// MsgPayForBlobResponse describes the response returned after the submission
// of a PayForBlob
type MsgPayForBlobResponse struct {
}

func (m *MsgPayForBlobResponse) Reset()         { *m = MsgPayForBlobResponse{} }
func (m *MsgPayForBlobResponse) String() string { return proto.CompactTextString(m) }
func (*MsgPayForBlobResponse) ProtoMessage()    {}
func (*MsgPayForBlobResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f945cb94fe124aae, []int{4}
}
func (m *MsgPayForBlobResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPayForBlobResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPayForBlobResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPayForBlobResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPayForBlobResponse.Merge(m, src)
}
func (m *MsgPayForBlobResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgPayForBlobResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPayForBlobResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPayForBlobResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgWirePayForBlob)(nil), "blob.MsgWirePayForBlob")
	proto.RegisterType((*MsgWirePayForBlobResponse)(nil), "blob.MsgWirePayForBlobResponse")
	proto.RegisterType((*ShareCommitAndSignature)(nil), "blob.ShareCommitAndSignature")
	proto.RegisterType((*MsgPayForBlob)(nil), "blob.MsgPayForBlob")
	proto.RegisterType((*MsgPayForBlobResponse)(nil), "blob.MsgPayForBlobResponse")
}

func init() { proto.RegisterFile("blob/tx.proto", fileDescriptor_f945cb94fe124aae) }

var fileDescriptor_f945cb94fe124aae = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x18, 0xcc, 0xb6, 0x51, 0x69, 0xbe, 0x26, 0x02, 0x96, 0x9f, 0xba, 0x49, 0xb1, 0xdc, 0x70, 0x31,
	0x07, 0x62, 0x54, 0x9e, 0x80, 0x22, 0x21, 0x40, 0x8a, 0x84, 0x1c, 0x09, 0x24, 0x2e, 0xd1, 0x3a,
	0xf9, 0xd8, 0xac, 0x14, 0xef, 0x67, 0xed, 0x6e, 0x51, 0xd3, 0x23, 0x4f, 0x80, 0xc4, 0xc3, 0xf0,
	0x0a, 0x1c, 0x2b, 0x71, 0xe1, 0x88, 0x12, 0xce, 0x3c, 0x03, 0xf2, 0x9a, 0xa6, 0x54, 0x35, 0x12,
	0x27, 0x6e, 0xb3, 0x33, 0x9e, 0xcf, 0xf3, 0x8d, 0xbd, 0xd0, 0xc9, 0xe6, 0x94, 0x25, 0xee, 0x64,
	0x50, 0x18, 0x72, 0xc4, 0x9b, 0xe5, 0xb1, 0x7b, 0x5b, 0x92, 0x24, 0x4f, 0x24, 0x25, 0xaa, 0xb4,
	0xee, 0xbe, 0x24, 0x92, 0x73, 0x4c, 0x44, 0xa1, 0x12, 0xa1, 0x35, 0x39, 0xe1, 0x14, 0x69, 0x5b,
	0xa9, 0xfd, 0x9f, 0x0c, 0x6e, 0x0e, 0xad, 0x7c, 0xa3, 0x0c, 0xbe, 0x12, 0x8b, 0x67, 0x64, 0x8e,
	0xe6, 0x94, 0xf1, 0xbb, 0xb0, 0x65, 0x95, 0xd4, 0x68, 0x02, 0x16, 0xb1, 0xb8, 0x95, 0xfe, 0x3e,
	0xf1, 0x03, 0x68, 0x6b, 0x91, 0xa3, 0x2d, 0xc4, 0x04, 0xc7, 0x6a, 0x1a, 0x6c, 0x44, 0x2c, 0x6e,
	0xa7, 0x3b, 0x6b, 0xee, 0xc5, 0x94, 0xf7, 0xa0, 0x55, 0x86, 0x19, 0x5b, 0x75, 0x8a, 0xc1, 0x66,
	0xc4, 0xe2, 0x66, 0xba, 0x5d, 0x12, 0x23, 0x75, 0x8a, 0x9c, 0x83, 0x4f, 0x1a, 0x34, 0xbd, 0xcf,
	0x63, 0xfe, 0x1c, 0x6e, 0xd8, 0x99, 0x30, 0x38, 0x9e, 0x50, 0x9e, 0x2b, 0x97, 0xa3, 0x76, 0xc1,
	0xb5, 0x88, 0xc5, 0x3b, 0x87, 0xf7, 0x06, 0xe5, 0x03, 0x83, 0x51, 0xa9, 0x3e, 0xf5, 0xe2, 0x13,
	0x3d, 0x1d, 0x29, 0xa9, 0x85, 0x3b, 0x36, 0x98, 0x5e, 0xb7, 0x17, 0x42, 0xe9, 0xe2, 0xf7, 0xa1,
	0x53, 0x4d, 0x7a, 0x8f, 0xc6, 0x2a, 0xd2, 0xc1, 0x76, 0xc4, 0xe2, 0x4e, 0xda, 0xf6, 0xe4, 0xeb,
	0x8a, 0xeb, 0xf7, 0x60, 0xef, 0xca, 0xbe, 0x29, 0xda, 0x82, 0xb4, 0xc5, 0x7e, 0x06, 0xbb, 0x7f,
	0x79, 0x1b, 0x7f, 0x50, 0x13, 0xb3, 0x5a, 0xff, 0x4a, 0x8e, 0x7d, 0x68, 0xd9, 0x73, 0x9f, 0xaf,
	0xa0, 0x9d, 0x5e, 0x10, 0xfd, 0xcf, 0x0c, 0x3a, 0x43, 0x2b, 0xff, 0x43, 0xdb, 0x75, 0x91, 0x9b,
	0xf5, 0x91, 0xff, 0xa9, 0xba, 0x5d, 0xb8, 0x73, 0x29, 0xf8, 0x79, 0x6d, 0x87, 0x33, 0xd8, 0x1c,
	0x5a, 0xc9, 0x05, 0xc0, 0x1f, 0x5b, 0xdd, 0xaa, 0xbe, 0xde, 0x25, 0x47, 0xb7, 0x57, 0x43, 0xae,
	0xdb, 0x3f, 0xf8, 0xf0, 0xf5, 0xc7, 0xa7, 0x8d, 0x1e, 0xdf, 0x4b, 0x26, 0x38, 0x47, 0xeb, 0x94,
	0x48, 0xfc, 0x6f, 0x5e, 0x88, 0xc5, 0x3b, 0x32, 0x25, 0x3c, 0x7a, 0xf9, 0x65, 0x19, 0xb2, 0xb3,
	0x65, 0xc8, 0xbe, 0x2f, 0x43, 0xf6, 0x71, 0x15, 0x36, 0xce, 0x56, 0x61, 0xe3, 0xdb, 0x2a, 0x6c,
	0xbc, 0x7d, 0x24, 0x95, 0x9b, 0x1d, 0x67, 0x83, 0x09, 0xe5, 0x6b, 0x3b, 0x19, 0xb9, 0xc6, 0x0f,
	0x45, 0x51, 0x24, 0x27, 0xd5, 0x40, 0xb7, 0x28, 0xd0, 0x66, 0x5b, 0xfe, 0x06, 0x3c, 0xfe, 0x15,
	0x00, 0x00, 0xff, 0xff, 0xee, 0xaa, 0x39, 0xce, 0x4c, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// PayForBlob allows the user to pay for the inclusion of a blob
	PayForBlob(ctx context.Context, in *MsgPayForBlob, opts ...grpc.CallOption) (*MsgPayForBlobResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) PayForBlob(ctx context.Context, in *MsgPayForBlob, opts ...grpc.CallOption) (*MsgPayForBlobResponse, error) {
	out := new(MsgPayForBlobResponse)
	err := c.cc.Invoke(ctx, "/blob.Msg/PayForBlob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// PayForBlob allows the user to pay for the inclusion of a blob
	PayForBlob(context.Context, *MsgPayForBlob) (*MsgPayForBlobResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) PayForBlob(ctx context.Context, req *MsgPayForBlob) (*MsgPayForBlobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayForBlob not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_PayForBlob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgPayForBlob)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).PayForBlob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blob.Msg/PayForBlob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).PayForBlob(ctx, req.(*MsgPayForBlob))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "blob.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PayForBlob",
			Handler:    _Msg_PayForBlob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blob/tx.proto",
}

func (m *MsgWirePayForBlob) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWirePayForBlob) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWirePayForBlob) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ShareVersion != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.ShareVersion))
		i--
		dAtA[i] = 0x40
	}
	if m.ShareCommitment != nil {
		{
			size, err := m.ShareCommitment.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Blob) > 0 {
		i -= len(m.Blob)
		copy(dAtA[i:], m.Blob)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Blob)))
		i--
		dAtA[i] = 0x22
	}
	if m.BlobSize != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.BlobSize))
		i--
		dAtA[i] = 0x18
	}
	if len(m.NamespaceId) > 0 {
		i -= len(m.NamespaceId)
		copy(dAtA[i:], m.NamespaceId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.NamespaceId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgWirePayForBlobResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWirePayForBlobResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWirePayForBlobResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ShareCommitAndSignature) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ShareCommitAndSignature) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ShareCommitAndSignature) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ShareCommitment) > 0 {
		i -= len(m.ShareCommitment)
		copy(dAtA[i:], m.ShareCommitment)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ShareCommitment)))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func (m *MsgPayForBlob) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPayForBlob) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPayForBlob) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ShareVersion != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.ShareVersion))
		i--
		dAtA[i] = 0x40
	}
	if len(m.ShareCommitment) > 0 {
		i -= len(m.ShareCommitment)
		copy(dAtA[i:], m.ShareCommitment)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ShareCommitment)))
		i--
		dAtA[i] = 0x22
	}
	if m.BlobSize != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.BlobSize))
		i--
		dAtA[i] = 0x18
	}
	if len(m.NamespaceId) > 0 {
		i -= len(m.NamespaceId)
		copy(dAtA[i:], m.NamespaceId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.NamespaceId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgPayForBlobResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPayForBlobResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPayForBlobResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgWirePayForBlob) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.NamespaceId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.BlobSize != 0 {
		n += 1 + sovTx(uint64(m.BlobSize))
	}
	l = len(m.Blob)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.ShareCommitment != nil {
		l = m.ShareCommitment.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if m.ShareVersion != 0 {
		n += 1 + sovTx(uint64(m.ShareVersion))
	}
	return n
}

func (m *MsgWirePayForBlobResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ShareCommitAndSignature) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ShareCommitment)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgPayForBlob) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.NamespaceId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.BlobSize != 0 {
		n += 1 + sovTx(uint64(m.BlobSize))
	}
	l = len(m.ShareCommitment)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.ShareVersion != 0 {
		n += 1 + sovTx(uint64(m.ShareVersion))
	}
	return n
}

func (m *MsgPayForBlobResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgWirePayForBlob) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgWirePayForBlob: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWirePayForBlob: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NamespaceId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NamespaceId = append(m.NamespaceId[:0], dAtA[iNdEx:postIndex]...)
			if m.NamespaceId == nil {
				m.NamespaceId = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlobSize", wireType)
			}
			m.BlobSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlobSize |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blob", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Blob = append(m.Blob[:0], dAtA[iNdEx:postIndex]...)
			if m.Blob == nil {
				m.Blob = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShareCommitment", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ShareCommitment == nil {
				m.ShareCommitment = &ShareCommitAndSignature{}
			}
			if err := m.ShareCommitment.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShareVersion", wireType)
			}
			m.ShareVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ShareVersion |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgWirePayForBlobResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgWirePayForBlobResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWirePayForBlobResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ShareCommitAndSignature) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ShareCommitAndSignature: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ShareCommitAndSignature: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShareCommitment", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ShareCommitment = append(m.ShareCommitment[:0], dAtA[iNdEx:postIndex]...)
			if m.ShareCommitment == nil {
				m.ShareCommitment = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgPayForBlob) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgPayForBlob: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPayForBlob: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NamespaceId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NamespaceId = append(m.NamespaceId[:0], dAtA[iNdEx:postIndex]...)
			if m.NamespaceId == nil {
				m.NamespaceId = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlobSize", wireType)
			}
			m.BlobSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlobSize |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShareCommitment", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ShareCommitment = append(m.ShareCommitment[:0], dAtA[iNdEx:postIndex]...)
			if m.ShareCommitment == nil {
				m.ShareCommitment = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShareVersion", wireType)
			}
			m.ShareVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ShareVersion |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgPayForBlobResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgPayForBlobResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPayForBlobResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
