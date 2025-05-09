// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: db_type_config.proto

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

type DbTypeConfig struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Postgres      *DbTypePostgresConfig  `protobuf:"bytes,1,opt,name=postgres,proto3" json:"postgres,omitempty"`
	Mysql         *DbTypeMysqlConfig     `protobuf:"bytes,2,opt,name=mysql,proto3" json:"mysql,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DbTypeConfig) Reset() {
	*x = DbTypeConfig{}
	mi := &file_db_type_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DbTypeConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DbTypeConfig) ProtoMessage() {}

func (x *DbTypeConfig) ProtoReflect() protoreflect.Message {
	mi := &file_db_type_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DbTypeConfig.ProtoReflect.Descriptor instead.
func (*DbTypeConfig) Descriptor() ([]byte, []int) {
	return file_db_type_config_proto_rawDescGZIP(), []int{0}
}

func (x *DbTypeConfig) GetPostgres() *DbTypePostgresConfig {
	if x != nil {
		return x.Postgres
	}
	return nil
}

func (x *DbTypeConfig) GetMysql() *DbTypeMysqlConfig {
	if x != nil {
		return x.Mysql
	}
	return nil
}

var File_db_type_config_proto protoreflect.FileDescriptor

const file_db_type_config_proto_rawDesc = "" +
	"\n" +
	"\x14db_type_config.proto\x12\x03nem\x1a\x1adb_type_mysql_config.proto\x1a\x1ddb_type_postgres_config.proto\"s\n" +
	"\fDbTypeConfig\x125\n" +
	"\bpostgres\x18\x01 \x01(\v2\x19.nem.DbTypePostgresConfigR\bpostgres\x12,\n" +
	"\x05mysql\x18\x02 \x01(\v2\x16.nem.DbTypeMysqlConfigR\x05mysqlB3\n" +
	"\x14github.com/nuzur/nemB\fDbTypeConfigP\x01Z\vnem/idl/genb\x06proto3"

var (
	file_db_type_config_proto_rawDescOnce sync.Once
	file_db_type_config_proto_rawDescData []byte
)

func file_db_type_config_proto_rawDescGZIP() []byte {
	file_db_type_config_proto_rawDescOnce.Do(func() {
		file_db_type_config_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_db_type_config_proto_rawDesc), len(file_db_type_config_proto_rawDesc)))
	})
	return file_db_type_config_proto_rawDescData
}

var file_db_type_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_db_type_config_proto_goTypes = []any{
	(*DbTypeConfig)(nil),         // 0: nem.DbTypeConfig
	(*DbTypePostgresConfig)(nil), // 1: nem.DbTypePostgresConfig
	(*DbTypeMysqlConfig)(nil),    // 2: nem.DbTypeMysqlConfig
}
var file_db_type_config_proto_depIdxs = []int32{
	1, // 0: nem.DbTypeConfig.postgres:type_name -> nem.DbTypePostgresConfig
	2, // 1: nem.DbTypeConfig.mysql:type_name -> nem.DbTypeMysqlConfig
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_db_type_config_proto_init() }
func file_db_type_config_proto_init() {
	if File_db_type_config_proto != nil {
		return
	}
	file_db_type_mysql_config_proto_init()
	file_db_type_postgres_config_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_db_type_config_proto_rawDesc), len(file_db_type_config_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_db_type_config_proto_goTypes,
		DependencyIndexes: file_db_type_config_proto_depIdxs,
		MessageInfos:      file_db_type_config_proto_msgTypes,
	}.Build()
	File_db_type_config_proto = out.File
	file_db_type_config_proto_goTypes = nil
	file_db_type_config_proto_depIdxs = nil
}
