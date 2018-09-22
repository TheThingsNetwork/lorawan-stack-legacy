// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: discovery/discovery.proto

package discovery // import "go.thethings.network/lorawan-stack-legacy/discovery"

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type GetRequest struct {
	ID                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ServiceName          string   `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()      { *m = GetRequest{} }
func (*GetRequest) ProtoMessage() {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_discovery_cada22fcf56e9ed8, []int{0}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(dst, src)
}
func (m *GetRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *GetRequest) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

type Announcement struct {
	ID                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ServiceName          string   `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	ServiceVersion       string   `protobuf:"bytes,3,opt,name=service_version,json=serviceVersion,proto3" json:"service_version,omitempty"`
	NetAddress           string   `protobuf:"bytes,11,opt,name=net_address,json=netAddress,proto3" json:"net_address,omitempty"`
	PublicKey            string   `protobuf:"bytes,12,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Certificate          string   `protobuf:"bytes,13,opt,name=certificate,proto3" json:"certificate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Announcement) Reset()      { *m = Announcement{} }
func (*Announcement) ProtoMessage() {}
func (*Announcement) Descriptor() ([]byte, []int) {
	return fileDescriptor_discovery_cada22fcf56e9ed8, []int{1}
}
func (m *Announcement) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Announcement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Announcement.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Announcement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Announcement.Merge(dst, src)
}
func (m *Announcement) XXX_Size() int {
	return m.Size()
}
func (m *Announcement) XXX_DiscardUnknown() {
	xxx_messageInfo_Announcement.DiscardUnknown(m)
}

var xxx_messageInfo_Announcement proto.InternalMessageInfo

func (m *Announcement) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Announcement) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *Announcement) GetServiceVersion() string {
	if m != nil {
		return m.ServiceVersion
	}
	return ""
}

func (m *Announcement) GetNetAddress() string {
	if m != nil {
		return m.NetAddress
	}
	return ""
}

func (m *Announcement) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

func (m *Announcement) GetCertificate() string {
	if m != nil {
		return m.Certificate
	}
	return ""
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "discovery.GetRequest")
	golang_proto.RegisterType((*GetRequest)(nil), "discovery.GetRequest")
	proto.RegisterType((*Announcement)(nil), "discovery.Announcement")
	golang_proto.RegisterType((*Announcement)(nil), "discovery.Announcement")
}
func (this *GetRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetRequest)
	if !ok {
		that2, ok := that.(GetRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ID != that1.ID {
		return false
	}
	if this.ServiceName != that1.ServiceName {
		return false
	}
	return true
}
func (this *Announcement) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Announcement)
	if !ok {
		that2, ok := that.(Announcement)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ID != that1.ID {
		return false
	}
	if this.ServiceName != that1.ServiceName {
		return false
	}
	if this.ServiceVersion != that1.ServiceVersion {
		return false
	}
	if this.NetAddress != that1.NetAddress {
		return false
	}
	if this.PublicKey != that1.PublicKey {
		return false
	}
	if this.Certificate != that1.Certificate {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Discovery service

type DiscoveryClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Announcement, error)
}

type discoveryClient struct {
	cc *grpc.ClientConn
}

func NewDiscoveryClient(cc *grpc.ClientConn) DiscoveryClient {
	return &discoveryClient{cc}
}

func (c *discoveryClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Announcement, error) {
	out := new(Announcement)
	err := c.cc.Invoke(ctx, "/discovery.Discovery/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Discovery service

type DiscoveryServer interface {
	Get(context.Context, *GetRequest) (*Announcement, error)
}

func RegisterDiscoveryServer(s *grpc.Server, srv DiscoveryServer) {
	s.RegisterService(&_Discovery_serviceDesc, srv)
}

func _Discovery_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discovery.Discovery/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Discovery_serviceDesc = grpc.ServiceDesc{
	ServiceName: "discovery.Discovery",
	HandlerType: (*DiscoveryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Discovery_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "discovery/discovery.proto",
}

func (m *GetRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDiscovery(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	if len(m.ServiceName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintDiscovery(dAtA, i, uint64(len(m.ServiceName)))
		i += copy(dAtA[i:], m.ServiceName)
	}
	return i, nil
}

func (m *Announcement) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Announcement) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDiscovery(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	if len(m.ServiceName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintDiscovery(dAtA, i, uint64(len(m.ServiceName)))
		i += copy(dAtA[i:], m.ServiceName)
	}
	if len(m.ServiceVersion) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintDiscovery(dAtA, i, uint64(len(m.ServiceVersion)))
		i += copy(dAtA[i:], m.ServiceVersion)
	}
	if len(m.NetAddress) > 0 {
		dAtA[i] = 0x5a
		i++
		i = encodeVarintDiscovery(dAtA, i, uint64(len(m.NetAddress)))
		i += copy(dAtA[i:], m.NetAddress)
	}
	if len(m.PublicKey) > 0 {
		dAtA[i] = 0x62
		i++
		i = encodeVarintDiscovery(dAtA, i, uint64(len(m.PublicKey)))
		i += copy(dAtA[i:], m.PublicKey)
	}
	if len(m.Certificate) > 0 {
		dAtA[i] = 0x6a
		i++
		i = encodeVarintDiscovery(dAtA, i, uint64(len(m.Certificate)))
		i += copy(dAtA[i:], m.Certificate)
	}
	return i, nil
}

func encodeVarintDiscovery(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedGetRequest(r randyDiscovery, easy bool) *GetRequest {
	this := &GetRequest{}
	this.ID = string(randStringDiscovery(r))
	this.ServiceName = string(randStringDiscovery(r))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedAnnouncement(r randyDiscovery, easy bool) *Announcement {
	this := &Announcement{}
	this.ID = string(randStringDiscovery(r))
	this.ServiceName = string(randStringDiscovery(r))
	this.ServiceVersion = string(randStringDiscovery(r))
	this.NetAddress = string(randStringDiscovery(r))
	this.PublicKey = string(randStringDiscovery(r))
	this.Certificate = string(randStringDiscovery(r))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyDiscovery interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneDiscovery(r randyDiscovery) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringDiscovery(r randyDiscovery) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneDiscovery(r)
	}
	return string(tmps)
}
func randUnrecognizedDiscovery(r randyDiscovery, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldDiscovery(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldDiscovery(dAtA []byte, r randyDiscovery, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateDiscovery(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateDiscovery(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateDiscovery(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateDiscovery(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateDiscovery(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateDiscovery(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateDiscovery(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *GetRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovDiscovery(uint64(l))
	}
	l = len(m.ServiceName)
	if l > 0 {
		n += 1 + l + sovDiscovery(uint64(l))
	}
	return n
}

func (m *Announcement) Size() (n int) {
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovDiscovery(uint64(l))
	}
	l = len(m.ServiceName)
	if l > 0 {
		n += 1 + l + sovDiscovery(uint64(l))
	}
	l = len(m.ServiceVersion)
	if l > 0 {
		n += 1 + l + sovDiscovery(uint64(l))
	}
	l = len(m.NetAddress)
	if l > 0 {
		n += 1 + l + sovDiscovery(uint64(l))
	}
	l = len(m.PublicKey)
	if l > 0 {
		n += 1 + l + sovDiscovery(uint64(l))
	}
	l = len(m.Certificate)
	if l > 0 {
		n += 1 + l + sovDiscovery(uint64(l))
	}
	return n
}

func sovDiscovery(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDiscovery(x uint64) (n int) {
	return sovDiscovery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *GetRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetRequest{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`ServiceName:` + fmt.Sprintf("%v", this.ServiceName) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Announcement) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Announcement{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`ServiceName:` + fmt.Sprintf("%v", this.ServiceName) + `,`,
		`ServiceVersion:` + fmt.Sprintf("%v", this.ServiceVersion) + `,`,
		`NetAddress:` + fmt.Sprintf("%v", this.NetAddress) + `,`,
		`PublicKey:` + fmt.Sprintf("%v", this.PublicKey) + `,`,
		`Certificate:` + fmt.Sprintf("%v", this.Certificate) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringDiscovery(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *GetRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDiscovery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscovery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDiscovery
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscovery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDiscovery
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDiscovery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDiscovery
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
func (m *Announcement) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDiscovery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Announcement: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Announcement: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscovery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDiscovery
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscovery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDiscovery
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscovery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDiscovery
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscovery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDiscovery
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NetAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscovery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDiscovery
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Certificate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscovery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDiscovery
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Certificate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDiscovery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDiscovery
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
func skipDiscovery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDiscovery
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
					return 0, ErrIntOverflowDiscovery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDiscovery
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthDiscovery
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDiscovery
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipDiscovery(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthDiscovery = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDiscovery   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("discovery/discovery.proto", fileDescriptor_discovery_cada22fcf56e9ed8)
}
func init() {
	golang_proto.RegisterFile("discovery/discovery.proto", fileDescriptor_discovery_cada22fcf56e9ed8)
}

var fileDescriptor_discovery_cada22fcf56e9ed8 = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xb1, 0x6f, 0x13, 0x31,
	0x14, 0xc6, 0xfd, 0x52, 0xa9, 0x52, 0x9c, 0x00, 0x92, 0x25, 0xe0, 0xa8, 0xc4, 0x6b, 0xe9, 0x02,
	0x4b, 0x2e, 0x12, 0x15, 0x7f, 0x40, 0xa3, 0x4a, 0x11, 0x42, 0x62, 0xc8, 0xc0, 0xc0, 0x12, 0x5d,
	0x7c, 0xaf, 0x17, 0x2b, 0x89, 0x5d, 0xee, 0x7c, 0xa9, 0xb2, 0x75, 0xec, 0xc8, 0xc8, 0xc8, 0x58,
	0xb6, 0x8e, 0x1d, 0x3b, 0x76, 0xec, 0xc0, 0xd0, 0x09, 0xf5, 0xec, 0xa5, 0x63, 0xc7, 0x8e, 0x08,
	0x5f, 0xdb, 0x64, 0x67, 0x7b, 0xef, 0xe7, 0x9f, 0x2d, 0xeb, 0xd3, 0xc7, 0x5f, 0xa5, 0xaa, 0x90,
	0x66, 0x4e, 0xf9, 0xa2, 0xfb, 0x38, 0xc5, 0x07, 0xb9, 0xb1, 0x46, 0x34, 0x1f, 0xc1, 0x46, 0x27,
	0x53, 0x76, 0x5c, 0x8e, 0x62, 0x69, 0x66, 0xdd, 0xcc, 0x64, 0xa6, 0x1b, 0x8c, 0x51, 0xb9, 0x1f,
	0xb6, 0xb0, 0x84, 0xa9, 0xbe, 0xb9, 0xdd, 0xe7, 0xbc, 0x4f, 0x76, 0x40, 0xdf, 0x4a, 0x2a, 0xac,
	0x78, 0xc1, 0x1b, 0x2a, 0x8d, 0x60, 0x0b, 0xde, 0x35, 0x7b, 0xeb, 0xee, 0xcf, 0x66, 0xe3, 0xe3,
	0xde, 0xa0, 0xa1, 0x52, 0xf1, 0x86, 0xb7, 0x0b, 0xca, 0xe7, 0x4a, 0xd2, 0x50, 0x27, 0x33, 0x8a,
	0x1a, 0xff, 0x8c, 0x41, 0xeb, 0x9e, 0x7d, 0x4e, 0x66, 0xb4, 0xfd, 0x1b, 0x78, 0x7b, 0x57, 0x6b,
	0x53, 0x6a, 0x49, 0x33, 0xd2, 0xff, 0xf3, 0x96, 0x78, 0xcb, 0x9f, 0x3d, 0x28, 0x73, 0xca, 0x0b,
	0x65, 0x74, 0xb4, 0x16, 0xac, 0xa7, 0xf7, 0xf8, 0x4b, 0x4d, 0xc5, 0x26, 0x6f, 0x69, 0xb2, 0xc3,
	0x24, 0x4d, 0x73, 0x2a, 0x8a, 0xa8, 0x15, 0x24, 0xae, 0xc9, 0xee, 0xd6, 0x44, 0xbc, 0xe6, 0xfc,
	0xa0, 0x1c, 0x4d, 0x95, 0x1c, 0x4e, 0x68, 0x11, 0xb5, 0xc3, 0x79, 0xb3, 0x26, 0x9f, 0x68, 0x21,
	0xb6, 0x78, 0x4b, 0x52, 0x6e, 0xd5, 0xbe, 0x92, 0x89, 0xa5, 0xe8, 0x49, 0xfd, 0x95, 0x15, 0xf4,
	0xbe, 0xc7, 0x9b, 0x7b, 0x0f, 0xd9, 0x8a, 0x0f, 0x7c, 0xad, 0x4f, 0x56, 0x3c, 0x8f, 0x97, 0xf9,
	0x2f, 0xc3, 0xdb, 0x78, 0xb9, 0x82, 0x57, 0x93, 0xe8, 0xfd, 0x82, 0x8b, 0x0a, 0xe1, 0xb2, 0x42,
	0xb8, 0xaa, 0x90, 0x5d, 0x57, 0xc8, 0x6e, 0x2a, 0x64, 0xb7, 0x15, 0xb2, 0xbb, 0x0a, 0xe1, 0xc8,
	0x21, 0x1c, 0x3b, 0x64, 0x27, 0x0e, 0xe1, 0xd4, 0x21, 0x3b, 0x73, 0xc8, 0xce, 0x1d, 0xb2, 0x0b,
	0x87, 0x70, 0xe9, 0x10, 0xae, 0x1c, 0xb2, 0x6b, 0x87, 0x70, 0xe3, 0x90, 0xdd, 0x3a, 0x84, 0x3b,
	0x87, 0xec, 0xc8, 0x23, 0x3b, 0xf6, 0x08, 0xdf, 0x3d, 0xb2, 0x1f, 0x1e, 0xe1, 0xa7, 0x47, 0x76,
	0xe2, 0x91, 0x9d, 0x7a, 0x84, 0x33, 0x8f, 0x70, 0xee, 0x11, 0xbe, 0xee, 0x64, 0x26, 0xb6, 0x63,
	0xb2, 0x63, 0xa5, 0xb3, 0x22, 0xd6, 0x64, 0x0f, 0x4d, 0x3e, 0xe9, 0x4e, 0x4d, 0x9e, 0x1c, 0x26,
	0xba, 0x53, 0xd8, 0x44, 0x4e, 0x3a, 0x53, 0xca, 0x12, 0xb9, 0xd2, 0xa7, 0xd1, 0x7a, 0xa8, 0xc5,
	0xce, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf7, 0xc8, 0xff, 0x69, 0x6d, 0x02, 0x00, 0x00,
}
