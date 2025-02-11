// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: field_type_url_config.proto

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

type FieldTypeURLConfig struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	AllowDomains      []string               `protobuf:"bytes,1,rep,name=allow_domains,json=allowDomains,proto3" json:"allow_domains,omitempty"`
	ExcludeDomains    []string               `protobuf:"bytes,2,rep,name=exclude_domains,json=excludeDomains,proto3" json:"exclude_domains,omitempty"`
	HttpsRequired     bool                   `protobuf:"varint,3,opt,name=https_required,json=httpsRequired,proto3" json:"https_required,omitempty"`
	AllowedExtensions []string               `protobuf:"bytes,4,rep,name=allowed_extensions,json=allowedExtensions,proto3" json:"allowed_extensions,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *FieldTypeURLConfig) Reset() {
	*x = FieldTypeURLConfig{}
	mi := &file_field_type_url_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FieldTypeURLConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldTypeURLConfig) ProtoMessage() {}

func (x *FieldTypeURLConfig) ProtoReflect() protoreflect.Message {
	mi := &file_field_type_url_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldTypeURLConfig.ProtoReflect.Descriptor instead.
func (*FieldTypeURLConfig) Descriptor() ([]byte, []int) {
	return file_field_type_url_config_proto_rawDescGZIP(), []int{0}
}

func (x *FieldTypeURLConfig) GetAllowDomains() []string {
	if x != nil {
		return x.AllowDomains
	}
	return nil
}

func (x *FieldTypeURLConfig) GetExcludeDomains() []string {
	if x != nil {
		return x.ExcludeDomains
	}
	return nil
}

func (x *FieldTypeURLConfig) GetHttpsRequired() bool {
	if x != nil {
		return x.HttpsRequired
	}
	return false
}

func (x *FieldTypeURLConfig) GetAllowedExtensions() []string {
	if x != nil {
		return x.AllowedExtensions
	}
	return nil
}

var File_field_type_url_config_proto protoreflect.FileDescriptor

var file_field_type_url_config_proto_rawDesc = string([]byte{
	0x0a, 0x1b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x75, 0x72, 0x6c,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6e,
	0x65, 0x6d, 0x22, 0xb8, 0x01, 0x0a, 0x12, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x55, 0x52, 0x4c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0c, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x27,
	0x0a, 0x0f, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x68, 0x74, 0x74, 0x70, 0x73,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0d, 0x68, 0x74, 0x74, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x2d,
	0x0a, 0x12, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x64, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x39, 0x0a,
	0x14, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x75, 0x7a, 0x75,
	0x72, 0x2f, 0x6e, 0x65, 0x6d, 0x42, 0x12, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x55, 0x52, 0x4c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x01, 0x5a, 0x0b, 0x6e, 0x65, 0x6d,
	0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_field_type_url_config_proto_rawDescOnce sync.Once
	file_field_type_url_config_proto_rawDescData []byte
)

func file_field_type_url_config_proto_rawDescGZIP() []byte {
	file_field_type_url_config_proto_rawDescOnce.Do(func() {
		file_field_type_url_config_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_field_type_url_config_proto_rawDesc), len(file_field_type_url_config_proto_rawDesc)))
	})
	return file_field_type_url_config_proto_rawDescData
}

var file_field_type_url_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_field_type_url_config_proto_goTypes = []any{
	(*FieldTypeURLConfig)(nil), // 0: nem.FieldTypeURLConfig
}
var file_field_type_url_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_field_type_url_config_proto_init() }
func file_field_type_url_config_proto_init() {
	if File_field_type_url_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_field_type_url_config_proto_rawDesc), len(file_field_type_url_config_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_field_type_url_config_proto_goTypes,
		DependencyIndexes: file_field_type_url_config_proto_depIdxs,
		MessageInfos:      file_field_type_url_config_proto_msgTypes,
	}.Build()
	File_field_type_url_config_proto = out.File
	file_field_type_url_config_proto_goTypes = nil
	file_field_type_url_config_proto_depIdxs = nil
}
