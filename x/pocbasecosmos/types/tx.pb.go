// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pocbasecosmos/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

type MsgAdminSpendCommunityPool struct {
	Initiator string                                   `protobuf:"bytes,1,opt,name=initiator,proto3" json:"initiator,omitempty"`
	ToAddress string                                   `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	Coins     github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
}

func (m *MsgAdminSpendCommunityPool) Reset()         { *m = MsgAdminSpendCommunityPool{} }
func (m *MsgAdminSpendCommunityPool) String() string { return proto.CompactTextString(m) }
func (*MsgAdminSpendCommunityPool) ProtoMessage()    {}
func (*MsgAdminSpendCommunityPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_75dc3e331b0b4a81, []int{0}
}
func (m *MsgAdminSpendCommunityPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAdminSpendCommunityPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAdminSpendCommunityPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAdminSpendCommunityPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAdminSpendCommunityPool.Merge(m, src)
}
func (m *MsgAdminSpendCommunityPool) XXX_Size() int {
	return m.Size()
}
func (m *MsgAdminSpendCommunityPool) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAdminSpendCommunityPool.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAdminSpendCommunityPool proto.InternalMessageInfo

func (m *MsgAdminSpendCommunityPool) GetInitiator() string {
	if m != nil {
		return m.Initiator
	}
	return ""
}

func (m *MsgAdminSpendCommunityPool) GetToAddress() string {
	if m != nil {
		return m.ToAddress
	}
	return ""
}

func (m *MsgAdminSpendCommunityPool) GetCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Coins
	}
	return nil
}

type MsgAdminSpendResponse struct {
}

func (m *MsgAdminSpendResponse) Reset()         { *m = MsgAdminSpendResponse{} }
func (m *MsgAdminSpendResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAdminSpendResponse) ProtoMessage()    {}
func (*MsgAdminSpendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_75dc3e331b0b4a81, []int{1}
}
func (m *MsgAdminSpendResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAdminSpendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAdminSpendResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAdminSpendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAdminSpendResponse.Merge(m, src)
}
func (m *MsgAdminSpendResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAdminSpendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAdminSpendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAdminSpendResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgAdminSpendCommunityPool)(nil), "rdpnd.pocbasecosmos.pocbasecosmos.MsgAdminSpendCommunityPool")
	proto.RegisterType((*MsgAdminSpendResponse)(nil), "rdpnd.pocbasecosmos.pocbasecosmos.MsgAdminSpendResponse")
}

func init() { proto.RegisterFile("pocbasecosmos/tx.proto", fileDescriptor_75dc3e331b0b4a81) }

var fileDescriptor_75dc3e331b0b4a81 = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xc1, 0x4e, 0x2a, 0x31,
	0x14, 0x9d, 0x3e, 0xf2, 0x5e, 0x42, 0xdf, 0x6e, 0xa2, 0x02, 0x13, 0x1d, 0x90, 0x15, 0x1b, 0x5a,
	0xc1, 0xc4, 0xb8, 0x71, 0x01, 0xac, 0x49, 0x08, 0xee, 0xdc, 0x90, 0x99, 0x69, 0x33, 0x36, 0x3a,
	0xbd, 0xcd, 0xb4, 0x18, 0xf8, 0x0b, 0xe3, 0xd2, 0x4f, 0xf0, 0x1b, 0xfc, 0x00, 0x96, 0x2c, 0x5d,
	0xa9, 0x81, 0x1f, 0x31, 0x33, 0x1d, 0xa2, 0x98, 0x10, 0xe3, 0xaa, 0xed, 0x3d, 0x3d, 0xe7, 0x9e,
	0x7b, 0x5a, 0x7c, 0xa0, 0x20, 0x0a, 0x03, 0xcd, 0x23, 0xd0, 0x09, 0x68, 0x6a, 0x66, 0x44, 0xa5,
	0x60, 0xc0, 0x3d, 0x4e, 0x99, 0x92, 0x8c, 0x6c, 0xa1, 0xdb, 0x27, 0x6f, 0x2f, 0x86, 0x18, 0xf2,
	0xdb, 0x34, 0xdb, 0x59, 0xa2, 0x57, 0xb3, 0xe8, 0xc4, 0x02, 0x1b, 0x62, 0x0e, 0xf9, 0x45, 0x93,
	0x4c, 0x83, 0xde, 0x75, 0x42, 0x6e, 0x82, 0x0e, 0x8d, 0x40, 0x48, 0x8b, 0x37, 0x9f, 0x11, 0xf6,
	0x86, 0x3a, 0xee, 0xb1, 0x44, 0xc8, 0x4b, 0xc5, 0x25, 0x1b, 0x40, 0x92, 0x4c, 0xa5, 0x30, 0xf3,
	0x11, 0xc0, 0xad, 0x7b, 0x88, 0xcb, 0x42, 0x0a, 0x23, 0x02, 0x03, 0x69, 0x15, 0x35, 0x50, 0xab,
	0x3c, 0xfe, 0x2c, 0xb8, 0x47, 0x18, 0x1b, 0x98, 0x04, 0x8c, 0xa5, 0x5c, 0xeb, 0xea, 0x1f, 0x0b,
	0x1b, 0xe8, 0xd9, 0x82, 0x1b, 0xe0, 0xbf, 0x59, 0x27, 0x5d, 0x2d, 0x35, 0x4a, 0xad, 0xff, 0xdd,
	0x1a, 0x29, 0x9c, 0x65, 0x5e, 0x48, 0xe1, 0x85, 0x0c, 0x40, 0xc8, 0xfe, 0xc9, 0xe2, 0xb5, 0xee,
	0x3c, 0xbd, 0xd5, 0x5b, 0xb1, 0x30, 0xd7, 0xd3, 0x90, 0x44, 0x90, 0x14, 0x63, 0x14, 0x4b, 0x5b,
	0xb3, 0x1b, 0x6a, 0xe6, 0x8a, 0xeb, 0x9c, 0xa0, 0xc7, 0x56, 0xb9, 0x59, 0xc1, 0xfb, 0x5b, 0xee,
	0xc7, 0x5c, 0x2b, 0x90, 0x9a, 0x77, 0x1f, 0x11, 0x2e, 0x0d, 0x75, 0xec, 0x3e, 0x20, 0x5c, 0xd9,
	0x35, 0xdc, 0x05, 0xf9, 0x31, 0x70, 0xb2, 0x3b, 0x1b, 0xef, 0xfc, 0xb7, 0xf4, 0x8d, 0xb9, 0xfe,
	0x68, 0xb1, 0xf2, 0xd1, 0x72, 0xe5, 0xa3, 0xf7, 0x95, 0x8f, 0xee, 0xd7, 0xbe, 0xb3, 0x5c, 0xfb,
	0xce, 0xcb, 0xda, 0x77, 0xae, 0xce, 0xbe, 0x04, 0x90, 0xab, 0x53, 0x05, 0x51, 0x3b, 0x13, 0x6c,
	0x17, 0x79, 0xcc, 0xe8, 0xb7, 0xdf, 0x93, 0x85, 0x12, 0xfe, 0xcb, 0x5f, 0xf3, 0xf4, 0x23, 0x00,
	0x00, 0xff, 0xff, 0x1c, 0xda, 0xe9, 0x94, 0x5b, 0x02, 0x00, 0x00,
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
	// this line is used by starport scaffolding # proto/tx/rpc
	AdminSpendCommunityPool(ctx context.Context, in *MsgAdminSpendCommunityPool, opts ...grpc.CallOption) (*MsgAdminSpendResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) AdminSpendCommunityPool(ctx context.Context, in *MsgAdminSpendCommunityPool, opts ...grpc.CallOption) (*MsgAdminSpendResponse, error) {
	out := new(MsgAdminSpendResponse)
	err := c.cc.Invoke(ctx, "/rdpnd.pocbasecosmos.pocbasecosmos.Msg/AdminSpendCommunityPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// this line is used by starport scaffolding # proto/tx/rpc
	AdminSpendCommunityPool(context.Context, *MsgAdminSpendCommunityPool) (*MsgAdminSpendResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) AdminSpendCommunityPool(ctx context.Context, req *MsgAdminSpendCommunityPool) (*MsgAdminSpendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminSpendCommunityPool not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_AdminSpendCommunityPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAdminSpendCommunityPool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AdminSpendCommunityPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rdpnd.pocbasecosmos.pocbasecosmos.Msg/AdminSpendCommunityPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AdminSpendCommunityPool(ctx, req.(*MsgAdminSpendCommunityPool))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rdpnd.pocbasecosmos.pocbasecosmos.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AdminSpendCommunityPool",
			Handler:    _Msg_AdminSpendCommunityPool_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pocbasecosmos/tx.proto",
}

func (m *MsgAdminSpendCommunityPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAdminSpendCommunityPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAdminSpendCommunityPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ToAddress) > 0 {
		i -= len(m.ToAddress)
		copy(dAtA[i:], m.ToAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ToAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Initiator) > 0 {
		i -= len(m.Initiator)
		copy(dAtA[i:], m.Initiator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Initiator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAdminSpendResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAdminSpendResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAdminSpendResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgAdminSpendCommunityPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Initiator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ToAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgAdminSpendResponse) Size() (n int) {
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
func (m *MsgAdminSpendCommunityPool) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgAdminSpendCommunityPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAdminSpendCommunityPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Initiator", wireType)
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
			m.Initiator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToAddress", wireType)
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
			m.ToAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
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
			m.Coins = append(m.Coins, types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
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
func (m *MsgAdminSpendResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgAdminSpendResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAdminSpendResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
