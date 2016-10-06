// Code generated by protoc-gen-gogo.
// source: github.com/TheThingsNetwork/ttn/api/protocol/lorawan/lorawan.proto
// DO NOT EDIT!

/*
	Package lorawan is a generated protocol buffer package.

	It is generated from these files:
		github.com/TheThingsNetwork/ttn/api/protocol/lorawan/lorawan.proto

	It has these top-level messages:
		Metadata
		TxConfiguration
		ActivationMetadata
*/
package lorawan

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import github_com_TheThingsNetwork_ttn_core_types "github.com/TheThingsNetwork/ttn/core/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Modulation int32

const (
	Modulation_LORA Modulation = 0
	Modulation_FSK  Modulation = 1
)

var Modulation_name = map[int32]string{
	0: "LORA",
	1: "FSK",
}
var Modulation_value = map[string]int32{
	"LORA": 0,
	"FSK":  1,
}

func (x Modulation) String() string {
	return proto.EnumName(Modulation_name, int32(x))
}
func (Modulation) EnumDescriptor() ([]byte, []int) { return fileDescriptorLorawan, []int{0} }

type Region int32

const (
	Region_EU_863_870 Region = 0
	Region_US_902_928 Region = 1
	Region_CN_779_787 Region = 2
	Region_EU_433     Region = 3
	Region_AU_915_928 Region = 4
	Region_CN_470_510 Region = 5
	Region_AS_923     Region = 6
	Region_SK_920_923 Region = 7
)

var Region_name = map[int32]string{
	0: "EU_863_870",
	1: "US_902_928",
	2: "CN_779_787",
	3: "EU_433",
	4: "AU_915_928",
	5: "CN_470_510",
	6: "AS_923",
	7: "SK_920_923",
}
var Region_value = map[string]int32{
	"EU_863_870": 0,
	"US_902_928": 1,
	"CN_779_787": 2,
	"EU_433":     3,
	"AU_915_928": 4,
	"CN_470_510": 5,
	"AS_923":     6,
	"SK_920_923": 7,
}

func (x Region) String() string {
	return proto.EnumName(Region_name, int32(x))
}
func (Region) EnumDescriptor() ([]byte, []int) { return fileDescriptorLorawan, []int{1} }

type Metadata struct {
	Modulation Modulation `protobuf:"varint,11,opt,name=modulation,proto3,enum=lorawan.Modulation" json:"modulation,omitempty"`
	DataRate   string     `protobuf:"bytes,12,opt,name=data_rate,json=dataRate,proto3" json:"data_rate,omitempty"`
	BitRate    uint32     `protobuf:"varint,13,opt,name=bit_rate,json=bitRate,proto3" json:"bit_rate,omitempty"`
	CodingRate string     `protobuf:"bytes,14,opt,name=coding_rate,json=codingRate,proto3" json:"coding_rate,omitempty"`
	FCnt       uint32     `protobuf:"varint,15,opt,name=f_cnt,json=fCnt,proto3" json:"f_cnt,omitempty"`
}

func (m *Metadata) Reset()                    { *m = Metadata{} }
func (m *Metadata) String() string            { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()               {}
func (*Metadata) Descriptor() ([]byte, []int) { return fileDescriptorLorawan, []int{0} }

type TxConfiguration struct {
	Modulation Modulation `protobuf:"varint,11,opt,name=modulation,proto3,enum=lorawan.Modulation" json:"modulation,omitempty"`
	DataRate   string     `protobuf:"bytes,12,opt,name=data_rate,json=dataRate,proto3" json:"data_rate,omitempty"`
	BitRate    uint32     `protobuf:"varint,13,opt,name=bit_rate,json=bitRate,proto3" json:"bit_rate,omitempty"`
	CodingRate string     `protobuf:"bytes,14,opt,name=coding_rate,json=codingRate,proto3" json:"coding_rate,omitempty"`
	FCnt       uint32     `protobuf:"varint,15,opt,name=f_cnt,json=fCnt,proto3" json:"f_cnt,omitempty"`
}

func (m *TxConfiguration) Reset()                    { *m = TxConfiguration{} }
func (m *TxConfiguration) String() string            { return proto.CompactTextString(m) }
func (*TxConfiguration) ProtoMessage()               {}
func (*TxConfiguration) Descriptor() ([]byte, []int) { return fileDescriptorLorawan, []int{1} }

type ActivationMetadata struct {
	AppEui      *github_com_TheThingsNetwork_ttn_core_types.AppEUI  `protobuf:"bytes,1,opt,name=app_eui,json=appEui,proto3,customtype=github.com/TheThingsNetwork/ttn/core/types.AppEUI" json:"app_eui,omitempty"`
	DevEui      *github_com_TheThingsNetwork_ttn_core_types.DevEUI  `protobuf:"bytes,2,opt,name=dev_eui,json=devEui,proto3,customtype=github.com/TheThingsNetwork/ttn/core/types.DevEUI" json:"dev_eui,omitempty"`
	DevAddr     *github_com_TheThingsNetwork_ttn_core_types.DevAddr `protobuf:"bytes,3,opt,name=dev_addr,json=devAddr,proto3,customtype=github.com/TheThingsNetwork/ttn/core/types.DevAddr" json:"dev_addr,omitempty"`
	NwkSKey     *github_com_TheThingsNetwork_ttn_core_types.NwkSKey `protobuf:"bytes,4,opt,name=nwk_s_key,json=nwkSKey,proto3,customtype=github.com/TheThingsNetwork/ttn/core/types.NwkSKey" json:"nwk_s_key,omitempty"`
	Rx1DrOffset uint32                                              `protobuf:"varint,11,opt,name=rx1_dr_offset,json=rx1DrOffset,proto3" json:"rx1_dr_offset,omitempty"`
	Rx2Dr       uint32                                              `protobuf:"varint,12,opt,name=rx2_dr,json=rx2Dr,proto3" json:"rx2_dr,omitempty"`
	RxDelay     uint32                                              `protobuf:"varint,13,opt,name=rx_delay,json=rxDelay,proto3" json:"rx_delay,omitempty"`
	CfList      []uint64                                            `protobuf:"varint,14,rep,name=cf_list,json=cfList" json:"cf_list,omitempty"`
}

func (m *ActivationMetadata) Reset()                    { *m = ActivationMetadata{} }
func (m *ActivationMetadata) String() string            { return proto.CompactTextString(m) }
func (*ActivationMetadata) ProtoMessage()               {}
func (*ActivationMetadata) Descriptor() ([]byte, []int) { return fileDescriptorLorawan, []int{2} }

func init() {
	proto.RegisterType((*Metadata)(nil), "lorawan.Metadata")
	proto.RegisterType((*TxConfiguration)(nil), "lorawan.TxConfiguration")
	proto.RegisterType((*ActivationMetadata)(nil), "lorawan.ActivationMetadata")
	proto.RegisterEnum("lorawan.Modulation", Modulation_name, Modulation_value)
	proto.RegisterEnum("lorawan.Region", Region_name, Region_value)
}
func (m *Metadata) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Metadata) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Modulation != 0 {
		data[i] = 0x58
		i++
		i = encodeVarintLorawan(data, i, uint64(m.Modulation))
	}
	if len(m.DataRate) > 0 {
		data[i] = 0x62
		i++
		i = encodeVarintLorawan(data, i, uint64(len(m.DataRate)))
		i += copy(data[i:], m.DataRate)
	}
	if m.BitRate != 0 {
		data[i] = 0x68
		i++
		i = encodeVarintLorawan(data, i, uint64(m.BitRate))
	}
	if len(m.CodingRate) > 0 {
		data[i] = 0x72
		i++
		i = encodeVarintLorawan(data, i, uint64(len(m.CodingRate)))
		i += copy(data[i:], m.CodingRate)
	}
	if m.FCnt != 0 {
		data[i] = 0x78
		i++
		i = encodeVarintLorawan(data, i, uint64(m.FCnt))
	}
	return i, nil
}

func (m *TxConfiguration) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *TxConfiguration) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Modulation != 0 {
		data[i] = 0x58
		i++
		i = encodeVarintLorawan(data, i, uint64(m.Modulation))
	}
	if len(m.DataRate) > 0 {
		data[i] = 0x62
		i++
		i = encodeVarintLorawan(data, i, uint64(len(m.DataRate)))
		i += copy(data[i:], m.DataRate)
	}
	if m.BitRate != 0 {
		data[i] = 0x68
		i++
		i = encodeVarintLorawan(data, i, uint64(m.BitRate))
	}
	if len(m.CodingRate) > 0 {
		data[i] = 0x72
		i++
		i = encodeVarintLorawan(data, i, uint64(len(m.CodingRate)))
		i += copy(data[i:], m.CodingRate)
	}
	if m.FCnt != 0 {
		data[i] = 0x78
		i++
		i = encodeVarintLorawan(data, i, uint64(m.FCnt))
	}
	return i, nil
}

func (m *ActivationMetadata) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ActivationMetadata) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.AppEui != nil {
		data[i] = 0xa
		i++
		i = encodeVarintLorawan(data, i, uint64(m.AppEui.Size()))
		n1, err := m.AppEui.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.DevEui != nil {
		data[i] = 0x12
		i++
		i = encodeVarintLorawan(data, i, uint64(m.DevEui.Size()))
		n2, err := m.DevEui.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.DevAddr != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintLorawan(data, i, uint64(m.DevAddr.Size()))
		n3, err := m.DevAddr.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.NwkSKey != nil {
		data[i] = 0x22
		i++
		i = encodeVarintLorawan(data, i, uint64(m.NwkSKey.Size()))
		n4, err := m.NwkSKey.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.Rx1DrOffset != 0 {
		data[i] = 0x58
		i++
		i = encodeVarintLorawan(data, i, uint64(m.Rx1DrOffset))
	}
	if m.Rx2Dr != 0 {
		data[i] = 0x60
		i++
		i = encodeVarintLorawan(data, i, uint64(m.Rx2Dr))
	}
	if m.RxDelay != 0 {
		data[i] = 0x68
		i++
		i = encodeVarintLorawan(data, i, uint64(m.RxDelay))
	}
	if len(m.CfList) > 0 {
		for _, num := range m.CfList {
			data[i] = 0x70
			i++
			i = encodeVarintLorawan(data, i, uint64(num))
		}
	}
	return i, nil
}

func encodeFixed64Lorawan(data []byte, offset int, v uint64) int {
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
func encodeFixed32Lorawan(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintLorawan(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *Metadata) Size() (n int) {
	var l int
	_ = l
	if m.Modulation != 0 {
		n += 1 + sovLorawan(uint64(m.Modulation))
	}
	l = len(m.DataRate)
	if l > 0 {
		n += 1 + l + sovLorawan(uint64(l))
	}
	if m.BitRate != 0 {
		n += 1 + sovLorawan(uint64(m.BitRate))
	}
	l = len(m.CodingRate)
	if l > 0 {
		n += 1 + l + sovLorawan(uint64(l))
	}
	if m.FCnt != 0 {
		n += 1 + sovLorawan(uint64(m.FCnt))
	}
	return n
}

func (m *TxConfiguration) Size() (n int) {
	var l int
	_ = l
	if m.Modulation != 0 {
		n += 1 + sovLorawan(uint64(m.Modulation))
	}
	l = len(m.DataRate)
	if l > 0 {
		n += 1 + l + sovLorawan(uint64(l))
	}
	if m.BitRate != 0 {
		n += 1 + sovLorawan(uint64(m.BitRate))
	}
	l = len(m.CodingRate)
	if l > 0 {
		n += 1 + l + sovLorawan(uint64(l))
	}
	if m.FCnt != 0 {
		n += 1 + sovLorawan(uint64(m.FCnt))
	}
	return n
}

func (m *ActivationMetadata) Size() (n int) {
	var l int
	_ = l
	if m.AppEui != nil {
		l = m.AppEui.Size()
		n += 1 + l + sovLorawan(uint64(l))
	}
	if m.DevEui != nil {
		l = m.DevEui.Size()
		n += 1 + l + sovLorawan(uint64(l))
	}
	if m.DevAddr != nil {
		l = m.DevAddr.Size()
		n += 1 + l + sovLorawan(uint64(l))
	}
	if m.NwkSKey != nil {
		l = m.NwkSKey.Size()
		n += 1 + l + sovLorawan(uint64(l))
	}
	if m.Rx1DrOffset != 0 {
		n += 1 + sovLorawan(uint64(m.Rx1DrOffset))
	}
	if m.Rx2Dr != 0 {
		n += 1 + sovLorawan(uint64(m.Rx2Dr))
	}
	if m.RxDelay != 0 {
		n += 1 + sovLorawan(uint64(m.RxDelay))
	}
	if len(m.CfList) > 0 {
		for _, e := range m.CfList {
			n += 1 + sovLorawan(uint64(e))
		}
	}
	return n
}

func sovLorawan(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozLorawan(x uint64) (n int) {
	return sovLorawan(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Metadata) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLorawan
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
			return fmt.Errorf("proto: Metadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Metadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Modulation", wireType)
			}
			m.Modulation = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLorawan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Modulation |= (Modulation(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLorawan
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
				return ErrInvalidLengthLorawan
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataRate = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BitRate", wireType)
			}
			m.BitRate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLorawan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.BitRate |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodingRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLorawan
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
				return ErrInvalidLengthLorawan
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CodingRate = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FCnt", wireType)
			}
			m.FCnt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLorawan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.FCnt |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLorawan(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLorawan
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
func (m *TxConfiguration) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLorawan
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
			return fmt.Errorf("proto: TxConfiguration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxConfiguration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Modulation", wireType)
			}
			m.Modulation = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLorawan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Modulation |= (Modulation(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLorawan
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
				return ErrInvalidLengthLorawan
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataRate = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BitRate", wireType)
			}
			m.BitRate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLorawan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.BitRate |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodingRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLorawan
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
				return ErrInvalidLengthLorawan
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CodingRate = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FCnt", wireType)
			}
			m.FCnt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLorawan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.FCnt |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLorawan(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLorawan
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
func (m *ActivationMetadata) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLorawan
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
			return fmt.Errorf("proto: ActivationMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ActivationMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppEui", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLorawan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthLorawan
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_TheThingsNetwork_ttn_core_types.AppEUI
			m.AppEui = &v
			if err := m.AppEui.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DevEui", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLorawan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthLorawan
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_TheThingsNetwork_ttn_core_types.DevEUI
			m.DevEui = &v
			if err := m.DevEui.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DevAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLorawan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthLorawan
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_TheThingsNetwork_ttn_core_types.DevAddr
			m.DevAddr = &v
			if err := m.DevAddr.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NwkSKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLorawan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthLorawan
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_TheThingsNetwork_ttn_core_types.NwkSKey
			m.NwkSKey = &v
			if err := m.NwkSKey.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rx1DrOffset", wireType)
			}
			m.Rx1DrOffset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLorawan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Rx1DrOffset |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rx2Dr", wireType)
			}
			m.Rx2Dr = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLorawan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Rx2Dr |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RxDelay", wireType)
			}
			m.RxDelay = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLorawan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.RxDelay |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CfList", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLorawan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CfList = append(m.CfList, v)
		default:
			iNdEx = preIndex
			skippy, err := skipLorawan(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLorawan
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
func skipLorawan(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLorawan
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
					return 0, ErrIntOverflowLorawan
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
					return 0, ErrIntOverflowLorawan
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
				return 0, ErrInvalidLengthLorawan
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLorawan
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
				next, err := skipLorawan(data[start:])
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
	ErrInvalidLengthLorawan = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLorawan   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/TheThingsNetwork/ttn/api/protocol/lorawan/lorawan.proto", fileDescriptorLorawan)
}

var fileDescriptorLorawan = []byte{
	// 601 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xd4, 0x93, 0xcf, 0x6e, 0xda, 0x4e,
	0x10, 0xc7, 0xe3, 0x40, 0x6c, 0x32, 0x09, 0xc4, 0x72, 0xf4, 0xd3, 0xcf, 0x6d, 0x25, 0x40, 0x39,
	0xa1, 0x48, 0xe5, 0x8f, 0x49, 0x02, 0x1c, 0x49, 0xa0, 0x52, 0x95, 0x84, 0xa8, 0x26, 0x9c, 0x57,
	0x8b, 0xbd, 0x76, 0x56, 0x10, 0xaf, 0xb5, 0x2c, 0x60, 0x6e, 0x7d, 0x8c, 0xbe, 0x44, 0x6f, 0x7d,
	0x88, 0xaa, 0xa7, 0x9e, 0x73, 0x88, 0xaa, 0xf4, 0x45, 0xaa, 0x5d, 0x87, 0x34, 0xb7, 0xaa, 0xb9,
	0xf5, 0xc4, 0x7c, 0xe7, 0x3b, 0xf3, 0xd9, 0x11, 0x33, 0x86, 0xd3, 0x90, 0x8a, 0x9b, 0xf9, 0xb8,
	0xea, 0xb1, 0xdb, 0xda, 0xf5, 0x0d, 0xb9, 0xbe, 0xa1, 0x51, 0x38, 0x1b, 0x10, 0xb1, 0x64, 0x7c,
	0x52, 0x13, 0x22, 0xaa, 0xe1, 0x98, 0xd6, 0x62, 0xce, 0x04, 0xf3, 0xd8, 0xb4, 0x36, 0x65, 0x1c,
	0x2f, 0x71, 0xb4, 0xfe, 0xad, 0x2a, 0xc3, 0x32, 0x1e, 0xe5, 0xeb, 0xb7, 0xcf, 0x60, 0x21, 0x0b,
	0x59, 0xda, 0x38, 0x9e, 0x07, 0x4a, 0x29, 0xa1, 0xa2, 0xb4, 0xef, 0xe0, 0xb3, 0x06, 0xb9, 0x4b,
	0x22, 0xb0, 0x8f, 0x05, 0xb6, 0x9a, 0x00, 0xb7, 0xcc, 0x9f, 0x4f, 0xb1, 0xa0, 0x2c, 0xb2, 0x77,
	0xca, 0x5a, 0xa5, 0xe0, 0xec, 0x57, 0xd7, 0x0f, 0x5d, 0x3e, 0x59, 0xee, 0xb3, 0x32, 0xeb, 0x0d,
	0x6c, 0xcb, 0x66, 0xc4, 0xb1, 0x20, 0xf6, 0x6e, 0x59, 0xab, 0x6c, 0xbb, 0x39, 0x99, 0x70, 0xb1,
	0x20, 0xd6, 0x2b, 0xc8, 0x8d, 0xa9, 0x48, 0xbd, 0x7c, 0x59, 0xab, 0xe4, 0x5d, 0x63, 0x4c, 0x85,
	0xb2, 0x4a, 0xb0, 0xe3, 0x31, 0x9f, 0x46, 0x61, 0xea, 0x16, 0x54, 0x27, 0xa4, 0x29, 0x55, 0xb0,
	0x0f, 0x5b, 0x01, 0xf2, 0x22, 0x61, 0xef, 0xa9, 0xc6, 0x6c, 0x70, 0x16, 0x89, 0x83, 0x2f, 0x1a,
	0xec, 0x5d, 0x27, 0x67, 0x2c, 0x0a, 0x68, 0x38, 0xe7, 0xe9, 0x04, 0xff, 0xc0, 0xd8, 0xdf, 0x32,
	0x60, 0x75, 0x3d, 0x41, 0x17, 0xea, 0xf1, 0xa7, 0x3f, 0x7c, 0x00, 0x06, 0x8e, 0x63, 0x44, 0xe6,
	0xd4, 0xd6, 0xca, 0x5a, 0x65, 0xf7, 0xf4, 0xf8, 0xee, 0xbe, 0xd4, 0xf8, 0xd3, 0x39, 0x78, 0x8c,
	0x93, 0x9a, 0x58, 0xc5, 0x64, 0x56, 0xed, 0xc6, 0x71, 0x7f, 0xf4, 0xde, 0xd5, 0x71, 0x1c, 0xf7,
	0xe7, 0x54, 0xf2, 0x7c, 0xb2, 0x50, 0xbc, 0xcd, 0x17, 0xf1, 0x7a, 0x64, 0xa1, 0x78, 0x3e, 0x59,
	0x48, 0xde, 0x07, 0xc8, 0x49, 0x1e, 0xf6, 0x7d, 0x6e, 0x67, 0x14, 0xf0, 0xe4, 0xee, 0xbe, 0xe4,
	0xfc, 0x1d, 0xb0, 0xeb, 0xfb, 0xdc, 0x95, 0x73, 0xc9, 0xc0, 0x72, 0x61, 0x3b, 0x5a, 0x4e, 0xd0,
	0x0c, 0x4d, 0xc8, 0xca, 0xce, 0xbe, 0x88, 0x39, 0x58, 0x4e, 0x86, 0xe7, 0x64, 0xe5, 0x1a, 0x51,
	0x1a, 0x58, 0x07, 0x90, 0xe7, 0x49, 0x03, 0xf9, 0x1c, 0xb1, 0x20, 0x98, 0x11, 0xa1, 0x6e, 0x20,
	0xef, 0xee, 0xf0, 0xa4, 0xd1, 0xe3, 0x57, 0x2a, 0x65, 0xfd, 0x07, 0x3a, 0x4f, 0x1c, 0xe4, 0x73,
	0xb5, 0xec, 0xbc, 0xbb, 0xc5, 0x13, 0xa7, 0xc7, 0xe5, 0xa6, 0x79, 0x82, 0x7c, 0x32, 0xc5, 0xab,
	0xf5, 0xa6, 0x79, 0xd2, 0x93, 0xd2, 0xfa, 0x1f, 0x0c, 0x2f, 0x40, 0x53, 0x3a, 0x13, 0x76, 0xa1,
	0x9c, 0xa9, 0x64, 0x5d, 0xdd, 0x0b, 0x2e, 0xe8, 0x4c, 0x1c, 0x96, 0x00, 0x7e, 0x1f, 0x95, 0x95,
	0x83, 0xec, 0xc5, 0x95, 0xdb, 0x35, 0x37, 0x2c, 0x03, 0x32, 0xef, 0x86, 0xe7, 0xa6, 0x76, 0xf8,
	0x51, 0x03, 0xdd, 0x25, 0xa1, 0x74, 0x0b, 0x00, 0xfd, 0x11, 0x6a, 0x9f, 0x34, 0x51, 0xbb, 0x55,
	0x37, 0x37, 0xa4, 0x1e, 0x0d, 0x51, 0xa7, 0xee, 0xa0, 0x8e, 0xd3, 0x36, 0x35, 0xa9, 0xcf, 0x06,
	0xa8, 0xd5, 0xea, 0xa0, 0x56, 0xbb, 0x65, 0x6e, 0x5a, 0x00, 0x7a, 0x7f, 0x84, 0x8e, 0x9a, 0x4d,
	0x33, 0x23, 0xbd, 0xee, 0x08, 0x75, 0x1a, 0xc7, 0xaa, 0x36, 0xfb, 0x58, 0x7b, 0xd4, 0xaa, 0xa3,
	0xe3, 0x46, 0xdd, 0xdc, 0x92, 0xb5, 0xdd, 0x21, 0xea, 0x38, 0x4d, 0x53, 0x97, 0xde, 0xf0, 0x1c,
	0x75, 0x9c, 0xba, 0xd2, 0xc6, 0xa9, 0xf9, 0xf5, 0xa1, 0xa8, 0x7d, 0x7f, 0x28, 0x6a, 0x3f, 0x1e,
	0x8a, 0xda, 0xa7, 0x9f, 0xc5, 0x8d, 0xb1, 0xae, 0x3e, 0xf8, 0xe6, 0xaf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x91, 0xac, 0x02, 0x91, 0x6e, 0x04, 0x00, 0x00,
}
