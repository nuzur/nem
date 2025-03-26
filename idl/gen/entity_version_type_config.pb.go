// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: entity_version_type_config.proto

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

type EntityVersionTypeConfig struct {
	state         protoimpl.MessageState         `protogen:"open.v1"`
	Field         *EntityVersionTypeFieldConfig  `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Entity        *EntityVersionTypeEntityConfig `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EntityVersionTypeConfig) Reset() {
	*x = EntityVersionTypeConfig{}
	mi := &file_entity_version_type_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EntityVersionTypeConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityVersionTypeConfig) ProtoMessage() {}

func (x *EntityVersionTypeConfig) ProtoReflect() protoreflect.Message {
	mi := &file_entity_version_type_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityVersionTypeConfig.ProtoReflect.Descriptor instead.
func (*EntityVersionTypeConfig) Descriptor() ([]byte, []int) {
	return file_entity_version_type_config_proto_rawDescGZIP(), []int{0}
}

func (x *EntityVersionTypeConfig) GetField() *EntityVersionTypeFieldConfig {
	if x != nil {
		return x.Field
	}
	return nil
}

func (x *EntityVersionTypeConfig) GetEntity() *EntityVersionTypeEntityConfig {
	if x != nil {
		return x.Entity
	}
	return nil
}

var File_entity_version_type_config_proto protoreflect.FileDescriptor

const file_entity_version_type_config_proto_rawDesc = "" +
	"\n" +
	" entity_version_type_config.proto\x12\x03nem\x1a'entity_version_type_entity_config.proto\x1a&entity_version_type_field_config.proto\"\x8e\x01\n" +
	"\x17EntityVersionTypeConfig\x127\n" +
	"\x05field\x18\x01 \x01(\v2!.nem.EntityVersionTypeFieldConfigR\x05field\x12:\n" +
	"\x06entity\x18\x02 \x01(\v2\".nem.EntityVersionTypeEntityConfigR\x06entityB>\n" +
	"\x14github.com/nuzur/nemB\x17EntityVersionTypeConfigP\x01Z\vnem/idl/genb\x06proto3"

var (
	file_entity_version_type_config_proto_rawDescOnce sync.Once
	file_entity_version_type_config_proto_rawDescData []byte
)

func file_entity_version_type_config_proto_rawDescGZIP() []byte {
	file_entity_version_type_config_proto_rawDescOnce.Do(func() {
		file_entity_version_type_config_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_entity_version_type_config_proto_rawDesc), len(file_entity_version_type_config_proto_rawDesc)))
	})
	return file_entity_version_type_config_proto_rawDescData
}

var file_entity_version_type_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_entity_version_type_config_proto_goTypes = []any{
	(*EntityVersionTypeConfig)(nil),       // 0: nem.EntityVersionTypeConfig
	(*EntityVersionTypeFieldConfig)(nil),  // 1: nem.EntityVersionTypeFieldConfig
	(*EntityVersionTypeEntityConfig)(nil), // 2: nem.EntityVersionTypeEntityConfig
}
var file_entity_version_type_config_proto_depIdxs = []int32{
	1, // 0: nem.EntityVersionTypeConfig.field:type_name -> nem.EntityVersionTypeFieldConfig
	2, // 1: nem.EntityVersionTypeConfig.entity:type_name -> nem.EntityVersionTypeEntityConfig
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_entity_version_type_config_proto_init() }
func file_entity_version_type_config_proto_init() {
	if File_entity_version_type_config_proto != nil {
		return
	}
	file_entity_version_type_entity_config_proto_init()
	file_entity_version_type_field_config_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_entity_version_type_config_proto_rawDesc), len(file_entity_version_type_config_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entity_version_type_config_proto_goTypes,
		DependencyIndexes: file_entity_version_type_config_proto_depIdxs,
		MessageInfos:      file_entity_version_type_config_proto_msgTypes,
	}.Build()
	File_entity_version_type_config_proto = out.File
	file_entity_version_type_config_proto_goTypes = nil
	file_entity_version_type_config_proto_depIdxs = nil
}
