// Code generated by protoc-gen-go.
// source: github.com/bradfitz/grpc-go16-demo/proto/service_proto/bigtable_service_messages.proto
// DO NOT EDIT!

/*
Package google_bigtable_v1 is a generated protocol buffer package.

It is generated from these files:
	github.com/bradfitz/grpc-go16-demo/proto/service_proto/bigtable_service_messages.proto
	github.com/bradfitz/grpc-go16-demo/proto/service_proto/bigtable_service.proto

It has these top-level messages:
	ReadRowsRequest
	ReadRowsResponse
	SampleRowKeysRequest
	SampleRowKeysResponse
	MutateRowRequest
	CheckAndMutateRowRequest
	CheckAndMutateRowResponse
	ReadModifyWriteRowRequest
*/
package google_bigtable_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_bigtable_v11 "github.com/bradfitz/grpc-go16-demo/proto/data_proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Request message for BigtableServer.ReadRows.
type ReadRowsRequest struct {
	// The unique name of the table from which to read.
	TableName string `protobuf:"bytes,1,opt,name=table_name" json:"table_name,omitempty"`
	// If neither row_key nor row_range is set, reads from all rows.
	//
	// Types that are valid to be assigned to Target:
	//	*ReadRowsRequest_RowKey
	//	*ReadRowsRequest_RowRange
	Target isReadRowsRequest_Target `protobuf_oneof:"target"`
	// The filter to apply to the contents of the specified row(s). If unset,
	// reads the entire table.
	Filter *google_bigtable_v11.RowFilter `protobuf:"bytes,5,opt,name=filter" json:"filter,omitempty"`
	// By default, rows are read sequentially, producing results which are
	// guaranteed to arrive in increasing row order. Setting
	// "allow_row_interleaving" to true allows multiple rows to be interleaved in
	// the response stream, which increases throughput but breaks this guarantee,
	// and may force the client to use more memory to buffer partially-received
	// rows. Cannot be set to true when specifying "num_rows_limit".
	AllowRowInterleaving bool `protobuf:"varint,6,opt,name=allow_row_interleaving" json:"allow_row_interleaving,omitempty"`
	// The read will terminate after committing to N rows' worth of results. The
	// default (zero) is to return all results.
	// Note that "allow_row_interleaving" cannot be set to true when this is set.
	NumRowsLimit int64 `protobuf:"varint,7,opt,name=num_rows_limit" json:"num_rows_limit,omitempty"`
}

func (m *ReadRowsRequest) Reset()                    { *m = ReadRowsRequest{} }
func (m *ReadRowsRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRowsRequest) ProtoMessage()               {}
func (*ReadRowsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isReadRowsRequest_Target interface {
	isReadRowsRequest_Target()
}

type ReadRowsRequest_RowKey struct {
	RowKey []byte `protobuf:"bytes,2,opt,name=row_key,proto3,oneof"`
}
type ReadRowsRequest_RowRange struct {
	RowRange *google_bigtable_v11.RowRange `protobuf:"bytes,3,opt,name=row_range,oneof"`
}

func (*ReadRowsRequest_RowKey) isReadRowsRequest_Target()   {}
func (*ReadRowsRequest_RowRange) isReadRowsRequest_Target() {}

func (m *ReadRowsRequest) GetTarget() isReadRowsRequest_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *ReadRowsRequest) GetRowKey() []byte {
	if x, ok := m.GetTarget().(*ReadRowsRequest_RowKey); ok {
		return x.RowKey
	}
	return nil
}

func (m *ReadRowsRequest) GetRowRange() *google_bigtable_v11.RowRange {
	if x, ok := m.GetTarget().(*ReadRowsRequest_RowRange); ok {
		return x.RowRange
	}
	return nil
}

func (m *ReadRowsRequest) GetFilter() *google_bigtable_v11.RowFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ReadRowsRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ReadRowsRequest_OneofMarshaler, _ReadRowsRequest_OneofUnmarshaler, _ReadRowsRequest_OneofSizer, []interface{}{
		(*ReadRowsRequest_RowKey)(nil),
		(*ReadRowsRequest_RowRange)(nil),
	}
}

func _ReadRowsRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ReadRowsRequest)
	// target
	switch x := m.Target.(type) {
	case *ReadRowsRequest_RowKey:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.RowKey)
	case *ReadRowsRequest_RowRange:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RowRange); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ReadRowsRequest.Target has unexpected type %T", x)
	}
	return nil
}

func _ReadRowsRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ReadRowsRequest)
	switch tag {
	case 2: // target.row_key
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Target = &ReadRowsRequest_RowKey{x}
		return true, err
	case 3: // target.row_range
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_bigtable_v11.RowRange)
		err := b.DecodeMessage(msg)
		m.Target = &ReadRowsRequest_RowRange{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ReadRowsRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ReadRowsRequest)
	// target
	switch x := m.Target.(type) {
	case *ReadRowsRequest_RowKey:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.RowKey)))
		n += len(x.RowKey)
	case *ReadRowsRequest_RowRange:
		s := proto.Size(x.RowRange)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Response message for BigtableService.ReadRows.
type ReadRowsResponse struct {
	// The key of the row for which we're receiving data.
	// Results will be received in increasing row key order, unless
	// "allow_row_interleaving" was specified in the request.
	RowKey []byte `protobuf:"bytes,1,opt,name=row_key,proto3" json:"row_key,omitempty"`
	// One or more chunks of the row specified by "row_key".
	Chunks []*ReadRowsResponse_Chunk `protobuf:"bytes,2,rep,name=chunks" json:"chunks,omitempty"`
}

func (m *ReadRowsResponse) Reset()                    { *m = ReadRowsResponse{} }
func (m *ReadRowsResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadRowsResponse) ProtoMessage()               {}
func (*ReadRowsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ReadRowsResponse) GetChunks() []*ReadRowsResponse_Chunk {
	if m != nil {
		return m.Chunks
	}
	return nil
}

// Specifies a piece of a row's contents returned as part of the read
// response stream.
type ReadRowsResponse_Chunk struct {
	// Types that are valid to be assigned to Chunk:
	//	*ReadRowsResponse_Chunk_RowContents
	//	*ReadRowsResponse_Chunk_ResetRow
	//	*ReadRowsResponse_Chunk_CommitRow
	Chunk isReadRowsResponse_Chunk_Chunk `protobuf_oneof:"chunk"`
}

func (m *ReadRowsResponse_Chunk) Reset()                    { *m = ReadRowsResponse_Chunk{} }
func (m *ReadRowsResponse_Chunk) String() string            { return proto.CompactTextString(m) }
func (*ReadRowsResponse_Chunk) ProtoMessage()               {}
func (*ReadRowsResponse_Chunk) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type isReadRowsResponse_Chunk_Chunk interface {
	isReadRowsResponse_Chunk_Chunk()
}

type ReadRowsResponse_Chunk_RowContents struct {
	RowContents *google_bigtable_v11.Family `protobuf:"bytes,1,opt,name=row_contents,oneof"`
}
type ReadRowsResponse_Chunk_ResetRow struct {
	ResetRow bool `protobuf:"varint,2,opt,name=reset_row,oneof"`
}
type ReadRowsResponse_Chunk_CommitRow struct {
	CommitRow bool `protobuf:"varint,3,opt,name=commit_row,oneof"`
}

func (*ReadRowsResponse_Chunk_RowContents) isReadRowsResponse_Chunk_Chunk() {}
func (*ReadRowsResponse_Chunk_ResetRow) isReadRowsResponse_Chunk_Chunk()    {}
func (*ReadRowsResponse_Chunk_CommitRow) isReadRowsResponse_Chunk_Chunk()   {}

func (m *ReadRowsResponse_Chunk) GetChunk() isReadRowsResponse_Chunk_Chunk {
	if m != nil {
		return m.Chunk
	}
	return nil
}

func (m *ReadRowsResponse_Chunk) GetRowContents() *google_bigtable_v11.Family {
	if x, ok := m.GetChunk().(*ReadRowsResponse_Chunk_RowContents); ok {
		return x.RowContents
	}
	return nil
}

func (m *ReadRowsResponse_Chunk) GetResetRow() bool {
	if x, ok := m.GetChunk().(*ReadRowsResponse_Chunk_ResetRow); ok {
		return x.ResetRow
	}
	return false
}

func (m *ReadRowsResponse_Chunk) GetCommitRow() bool {
	if x, ok := m.GetChunk().(*ReadRowsResponse_Chunk_CommitRow); ok {
		return x.CommitRow
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ReadRowsResponse_Chunk) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ReadRowsResponse_Chunk_OneofMarshaler, _ReadRowsResponse_Chunk_OneofUnmarshaler, _ReadRowsResponse_Chunk_OneofSizer, []interface{}{
		(*ReadRowsResponse_Chunk_RowContents)(nil),
		(*ReadRowsResponse_Chunk_ResetRow)(nil),
		(*ReadRowsResponse_Chunk_CommitRow)(nil),
	}
}

func _ReadRowsResponse_Chunk_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ReadRowsResponse_Chunk)
	// chunk
	switch x := m.Chunk.(type) {
	case *ReadRowsResponse_Chunk_RowContents:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RowContents); err != nil {
			return err
		}
	case *ReadRowsResponse_Chunk_ResetRow:
		t := uint64(0)
		if x.ResetRow {
			t = 1
		}
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *ReadRowsResponse_Chunk_CommitRow:
		t := uint64(0)
		if x.CommitRow {
			t = 1
		}
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case nil:
	default:
		return fmt.Errorf("ReadRowsResponse_Chunk.Chunk has unexpected type %T", x)
	}
	return nil
}

func _ReadRowsResponse_Chunk_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ReadRowsResponse_Chunk)
	switch tag {
	case 1: // chunk.row_contents
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_bigtable_v11.Family)
		err := b.DecodeMessage(msg)
		m.Chunk = &ReadRowsResponse_Chunk_RowContents{msg}
		return true, err
	case 2: // chunk.reset_row
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Chunk = &ReadRowsResponse_Chunk_ResetRow{x != 0}
		return true, err
	case 3: // chunk.commit_row
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Chunk = &ReadRowsResponse_Chunk_CommitRow{x != 0}
		return true, err
	default:
		return false, nil
	}
}

func _ReadRowsResponse_Chunk_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ReadRowsResponse_Chunk)
	// chunk
	switch x := m.Chunk.(type) {
	case *ReadRowsResponse_Chunk_RowContents:
		s := proto.Size(x.RowContents)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ReadRowsResponse_Chunk_ResetRow:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += 1
	case *ReadRowsResponse_Chunk_CommitRow:
		n += proto.SizeVarint(3<<3 | proto.WireVarint)
		n += 1
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Request message for BigtableService.SampleRowKeys.
type SampleRowKeysRequest struct {
	// The unique name of the table from which to sample row keys.
	TableName string `protobuf:"bytes,1,opt,name=table_name" json:"table_name,omitempty"`
}

func (m *SampleRowKeysRequest) Reset()                    { *m = SampleRowKeysRequest{} }
func (m *SampleRowKeysRequest) String() string            { return proto.CompactTextString(m) }
func (*SampleRowKeysRequest) ProtoMessage()               {}
func (*SampleRowKeysRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// Response message for BigtableService.SampleRowKeys.
type SampleRowKeysResponse struct {
	// Sorted streamed sequence of sample row keys in the table. The table might
	// have contents before the first row key in the list and after the last one,
	// but a key containing the empty string indicates "end of table" and will be
	// the last response given, if present.
	// Note that row keys in this list may not have ever been written to or read
	// from, and users should therefore not make any assumptions about the row key
	// structure that are specific to their use case.
	RowKey []byte `protobuf:"bytes,1,opt,name=row_key,proto3" json:"row_key,omitempty"`
	// Approximate total storage space used by all rows in the table which precede
	// "row_key". Buffering the contents of all rows between two subsequent
	// samples would require space roughly equal to the difference in their
	// "offset_bytes" fields.
	OffsetBytes int64 `protobuf:"varint,2,opt,name=offset_bytes" json:"offset_bytes,omitempty"`
}

func (m *SampleRowKeysResponse) Reset()                    { *m = SampleRowKeysResponse{} }
func (m *SampleRowKeysResponse) String() string            { return proto.CompactTextString(m) }
func (*SampleRowKeysResponse) ProtoMessage()               {}
func (*SampleRowKeysResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// Request message for BigtableService.MutateRow.
type MutateRowRequest struct {
	// The unique name of the table to which the mutation should be applied.
	TableName string `protobuf:"bytes,1,opt,name=table_name" json:"table_name,omitempty"`
	// The key of the row to which the mutation should be applied.
	RowKey []byte `protobuf:"bytes,2,opt,name=row_key,proto3" json:"row_key,omitempty"`
	// Changes to be atomically applied to the specified row. Entries are applied
	// in order, meaning that earlier mutations can be masked by later ones.
	// Must contain at least one entry and at most 100000.
	Mutations []*google_bigtable_v11.Mutation `protobuf:"bytes,3,rep,name=mutations" json:"mutations,omitempty"`
}

func (m *MutateRowRequest) Reset()                    { *m = MutateRowRequest{} }
func (m *MutateRowRequest) String() string            { return proto.CompactTextString(m) }
func (*MutateRowRequest) ProtoMessage()               {}
func (*MutateRowRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *MutateRowRequest) GetMutations() []*google_bigtable_v11.Mutation {
	if m != nil {
		return m.Mutations
	}
	return nil
}

// Request message for BigtableService.CheckAndMutateRowRequest
type CheckAndMutateRowRequest struct {
	// The unique name of the table to which the conditional mutation should be
	// applied.
	TableName string `protobuf:"bytes,1,opt,name=table_name" json:"table_name,omitempty"`
	// The key of the row to which the conditional mutation should be applied.
	RowKey []byte `protobuf:"bytes,2,opt,name=row_key,proto3" json:"row_key,omitempty"`
	// The filter to be applied to the contents of the specified row. Depending
	// on whether or not any results are yielded, either "true_mutations" or
	// "false_mutations" will be executed. If unset, checks that the row contains
	// any values at all.
	PredicateFilter *google_bigtable_v11.RowFilter `protobuf:"bytes,6,opt,name=predicate_filter" json:"predicate_filter,omitempty"`
	// Changes to be atomically applied to the specified row if "predicate_filter"
	// yields at least one cell when applied to "row_key". Entries are applied in
	// order, meaning that earlier mutations can be masked by later ones.
	// Must contain at least one entry if "false_mutations" is empty, and at most
	// 100000.
	TrueMutations []*google_bigtable_v11.Mutation `protobuf:"bytes,4,rep,name=true_mutations" json:"true_mutations,omitempty"`
	// Changes to be atomically applied to the specified row if "predicate_filter"
	// does not yield any cells when applied to "row_key". Entries are applied in
	// order, meaning that earlier mutations can be masked by later ones.
	// Must contain at least one entry if "true_mutations" is empty, and at most
	// 100000.
	FalseMutations []*google_bigtable_v11.Mutation `protobuf:"bytes,5,rep,name=false_mutations" json:"false_mutations,omitempty"`
}

func (m *CheckAndMutateRowRequest) Reset()                    { *m = CheckAndMutateRowRequest{} }
func (m *CheckAndMutateRowRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckAndMutateRowRequest) ProtoMessage()               {}
func (*CheckAndMutateRowRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CheckAndMutateRowRequest) GetPredicateFilter() *google_bigtable_v11.RowFilter {
	if m != nil {
		return m.PredicateFilter
	}
	return nil
}

func (m *CheckAndMutateRowRequest) GetTrueMutations() []*google_bigtable_v11.Mutation {
	if m != nil {
		return m.TrueMutations
	}
	return nil
}

func (m *CheckAndMutateRowRequest) GetFalseMutations() []*google_bigtable_v11.Mutation {
	if m != nil {
		return m.FalseMutations
	}
	return nil
}

// Response message for BigtableService.CheckAndMutateRowRequest.
type CheckAndMutateRowResponse struct {
	// Whether or not the request's "predicate_filter" yielded any results for
	// the specified row.
	PredicateMatched bool `protobuf:"varint,1,opt,name=predicate_matched" json:"predicate_matched,omitempty"`
}

func (m *CheckAndMutateRowResponse) Reset()                    { *m = CheckAndMutateRowResponse{} }
func (m *CheckAndMutateRowResponse) String() string            { return proto.CompactTextString(m) }
func (*CheckAndMutateRowResponse) ProtoMessage()               {}
func (*CheckAndMutateRowResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

// Request message for BigtableService.ReadModifyWriteRowRequest.
type ReadModifyWriteRowRequest struct {
	// The unique name of the table to which the read/modify/write rules should be
	// applied.
	TableName string `protobuf:"bytes,1,opt,name=table_name" json:"table_name,omitempty"`
	// The key of the row to which the read/modify/write rules should be applied.
	RowKey []byte `protobuf:"bytes,2,opt,name=row_key,proto3" json:"row_key,omitempty"`
	// Rules specifying how the specified row's contents are to be transformed
	// into writes. Entries are applied in order, meaning that earlier rules will
	// affect the results of later ones.
	Rules []*google_bigtable_v11.ReadModifyWriteRule `protobuf:"bytes,3,rep,name=rules" json:"rules,omitempty"`
}

func (m *ReadModifyWriteRowRequest) Reset()                    { *m = ReadModifyWriteRowRequest{} }
func (m *ReadModifyWriteRowRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadModifyWriteRowRequest) ProtoMessage()               {}
func (*ReadModifyWriteRowRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ReadModifyWriteRowRequest) GetRules() []*google_bigtable_v11.ReadModifyWriteRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

func init() {
	proto.RegisterType((*ReadRowsRequest)(nil), "google.bigtable.v1.ReadRowsRequest")
	proto.RegisterType((*ReadRowsResponse)(nil), "google.bigtable.v1.ReadRowsResponse")
	proto.RegisterType((*ReadRowsResponse_Chunk)(nil), "google.bigtable.v1.ReadRowsResponse.Chunk")
	proto.RegisterType((*SampleRowKeysRequest)(nil), "google.bigtable.v1.SampleRowKeysRequest")
	proto.RegisterType((*SampleRowKeysResponse)(nil), "google.bigtable.v1.SampleRowKeysResponse")
	proto.RegisterType((*MutateRowRequest)(nil), "google.bigtable.v1.MutateRowRequest")
	proto.RegisterType((*CheckAndMutateRowRequest)(nil), "google.bigtable.v1.CheckAndMutateRowRequest")
	proto.RegisterType((*CheckAndMutateRowResponse)(nil), "google.bigtable.v1.CheckAndMutateRowResponse")
	proto.RegisterType((*ReadModifyWriteRowRequest)(nil), "google.bigtable.v1.ReadModifyWriteRowRequest")
}

var fileDescriptor0 = []byte{
	// 574 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xd1, 0x6e, 0x12, 0x41,
	0x14, 0x75, 0x5d, 0xd9, 0xd2, 0x5b, 0x52, 0xda, 0xb1, 0x36, 0x5b, 0x52, 0x8d, 0xd9, 0x17, 0x4d,
	0x13, 0x97, 0x14, 0xb5, 0x1a, 0x1f, 0x4c, 0xa4, 0x49, 0xd3, 0xc4, 0x90, 0x34, 0xf4, 0xa1, 0x8f,
	0x64, 0xd8, 0xbd, 0x2c, 0x13, 0x66, 0x77, 0x70, 0x67, 0x96, 0xca, 0x5f, 0xfa, 0x01, 0xbe, 0xfa,
	0x1f, 0xce, 0x0c, 0x4b, 0x69, 0x10, 0x2a, 0xfa, 0x40, 0x02, 0x67, 0xce, 0xbd, 0xf7, 0x9c, 0x3b,
	0x87, 0x81, 0x9b, 0x44, 0x88, 0x84, 0x63, 0x98, 0x08, 0x4e, 0xb3, 0x24, 0x14, 0x79, 0xd2, 0x8c,
	0xb8, 0x28, 0xe2, 0x66, 0x9f, 0x25, 0x8a, 0xf6, 0x39, 0x36, 0x59, 0xa6, 0x30, 0xcf, 0x28, 0x6f,
	0x4a, 0xcc, 0x27, 0x2c, 0xc2, 0xde, 0x38, 0x17, 0x4a, 0xdc, 0x9d, 0xf7, 0xe6, 0x70, 0x8a, 0x52,
	0xd2, 0x04, 0x65, 0x68, 0xcf, 0x09, 0x29, 0x1b, 0xcf, 0x79, 0xe1, 0xe4, 0xb4, 0x71, 0xb9, 0xf9,
	0xb0, 0x98, 0x2a, 0xba, 0x3c, 0xc9, 0x60, 0xb3, 0xee, 0xc1, 0x4f, 0x07, 0xea, 0x5d, 0xa4, 0x71,
	0x57, 0xdc, 0xca, 0x2e, 0x7e, 0x2b, 0x50, 0x2a, 0x42, 0x00, 0x66, 0xbc, 0x8c, 0xa6, 0xe8, 0x3b,
	0x2f, 0x9d, 0xd7, 0xdb, 0x64, 0x1f, 0xb6, 0x72, 0x71, 0xdb, 0x1b, 0xe1, 0xd4, 0x7f, 0xac, 0x81,
	0xda, 0xe5, 0x23, 0x72, 0x0a, 0xdb, 0x06, 0xca, 0xb5, 0x02, 0xf4, 0x5d, 0x0d, 0xee, 0xb4, 0x8e,
	0xc3, 0x3f, 0xc5, 0x86, 0xba, 0x75, 0xd7, 0x70, 0x74, 0xc9, 0x1b, 0xf0, 0x06, 0x8c, 0x6b, 0x65,
	0x7e, 0xc5, 0xf2, 0x9f, 0xaf, 0xe1, 0x5f, 0x58, 0x12, 0x79, 0x01, 0x87, 0x94, 0x73, 0x33, 0x43,
	0x7f, 0xac, 0x23, 0x8e, 0x74, 0xc2, 0xb2, 0xc4, 0xf7, 0x74, 0x79, 0x95, 0x1c, 0xc2, 0x6e, 0x56,
	0xa4, 0xe6, 0x54, 0xf6, 0x38, 0x4b, 0x99, 0xf2, 0xb7, 0x34, 0xee, 0xb6, 0xab, 0xe0, 0x29, 0x9a,
	0x27, 0xa8, 0x82, 0x1f, 0x0e, 0xec, 0x2d, 0xec, 0xc9, 0xb1, 0xc8, 0x24, 0x92, 0xfa, 0xc2, 0x8b,
	0x31, 0x57, 0x23, 0x9f, 0xc0, 0x8b, 0x86, 0x45, 0x36, 0x92, 0xda, 0x9b, 0xab, 0x65, 0x9d, 0xac,
	0x94, 0xb5, 0xd4, 0x26, 0x3c, 0x37, 0x25, 0x0d, 0x01, 0x15, 0xfb, 0x85, 0xb4, 0xa0, 0x66, 0xba,
	0x46, 0x42, 0xeb, 0xcc, 0x94, 0xb4, 0xad, 0x77, 0x5a, 0x8d, 0x55, 0xad, 0x2e, 0x68, 0xca, 0xf8,
	0x54, 0xef, 0xe3, 0xa9, 0x5e, 0x21, 0x4a, 0x54, 0xc6, 0x82, 0xdd, 0x6b, 0x55, 0x83, 0x07, 0x00,
	0x91, 0x48, 0xb5, 0x1b, 0x8b, 0xba, 0x33, 0xb4, 0xbd, 0x05, 0x15, 0xab, 0x31, 0x38, 0x81, 0x83,
	0x6b, 0x9a, 0x8e, 0x39, 0x6a, 0x31, 0x5f, 0x71, 0xfa, 0xd0, 0xad, 0x05, 0x9f, 0xe1, 0xd9, 0x12,
	0x77, 0xdd, 0x0a, 0x0e, 0xa0, 0x26, 0x06, 0x03, 0x23, 0xa5, 0x3f, 0x55, 0x28, 0xad, 0x18, 0x37,
	0x18, 0xc2, 0x5e, 0xa7, 0x50, 0x54, 0x99, 0xfa, 0x87, 0xd2, 0x51, 0x5f, 0x4a, 0x07, 0x69, 0xc2,
	0x76, 0x6a, 0x0a, 0x99, 0x9e, 0xa6, 0x2d, 0xb8, 0xeb, 0xb2, 0xd1, 0x29, 0x49, 0xc1, 0x2f, 0x07,
	0xfc, 0xf3, 0x21, 0x46, 0xa3, 0x2f, 0x59, 0xfc, 0x7f, 0x23, 0x3f, 0xc0, 0xde, 0x38, 0xc7, 0x98,
	0x45, 0xba, 0xb6, 0x57, 0xa6, 0xcc, 0xdb, 0x24, 0x65, 0xef, 0x60, 0x57, 0xe5, 0x85, 0xfe, 0xdf,
	0xdd, 0x09, 0x7e, 0xf2, 0x77, 0xc1, 0xe4, 0x3d, 0xd4, 0x07, 0x94, 0xcb, 0xfb, 0x65, 0x95, 0x0d,
	0x7c, 0x9e, 0xc1, 0xd1, 0x0a, 0x9b, 0xe5, 0xad, 0x1c, 0xc1, 0xfe, 0xc2, 0x42, 0x4a, 0x55, 0x34,
	0xc4, 0xd8, 0xda, 0xad, 0x06, 0xdf, 0xe1, 0xc8, 0x04, 0xb0, 0x23, 0x62, 0x36, 0x98, 0xde, 0xe4,
	0xec, 0xdf, 0xf7, 0x73, 0x06, 0x95, 0xbc, 0xe0, 0x38, 0xbf, 0x8e, 0x57, 0xeb, 0x32, 0x7e, 0x7f,
	0x84, 0xe6, 0xb7, 0x3f, 0xc2, 0xa1, 0x8e, 0xe3, 0x0a, 0x76, 0xfb, 0xb8, 0x5d, 0xfe, 0xb8, 0x9e,
	0xbd, 0x5c, 0x9d, 0xf2, 0xe1, 0xba, 0x32, 0x2f, 0xcb, 0x95, 0xd3, 0xf7, 0xec, 0x13, 0xf3, 0xf6,
	0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xef, 0x18, 0xfb, 0x81, 0x1b, 0x05, 0x00, 0x00,
}
