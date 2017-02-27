// Code generated by protoc-gen-gogo.
// source: stat.proto
// DO NOT EDIT!

/*
	Package diffcopy is a generated protocol buffer package.

	It is generated from these files:
		stat.proto

	It has these top-level messages:
		Stat
*/
package diffcopy

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Stat struct {
	Path    string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Mode    uint32 `protobuf:"varint,2,opt,name=mode,proto3" json:"mode,omitempty"`
	Uid     uint32 `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
	Gid     uint32 `protobuf:"varint,4,opt,name=gid,proto3" json:"gid,omitempty"`
	Size_   int64  `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	ModTime int64  `protobuf:"varint,6,opt,name=modTime,proto3" json:"modTime,omitempty"`
	// int32 typeflag = 7;
	Linkname string `protobuf:"bytes,7,opt,name=linkname,proto3" json:"linkname,omitempty"`
	Devmajor int64  `protobuf:"varint,8,opt,name=devmajor,proto3" json:"devmajor,omitempty"`
	Devminor int64  `protobuf:"varint,9,opt,name=devminor,proto3" json:"devminor,omitempty"`
}

func (m *Stat) Reset()         { *m = Stat{} }
func (m *Stat) String() string { return proto.CompactTextString(m) }
func (*Stat) ProtoMessage()    {}

func init() {
	proto.RegisterType((*Stat)(nil), "diffcopy.Stat")
}
func (m *Stat) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Stat) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Path) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintStat(data, i, uint64(len(m.Path)))
		i += copy(data[i:], m.Path)
	}
	if m.Mode != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintStat(data, i, uint64(m.Mode))
	}
	if m.Uid != 0 {
		data[i] = 0x18
		i++
		i = encodeVarintStat(data, i, uint64(m.Uid))
	}
	if m.Gid != 0 {
		data[i] = 0x20
		i++
		i = encodeVarintStat(data, i, uint64(m.Gid))
	}
	if m.Size_ != 0 {
		data[i] = 0x28
		i++
		i = encodeVarintStat(data, i, uint64(m.Size_))
	}
	if m.ModTime != 0 {
		data[i] = 0x30
		i++
		i = encodeVarintStat(data, i, uint64(m.ModTime))
	}
	if len(m.Linkname) > 0 {
		data[i] = 0x3a
		i++
		i = encodeVarintStat(data, i, uint64(len(m.Linkname)))
		i += copy(data[i:], m.Linkname)
	}
	if m.Devmajor != 0 {
		data[i] = 0x40
		i++
		i = encodeVarintStat(data, i, uint64(m.Devmajor))
	}
	if m.Devminor != 0 {
		data[i] = 0x48
		i++
		i = encodeVarintStat(data, i, uint64(m.Devminor))
	}
	return i, nil
}

func encodeFixed64Stat(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Stat(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintStat(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *Stat) Size() (n int) {
	var l int
	_ = l
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovStat(uint64(l))
	}
	if m.Mode != 0 {
		n += 1 + sovStat(uint64(m.Mode))
	}
	if m.Uid != 0 {
		n += 1 + sovStat(uint64(m.Uid))
	}
	if m.Gid != 0 {
		n += 1 + sovStat(uint64(m.Gid))
	}
	if m.Size_ != 0 {
		n += 1 + sovStat(uint64(m.Size_))
	}
	if m.ModTime != 0 {
		n += 1 + sovStat(uint64(m.ModTime))
	}
	l = len(m.Linkname)
	if l > 0 {
		n += 1 + l + sovStat(uint64(l))
	}
	if m.Devmajor != 0 {
		n += 1 + sovStat(uint64(m.Devmajor))
	}
	if m.Devminor != 0 {
		n += 1 + sovStat(uint64(m.Devminor))
	}
	return n
}

func sovStat(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozStat(x uint64) (n int) {
	return sovStat(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Stat) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStat
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Stat: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Stat: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStat
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mode", wireType)
			}
			m.Mode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Mode |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			m.Uid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Uid |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gid", wireType)
			}
			m.Gid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Gid |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			m.Size_ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Size_ |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModTime", wireType)
			}
			m.ModTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.ModTime |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Linkname", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStat
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Linkname = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Devmajor", wireType)
			}
			m.Devmajor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Devmajor |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Devminor", wireType)
			}
			m.Devminor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Devminor |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStat(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStat
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
func skipStat(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStat
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowStat
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowStat
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthStat
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowStat
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipStat(data[start:])
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
	ErrInvalidLengthStat = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStat   = fmt.Errorf("proto: integer overflow")
)
