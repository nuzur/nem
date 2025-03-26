// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: data_change_field_value.proto

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

type DataChangeFieldValue struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FieldUuid     string                 `protobuf:"bytes,1,opt,name=field_uuid,json=fieldUuid,proto3" json:"field_uuid,omitempty"`
	Value         string                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DataChangeFieldValue) Reset() {
	*x = DataChangeFieldValue{}
	mi := &file_data_change_field_value_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataChangeFieldValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataChangeFieldValue) ProtoMessage() {}

func (x *DataChangeFieldValue) ProtoReflect() protoreflect.Message {
	mi := &file_data_change_field_value_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataChangeFieldValue.ProtoReflect.Descriptor instead.
func (*DataChangeFieldValue) Descriptor() ([]byte, []int) {
	return file_data_change_field_value_proto_rawDescGZIP(), []int{0}
}

func (x *DataChangeFieldValue) GetFieldUuid() string {
	if x != nil {
		return x.FieldUuid
	}
	return ""
}

func (x *DataChangeFieldValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_data_change_field_value_proto protoreflect.FileDescriptor

const file_data_change_field_value_proto_rawDesc = "" +
	"\n" +
	"\x1ddata_change_field_value.proto\x12\x03nem\"K\n" +
	"\x14DataChangeFieldValue\x12\x1d\n" +
	"\n" +
	"field_uuid\x18\x01 \x01(\tR\tfieldUuid\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05valueB;\n" +
	"\x14github.com/nuzur/nemB\x14DataChangeFieldValueP\x01Z\vnem/idl/genb\x06proto3"

var (
	file_data_change_field_value_proto_rawDescOnce sync.Once
	file_data_change_field_value_proto_rawDescData []byte
)

func file_data_change_field_value_proto_rawDescGZIP() []byte {
	file_data_change_field_value_proto_rawDescOnce.Do(func() {
		file_data_change_field_value_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_data_change_field_value_proto_rawDesc), len(file_data_change_field_value_proto_rawDesc)))
	})
	return file_data_change_field_value_proto_rawDescData
}

var file_data_change_field_value_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_data_change_field_value_proto_goTypes = []any{
	(*DataChangeFieldValue)(nil), // 0: nem.DataChangeFieldValue
}
var file_data_change_field_value_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_data_change_field_value_proto_init() }
func file_data_change_field_value_proto_init() {
	if File_data_change_field_value_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_data_change_field_value_proto_rawDesc), len(file_data_change_field_value_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_data_change_field_value_proto_goTypes,
		DependencyIndexes: file_data_change_field_value_proto_depIdxs,
		MessageInfos:      file_data_change_field_value_proto_msgTypes,
	}.Build()
	File_data_change_field_value_proto = out.File
	file_data_change_field_value_proto_goTypes = nil
	file_data_change_field_value_proto_depIdxs = nil
}
