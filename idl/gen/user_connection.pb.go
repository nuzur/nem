// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: user_connection.proto

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

type UserConnectionStatus int32

const (
	UserConnectionStatus_USER_CONNECTION_STATUS_INVALID  UserConnectionStatus = 0
	UserConnectionStatus_USER_CONNECTION_STATUS_ACTIVE   UserConnectionStatus = 1
	UserConnectionStatus_USER_CONNECTION_STATUS_DISABLED UserConnectionStatus = 2
)

// Enum value maps for UserConnectionStatus.
var (
	UserConnectionStatus_name = map[int32]string{
		0: "USER_CONNECTION_STATUS_INVALID",
		1: "USER_CONNECTION_STATUS_ACTIVE",
		2: "USER_CONNECTION_STATUS_DISABLED",
	}
	UserConnectionStatus_value = map[string]int32{
		"USER_CONNECTION_STATUS_INVALID":  0,
		"USER_CONNECTION_STATUS_ACTIVE":   1,
		"USER_CONNECTION_STATUS_DISABLED": 2,
	}
)

func (x UserConnectionStatus) Enum() *UserConnectionStatus {
	p := new(UserConnectionStatus)
	*p = x
	return p
}

func (x UserConnectionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserConnectionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_user_connection_proto_enumTypes[0].Descriptor()
}

func (UserConnectionStatus) Type() protoreflect.EnumType {
	return &file_user_connection_proto_enumTypes[0]
}

func (x UserConnectionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserConnectionStatus.Descriptor instead.
func (UserConnectionStatus) EnumDescriptor() ([]byte, []int) {
	return file_user_connection_proto_rawDescGZIP(), []int{0}
}

type UserConnectionType int32

const (
	UserConnectionType_USER_CONNECTION_TYPE_INVALID UserConnectionType = 0
	UserConnectionType_USER_CONNECTION_TYPE_LOCAL   UserConnectionType = 1
	UserConnectionType_USER_CONNECTION_TYPE_REMOTE  UserConnectionType = 2
)

// Enum value maps for UserConnectionType.
var (
	UserConnectionType_name = map[int32]string{
		0: "USER_CONNECTION_TYPE_INVALID",
		1: "USER_CONNECTION_TYPE_LOCAL",
		2: "USER_CONNECTION_TYPE_REMOTE",
	}
	UserConnectionType_value = map[string]int32{
		"USER_CONNECTION_TYPE_INVALID": 0,
		"USER_CONNECTION_TYPE_LOCAL":   1,
		"USER_CONNECTION_TYPE_REMOTE":  2,
	}
)

func (x UserConnectionType) Enum() *UserConnectionType {
	p := new(UserConnectionType)
	*p = x
	return p
}

func (x UserConnectionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserConnectionType) Descriptor() protoreflect.EnumDescriptor {
	return file_user_connection_proto_enumTypes[1].Descriptor()
}

func (UserConnectionType) Type() protoreflect.EnumType {
	return &file_user_connection_proto_enumTypes[1]
}

func (x UserConnectionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserConnectionType.Descriptor instead.
func (UserConnectionType) EnumDescriptor() ([]byte, []int) {
	return file_user_connection_proto_rawDescGZIP(), []int{1}
}

type UserConnection struct {
	state              protoimpl.MessageState     `protogen:"open.v1"`
	Uuid               string                     `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	UserUuid           string                     `protobuf:"bytes,2,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	ProjectUuid        string                     `protobuf:"bytes,3,opt,name=project_uuid,json=projectUuid,proto3" json:"project_uuid,omitempty"`
	ProjectVersionUuid string                     `protobuf:"bytes,4,opt,name=project_version_uuid,json=projectVersionUuid,proto3" json:"project_version_uuid,omitempty"`
	Type               UserConnectionType         `protobuf:"varint,5,opt,name=type,proto3,enum=nem.UserConnectionType" json:"type,omitempty"`
	TypeConfig         *UserConnectionTypeConfig  `protobuf:"bytes,6,opt,name=type_config,json=typeConfig,proto3" json:"type_config,omitempty"`
	DbSchema           string                     `protobuf:"bytes,7,opt,name=db_schema,json=dbSchema,proto3" json:"db_schema,omitempty"`
	Executions         []*UserConnectionExecution `protobuf:"bytes,8,rep,name=executions,proto3" json:"executions,omitempty"`
	Status             UserConnectionStatus       `protobuf:"varint,9,opt,name=status,proto3,enum=nem.UserConnectionStatus" json:"status,omitempty"`
	CreatedAt          *timestamppb.Timestamp     `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt          *timestamppb.Timestamp     `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *UserConnection) Reset() {
	*x = UserConnection{}
	mi := &file_user_connection_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserConnection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserConnection) ProtoMessage() {}

func (x *UserConnection) ProtoReflect() protoreflect.Message {
	mi := &file_user_connection_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserConnection.ProtoReflect.Descriptor instead.
func (*UserConnection) Descriptor() ([]byte, []int) {
	return file_user_connection_proto_rawDescGZIP(), []int{0}
}

func (x *UserConnection) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *UserConnection) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *UserConnection) GetProjectUuid() string {
	if x != nil {
		return x.ProjectUuid
	}
	return ""
}

func (x *UserConnection) GetProjectVersionUuid() string {
	if x != nil {
		return x.ProjectVersionUuid
	}
	return ""
}

func (x *UserConnection) GetType() UserConnectionType {
	if x != nil {
		return x.Type
	}
	return UserConnectionType_USER_CONNECTION_TYPE_INVALID
}

func (x *UserConnection) GetTypeConfig() *UserConnectionTypeConfig {
	if x != nil {
		return x.TypeConfig
	}
	return nil
}

func (x *UserConnection) GetDbSchema() string {
	if x != nil {
		return x.DbSchema
	}
	return ""
}

func (x *UserConnection) GetExecutions() []*UserConnectionExecution {
	if x != nil {
		return x.Executions
	}
	return nil
}

func (x *UserConnection) GetStatus() UserConnectionStatus {
	if x != nil {
		return x.Status
	}
	return UserConnectionStatus_USER_CONNECTION_STATUS_INVALID
}

func (x *UserConnection) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UserConnection) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_user_connection_proto protoreflect.FileDescriptor

const file_user_connection_proto_rawDesc = "" +
	"\n" +
	"\x15user_connection.proto\x12\x03nem\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1fuser_connection_execution.proto\x1a!user_connection_type_config.proto\"\x87\x04\n" +
	"\x0eUserConnection\x12\x12\n" +
	"\x04uuid\x18\x01 \x01(\tR\x04uuid\x12\x1b\n" +
	"\tuser_uuid\x18\x02 \x01(\tR\buserUuid\x12!\n" +
	"\fproject_uuid\x18\x03 \x01(\tR\vprojectUuid\x120\n" +
	"\x14project_version_uuid\x18\x04 \x01(\tR\x12projectVersionUuid\x12+\n" +
	"\x04type\x18\x05 \x01(\x0e2\x17.nem.UserConnectionTypeR\x04type\x12>\n" +
	"\vtype_config\x18\x06 \x01(\v2\x1d.nem.UserConnectionTypeConfigR\n" +
	"typeConfig\x12\x1b\n" +
	"\tdb_schema\x18\a \x01(\tR\bdbSchema\x12<\n" +
	"\n" +
	"executions\x18\b \x03(\v2\x1c.nem.UserConnectionExecutionR\n" +
	"executions\x121\n" +
	"\x06status\x18\t \x01(\x0e2\x19.nem.UserConnectionStatusR\x06status\x129\n" +
	"\n" +
	"created_at\x18\n" +
	" \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\v \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt*\x82\x01\n" +
	"\x14UserConnectionStatus\x12\"\n" +
	"\x1eUSER_CONNECTION_STATUS_INVALID\x10\x00\x12!\n" +
	"\x1dUSER_CONNECTION_STATUS_ACTIVE\x10\x01\x12#\n" +
	"\x1fUSER_CONNECTION_STATUS_DISABLED\x10\x02*w\n" +
	"\x12UserConnectionType\x12 \n" +
	"\x1cUSER_CONNECTION_TYPE_INVALID\x10\x00\x12\x1e\n" +
	"\x1aUSER_CONNECTION_TYPE_LOCAL\x10\x01\x12\x1f\n" +
	"\x1bUSER_CONNECTION_TYPE_REMOTE\x10\x02B5\n" +
	"\x14github.com/nuzur/nemB\x0eUserConnectionP\x01Z\vnem/idl/genb\x06proto3"

var (
	file_user_connection_proto_rawDescOnce sync.Once
	file_user_connection_proto_rawDescData []byte
)

func file_user_connection_proto_rawDescGZIP() []byte {
	file_user_connection_proto_rawDescOnce.Do(func() {
		file_user_connection_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_user_connection_proto_rawDesc), len(file_user_connection_proto_rawDesc)))
	})
	return file_user_connection_proto_rawDescData
}

var file_user_connection_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_user_connection_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_user_connection_proto_goTypes = []any{
	(UserConnectionStatus)(0),        // 0: nem.UserConnectionStatus
	(UserConnectionType)(0),          // 1: nem.UserConnectionType
	(*UserConnection)(nil),           // 2: nem.UserConnection
	(*UserConnectionTypeConfig)(nil), // 3: nem.UserConnectionTypeConfig
	(*UserConnectionExecution)(nil),  // 4: nem.UserConnectionExecution
	(*timestamppb.Timestamp)(nil),    // 5: google.protobuf.Timestamp
}
var file_user_connection_proto_depIdxs = []int32{
	1, // 0: nem.UserConnection.type:type_name -> nem.UserConnectionType
	3, // 1: nem.UserConnection.type_config:type_name -> nem.UserConnectionTypeConfig
	4, // 2: nem.UserConnection.executions:type_name -> nem.UserConnectionExecution
	0, // 3: nem.UserConnection.status:type_name -> nem.UserConnectionStatus
	5, // 4: nem.UserConnection.created_at:type_name -> google.protobuf.Timestamp
	5, // 5: nem.UserConnection.updated_at:type_name -> google.protobuf.Timestamp
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_user_connection_proto_init() }
func file_user_connection_proto_init() {
	if File_user_connection_proto != nil {
		return
	}
	file_user_connection_execution_proto_init()
	file_user_connection_type_config_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_user_connection_proto_rawDesc), len(file_user_connection_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_connection_proto_goTypes,
		DependencyIndexes: file_user_connection_proto_depIdxs,
		EnumInfos:         file_user_connection_proto_enumTypes,
		MessageInfos:      file_user_connection_proto_msgTypes,
	}.Build()
	File_user_connection_proto = out.File
	file_user_connection_proto_goTypes = nil
	file_user_connection_proto_depIdxs = nil
}
