// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: warden/warden/v1beta3/signature.proto

package v1beta3

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	proto "github.com/cosmos/gogoproto/proto"
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

// SignRequestStatus indicates the status of a signature request.
//
// The possible state transitions are:
//   - PENDING -> FULFILLED
//   - PENDING -> REJECTED
type SignRequestStatus int32

const (
	// The request is missing the status field.
	SignRequestStatus_SIGN_REQUEST_STATUS_UNSPECIFIED SignRequestStatus = 0
	// The request is waiting to be fulfilled. This is the initial state of a
	// request.
	SignRequestStatus_SIGN_REQUEST_STATUS_PENDING SignRequestStatus = 1
	// The request was fulfilled. This is a final state for a request.
	SignRequestStatus_SIGN_REQUEST_STATUS_FULFILLED SignRequestStatus = 2
	// The request was rejected. This is a final state for a request.
	SignRequestStatus_SIGN_REQUEST_STATUS_REJECTED SignRequestStatus = 3
)

var SignRequestStatus_name = map[int32]string{
	0: "SIGN_REQUEST_STATUS_UNSPECIFIED",
	1: "SIGN_REQUEST_STATUS_PENDING",
	2: "SIGN_REQUEST_STATUS_FULFILLED",
	3: "SIGN_REQUEST_STATUS_REJECTED",
}

var SignRequestStatus_value = map[string]int32{
	"SIGN_REQUEST_STATUS_UNSPECIFIED": 0,
	"SIGN_REQUEST_STATUS_PENDING":     1,
	"SIGN_REQUEST_STATUS_FULFILLED":   2,
	"SIGN_REQUEST_STATUS_REJECTED":    3,
}

func (x SignRequestStatus) String() string {
	return proto.EnumName(SignRequestStatus_name, int32(x))
}

func (SignRequestStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_950b6dc74b0989bd, []int{0}
}

// SignRequest is the request from a user (creator) to a Keychain to sign a
// message (data_for_signing).
//
// Once that the Keychain has received the request, it will either fulfill it
// or reject it. The result of the request will be stored in the result field.
type SignRequest struct {
	// Unique id of the request.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Address of the creator of the request.
	Creator string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	// Key ID of the key to be used for signing.
	KeyId uint64 `protobuf:"varint,3,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	// Data to be signed.
	DataForSigning []byte `protobuf:"bytes,4,opt,name=data_for_signing,json=dataForSigning,proto3" json:"data_for_signing,omitempty"`
	// Status of the request.
	Status SignRequestStatus `protobuf:"varint,5,opt,name=status,proto3,enum=warden.warden.v1beta3.SignRequestStatus" json:"status,omitempty"`
	// Result of the request, depending on the status:
	//   If pending, this field is empty.
	//   If approved, this field contains the signed data.
	//   If rejected, this field contains the reason.
	//
	// Types that are valid to be assigned to Result:
	//	*SignRequest_SignedData
	//	*SignRequest_RejectReason
	Result        isSignRequest_Result `protobuf_oneof:"result"`
	EncryptionKey []byte               `protobuf:"bytes,8,opt,name=encryption_key,json=encryptionKey,proto3" json:"encryption_key,omitempty"`
}

func (m *SignRequest) Reset()         { *m = SignRequest{} }
func (m *SignRequest) String() string { return proto.CompactTextString(m) }
func (*SignRequest) ProtoMessage()    {}
func (*SignRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_950b6dc74b0989bd, []int{0}
}
func (m *SignRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignRequest.Merge(m, src)
}
func (m *SignRequest) XXX_Size() int {
	return m.Size()
}
func (m *SignRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignRequest proto.InternalMessageInfo

type isSignRequest_Result interface {
	isSignRequest_Result()
	MarshalTo([]byte) (int, error)
	Size() int
}

type SignRequest_SignedData struct {
	SignedData []byte `protobuf:"bytes,6,opt,name=signed_data,json=signedData,proto3,oneof" json:"signed_data,omitempty"`
}
type SignRequest_RejectReason struct {
	RejectReason string `protobuf:"bytes,7,opt,name=reject_reason,json=rejectReason,proto3,oneof" json:"reject_reason,omitempty"`
}

func (*SignRequest_SignedData) isSignRequest_Result()   {}
func (*SignRequest_RejectReason) isSignRequest_Result() {}

func (m *SignRequest) GetResult() isSignRequest_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *SignRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SignRequest) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *SignRequest) GetKeyId() uint64 {
	if m != nil {
		return m.KeyId
	}
	return 0
}

func (m *SignRequest) GetDataForSigning() []byte {
	if m != nil {
		return m.DataForSigning
	}
	return nil
}

func (m *SignRequest) GetStatus() SignRequestStatus {
	if m != nil {
		return m.Status
	}
	return SignRequestStatus_SIGN_REQUEST_STATUS_UNSPECIFIED
}

func (m *SignRequest) GetSignedData() []byte {
	if x, ok := m.GetResult().(*SignRequest_SignedData); ok {
		return x.SignedData
	}
	return nil
}

func (m *SignRequest) GetRejectReason() string {
	if x, ok := m.GetResult().(*SignRequest_RejectReason); ok {
		return x.RejectReason
	}
	return ""
}

func (m *SignRequest) GetEncryptionKey() []byte {
	if m != nil {
		return m.EncryptionKey
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SignRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SignRequest_SignedData)(nil),
		(*SignRequest_RejectReason)(nil),
	}
}

func init() {
	proto.RegisterEnum("warden.warden.v1beta3.SignRequestStatus", SignRequestStatus_name, SignRequestStatus_value)
	proto.RegisterType((*SignRequest)(nil), "warden.warden.v1beta3.SignRequest")
}

func init() {
	proto.RegisterFile("warden/warden/v1beta3/signature.proto", fileDescriptor_950b6dc74b0989bd)
}

var fileDescriptor_950b6dc74b0989bd = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4f, 0x6e, 0xd3, 0x40,
	0x14, 0xc6, 0x3d, 0x6e, 0xeb, 0x96, 0xd7, 0x34, 0x38, 0x23, 0x55, 0xb2, 0xf8, 0xe3, 0x1a, 0x50,
	0x25, 0xab, 0x12, 0x8e, 0x68, 0x37, 0x2c, 0x21, 0xb5, 0x93, 0x1a, 0xa2, 0xa8, 0x8c, 0x93, 0x0d,
	0x1b, 0x6b, 0x62, 0x0f, 0xc1, 0xa4, 0xf5, 0x84, 0xf1, 0x18, 0xf0, 0x2d, 0x38, 0x01, 0x67, 0x40,
	0x9c, 0x82, 0x65, 0x97, 0xb0, 0x43, 0xc9, 0x82, 0x6b, 0x20, 0x3b, 0x0e, 0x54, 0x34, 0x1b, 0x8f,
	0xbf, 0xef, 0xfd, 0x66, 0xde, 0xf7, 0xa4, 0x07, 0x87, 0x1f, 0xa9, 0x88, 0x59, 0xda, 0xae, 0x8f,
	0x0f, 0x4f, 0xc6, 0x4c, 0xd2, 0x93, 0x76, 0x96, 0x4c, 0x52, 0x2a, 0x73, 0xc1, 0x9c, 0x99, 0xe0,
	0x92, 0xe3, 0xfd, 0x65, 0xdd, 0xa9, 0x8f, 0x1a, 0xbb, 0xd3, 0xa2, 0x97, 0x49, 0xca, 0xdb, 0xd5,
	0x77, 0x49, 0x3e, 0xfc, 0xa9, 0xc2, 0x6e, 0x90, 0x4c, 0x52, 0xc2, 0xde, 0xe7, 0x2c, 0x93, 0xb8,
	0x09, 0x6a, 0x12, 0x1b, 0xc8, 0x42, 0xf6, 0x26, 0x51, 0x93, 0x18, 0x1b, 0xb0, 0x1d, 0x09, 0x46,
	0x25, 0x17, 0x86, 0x6a, 0x21, 0xfb, 0x16, 0x59, 0x49, 0xbc, 0x0f, 0xda, 0x94, 0x15, 0x61, 0x12,
	0x1b, 0x1b, 0x15, 0xbd, 0x35, 0x65, 0x85, 0x1f, 0x63, 0x1b, 0xf4, 0x98, 0x4a, 0x1a, 0xbe, 0xe1,
	0x22, 0x2c, 0x63, 0x25, 0xe9, 0xc4, 0xd8, 0xb4, 0x90, 0xdd, 0x20, 0xcd, 0xd2, 0xef, 0x72, 0x11,
	0x2c, 0x5d, 0xfc, 0x0c, 0xb4, 0x4c, 0x52, 0x99, 0x67, 0xc6, 0x96, 0x85, 0xec, 0xe6, 0xb1, 0xed,
	0xac, 0x4d, 0xed, 0x5c, 0x8b, 0x17, 0x54, 0x3c, 0xa9, 0xef, 0xe1, 0x63, 0xd8, 0x2d, 0x5b, 0xb0,
	0x38, 0x2c, 0x9f, 0x36, 0xb4, 0xb2, 0x4d, 0xe7, 0xf6, 0xb7, 0xdf, 0x5f, 0x8f, 0x20, 0xa8, 0x7c,
	0x97, 0x4a, 0x7a, 0xa6, 0x10, 0xc8, 0xfe, 0x2a, 0xfc, 0x14, 0xf6, 0x04, 0x7b, 0xc7, 0x22, 0x19,
	0x0a, 0x46, 0x33, 0x9e, 0x1a, 0xdb, 0xe5, 0x58, 0x9d, 0x56, 0x79, 0xab, 0x41, 0xaa, 0x0a, 0xa9,
	0x0a, 0x67, 0x0a, 0x69, 0x88, 0x6b, 0x1a, 0x1f, 0x42, 0x93, 0xa5, 0x91, 0x28, 0x66, 0x32, 0xe1,
	0x69, 0x38, 0x65, 0x85, 0xb1, 0x53, 0xcd, 0xb5, 0xf7, 0xcf, 0x7d, 0xc9, 0x8a, 0xce, 0x0e, 0x68,
	0x82, 0x65, 0xf9, 0x85, 0x3c, 0xfa, 0x82, 0xa0, 0x75, 0x23, 0x3c, 0x7e, 0x04, 0x07, 0x81, 0xdf,
	0x1b, 0x84, 0xc4, 0x7b, 0x35, 0xf2, 0x82, 0x61, 0x18, 0x0c, 0x9f, 0x0f, 0x47, 0x41, 0x38, 0x1a,
	0x04, 0xe7, 0xde, 0xa9, 0xdf, 0xf5, 0x3d, 0x57, 0x57, 0xf0, 0x01, 0xdc, 0x5d, 0x07, 0x9d, 0x7b,
	0x03, 0xd7, 0x1f, 0xf4, 0x74, 0x84, 0x1f, 0xc0, 0xfd, 0x75, 0x40, 0x77, 0xd4, 0xef, 0xfa, 0xfd,
	0xbe, 0xe7, 0xea, 0x2a, 0xb6, 0xe0, 0xde, 0x3a, 0x84, 0x78, 0x2f, 0xbc, 0xd3, 0xa1, 0xe7, 0xea,
	0x1b, 0x1d, 0xfa, 0x7d, 0x6e, 0xa2, 0xab, 0xb9, 0x89, 0x7e, 0xcd, 0x4d, 0xf4, 0x79, 0x61, 0x2a,
	0x57, 0x0b, 0x53, 0xf9, 0xb1, 0x30, 0x95, 0xd7, 0xbd, 0x49, 0x22, 0xdf, 0xe6, 0x63, 0x27, 0xe2,
	0x97, 0xf5, 0xae, 0x3d, 0xae, 0xf6, 0x25, 0xe2, 0x17, 0xb5, 0xfe, 0x4f, 0xb6, 0x3f, 0xad, 0x7e,
	0x64, 0x31, 0x63, 0xd9, 0x6a, 0x33, 0xc7, 0x5a, 0xc5, 0x9d, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff,
	0x75, 0xf5, 0xe6, 0x8f, 0xb9, 0x02, 0x00, 0x00,
}

func (m *SignRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EncryptionKey) > 0 {
		i -= len(m.EncryptionKey)
		copy(dAtA[i:], m.EncryptionKey)
		i = encodeVarintSignature(dAtA, i, uint64(len(m.EncryptionKey)))
		i--
		dAtA[i] = 0x42
	}
	if m.Result != nil {
		{
			size := m.Result.Size()
			i -= size
			if _, err := m.Result.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if m.Status != 0 {
		i = encodeVarintSignature(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x28
	}
	if len(m.DataForSigning) > 0 {
		i -= len(m.DataForSigning)
		copy(dAtA[i:], m.DataForSigning)
		i = encodeVarintSignature(dAtA, i, uint64(len(m.DataForSigning)))
		i--
		dAtA[i] = 0x22
	}
	if m.KeyId != 0 {
		i = encodeVarintSignature(dAtA, i, uint64(m.KeyId))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintSignature(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintSignature(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SignRequest_SignedData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignRequest_SignedData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.SignedData != nil {
		i -= len(m.SignedData)
		copy(dAtA[i:], m.SignedData)
		i = encodeVarintSignature(dAtA, i, uint64(len(m.SignedData)))
		i--
		dAtA[i] = 0x32
	}
	return len(dAtA) - i, nil
}
func (m *SignRequest_RejectReason) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignRequest_RejectReason) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.RejectReason)
	copy(dAtA[i:], m.RejectReason)
	i = encodeVarintSignature(dAtA, i, uint64(len(m.RejectReason)))
	i--
	dAtA[i] = 0x3a
	return len(dAtA) - i, nil
}
func encodeVarintSignature(dAtA []byte, offset int, v uint64) int {
	offset -= sovSignature(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SignRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovSignature(uint64(m.Id))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovSignature(uint64(l))
	}
	if m.KeyId != 0 {
		n += 1 + sovSignature(uint64(m.KeyId))
	}
	l = len(m.DataForSigning)
	if l > 0 {
		n += 1 + l + sovSignature(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovSignature(uint64(m.Status))
	}
	if m.Result != nil {
		n += m.Result.Size()
	}
	l = len(m.EncryptionKey)
	if l > 0 {
		n += 1 + l + sovSignature(uint64(l))
	}
	return n
}

func (m *SignRequest_SignedData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SignedData != nil {
		l = len(m.SignedData)
		n += 1 + l + sovSignature(uint64(l))
	}
	return n
}
func (m *SignRequest_RejectReason) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RejectReason)
	n += 1 + l + sovSignature(uint64(l))
	return n
}

func sovSignature(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSignature(x uint64) (n int) {
	return sovSignature(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SignRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSignature
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
			return fmt.Errorf("proto: SignRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignature
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignature
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
				return ErrInvalidLengthSignature
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSignature
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyId", wireType)
			}
			m.KeyId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignature
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeyId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataForSigning", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignature
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
				return ErrInvalidLengthSignature
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSignature
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataForSigning = append(m.DataForSigning[:0], dAtA[iNdEx:postIndex]...)
			if m.DataForSigning == nil {
				m.DataForSigning = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignature
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= SignRequestStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignedData", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignature
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
				return ErrInvalidLengthSignature
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSignature
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, postIndex-iNdEx)
			copy(v, dAtA[iNdEx:postIndex])
			m.Result = &SignRequest_SignedData{v}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RejectReason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignature
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
				return ErrInvalidLengthSignature
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSignature
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Result = &SignRequest_RejectReason{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EncryptionKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignature
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
				return ErrInvalidLengthSignature
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSignature
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EncryptionKey = append(m.EncryptionKey[:0], dAtA[iNdEx:postIndex]...)
			if m.EncryptionKey == nil {
				m.EncryptionKey = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSignature(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSignature
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
func skipSignature(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSignature
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
					return 0, ErrIntOverflowSignature
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
					return 0, ErrIntOverflowSignature
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
				return 0, ErrInvalidLengthSignature
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSignature
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSignature
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSignature        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSignature          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSignature = fmt.Errorf("proto: unexpected end of group")
)