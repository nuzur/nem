// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: index.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IndexStatus int32

const (
	IndexStatus_INDEX_STATUS_INVALID  IndexStatus = 0
	IndexStatus_INDEX_STATUS_ACTIVE   IndexStatus = 1
	IndexStatus_INDEX_STATUS_DISABLED IndexStatus = 2
)

// Enum value maps for IndexStatus.
var (
	IndexStatus_name = map[int32]string{
		0: "INDEX_STATUS_INVALID",
		1: "INDEX_STATUS_ACTIVE",
		2: "INDEX_STATUS_DISABLED",
	}
	IndexStatus_value = map[string]int32{
		"INDEX_STATUS_INVALID":  0,
		"INDEX_STATUS_ACTIVE":   1,
		"INDEX_STATUS_DISABLED": 2,
	}
)

func (x IndexStatus) Enum() *IndexStatus {
	p := new(IndexStatus)
	*p = x
	return p
}

func (x IndexStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IndexStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_index_proto_enumTypes[0].Descriptor()
}

func (IndexStatus) Type() protoreflect.EnumType {
	return &file_index_proto_enumTypes[0]
}

func (x IndexStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IndexStatus.Descriptor instead.
func (IndexStatus) EnumDescriptor() ([]byte, []int) {
	return file_index_proto_rawDescGZIP(), []int{0}
}

type IndexType int32

const (
	IndexType_INDEX_TYPE_INVALID  IndexType = 0
	IndexType_INDEX_TYPE_INDEX    IndexType = 1
	IndexType_INDEX_TYPE_PRIMARY  IndexType = 2
	IndexType_INDEX_TYPE_UNIQUE   IndexType = 3
	IndexType_INDEX_TYPE_FULLTEXT IndexType = 4
)

// Enum value maps for IndexType.
var (
	IndexType_name = map[int32]string{
		0: "INDEX_TYPE_INVALID",
		1: "INDEX_TYPE_INDEX",
		2: "INDEX_TYPE_PRIMARY",
		3: "INDEX_TYPE_UNIQUE",
		4: "INDEX_TYPE_FULLTEXT",
	}
	IndexType_value = map[string]int32{
		"INDEX_TYPE_INVALID":  0,
		"INDEX_TYPE_INDEX":    1,
		"INDEX_TYPE_PRIMARY":  2,
		"INDEX_TYPE_UNIQUE":   3,
		"INDEX_TYPE_FULLTEXT": 4,
	}
)

func (x IndexType) Enum() *IndexType {
	p := new(IndexType)
	*p = x
	return p
}

func (x IndexType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IndexType) Descriptor() protoreflect.EnumDescriptor {
	return file_index_proto_enumTypes[1].Descriptor()
}

func (IndexType) Type() protoreflect.EnumType {
	return &file_index_proto_enumTypes[1]
}

func (x IndexType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IndexType.Descriptor instead.
func (IndexType) EnumDescriptor() ([]byte, []int) {
	return file_index_proto_rawDescGZIP(), []int{1}
}

type Index struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uuid          string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Identifier    string                 `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Type          IndexType              `protobuf:"varint,3,opt,name=type,proto3,enum=nem.IndexType" json:"type,omitempty"`
	Fields        []*IndexField          `protobuf:"bytes,4,rep,name=fields,proto3" json:"fields,omitempty"`
	Status        IndexStatus            `protobuf:"varint,5,opt,name=status,proto3,enum=nem.IndexStatus" json:"status,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CreatedByUuid string                 `protobuf:"bytes,8,opt,name=created_by_uuid,json=createdByUuid,proto3" json:"created_by_uuid,omitempty"`
	UpdatedByUuid string                 `protobuf:"bytes,9,opt,name=updated_by_uuid,json=updatedByUuid,proto3" json:"updated_by_uuid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Index) Reset() {
	*x = Index{}
	mi := &file_index_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Index) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Index) ProtoMessage() {}

func (x *Index) ProtoReflect() protoreflect.Message {
	mi := &file_index_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Index.ProtoReflect.Descriptor instead.
func (*Index) Descriptor() ([]byte, []int) {
	return file_index_proto_rawDescGZIP(), []int{0}
}

func (x *Index) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Index) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *Index) GetType() IndexType {
	if x != nil {
		return x.Type
	}
	return IndexType_INDEX_TYPE_INVALID
}

func (x *Index) GetFields() []*IndexField {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *Index) GetStatus() IndexStatus {
	if x != nil {
		return x.Status
	}
	return IndexStatus_INDEX_STATUS_INVALID
}

func (x *Index) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Index) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Index) GetCreatedByUuid() string {
	if x != nil {
		return x.CreatedByUuid
	}
	return ""
}

func (x *Index) GetUpdatedByUuid() string {
	if x != nil {
		return x.UpdatedByUuid
	}
	return ""
}

var File_index_proto protoreflect.FileDescriptor

const file_index_proto_rawDesc = "" +
	"\n" +
	"\vindex.proto\x12\x03nem\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x11index_field.proto\"\xf8\x02\n" +
	"\x05Index\x12\x12\n" +
	"\x04uuid\x18\x01 \x01(\tR\x04uuid\x12\x1e\n" +
	"\n" +
	"identifier\x18\x02 \x01(\tR\n" +
	"identifier\x12\"\n" +
	"\x04type\x18\x03 \x01(\x0e2\x0e.nem.IndexTypeR\x04type\x12'\n" +
	"\x06fields\x18\x04 \x03(\v2\x0f.nem.IndexFieldR\x06fields\x12(\n" +
	"\x06status\x18\x05 \x01(\x0e2\x10.nem.IndexStatusR\x06status\x129\n" +
	"\n" +
	"created_at\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\a \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\x12&\n" +
	"\x0fcreated_by_uuid\x18\b \x01(\tR\rcreatedByUuid\x12&\n" +
	"\x0fupdated_by_uuid\x18\t \x01(\tR\rupdatedByUuid*[\n" +
	"\vIndexStatus\x12\x18\n" +
	"\x14INDEX_STATUS_INVALID\x10\x00\x12\x17\n" +
	"\x13INDEX_STATUS_ACTIVE\x10\x01\x12\x19\n" +
	"\x15INDEX_STATUS_DISABLED\x10\x02*\x81\x01\n" +
	"\tIndexType\x12\x16\n" +
	"\x12INDEX_TYPE_INVALID\x10\x00\x12\x14\n" +
	"\x10INDEX_TYPE_INDEX\x10\x01\x12\x16\n" +
	"\x12INDEX_TYPE_PRIMARY\x10\x02\x12\x15\n" +
	"\x11INDEX_TYPE_UNIQUE\x10\x03\x12\x17\n" +
	"\x13INDEX_TYPE_FULLTEXT\x10\x04B,\n" +
	"\x14github.com/nuzur/nemB\x05IndexP\x01Z\vnem/idl/genb\x06proto3"

var (
	file_index_proto_rawDescOnce sync.Once
	file_index_proto_rawDescData []byte
)

func file_index_proto_rawDescGZIP() []byte {
	file_index_proto_rawDescOnce.Do(func() {
		file_index_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_index_proto_rawDesc), len(file_index_proto_rawDesc)))
	})
	return file_index_proto_rawDescData
}

var file_index_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_index_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_index_proto_goTypes = []any{
	(IndexStatus)(0),              // 0: nem.IndexStatus
	(IndexType)(0),                // 1: nem.IndexType
	(*Index)(nil),                 // 2: nem.Index
	(*IndexField)(nil),            // 3: nem.IndexField
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_index_proto_depIdxs = []int32{
	1, // 0: nem.Index.type:type_name -> nem.IndexType
	3, // 1: nem.Index.fields:type_name -> nem.IndexField
	0, // 2: nem.Index.status:type_name -> nem.IndexStatus
	4, // 3: nem.Index.created_at:type_name -> google.protobuf.Timestamp
	4, // 4: nem.Index.updated_at:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_index_proto_init() }
func file_index_proto_init() {
	if File_index_proto != nil {
		return
	}
	file_index_field_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_index_proto_rawDesc), len(file_index_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_index_proto_goTypes,
		DependencyIndexes: file_index_proto_depIdxs,
		EnumInfos:         file_index_proto_enumTypes,
		MessageInfos:      file_index_proto_msgTypes,
	}.Build()
	File_index_proto = out.File
	file_index_proto_goTypes = nil
	file_index_proto_depIdxs = nil
}
