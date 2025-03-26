// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: project_extension.proto

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

type ProjectExtensionStatus int32

const (
	ProjectExtensionStatus_PROJECT_EXTENSION_STATUS_INVALID  ProjectExtensionStatus = 0
	ProjectExtensionStatus_PROJECT_EXTENSION_STATUS_ACTIVE   ProjectExtensionStatus = 1
	ProjectExtensionStatus_PROJECT_EXTENSION_STATUS_DISABLED ProjectExtensionStatus = 2
)

// Enum value maps for ProjectExtensionStatus.
var (
	ProjectExtensionStatus_name = map[int32]string{
		0: "PROJECT_EXTENSION_STATUS_INVALID",
		1: "PROJECT_EXTENSION_STATUS_ACTIVE",
		2: "PROJECT_EXTENSION_STATUS_DISABLED",
	}
	ProjectExtensionStatus_value = map[string]int32{
		"PROJECT_EXTENSION_STATUS_INVALID":  0,
		"PROJECT_EXTENSION_STATUS_ACTIVE":   1,
		"PROJECT_EXTENSION_STATUS_DISABLED": 2,
	}
)

func (x ProjectExtensionStatus) Enum() *ProjectExtensionStatus {
	p := new(ProjectExtensionStatus)
	*p = x
	return p
}

func (x ProjectExtensionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProjectExtensionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_project_extension_proto_enumTypes[0].Descriptor()
}

func (ProjectExtensionStatus) Type() protoreflect.EnumType {
	return &file_project_extension_proto_enumTypes[0]
}

func (x ProjectExtensionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProjectExtensionStatus.Descriptor instead.
func (ProjectExtensionStatus) EnumDescriptor() ([]byte, []int) {
	return file_project_extension_proto_rawDescGZIP(), []int{0}
}

type ProjectExtension struct {
	state                     protoimpl.MessageState `protogen:"open.v1"`
	Uuid                      string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	ExtensionUuid             string                 `protobuf:"bytes,2,opt,name=extension_uuid,json=extensionUuid,proto3" json:"extension_uuid,omitempty"`
	ExtensionVersionUuid      string                 `protobuf:"bytes,3,opt,name=extension_version_uuid,json=extensionVersionUuid,proto3" json:"extension_version_uuid,omitempty"`
	ConfigurationEntityValues string                 `protobuf:"bytes,4,opt,name=configuration_entity_values,json=configurationEntityValues,proto3" json:"configuration_entity_values,omitempty"`
	Blocking                  bool                   `protobuf:"varint,5,opt,name=blocking,proto3" json:"blocking,omitempty"`
	Status                    ProjectExtensionStatus `protobuf:"varint,6,opt,name=status,proto3,enum=nem.ProjectExtensionStatus" json:"status,omitempty"`
	CreatedAt                 *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt                 *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CreatedByUuid             string                 `protobuf:"bytes,9,opt,name=created_by_uuid,json=createdByUuid,proto3" json:"created_by_uuid,omitempty"`
	UpdatedByUuid             string                 `protobuf:"bytes,10,opt,name=updated_by_uuid,json=updatedByUuid,proto3" json:"updated_by_uuid,omitempty"`
	unknownFields             protoimpl.UnknownFields
	sizeCache                 protoimpl.SizeCache
}

func (x *ProjectExtension) Reset() {
	*x = ProjectExtension{}
	mi := &file_project_extension_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProjectExtension) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectExtension) ProtoMessage() {}

func (x *ProjectExtension) ProtoReflect() protoreflect.Message {
	mi := &file_project_extension_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectExtension.ProtoReflect.Descriptor instead.
func (*ProjectExtension) Descriptor() ([]byte, []int) {
	return file_project_extension_proto_rawDescGZIP(), []int{0}
}

func (x *ProjectExtension) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *ProjectExtension) GetExtensionUuid() string {
	if x != nil {
		return x.ExtensionUuid
	}
	return ""
}

func (x *ProjectExtension) GetExtensionVersionUuid() string {
	if x != nil {
		return x.ExtensionVersionUuid
	}
	return ""
}

func (x *ProjectExtension) GetConfigurationEntityValues() string {
	if x != nil {
		return x.ConfigurationEntityValues
	}
	return ""
}

func (x *ProjectExtension) GetBlocking() bool {
	if x != nil {
		return x.Blocking
	}
	return false
}

func (x *ProjectExtension) GetStatus() ProjectExtensionStatus {
	if x != nil {
		return x.Status
	}
	return ProjectExtensionStatus_PROJECT_EXTENSION_STATUS_INVALID
}

func (x *ProjectExtension) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ProjectExtension) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *ProjectExtension) GetCreatedByUuid() string {
	if x != nil {
		return x.CreatedByUuid
	}
	return ""
}

func (x *ProjectExtension) GetUpdatedByUuid() string {
	if x != nil {
		return x.UpdatedByUuid
	}
	return ""
}

var File_project_extension_proto protoreflect.FileDescriptor

const file_project_extension_proto_rawDesc = "" +
	"\n" +
	"\x17project_extension.proto\x12\x03nem\x1a\x1fgoogle/protobuf/timestamp.proto\"\xda\x03\n" +
	"\x10ProjectExtension\x12\x12\n" +
	"\x04uuid\x18\x01 \x01(\tR\x04uuid\x12%\n" +
	"\x0eextension_uuid\x18\x02 \x01(\tR\rextensionUuid\x124\n" +
	"\x16extension_version_uuid\x18\x03 \x01(\tR\x14extensionVersionUuid\x12>\n" +
	"\x1bconfiguration_entity_values\x18\x04 \x01(\tR\x19configurationEntityValues\x12\x1a\n" +
	"\bblocking\x18\x05 \x01(\bR\bblocking\x123\n" +
	"\x06status\x18\x06 \x01(\x0e2\x1b.nem.ProjectExtensionStatusR\x06status\x129\n" +
	"\n" +
	"created_at\x18\a \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\b \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\x12&\n" +
	"\x0fcreated_by_uuid\x18\t \x01(\tR\rcreatedByUuid\x12&\n" +
	"\x0fupdated_by_uuid\x18\n" +
	" \x01(\tR\rupdatedByUuid*\x8a\x01\n" +
	"\x16ProjectExtensionStatus\x12$\n" +
	" PROJECT_EXTENSION_STATUS_INVALID\x10\x00\x12#\n" +
	"\x1fPROJECT_EXTENSION_STATUS_ACTIVE\x10\x01\x12%\n" +
	"!PROJECT_EXTENSION_STATUS_DISABLED\x10\x02B7\n" +
	"\x14github.com/nuzur/nemB\x10ProjectExtensionP\x01Z\vnem/idl/genb\x06proto3"

var (
	file_project_extension_proto_rawDescOnce sync.Once
	file_project_extension_proto_rawDescData []byte
)

func file_project_extension_proto_rawDescGZIP() []byte {
	file_project_extension_proto_rawDescOnce.Do(func() {
		file_project_extension_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_project_extension_proto_rawDesc), len(file_project_extension_proto_rawDesc)))
	})
	return file_project_extension_proto_rawDescData
}

var file_project_extension_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_project_extension_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_project_extension_proto_goTypes = []any{
	(ProjectExtensionStatus)(0),   // 0: nem.ProjectExtensionStatus
	(*ProjectExtension)(nil),      // 1: nem.ProjectExtension
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_project_extension_proto_depIdxs = []int32{
	0, // 0: nem.ProjectExtension.status:type_name -> nem.ProjectExtensionStatus
	2, // 1: nem.ProjectExtension.created_at:type_name -> google.protobuf.Timestamp
	2, // 2: nem.ProjectExtension.updated_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_project_extension_proto_init() }
func file_project_extension_proto_init() {
	if File_project_extension_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_project_extension_proto_rawDesc), len(file_project_extension_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_extension_proto_goTypes,
		DependencyIndexes: file_project_extension_proto_depIdxs,
		EnumInfos:         file_project_extension_proto_enumTypes,
		MessageInfos:      file_project_extension_proto_msgTypes,
	}.Build()
	File_project_extension_proto = out.File
	file_project_extension_proto_goTypes = nil
	file_project_extension_proto_depIdxs = nil
}
