// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: extension_version.proto

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

type ExtensionVersionExecutionMode int32

const (
	ExtensionVersionExecutionMode_EXTENSION_VERSION_EXECUTION_MODE_INVALID   ExtensionVersionExecutionMode = 0
	ExtensionVersionExecutionMode_EXTENSION_VERSION_EXECUTION_MODE_MANUAL    ExtensionVersionExecutionMode = 1
	ExtensionVersionExecutionMode_EXTENSION_VERSION_EXECUTION_MODE_AUTOMATIC ExtensionVersionExecutionMode = 2
)

// Enum value maps for ExtensionVersionExecutionMode.
var (
	ExtensionVersionExecutionMode_name = map[int32]string{
		0: "EXTENSION_VERSION_EXECUTION_MODE_INVALID",
		1: "EXTENSION_VERSION_EXECUTION_MODE_MANUAL",
		2: "EXTENSION_VERSION_EXECUTION_MODE_AUTOMATIC",
	}
	ExtensionVersionExecutionMode_value = map[string]int32{
		"EXTENSION_VERSION_EXECUTION_MODE_INVALID":   0,
		"EXTENSION_VERSION_EXECUTION_MODE_MANUAL":    1,
		"EXTENSION_VERSION_EXECUTION_MODE_AUTOMATIC": 2,
	}
)

func (x ExtensionVersionExecutionMode) Enum() *ExtensionVersionExecutionMode {
	p := new(ExtensionVersionExecutionMode)
	*p = x
	return p
}

func (x ExtensionVersionExecutionMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExtensionVersionExecutionMode) Descriptor() protoreflect.EnumDescriptor {
	return file_extension_version_proto_enumTypes[0].Descriptor()
}

func (ExtensionVersionExecutionMode) Type() protoreflect.EnumType {
	return &file_extension_version_proto_enumTypes[0]
}

func (x ExtensionVersionExecutionMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExtensionVersionExecutionMode.Descriptor instead.
func (ExtensionVersionExecutionMode) EnumDescriptor() ([]byte, []int) {
	return file_extension_version_proto_rawDescGZIP(), []int{0}
}

type ExtensionVersionReviewStatus int32

const (
	ExtensionVersionReviewStatus_EXTENSION_VERSION_REVIEW_STATUS_INVALID   ExtensionVersionReviewStatus = 0
	ExtensionVersionReviewStatus_EXTENSION_VERSION_REVIEW_STATUS_DRAFT     ExtensionVersionReviewStatus = 1
	ExtensionVersionReviewStatus_EXTENSION_VERSION_REVIEW_STATUS_IN_REVIEW ExtensionVersionReviewStatus = 2
	ExtensionVersionReviewStatus_EXTENSION_VERSION_REVIEW_STATUS_APPROVED  ExtensionVersionReviewStatus = 3
	ExtensionVersionReviewStatus_EXTENSION_VERSION_REVIEW_STATUS_REJECTED  ExtensionVersionReviewStatus = 4
	ExtensionVersionReviewStatus_EXTENSION_VERSION_REVIEW_STATUS_PUBLISHED ExtensionVersionReviewStatus = 5
)

// Enum value maps for ExtensionVersionReviewStatus.
var (
	ExtensionVersionReviewStatus_name = map[int32]string{
		0: "EXTENSION_VERSION_REVIEW_STATUS_INVALID",
		1: "EXTENSION_VERSION_REVIEW_STATUS_DRAFT",
		2: "EXTENSION_VERSION_REVIEW_STATUS_IN_REVIEW",
		3: "EXTENSION_VERSION_REVIEW_STATUS_APPROVED",
		4: "EXTENSION_VERSION_REVIEW_STATUS_REJECTED",
		5: "EXTENSION_VERSION_REVIEW_STATUS_PUBLISHED",
	}
	ExtensionVersionReviewStatus_value = map[string]int32{
		"EXTENSION_VERSION_REVIEW_STATUS_INVALID":   0,
		"EXTENSION_VERSION_REVIEW_STATUS_DRAFT":     1,
		"EXTENSION_VERSION_REVIEW_STATUS_IN_REVIEW": 2,
		"EXTENSION_VERSION_REVIEW_STATUS_APPROVED":  3,
		"EXTENSION_VERSION_REVIEW_STATUS_REJECTED":  4,
		"EXTENSION_VERSION_REVIEW_STATUS_PUBLISHED": 5,
	}
)

func (x ExtensionVersionReviewStatus) Enum() *ExtensionVersionReviewStatus {
	p := new(ExtensionVersionReviewStatus)
	*p = x
	return p
}

func (x ExtensionVersionReviewStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExtensionVersionReviewStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_extension_version_proto_enumTypes[1].Descriptor()
}

func (ExtensionVersionReviewStatus) Type() protoreflect.EnumType {
	return &file_extension_version_proto_enumTypes[1]
}

func (x ExtensionVersionReviewStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExtensionVersionReviewStatus.Descriptor instead.
func (ExtensionVersionReviewStatus) EnumDescriptor() ([]byte, []int) {
	return file_extension_version_proto_rawDescGZIP(), []int{1}
}

type ExtensionVersionStatus int32

const (
	ExtensionVersionStatus_EXTENSION_VERSION_STATUS_INVALID  ExtensionVersionStatus = 0
	ExtensionVersionStatus_EXTENSION_VERSION_STATUS_ACTIVE   ExtensionVersionStatus = 1
	ExtensionVersionStatus_EXTENSION_VERSION_STATUS_DISABLED ExtensionVersionStatus = 2
)

// Enum value maps for ExtensionVersionStatus.
var (
	ExtensionVersionStatus_name = map[int32]string{
		0: "EXTENSION_VERSION_STATUS_INVALID",
		1: "EXTENSION_VERSION_STATUS_ACTIVE",
		2: "EXTENSION_VERSION_STATUS_DISABLED",
	}
	ExtensionVersionStatus_value = map[string]int32{
		"EXTENSION_VERSION_STATUS_INVALID":  0,
		"EXTENSION_VERSION_STATUS_ACTIVE":   1,
		"EXTENSION_VERSION_STATUS_DISABLED": 2,
	}
)

func (x ExtensionVersionStatus) Enum() *ExtensionVersionStatus {
	p := new(ExtensionVersionStatus)
	*p = x
	return p
}

func (x ExtensionVersionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExtensionVersionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_extension_version_proto_enumTypes[2].Descriptor()
}

func (ExtensionVersionStatus) Type() protoreflect.EnumType {
	return &file_extension_version_proto_enumTypes[2]
}

func (x ExtensionVersionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExtensionVersionStatus.Descriptor instead.
func (ExtensionVersionStatus) EnumDescriptor() ([]byte, []int) {
	return file_extension_version_proto_rawDescGZIP(), []int{2}
}

type ExtensionVersion struct {
	state               protoimpl.MessageState          `protogen:"open.v1"`
	Uuid                string                          `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Version             int64                           `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	ExtensionUuid       string                          `protobuf:"bytes,3,opt,name=extension_uuid,json=extensionUuid,proto3" json:"extension_uuid,omitempty"`
	DisplayVersion      string                          `protobuf:"bytes,4,opt,name=display_version,json=displayVersion,proto3" json:"display_version,omitempty"`
	Description         string                          `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	RepositoryTag       string                          `protobuf:"bytes,6,opt,name=repository_tag,json=repositoryTag,proto3" json:"repository_tag,omitempty"`
	ConfigurationEntity string                          `protobuf:"bytes,7,opt,name=configuration_entity,json=configurationEntity,proto3" json:"configuration_entity,omitempty"`
	ExecutionMode       []ExtensionVersionExecutionMode `protobuf:"varint,8,rep,packed,name=execution_mode,json=executionMode,proto3,enum=nem.ExtensionVersionExecutionMode" json:"execution_mode,omitempty"`
	ReviewStatus        ExtensionVersionReviewStatus    `protobuf:"varint,9,opt,name=review_status,json=reviewStatus,proto3,enum=nem.ExtensionVersionReviewStatus" json:"review_status,omitempty"`
	Status              ExtensionVersionStatus          `protobuf:"varint,10,opt,name=status,proto3,enum=nem.ExtensionVersionStatus" json:"status,omitempty"`
	CreatedAt           *timestamppb.Timestamp          `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt           *timestamppb.Timestamp          `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CreatedByUuid       string                          `protobuf:"bytes,13,opt,name=created_by_uuid,json=createdByUuid,proto3" json:"created_by_uuid,omitempty"`
	UpdatedByUuid       string                          `protobuf:"bytes,14,opt,name=updated_by_uuid,json=updatedByUuid,proto3" json:"updated_by_uuid,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *ExtensionVersion) Reset() {
	*x = ExtensionVersion{}
	mi := &file_extension_version_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExtensionVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtensionVersion) ProtoMessage() {}

func (x *ExtensionVersion) ProtoReflect() protoreflect.Message {
	mi := &file_extension_version_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtensionVersion.ProtoReflect.Descriptor instead.
func (*ExtensionVersion) Descriptor() ([]byte, []int) {
	return file_extension_version_proto_rawDescGZIP(), []int{0}
}

func (x *ExtensionVersion) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *ExtensionVersion) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *ExtensionVersion) GetExtensionUuid() string {
	if x != nil {
		return x.ExtensionUuid
	}
	return ""
}

func (x *ExtensionVersion) GetDisplayVersion() string {
	if x != nil {
		return x.DisplayVersion
	}
	return ""
}

func (x *ExtensionVersion) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ExtensionVersion) GetRepositoryTag() string {
	if x != nil {
		return x.RepositoryTag
	}
	return ""
}

func (x *ExtensionVersion) GetConfigurationEntity() string {
	if x != nil {
		return x.ConfigurationEntity
	}
	return ""
}

func (x *ExtensionVersion) GetExecutionMode() []ExtensionVersionExecutionMode {
	if x != nil {
		return x.ExecutionMode
	}
	return nil
}

func (x *ExtensionVersion) GetReviewStatus() ExtensionVersionReviewStatus {
	if x != nil {
		return x.ReviewStatus
	}
	return ExtensionVersionReviewStatus_EXTENSION_VERSION_REVIEW_STATUS_INVALID
}

func (x *ExtensionVersion) GetStatus() ExtensionVersionStatus {
	if x != nil {
		return x.Status
	}
	return ExtensionVersionStatus_EXTENSION_VERSION_STATUS_INVALID
}

func (x *ExtensionVersion) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ExtensionVersion) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *ExtensionVersion) GetCreatedByUuid() string {
	if x != nil {
		return x.CreatedByUuid
	}
	return ""
}

func (x *ExtensionVersion) GetUpdatedByUuid() string {
	if x != nil {
		return x.UpdatedByUuid
	}
	return ""
}

var File_extension_version_proto protoreflect.FileDescriptor

const file_extension_version_proto_rawDesc = "" +
	"\n" +
	"\x17extension_version.proto\x12\x03nem\x1a\x1fgoogle/protobuf/timestamp.proto\"\x9a\x05\n" +
	"\x10ExtensionVersion\x12\x12\n" +
	"\x04uuid\x18\x01 \x01(\tR\x04uuid\x12\x18\n" +
	"\aversion\x18\x02 \x01(\x03R\aversion\x12%\n" +
	"\x0eextension_uuid\x18\x03 \x01(\tR\rextensionUuid\x12'\n" +
	"\x0fdisplay_version\x18\x04 \x01(\tR\x0edisplayVersion\x12 \n" +
	"\vdescription\x18\x05 \x01(\tR\vdescription\x12%\n" +
	"\x0erepository_tag\x18\x06 \x01(\tR\rrepositoryTag\x121\n" +
	"\x14configuration_entity\x18\a \x01(\tR\x13configurationEntity\x12I\n" +
	"\x0eexecution_mode\x18\b \x03(\x0e2\".nem.ExtensionVersionExecutionModeR\rexecutionMode\x12F\n" +
	"\rreview_status\x18\t \x01(\x0e2!.nem.ExtensionVersionReviewStatusR\freviewStatus\x123\n" +
	"\x06status\x18\n" +
	" \x01(\x0e2\x1b.nem.ExtensionVersionStatusR\x06status\x129\n" +
	"\n" +
	"created_at\x18\v \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\f \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\x12&\n" +
	"\x0fcreated_by_uuid\x18\r \x01(\tR\rcreatedByUuid\x12&\n" +
	"\x0fupdated_by_uuid\x18\x0e \x01(\tR\rupdatedByUuid*\xaa\x01\n" +
	"\x1dExtensionVersionExecutionMode\x12,\n" +
	"(EXTENSION_VERSION_EXECUTION_MODE_INVALID\x10\x00\x12+\n" +
	"'EXTENSION_VERSION_EXECUTION_MODE_MANUAL\x10\x01\x12.\n" +
	"*EXTENSION_VERSION_EXECUTION_MODE_AUTOMATIC\x10\x02*\xb0\x02\n" +
	"\x1cExtensionVersionReviewStatus\x12+\n" +
	"'EXTENSION_VERSION_REVIEW_STATUS_INVALID\x10\x00\x12)\n" +
	"%EXTENSION_VERSION_REVIEW_STATUS_DRAFT\x10\x01\x12-\n" +
	")EXTENSION_VERSION_REVIEW_STATUS_IN_REVIEW\x10\x02\x12,\n" +
	"(EXTENSION_VERSION_REVIEW_STATUS_APPROVED\x10\x03\x12,\n" +
	"(EXTENSION_VERSION_REVIEW_STATUS_REJECTED\x10\x04\x12-\n" +
	")EXTENSION_VERSION_REVIEW_STATUS_PUBLISHED\x10\x05*\x8a\x01\n" +
	"\x16ExtensionVersionStatus\x12$\n" +
	" EXTENSION_VERSION_STATUS_INVALID\x10\x00\x12#\n" +
	"\x1fEXTENSION_VERSION_STATUS_ACTIVE\x10\x01\x12%\n" +
	"!EXTENSION_VERSION_STATUS_DISABLED\x10\x02B7\n" +
	"\x14github.com/nuzur/nemB\x10ExtensionVersionP\x01Z\vnem/idl/genb\x06proto3"

var (
	file_extension_version_proto_rawDescOnce sync.Once
	file_extension_version_proto_rawDescData []byte
)

func file_extension_version_proto_rawDescGZIP() []byte {
	file_extension_version_proto_rawDescOnce.Do(func() {
		file_extension_version_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_extension_version_proto_rawDesc), len(file_extension_version_proto_rawDesc)))
	})
	return file_extension_version_proto_rawDescData
}

var file_extension_version_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_extension_version_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_extension_version_proto_goTypes = []any{
	(ExtensionVersionExecutionMode)(0), // 0: nem.ExtensionVersionExecutionMode
	(ExtensionVersionReviewStatus)(0),  // 1: nem.ExtensionVersionReviewStatus
	(ExtensionVersionStatus)(0),        // 2: nem.ExtensionVersionStatus
	(*ExtensionVersion)(nil),           // 3: nem.ExtensionVersion
	(*timestamppb.Timestamp)(nil),      // 4: google.protobuf.Timestamp
}
var file_extension_version_proto_depIdxs = []int32{
	0, // 0: nem.ExtensionVersion.execution_mode:type_name -> nem.ExtensionVersionExecutionMode
	1, // 1: nem.ExtensionVersion.review_status:type_name -> nem.ExtensionVersionReviewStatus
	2, // 2: nem.ExtensionVersion.status:type_name -> nem.ExtensionVersionStatus
	4, // 3: nem.ExtensionVersion.created_at:type_name -> google.protobuf.Timestamp
	4, // 4: nem.ExtensionVersion.updated_at:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_extension_version_proto_init() }
func file_extension_version_proto_init() {
	if File_extension_version_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_extension_version_proto_rawDesc), len(file_extension_version_proto_rawDesc)),
			NumEnums:      3,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_extension_version_proto_goTypes,
		DependencyIndexes: file_extension_version_proto_depIdxs,
		EnumInfos:         file_extension_version_proto_enumTypes,
		MessageInfos:      file_extension_version_proto_msgTypes,
	}.Build()
	File_extension_version_proto = out.File
	file_extension_version_proto_goTypes = nil
	file_extension_version_proto_depIdxs = nil
}
