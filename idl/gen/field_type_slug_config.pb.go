// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: field_type_slug_config.proto

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

type FieldTypeSlugConfig struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	MinSize           int64                  `protobuf:"varint,1,opt,name=min_size,json=minSize,proto3" json:"min_size,omitempty"`
	MaxSize           int64                  `protobuf:"varint,2,opt,name=max_size,json=maxSize,proto3" json:"max_size,omitempty"`
	RegexValidation   string                 `protobuf:"bytes,3,opt,name=regex_validation,json=regexValidation,proto3" json:"regex_validation,omitempty"`
	BasedOnFieldUuids []string               `protobuf:"bytes,4,rep,name=based_on_field_uuids,json=basedOnFieldUuids,proto3" json:"based_on_field_uuids,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *FieldTypeSlugConfig) Reset() {
	*x = FieldTypeSlugConfig{}
	mi := &file_field_type_slug_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FieldTypeSlugConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldTypeSlugConfig) ProtoMessage() {}

func (x *FieldTypeSlugConfig) ProtoReflect() protoreflect.Message {
	mi := &file_field_type_slug_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldTypeSlugConfig.ProtoReflect.Descriptor instead.
func (*FieldTypeSlugConfig) Descriptor() ([]byte, []int) {
	return file_field_type_slug_config_proto_rawDescGZIP(), []int{0}
}

func (x *FieldTypeSlugConfig) GetMinSize() int64 {
	if x != nil {
		return x.MinSize
	}
	return 0
}

func (x *FieldTypeSlugConfig) GetMaxSize() int64 {
	if x != nil {
		return x.MaxSize
	}
	return 0
}

func (x *FieldTypeSlugConfig) GetRegexValidation() string {
	if x != nil {
		return x.RegexValidation
	}
	return ""
}

func (x *FieldTypeSlugConfig) GetBasedOnFieldUuids() []string {
	if x != nil {
		return x.BasedOnFieldUuids
	}
	return nil
}

var File_field_type_slug_config_proto protoreflect.FileDescriptor

var file_field_type_slug_config_proto_rawDesc = string([]byte{
	0x0a, 0x1c, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x73, 0x6c, 0x75,
	0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03,
	0x6e, 0x65, 0x6d, 0x22, 0xa7, 0x01, 0x0a, 0x13, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x53, 0x6c, 0x75, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x6d,
	0x69, 0x6e, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d,
	0x69, 0x6e, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x67, 0x65, 0x78, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x67,
	0x65, 0x78, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x14,
	0x62, 0x61, 0x73, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x75,
	0x75, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x62, 0x61, 0x73, 0x65,
	0x64, 0x4f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x55, 0x75, 0x69, 0x64, 0x73, 0x42, 0x3a, 0x0a,
	0x14, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x75, 0x7a, 0x75,
	0x72, 0x2f, 0x6e, 0x65, 0x6d, 0x42, 0x13, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x53, 0x6c, 0x75, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x01, 0x5a, 0x0b, 0x6e, 0x65,
	0x6d, 0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_field_type_slug_config_proto_rawDescOnce sync.Once
	file_field_type_slug_config_proto_rawDescData []byte
)

func file_field_type_slug_config_proto_rawDescGZIP() []byte {
	file_field_type_slug_config_proto_rawDescOnce.Do(func() {
		file_field_type_slug_config_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_field_type_slug_config_proto_rawDesc), len(file_field_type_slug_config_proto_rawDesc)))
	})
	return file_field_type_slug_config_proto_rawDescData
}

var file_field_type_slug_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_field_type_slug_config_proto_goTypes = []any{
	(*FieldTypeSlugConfig)(nil), // 0: nem.FieldTypeSlugConfig
}
var file_field_type_slug_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_field_type_slug_config_proto_init() }
func file_field_type_slug_config_proto_init() {
	if File_field_type_slug_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_field_type_slug_config_proto_rawDesc), len(file_field_type_slug_config_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_field_type_slug_config_proto_goTypes,
		DependencyIndexes: file_field_type_slug_config_proto_depIdxs,
		MessageInfos:      file_field_type_slug_config_proto_msgTypes,
	}.Build()
	File_field_type_slug_config_proto = out.File
	file_field_type_slug_config_proto_goTypes = nil
	file_field_type_slug_config_proto_depIdxs = nil
}
