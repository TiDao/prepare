// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: common/transaction.proto

package common

import (
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

// a transaction includes request and its result
type Transaction struct {
	// header of the transaction
	Header *TxHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// payload of the request
	RequestPayload []byte `protobuf:"bytes,2,opt,name=request_payload,json=requestPayload,proto3" json:"request_payload,omitempty"`
	// signature of request bytes(including header and payload)
	RequestSignature []byte `protobuf:"bytes,3,opt,name=request_signature,json=requestSignature,proto3" json:"request_signature,omitempty"`
	// result of the transaction, can be marshalled according to tx_type in header
	Result *Result `protobuf:"bytes,4,opt,name=result,proto3" json:"result,omitempty"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6296da495f91d72, []int{0}
}
func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return m.Size()
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetHeader() *TxHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Transaction) GetRequestPayload() []byte {
	if m != nil {
		return m.RequestPayload
	}
	return nil
}

func (m *Transaction) GetRequestSignature() []byte {
	if m != nil {
		return m.RequestSignature
	}
	return nil
}

func (m *Transaction) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

// transactioninfo inclde transaction and its block height hash and tx index
type TransactionInfo struct {
	// transaction raw data
	Transaction *Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	// block height
	BlockHeight uint64 `protobuf:"varint,2,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	// block hash
	BlockHash []byte `protobuf:"bytes,3,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	// transaction index in block
	TxIndex uint32 `protobuf:"varint,4,opt,name=tx_index,json=txIndex,proto3" json:"tx_index,omitempty"`
}

func (m *TransactionInfo) Reset()         { *m = TransactionInfo{} }
func (m *TransactionInfo) String() string { return proto.CompactTextString(m) }
func (*TransactionInfo) ProtoMessage()    {}
func (*TransactionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6296da495f91d72, []int{1}
}
func (m *TransactionInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TransactionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransactionInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TransactionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionInfo.Merge(m, src)
}
func (m *TransactionInfo) XXX_Size() int {
	return m.Size()
}
func (m *TransactionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionInfo proto.InternalMessageInfo

func (m *TransactionInfo) GetTransaction() *Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func (m *TransactionInfo) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *TransactionInfo) GetBlockHash() []byte {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

func (m *TransactionInfo) GetTxIndex() uint32 {
	if m != nil {
		return m.TxIndex
	}
	return 0
}

func init() {
	proto.RegisterType((*Transaction)(nil), "common.Transaction")
	proto.RegisterType((*TransactionInfo)(nil), "common.TransactionInfo")
}

func init() { proto.RegisterFile("common/transaction.proto", fileDescriptor_f6296da495f91d72) }

var fileDescriptor_f6296da495f91d72 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xc1, 0x4e, 0xf2, 0x40,
	0x14, 0x85, 0x99, 0xff, 0x27, 0x55, 0x6f, 0x11, 0x70, 0xd0, 0xa4, 0x92, 0xd8, 0x20, 0x0b, 0x25,
	0x31, 0xb6, 0x89, 0xc6, 0x17, 0xc0, 0x0d, 0xec, 0x4c, 0x65, 0xe5, 0x86, 0x4c, 0xcb, 0xd8, 0x36,
	0xc0, 0x4c, 0x9d, 0x0e, 0x49, 0x79, 0x0b, 0x5f, 0xc2, 0x47, 0xf0, 0x1d, 0x5c, 0xb2, 0x74, 0x69,
	0xe0, 0x45, 0x8c, 0x33, 0x03, 0x74, 0x79, 0xbf, 0x73, 0x7a, 0xef, 0x39, 0x1d, 0x70, 0x22, 0x3e,
	0x9f, 0x73, 0xe6, 0x4b, 0x41, 0x58, 0x4e, 0x22, 0x99, 0x72, 0xe6, 0x65, 0x82, 0x4b, 0x8e, 0x2d,
	0xad, 0xb4, 0x4f, 0x8d, 0x43, 0xd0, 0xb7, 0x05, 0xcd, 0xa5, 0x56, 0xdb, 0xad, 0x1d, 0xcd, 0x17,
	0x33, 0x03, 0xbb, 0x9f, 0x08, 0xec, 0xd1, 0x7e, 0x11, 0xee, 0x81, 0x95, 0x50, 0x32, 0xa1, 0xc2,
	0x41, 0x1d, 0xd4, 0xb3, 0xef, 0x9a, 0x9e, 0xfe, 0xca, 0x1b, 0x15, 0x03, 0xc5, 0x03, 0xa3, 0xe3,
	0x6b, 0x68, 0x98, 0xfd, 0xe3, 0x8c, 0x2c, 0x67, 0x9c, 0x4c, 0x9c, 0x7f, 0x1d, 0xd4, 0xab, 0x05,
	0x75, 0x83, 0x9f, 0x34, 0xc5, 0x37, 0x70, 0xb2, 0x35, 0xe6, 0x69, 0xcc, 0x88, 0x5c, 0x08, 0xea,
	0xfc, 0x57, 0xd6, 0xa6, 0x11, 0x9e, 0xb7, 0x1c, 0x5f, 0x81, 0xa5, 0xf3, 0x39, 0x55, 0x75, 0xbf,
	0xbe, 0xbd, 0x1f, 0x28, 0x1a, 0x18, 0xb5, 0xfb, 0x81, 0xa0, 0x51, 0xca, 0x3d, 0x64, 0xaf, 0x1c,
	0x3f, 0x80, 0x5d, 0xfa, 0x27, 0xa6, 0x40, 0x6b, 0x57, 0x60, 0x2f, 0x05, 0x65, 0x1f, 0xbe, 0x84,
	0x5a, 0x38, 0xe3, 0xd1, 0x74, 0x9c, 0xd0, 0x34, 0x4e, 0xa4, 0x6a, 0x51, 0x0d, 0x6c, 0xc5, 0x06,
	0x0a, 0xe1, 0x0b, 0x00, 0x63, 0x21, 0x79, 0x62, 0xb2, 0x1f, 0x69, 0x03, 0xc9, 0x13, 0x7c, 0x0e,
	0x87, 0xb2, 0x18, 0xa7, 0x6c, 0x42, 0x0b, 0x15, 0xfb, 0x38, 0x38, 0x90, 0xc5, 0xf0, 0x6f, 0xec,
	0x2f, 0xbf, 0xd6, 0x2e, 0x5a, 0xad, 0x5d, 0xf4, 0xb3, 0x76, 0xd1, 0xfb, 0xc6, 0xad, 0xac, 0x36,
	0x6e, 0xe5, 0x7b, 0xe3, 0x56, 0xc0, 0xe1, 0x22, 0xf6, 0xa2, 0x84, 0xa4, 0x6c, 0x4e, 0xa6, 0x54,
	0x78, 0x59, 0x68, 0x92, 0xf6, 0xcf, 0x1e, 0x77, 0xb4, 0x14, 0xfa, 0xa5, 0x6c, 0xe6, 0x22, 0xf6,
	0xf7, 0xe3, 0x6d, 0xcc, 0xfd, 0x2c, 0xf4, 0xd5, 0x8b, 0xc6, 0xdc, 0xd7, 0x6b, 0x42, 0x4b, 0xcd,
	0xf7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x09, 0x45, 0xcf, 0x30, 0x02, 0x00, 0x00,
}

func (m *Transaction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Transaction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Transaction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Result != nil {
		{
			size, err := m.Result.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTransaction(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.RequestSignature) > 0 {
		i -= len(m.RequestSignature)
		copy(dAtA[i:], m.RequestSignature)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.RequestSignature)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.RequestPayload) > 0 {
		i -= len(m.RequestPayload)
		copy(dAtA[i:], m.RequestPayload)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.RequestPayload)))
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
			i = encodeVarintTransaction(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TransactionInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TransactionInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TransactionInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TxIndex != 0 {
		i = encodeVarintTransaction(dAtA, i, uint64(m.TxIndex))
		i--
		dAtA[i] = 0x20
	}
	if len(m.BlockHash) > 0 {
		i -= len(m.BlockHash)
		copy(dAtA[i:], m.BlockHash)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.BlockHash)))
		i--
		dAtA[i] = 0x1a
	}
	if m.BlockHeight != 0 {
		i = encodeVarintTransaction(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x10
	}
	if m.Transaction != nil {
		{
			size, err := m.Transaction.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTransaction(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTransaction(dAtA []byte, offset int, v uint64) int {
	offset -= sovTransaction(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Transaction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovTransaction(uint64(l))
	}
	l = len(m.RequestPayload)
	if l > 0 {
		n += 1 + l + sovTransaction(uint64(l))
	}
	l = len(m.RequestSignature)
	if l > 0 {
		n += 1 + l + sovTransaction(uint64(l))
	}
	if m.Result != nil {
		l = m.Result.Size()
		n += 1 + l + sovTransaction(uint64(l))
	}
	return n
}

func (m *TransactionInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Transaction != nil {
		l = m.Transaction.Size()
		n += 1 + l + sovTransaction(uint64(l))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovTransaction(uint64(m.BlockHeight))
	}
	l = len(m.BlockHash)
	if l > 0 {
		n += 1 + l + sovTransaction(uint64(l))
	}
	if m.TxIndex != 0 {
		n += 1 + sovTransaction(uint64(m.TxIndex))
	}
	return n
}

func sovTransaction(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTransaction(x uint64) (n int) {
	return sovTransaction(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Transaction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransaction
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
			return fmt.Errorf("proto: Transaction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Transaction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &TxHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestPayload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestPayload = append(m.RequestPayload[:0], dAtA[iNdEx:postIndex]...)
			if m.RequestPayload == nil {
				m.RequestPayload = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestSignature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestSignature = append(m.RequestSignature[:0], dAtA[iNdEx:postIndex]...)
			if m.RequestSignature == nil {
				m.RequestSignature = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Result == nil {
				m.Result = &Result{}
			}
			if err := m.Result.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTransaction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTransaction
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
func (m *TransactionInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransaction
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
			return fmt.Errorf("proto: TransactionInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TransactionInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transaction", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Transaction == nil {
				m.Transaction = &Transaction{}
			}
			if err := m.Transaction.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockHash = append(m.BlockHash[:0], dAtA[iNdEx:postIndex]...)
			if m.BlockHash == nil {
				m.BlockHash = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxIndex", wireType)
			}
			m.TxIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TxIndex |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTransaction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTransaction
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
func skipTransaction(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTransaction
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
					return 0, ErrIntOverflowTransaction
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
					return 0, ErrIntOverflowTransaction
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
				return 0, ErrInvalidLengthTransaction
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTransaction
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTransaction
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTransaction        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTransaction          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTransaction = fmt.Errorf("proto: unexpected end of group")
)