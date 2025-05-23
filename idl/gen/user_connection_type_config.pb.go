// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: user_connection_type_config.proto

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

type UserConnectionTypeConfig struct {
	state         protoimpl.MessageState      `protogen:"open.v1"`
	Local         *UserConnectionLocalConfig  `protobuf:"bytes,1,opt,name=local,proto3" json:"local,omitempty"`
	Remote        *UserConnectionRemoteConfig `protobuf:"bytes,2,opt,name=remote,proto3" json:"remote,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserConnectionTypeConfig) Reset() {
	*x = UserConnectionTypeConfig{}
	mi := &file_user_connection_type_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserConnectionTypeConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserConnectionTypeConfig) ProtoMessage() {}

func (x *UserConnectionTypeConfig) ProtoReflect() protoreflect.Message {
	mi := &file_user_connection_type_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserConnectionTypeConfig.ProtoReflect.Descriptor instead.
func (*UserConnectionTypeConfig) Descriptor() ([]byte, []int) {
	return file_user_connection_type_config_proto_rawDescGZIP(), []int{0}
}

func (x *UserConnectionTypeConfig) GetLocal() *UserConnectionLocalConfig {
	if x != nil {
		return x.Local
	}
	return nil
}

func (x *UserConnectionTypeConfig) GetRemote() *UserConnectionRemoteConfig {
	if x != nil {
		return x.Remote
	}
	return nil
}

var File_user_connection_type_config_proto protoreflect.FileDescriptor

const file_user_connection_type_config_proto_rawDesc = "" +
	"\n" +
	"!user_connection_type_config.proto\x12\x03nem\x1a\"user_connection_local_config.proto\x1a#user_connection_remote_config.proto\"\x89\x01\n" +
	"\x18UserConnectionTypeConfig\x124\n" +
	"\x05local\x18\x01 \x01(\v2\x1e.nem.UserConnectionLocalConfigR\x05local\x127\n" +
	"\x06remote\x18\x02 \x01(\v2\x1f.nem.UserConnectionRemoteConfigR\x06remoteB?\n" +
	"\x14github.com/nuzur/nemB\x18UserConnectionTypeConfigP\x01Z\vnem/idl/genb\x06proto3"

var (
	file_user_connection_type_config_proto_rawDescOnce sync.Once
	file_user_connection_type_config_proto_rawDescData []byte
)

func file_user_connection_type_config_proto_rawDescGZIP() []byte {
	file_user_connection_type_config_proto_rawDescOnce.Do(func() {
		file_user_connection_type_config_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_user_connection_type_config_proto_rawDesc), len(file_user_connection_type_config_proto_rawDesc)))
	})
	return file_user_connection_type_config_proto_rawDescData
}

var file_user_connection_type_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_user_connection_type_config_proto_goTypes = []any{
	(*UserConnectionTypeConfig)(nil),   // 0: nem.UserConnectionTypeConfig
	(*UserConnectionLocalConfig)(nil),  // 1: nem.UserConnectionLocalConfig
	(*UserConnectionRemoteConfig)(nil), // 2: nem.UserConnectionRemoteConfig
}
var file_user_connection_type_config_proto_depIdxs = []int32{
	1, // 0: nem.UserConnectionTypeConfig.local:type_name -> nem.UserConnectionLocalConfig
	2, // 1: nem.UserConnectionTypeConfig.remote:type_name -> nem.UserConnectionRemoteConfig
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_user_connection_type_config_proto_init() }
func file_user_connection_type_config_proto_init() {
	if File_user_connection_type_config_proto != nil {
		return
	}
	file_user_connection_local_config_proto_init()
	file_user_connection_remote_config_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_user_connection_type_config_proto_rawDesc), len(file_user_connection_type_config_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_connection_type_config_proto_goTypes,
		DependencyIndexes: file_user_connection_type_config_proto_depIdxs,
		MessageInfos:      file_user_connection_type_config_proto_msgTypes,
	}.Build()
	File_user_connection_type_config_proto = out.File
	file_user_connection_type_config_proto_goTypes = nil
	file_user_connection_type_config_proto_depIdxs = nil
}
