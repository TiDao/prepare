// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: store/store.proto

package store

import (
	common "chainmaker.org/chainmaker-go/pb/protogo/common"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type DbType int32

const (
	DbType_INVALID_DB     DbType = 0
	DbType_BLOCK_DB       DbType = 1
	DbType_BLOCK_INDEX_DB DbType = 2
	DbType_TX_DB          DbType = 3
	DbType_TX_INDEX_DB    DbType = 4
	DbType_SOFT_DB        DbType = 5
	DbType_STATE_DB       DbType = 6
	DbType_READ_WRITE_DB  DbType = 7
)

var DbType_name = map[int32]string{
	0: "INVALID_DB",
	1: "BLOCK_DB",
	2: "BLOCK_INDEX_DB",
	3: "TX_DB",
	4: "TX_INDEX_DB",
	5: "SOFT_DB",
	6: "STATE_DB",
	7: "READ_WRITE_DB",
}

var DbType_value = map[string]int32{
	"INVALID_DB":     0,
	"BLOCK_DB":       1,
	"BLOCK_INDEX_DB": 2,
	"TX_DB":          3,
	"TX_INDEX_DB":    4,
	"SOFT_DB":        5,
	"STATE_DB":       6,
	"READ_WRITE_DB":  7,
}

func (x DbType) String() string {
	return proto.EnumName(DbType_name, int32(x))
}

func (DbType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8549980b097f750b, []int{0}
}

// block structure used in serialization
type SerializedBlock struct {
	// header of block
	Header *common.BlockHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// transaction execution sequence of the block, described by DAG
	Dag *common.DAG `protobuf:"bytes,2,opt,name=dag,proto3" json:"dag,omitempty"`
	// transaction id list within the block
	TxIds []string `protobuf:"bytes,3,rep,name=tx_ids,json=txIds,proto3" json:"tx_ids,omitempty"`
	// block additional data, not included in block hash calculation
	AdditionalData *common.AdditionalData `protobuf:"bytes,4,opt,name=additional_data,json=additionalData,proto3" json:"additional_data,omitempty"`
}

func (m *SerializedBlock) Reset()         { *m = SerializedBlock{} }
func (m *SerializedBlock) String() string { return proto.CompactTextString(m) }
func (*SerializedBlock) ProtoMessage()    {}
func (*SerializedBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_8549980b097f750b, []int{0}
}
func (m *SerializedBlock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SerializedBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SerializedBlock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SerializedBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SerializedBlock.Merge(m, src)
}
func (m *SerializedBlock) XXX_Size() int {
	return m.Size()
}
func (m *SerializedBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_SerializedBlock.DiscardUnknown(m)
}

var xxx_messageInfo_SerializedBlock proto.InternalMessageInfo

func (m *SerializedBlock) GetHeader() *common.BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SerializedBlock) GetDag() *common.DAG {
	if m != nil {
		return m.Dag
	}
	return nil
}

func (m *SerializedBlock) GetTxIds() []string {
	if m != nil {
		return m.TxIds
	}
	return nil
}

func (m *SerializedBlock) GetAdditionalData() *common.AdditionalData {
	if m != nil {
		return m.AdditionalData
	}
	return nil
}

// Block and its read/write set information
type BlockWithRWSet struct {
	// block data
	Block *common.Block `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	// transaction read/write set of blocks
	TxRWSets []*common.TxRWSet `protobuf:"bytes,2,rep,name=txRWSets,proto3" json:"txRWSets,omitempty"`
	// contract event info
	ContractEvents []*common.ContractEvent `protobuf:"bytes,3,rep,name=contract_events,json=contractEvents,proto3" json:"contract_events,omitempty"`
}

func (m *BlockWithRWSet) Reset()         { *m = BlockWithRWSet{} }
func (m *BlockWithRWSet) String() string { return proto.CompactTextString(m) }
func (*BlockWithRWSet) ProtoMessage()    {}
func (*BlockWithRWSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_8549980b097f750b, []int{1}
}
func (m *BlockWithRWSet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockWithRWSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockWithRWSet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockWithRWSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockWithRWSet.Merge(m, src)
}
func (m *BlockWithRWSet) XXX_Size() int {
	return m.Size()
}
func (m *BlockWithRWSet) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockWithRWSet.DiscardUnknown(m)
}

var xxx_messageInfo_BlockWithRWSet proto.InternalMessageInfo

func (m *BlockWithRWSet) GetBlock() *common.Block {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *BlockWithRWSet) GetTxRWSets() []*common.TxRWSet {
	if m != nil {
		return m.TxRWSets
	}
	return nil
}

func (m *BlockWithRWSet) GetContractEvents() []*common.ContractEvent {
	if m != nil {
		return m.ContractEvents
	}
	return nil
}

func init() {
	proto.RegisterEnum("store.DbType", DbType_name, DbType_value)
	proto.RegisterType((*SerializedBlock)(nil), "store.SerializedBlock")
	proto.RegisterType((*BlockWithRWSet)(nil), "store.BlockWithRWSet")
}

func init() { proto.RegisterFile("store/store.proto", fileDescriptor_8549980b097f750b) }

var fileDescriptor_8549980b097f750b = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0x41, 0x6e, 0x9b, 0x40,
	0x14, 0x86, 0x8d, 0x09, 0x24, 0x79, 0x34, 0xd0, 0xbc, 0x28, 0x15, 0xaa, 0x54, 0x64, 0xa5, 0x1b,
	0xab, 0x51, 0x8c, 0xe4, 0xee, 0x5b, 0xe1, 0xe0, 0xa6, 0xa8, 0x51, 0x22, 0x8d, 0x51, 0x1d, 0x75,
	0x83, 0xc6, 0x30, 0xb2, 0x51, 0x6c, 0xc6, 0x82, 0x69, 0xeb, 0x76, 0xdd, 0x03, 0xf4, 0x0c, 0xbd,
	0x43, 0xef, 0xd0, 0x65, 0x96, 0x5d, 0x56, 0xf6, 0x45, 0x2a, 0x06, 0x70, 0x92, 0x0d, 0xe2, 0xfb,
	0xbf, 0x7f, 0x9e, 0x1e, 0x68, 0xe0, 0xb0, 0x10, 0x3c, 0x67, 0xae, 0x7c, 0xf6, 0x96, 0x39, 0x17,
	0x1c, 0x35, 0x09, 0xcf, 0x31, 0xe6, 0x8b, 0x05, 0xcf, 0xdc, 0xc9, 0x9c, 0xc7, 0xb7, 0x95, 0xda,
	0x66, 0xf9, 0xd7, 0x82, 0x89, 0x3a, 0x3b, 0x6a, 0x32, 0x56, 0x7c, 0x9e, 0xd7, 0xe1, 0xc9, 0x6f,
	0x05, 0xac, 0x11, 0xcb, 0x53, 0x3a, 0x4f, 0xbf, 0xb3, 0x64, 0x50, 0x8e, 0xc0, 0x53, 0xd0, 0x67,
	0x8c, 0x26, 0x2c, 0xb7, 0x95, 0x8e, 0xd2, 0x35, 0xfa, 0x47, 0xbd, 0xea, 0x64, 0x4f, 0xea, 0xf7,
	0x52, 0x91, 0xba, 0x82, 0x2f, 0x40, 0x4d, 0xe8, 0xd4, 0x6e, 0xcb, 0xa6, 0xd1, 0x34, 0x7d, 0xef,
	0x82, 0x94, 0x39, 0x1e, 0x83, 0x2e, 0x56, 0x51, 0x9a, 0x14, 0xb6, 0xda, 0x51, 0xbb, 0xfb, 0x44,
	0x13, 0xab, 0x20, 0x29, 0xf0, 0x2d, 0x58, 0x34, 0x49, 0x52, 0x91, 0xf2, 0x8c, 0xce, 0xa3, 0x84,
	0x0a, 0x6a, 0xef, 0xc8, 0x09, 0xcf, 0x9a, 0x09, 0xde, 0x56, 0xfb, 0x54, 0x50, 0x62, 0xd2, 0x47,
	0x7c, 0xf2, 0x4b, 0x01, 0x53, 0xae, 0x33, 0x4e, 0xc5, 0x8c, 0x8c, 0x47, 0x4c, 0xe0, 0x4b, 0xd0,
	0xe4, 0x2f, 0xa8, 0xb7, 0x3e, 0x78, 0xb4, 0x35, 0xa9, 0x1c, 0x9e, 0xc2, 0x9e, 0x58, 0xc9, 0x7e,
	0x61, 0xb7, 0x3b, 0x6a, 0xd7, 0xe8, 0x5b, 0x4d, 0x2f, 0xac, 0x72, 0xb2, 0x2d, 0xe0, 0x1b, 0xb0,
	0x62, 0x9e, 0x89, 0x9c, 0xc6, 0x22, 0x62, 0x5f, 0x58, 0x26, 0xaa, 0xaf, 0x30, 0xfa, 0xc7, 0xcd,
	0x99, 0xf3, 0x5a, 0x0f, 0x4b, 0x4b, 0xcc, 0xf8, 0x21, 0x16, 0xaf, 0x7e, 0x28, 0xa0, 0xfb, 0x93,
	0xf0, 0xdb, 0x92, 0xa1, 0x09, 0x10, 0x5c, 0x7d, 0xf4, 0x2e, 0x03, 0x3f, 0xf2, 0x07, 0x4f, 0x5b,
	0xf8, 0x04, 0xf6, 0x06, 0x97, 0xd7, 0xe7, 0x1f, 0x4a, 0x52, 0x10, 0xc1, 0xac, 0x28, 0xb8, 0xf2,
	0x87, 0x37, 0x65, 0xd6, 0xc6, 0x7d, 0xd0, 0x42, 0xf9, 0xaa, 0xa2, 0x05, 0x46, 0x78, 0x73, 0xef,
	0x76, 0xd0, 0x80, 0xdd, 0xd1, 0xf5, 0xbb, 0xb0, 0x04, 0xad, 0x1c, 0x35, 0x0a, 0xbd, 0x70, 0x58,
	0x92, 0x8e, 0x87, 0x70, 0x40, 0x86, 0x9e, 0x1f, 0x8d, 0x49, 0x50, 0x45, 0xbb, 0x83, 0x8b, 0x3f,
	0x6b, 0x47, 0xb9, 0x5b, 0x3b, 0xca, 0xbf, 0xb5, 0xa3, 0xfc, 0xdc, 0x38, 0xad, 0xbb, 0x8d, 0xd3,
	0xfa, 0xbb, 0x71, 0x5a, 0x9f, 0xce, 0xe2, 0x19, 0x4d, 0xb3, 0x05, 0xbd, 0x65, 0x79, 0x8f, 0xe7,
	0x53, 0xf7, 0x1e, 0xcf, 0xa6, 0xdc, 0x5d, 0x4e, 0x5c, 0x79, 0x49, 0xa6, 0xbc, 0xba, 0x76, 0x13,
	0x5d, 0xe2, 0xeb, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc0, 0xee, 0xf7, 0x36, 0x8c, 0x02, 0x00,
	0x00,
}

func (m *SerializedBlock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SerializedBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SerializedBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AdditionalData != nil {
		{
			size, err := m.AdditionalData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintStore(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.TxIds) > 0 {
		for iNdEx := len(m.TxIds) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.TxIds[iNdEx])
			copy(dAtA[i:], m.TxIds[iNdEx])
			i = encodeVarintStore(dAtA, i, uint64(len(m.TxIds[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Dag != nil {
		{
			size, err := m.Dag.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintStore(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintStore(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BlockWithRWSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockWithRWSet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockWithRWSet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ContractEvents) > 0 {
		for iNdEx := len(m.ContractEvents) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ContractEvents[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStore(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.TxRWSets) > 0 {
		for iNdEx := len(m.TxRWSets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TxRWSets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStore(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Block != nil {
		{
			size, err := m.Block.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintStore(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStore(dAtA []byte, offset int, v uint64) int {
	offset -= sovStore(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SerializedBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovStore(uint64(l))
	}
	if m.Dag != nil {
		l = m.Dag.Size()
		n += 1 + l + sovStore(uint64(l))
	}
	if len(m.TxIds) > 0 {
		for _, s := range m.TxIds {
			l = len(s)
			n += 1 + l + sovStore(uint64(l))
		}
	}
	if m.AdditionalData != nil {
		l = m.AdditionalData.Size()
		n += 1 + l + sovStore(uint64(l))
	}
	return n
}

func (m *BlockWithRWSet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Block != nil {
		l = m.Block.Size()
		n += 1 + l + sovStore(uint64(l))
	}
	if len(m.TxRWSets) > 0 {
		for _, e := range m.TxRWSets {
			l = e.Size()
			n += 1 + l + sovStore(uint64(l))
		}
	}
	if len(m.ContractEvents) > 0 {
		for _, e := range m.ContractEvents {
			l = e.Size()
			n += 1 + l + sovStore(uint64(l))
		}
	}
	return n
}

func sovStore(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStore(x uint64) (n int) {
	return sovStore(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SerializedBlock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStore
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
			return fmt.Errorf("proto: SerializedBlock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SerializedBlock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &common.BlockHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dag", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Dag == nil {
				m.Dag = &common.DAG{}
			}
			if err := m.Dag.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxIds", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxIds = append(m.TxIds, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdditionalData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AdditionalData == nil {
				m.AdditionalData = &common.AdditionalData{}
			}
			if err := m.AdditionalData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStore(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStore
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
func (m *BlockWithRWSet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStore
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
			return fmt.Errorf("proto: BlockWithRWSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockWithRWSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Block == nil {
				m.Block = &common.Block{}
			}
			if err := m.Block.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxRWSets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxRWSets = append(m.TxRWSets, &common.TxRWSet{})
			if err := m.TxRWSets[len(m.TxRWSets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractEvents", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractEvents = append(m.ContractEvents, &common.ContractEvent{})
			if err := m.ContractEvents[len(m.ContractEvents)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStore(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStore
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
func skipStore(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStore
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
					return 0, ErrIntOverflowStore
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
					return 0, ErrIntOverflowStore
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
				return 0, ErrInvalidLengthStore
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStore
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStore
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStore        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStore          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStore = fmt.Errorf("proto: unexpected end of group")
)
