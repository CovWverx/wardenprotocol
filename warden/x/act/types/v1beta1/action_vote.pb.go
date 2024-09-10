// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: warden/act/v1beta1/action_vote.proto

package v1beta1

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Type of a vote.
type ActionVoteType int32

const (
	// Unspecified vote type.
	ActionVoteType_VOTE_TYPE_UNSPECIFIED ActionVoteType = 0
	// Positive vote for an action.
	ActionVoteType_VOTE_TYPE_APPROVED ActionVoteType = 1
	// Negative vote for an action.
	ActionVoteType_VOTE_TYPE_REJECTED ActionVoteType = 2
)

var ActionVoteType_name = map[int32]string{
	0: "VOTE_TYPE_UNSPECIFIED",
	1: "VOTE_TYPE_APPROVED",
	2: "VOTE_TYPE_REJECTED",
}

var ActionVoteType_value = map[string]int32{
	"VOTE_TYPE_UNSPECIFIED": 0,
	"VOTE_TYPE_APPROVED":    1,
	"VOTE_TYPE_REJECTED":    2,
}

func (x ActionVoteType) String() string {
	return proto.EnumName(ActionVoteType_name, int32(x))
}

func (ActionVoteType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b33d294255f825be, []int{0}
}

type ActionVote struct {
	// participant is the address of the voter.
	Participant string `protobuf:"bytes,1,opt,name=participant,proto3" json:"participant,omitempty"`
	// voted_at is a timestamp specifying when the voter voted on the action.
	VotedAt time.Time `protobuf:"bytes,2,opt,name=voted_at,json=votedAt,proto3,stdtime" json:"voted_at"`
	// vote is the type of the vote.
	VoteType ActionVoteType `protobuf:"varint,3,opt,name=vote_type,json=voteType,proto3,enum=warden.act.v1beta1.ActionVoteType" json:"vote_type,omitempty"`
}

func (m *ActionVote) Reset()         { *m = ActionVote{} }
func (m *ActionVote) String() string { return proto.CompactTextString(m) }
func (*ActionVote) ProtoMessage()    {}
func (*ActionVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_b33d294255f825be, []int{0}
}
func (m *ActionVote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ActionVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ActionVote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ActionVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionVote.Merge(m, src)
}
func (m *ActionVote) XXX_Size() int {
	return m.Size()
}
func (m *ActionVote) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionVote.DiscardUnknown(m)
}

var xxx_messageInfo_ActionVote proto.InternalMessageInfo

func (m *ActionVote) GetParticipant() string {
	if m != nil {
		return m.Participant
	}
	return ""
}

func (m *ActionVote) GetVotedAt() time.Time {
	if m != nil {
		return m.VotedAt
	}
	return time.Time{}
}

func (m *ActionVote) GetVoteType() ActionVoteType {
	if m != nil {
		return m.VoteType
	}
	return ActionVoteType_VOTE_TYPE_UNSPECIFIED
}

func init() {
	proto.RegisterEnum("warden.act.v1beta1.ActionVoteType", ActionVoteType_name, ActionVoteType_value)
	proto.RegisterType((*ActionVote)(nil), "warden.act.v1beta1.ActionVote")
}

func init() {
	proto.RegisterFile("warden/act/v1beta1/action_vote.proto", fileDescriptor_b33d294255f825be)
}

var fileDescriptor_b33d294255f825be = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x29, 0x4f, 0x2c, 0x4a,
	0x49, 0xcd, 0xd3, 0x4f, 0x4c, 0x2e, 0xd1, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0x04, 0xb1,
	0x33, 0xf3, 0xf3, 0xe2, 0xcb, 0xf2, 0x4b, 0x52, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x84,
	0x20, 0xaa, 0xf4, 0x12, 0x93, 0x4b, 0xf4, 0xa0, 0xaa, 0xa4, 0x04, 0x13, 0x73, 0x33, 0xf3, 0xf2,
	0xf5, 0xc1, 0x24, 0x44, 0x99, 0x94, 0x48, 0x7a, 0x7e, 0x7a, 0x3e, 0x98, 0xa9, 0x0f, 0x62, 0x41,
	0x45, 0xe5, 0xd3, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0xc1, 0xbc, 0xa4, 0xd2, 0x34, 0xfd, 0x92,
	0xcc, 0xdc, 0xd4, 0xe2, 0x92, 0xc4, 0xdc, 0x02, 0x88, 0x02, 0xa5, 0xad, 0x8c, 0x5c, 0x5c, 0x8e,
	0x60, 0x3b, 0xc3, 0xf2, 0x4b, 0x52, 0x85, 0x14, 0xb8, 0xb8, 0x0b, 0x12, 0x8b, 0x4a, 0x32, 0x93,
	0x33, 0x0b, 0x12, 0xf3, 0x4a, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x90, 0x85, 0x84, 0x5c,
	0xb8, 0x38, 0x40, 0x8e, 0x4b, 0x89, 0x4f, 0x2c, 0x91, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x36, 0x92,
	0xd2, 0x83, 0x58, 0xa2, 0x07, 0xb3, 0x44, 0x2f, 0x04, 0x66, 0x89, 0x13, 0xef, 0x89, 0x7b, 0xf2,
	0x0c, 0x13, 0xee, 0xcb, 0x33, 0xae, 0x78, 0xbe, 0x41, 0x8b, 0x31, 0x88, 0x1d, 0xac, 0xd5, 0xb1,
	0x44, 0xc8, 0x9e, 0x8b, 0x13, 0xc4, 0x8c, 0x2f, 0xa9, 0x2c, 0x48, 0x95, 0x60, 0x56, 0x60, 0xd4,
	0xe0, 0x33, 0x52, 0xd2, 0xc3, 0xf4, 0xa8, 0x1e, 0xc2, 0x69, 0x21, 0x95, 0x05, 0xa9, 0x41, 0x60,
	0xab, 0x41, 0x2c, 0xad, 0x68, 0x2e, 0x3e, 0x54, 0x39, 0x21, 0x49, 0x2e, 0xd1, 0x30, 0xff, 0x10,
	0xd7, 0xf8, 0x90, 0xc8, 0x00, 0xd7, 0xf8, 0x50, 0xbf, 0xe0, 0x00, 0x57, 0x67, 0x4f, 0x37, 0x4f,
	0x57, 0x17, 0x01, 0x06, 0x21, 0x31, 0x2e, 0x21, 0x84, 0x94, 0x63, 0x40, 0x40, 0x90, 0x7f, 0x98,
	0xab, 0x8b, 0x00, 0x23, 0xaa, 0x78, 0x90, 0xab, 0x97, 0xab, 0x73, 0x88, 0xab, 0x8b, 0x00, 0x93,
	0x53, 0xdc, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1,
	0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0xb9, 0xa4, 0x67, 0x96,
	0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x43, 0x9c, 0xab, 0x0b, 0xf6, 0x75, 0x72, 0x7e,
	0x0e, 0x94, 0x8f, 0xc6, 0xd5, 0xaf, 0x00, 0x47, 0x2f, 0xc8, 0x9f, 0xc5, 0xb0, 0x48, 0x4e, 0x62,
	0x03, 0x2b, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x64, 0xe6, 0x0a, 0x01, 0x02, 0x00,
	0x00,
}

func (m *ActionVote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ActionVote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ActionVote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.VoteType != 0 {
		i = encodeVarintActionVote(dAtA, i, uint64(m.VoteType))
		i--
		dAtA[i] = 0x18
	}
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.VotedAt, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.VotedAt):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintActionVote(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	if len(m.Participant) > 0 {
		i -= len(m.Participant)
		copy(dAtA[i:], m.Participant)
		i = encodeVarintActionVote(dAtA, i, uint64(len(m.Participant)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintActionVote(dAtA []byte, offset int, v uint64) int {
	offset -= sovActionVote(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ActionVote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Participant)
	if l > 0 {
		n += 1 + l + sovActionVote(uint64(l))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.VotedAt)
	n += 1 + l + sovActionVote(uint64(l))
	if m.VoteType != 0 {
		n += 1 + sovActionVote(uint64(m.VoteType))
	}
	return n
}

func sovActionVote(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozActionVote(x uint64) (n int) {
	return sovActionVote(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ActionVote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowActionVote
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
			return fmt.Errorf("proto: ActionVote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ActionVote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Participant", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionVote
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
				return ErrInvalidLengthActionVote
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActionVote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Participant = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionVote
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
				return ErrInvalidLengthActionVote
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthActionVote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.VotedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteType", wireType)
			}
			m.VoteType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionVote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VoteType |= ActionVoteType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipActionVote(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthActionVote
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
func skipActionVote(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowActionVote
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
					return 0, ErrIntOverflowActionVote
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
					return 0, ErrIntOverflowActionVote
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
				return 0, ErrInvalidLengthActionVote
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupActionVote
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthActionVote
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthActionVote        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowActionVote          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupActionVote = fmt.Errorf("proto: unexpected end of group")
)
