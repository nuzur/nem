// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: change_request_data_change.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type ChangeRequestDataChangeDataChangeType int32

const (
	ChangeRequestDataChangeDataChangeType_CHANGE_REQUEST_DATA_CHANGE_DATA_CHANGE_TYPE_INVALID       ChangeRequestDataChangeDataChangeType = 0
	ChangeRequestDataChangeDataChangeType_CHANGE_REQUEST_DATA_CHANGE_DATA_CHANGE_TYPE_UPDATE_FIELD  ChangeRequestDataChangeDataChangeType = 1
	ChangeRequestDataChangeDataChangeType_CHANGE_REQUEST_DATA_CHANGE_DATA_CHANGE_TYPE_CREATE_RECORD ChangeRequestDataChangeDataChangeType = 2
	ChangeRequestDataChangeDataChangeType_CHANGE_REQUEST_DATA_CHANGE_DATA_CHANGE_TYPE_DELETE_RECORD ChangeRequestDataChangeDataChangeType = 3
)

// Enum value maps for ChangeRequestDataChangeDataChangeType.
var (
	ChangeRequestDataChangeDataChangeType_name = map[int32]string{
		0: "CHANGE_REQUEST_DATA_CHANGE_DATA_CHANGE_TYPE_INVALID",
		1: "CHANGE_REQUEST_DATA_CHANGE_DATA_CHANGE_TYPE_UPDATE_FIELD",
		2: "CHANGE_REQUEST_DATA_CHANGE_DATA_CHANGE_TYPE_CREATE_RECORD",
		3: "CHANGE_REQUEST_DATA_CHANGE_DATA_CHANGE_TYPE_DELETE_RECORD",
	}
	ChangeRequestDataChangeDataChangeType_value = map[string]int32{
		"CHANGE_REQUEST_DATA_CHANGE_DATA_CHANGE_TYPE_INVALID":       0,
		"CHANGE_REQUEST_DATA_CHANGE_DATA_CHANGE_TYPE_UPDATE_FIELD":  1,
		"CHANGE_REQUEST_DATA_CHANGE_DATA_CHANGE_TYPE_CREATE_RECORD": 2,
		"CHANGE_REQUEST_DATA_CHANGE_DATA_CHANGE_TYPE_DELETE_RECORD": 3,
	}
)

func (x ChangeRequestDataChangeDataChangeType) Enum() *ChangeRequestDataChangeDataChangeType {
	p := new(ChangeRequestDataChangeDataChangeType)
	*p = x
	return p
}

func (x ChangeRequestDataChangeDataChangeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChangeRequestDataChangeDataChangeType) Descriptor() protoreflect.EnumDescriptor {
	return file_change_request_data_change_proto_enumTypes[0].Descriptor()
}

func (ChangeRequestDataChangeDataChangeType) Type() protoreflect.EnumType {
	return &file_change_request_data_change_proto_enumTypes[0]
}

func (x ChangeRequestDataChangeDataChangeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChangeRequestDataChangeDataChangeType.Descriptor instead.
func (ChangeRequestDataChangeDataChangeType) EnumDescriptor() ([]byte, []int) {
	return file_change_request_data_change_proto_rawDescGZIP(), []int{0}
}

type ChangeRequestDataChange struct {
	state                protoimpl.MessageState                `protogen:"open.v1"`
	Uuid                 string                                `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	DataChangeType       ChangeRequestDataChangeDataChangeType `protobuf:"varint,2,opt,name=data_change_type,json=dataChangeType,proto3,enum=nem.ChangeRequestDataChangeDataChangeType" json:"data_change_type,omitempty"`
	DataChangeTypeConfig *DataChangeTypeConfig                 `protobuf:"bytes,3,opt,name=data_change_type_config,json=dataChangeTypeConfig,proto3" json:"data_change_type_config,omitempty"`
	Order                int64                                 `protobuf:"varint,4,opt,name=order,proto3" json:"order,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *ChangeRequestDataChange) Reset() {
	*x = ChangeRequestDataChange{}
	mi := &file_change_request_data_change_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChangeRequestDataChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeRequestDataChange) ProtoMessage() {}

func (x *ChangeRequestDataChange) ProtoReflect() protoreflect.Message {
	mi := &file_change_request_data_change_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeRequestDataChange.ProtoReflect.Descriptor instead.
func (*ChangeRequestDataChange) Descriptor() ([]byte, []int) {
	return file_change_request_data_change_proto_rawDescGZIP(), []int{0}
}

func (x *ChangeRequestDataChange) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *ChangeRequestDataChange) GetDataChangeType() ChangeRequestDataChangeDataChangeType {
	if x != nil {
		return x.DataChangeType
	}
	return ChangeRequestDataChangeDataChangeType_CHANGE_REQUEST_DATA_CHANGE_DATA_CHANGE_TYPE_INVALID
}

func (x *ChangeRequestDataChange) GetDataChangeTypeConfig() *DataChangeTypeConfig {
	if x != nil {
		return x.DataChangeTypeConfig
	}
	return nil
}

func (x *ChangeRequestDataChange) GetOrder() int64 {
	if x != nil {
		return x.Order
	}
	return 0
}

var File_change_request_data_change_proto protoreflect.FileDescriptor

const file_change_request_data_change_proto_rawDesc = "" +
	"\n" +
	" change_request_data_change.proto\x12\x03nem\x1a\x1ddata_change_type_config.proto\"\xeb\x01\n" +
	"\x17ChangeRequestDataChange\x12\x12\n" +
	"\x04uuid\x18\x01 \x01(\tR\x04uuid\x12T\n" +
	"\x10data_change_type\x18\x02 \x01(\x0e2*.nem.ChangeRequestDataChangeDataChangeTypeR\x0edataChangeType\x12P\n" +
	"\x17data_change_type_config\x18\x03 \x01(\v2\x19.nem.DataChangeTypeConfigR\x14dataChangeTypeConfig\x12\x14\n" +
	"\x05order\x18\x04 \x01(\x03R\x05order*\x9c\x02\n" +
	"%ChangeRequestDataChangeDataChangeType\x127\n" +
	"3CHANGE_REQUEST_DATA_CHANGE_DATA_CHANGE_TYPE_INVALID\x10\x00\x12<\n" +
	"8CHANGE_REQUEST_DATA_CHANGE_DATA_CHANGE_TYPE_UPDATE_FIELD\x10\x01\x12=\n" +
	"9CHANGE_REQUEST_DATA_CHANGE_DATA_CHANGE_TYPE_CREATE_RECORD\x10\x02\x12=\n" +
	"9CHANGE_REQUEST_DATA_CHANGE_DATA_CHANGE_TYPE_DELETE_RECORD\x10\x03B>\n" +
	"\x14github.com/nuzur/nemB\x17ChangeRequestDataChangeP\x01Z\vnem/idl/genb\x06proto3"

var (
	file_change_request_data_change_proto_rawDescOnce sync.Once
	file_change_request_data_change_proto_rawDescData []byte
)

func file_change_request_data_change_proto_rawDescGZIP() []byte {
	file_change_request_data_change_proto_rawDescOnce.Do(func() {
		file_change_request_data_change_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_change_request_data_change_proto_rawDesc), len(file_change_request_data_change_proto_rawDesc)))
	})
	return file_change_request_data_change_proto_rawDescData
}

var file_change_request_data_change_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_change_request_data_change_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_change_request_data_change_proto_goTypes = []any{
	(ChangeRequestDataChangeDataChangeType)(0), // 0: nem.ChangeRequestDataChangeDataChangeType
	(*ChangeRequestDataChange)(nil),            // 1: nem.ChangeRequestDataChange
	(*DataChangeTypeConfig)(nil),               // 2: nem.DataChangeTypeConfig
}
var file_change_request_data_change_proto_depIdxs = []int32{
	0, // 0: nem.ChangeRequestDataChange.data_change_type:type_name -> nem.ChangeRequestDataChangeDataChangeType
	2, // 1: nem.ChangeRequestDataChange.data_change_type_config:type_name -> nem.DataChangeTypeConfig
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_change_request_data_change_proto_init() }
func file_change_request_data_change_proto_init() {
	if File_change_request_data_change_proto != nil {
		return
	}
	file_data_change_type_config_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_change_request_data_change_proto_rawDesc), len(file_change_request_data_change_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_change_request_data_change_proto_goTypes,
		DependencyIndexes: file_change_request_data_change_proto_depIdxs,
		EnumInfos:         file_change_request_data_change_proto_enumTypes,
		MessageInfos:      file_change_request_data_change_proto_msgTypes,
	}.Build()
	File_change_request_data_change_proto = out.File
	file_change_request_data_change_proto_goTypes = nil
	file_change_request_data_change_proto_depIdxs = nil
}
