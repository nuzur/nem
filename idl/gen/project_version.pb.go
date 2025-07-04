// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: project_version.proto

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

type ProjectVersionReviewStatus int32

const (
	ProjectVersionReviewStatus_PROJECT_VERSION_REVIEW_STATUS_INVALID   ProjectVersionReviewStatus = 0
	ProjectVersionReviewStatus_PROJECT_VERSION_REVIEW_STATUS_DRAFT     ProjectVersionReviewStatus = 1
	ProjectVersionReviewStatus_PROJECT_VERSION_REVIEW_STATUS_IN_REVIEW ProjectVersionReviewStatus = 2
	ProjectVersionReviewStatus_PROJECT_VERSION_REVIEW_STATUS_APPROVED  ProjectVersionReviewStatus = 3
	ProjectVersionReviewStatus_PROJECT_VERSION_REVIEW_STATUS_REJECTED  ProjectVersionReviewStatus = 4
	ProjectVersionReviewStatus_PROJECT_VERSION_REVIEW_STATUS_PUBLISHED ProjectVersionReviewStatus = 5
)

// Enum value maps for ProjectVersionReviewStatus.
var (
	ProjectVersionReviewStatus_name = map[int32]string{
		0: "PROJECT_VERSION_REVIEW_STATUS_INVALID",
		1: "PROJECT_VERSION_REVIEW_STATUS_DRAFT",
		2: "PROJECT_VERSION_REVIEW_STATUS_IN_REVIEW",
		3: "PROJECT_VERSION_REVIEW_STATUS_APPROVED",
		4: "PROJECT_VERSION_REVIEW_STATUS_REJECTED",
		5: "PROJECT_VERSION_REVIEW_STATUS_PUBLISHED",
	}
	ProjectVersionReviewStatus_value = map[string]int32{
		"PROJECT_VERSION_REVIEW_STATUS_INVALID":   0,
		"PROJECT_VERSION_REVIEW_STATUS_DRAFT":     1,
		"PROJECT_VERSION_REVIEW_STATUS_IN_REVIEW": 2,
		"PROJECT_VERSION_REVIEW_STATUS_APPROVED":  3,
		"PROJECT_VERSION_REVIEW_STATUS_REJECTED":  4,
		"PROJECT_VERSION_REVIEW_STATUS_PUBLISHED": 5,
	}
)

func (x ProjectVersionReviewStatus) Enum() *ProjectVersionReviewStatus {
	p := new(ProjectVersionReviewStatus)
	*p = x
	return p
}

func (x ProjectVersionReviewStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProjectVersionReviewStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_project_version_proto_enumTypes[0].Descriptor()
}

func (ProjectVersionReviewStatus) Type() protoreflect.EnumType {
	return &file_project_version_proto_enumTypes[0]
}

func (x ProjectVersionReviewStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProjectVersionReviewStatus.Descriptor instead.
func (ProjectVersionReviewStatus) EnumDescriptor() ([]byte, []int) {
	return file_project_version_proto_rawDescGZIP(), []int{0}
}

type ProjectVersionStatus int32

const (
	ProjectVersionStatus_PROJECT_VERSION_STATUS_INVALID  ProjectVersionStatus = 0
	ProjectVersionStatus_PROJECT_VERSION_STATUS_ACTIVE   ProjectVersionStatus = 1
	ProjectVersionStatus_PROJECT_VERSION_STATUS_DISABLED ProjectVersionStatus = 2
)

// Enum value maps for ProjectVersionStatus.
var (
	ProjectVersionStatus_name = map[int32]string{
		0: "PROJECT_VERSION_STATUS_INVALID",
		1: "PROJECT_VERSION_STATUS_ACTIVE",
		2: "PROJECT_VERSION_STATUS_DISABLED",
	}
	ProjectVersionStatus_value = map[string]int32{
		"PROJECT_VERSION_STATUS_INVALID":  0,
		"PROJECT_VERSION_STATUS_ACTIVE":   1,
		"PROJECT_VERSION_STATUS_DISABLED": 2,
	}
)

func (x ProjectVersionStatus) Enum() *ProjectVersionStatus {
	p := new(ProjectVersionStatus)
	*p = x
	return p
}

func (x ProjectVersionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProjectVersionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_project_version_proto_enumTypes[1].Descriptor()
}

func (ProjectVersionStatus) Type() protoreflect.EnumType {
	return &file_project_version_proto_enumTypes[1]
}

func (x ProjectVersionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProjectVersionStatus.Descriptor instead.
func (ProjectVersionStatus) EnumDescriptor() ([]byte, []int) {
	return file_project_version_proto_rawDescGZIP(), []int{1}
}

type ProjectVersion struct {
	state           protoimpl.MessageState      `protogen:"open.v1"`
	Uuid            string                      `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Version         int64                       `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	Identifier      string                      `protobuf:"bytes,3,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Description     string                      `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	ProjectUuid     string                      `protobuf:"bytes,5,opt,name=project_uuid,json=projectUuid,proto3" json:"project_uuid,omitempty"`
	Entities        []*Entity                   `protobuf:"bytes,6,rep,name=entities,proto3" json:"entities,omitempty"`
	Relationships   []*Relationship             `protobuf:"bytes,7,rep,name=relationships,proto3" json:"relationships,omitempty"`
	Enums           []*Enum                     `protobuf:"bytes,8,rep,name=enums,proto3" json:"enums,omitempty"`
	Services        []*Service                  `protobuf:"bytes,9,rep,name=services,proto3" json:"services,omitempty"`
	BaseVersionUuid string                      `protobuf:"bytes,10,opt,name=base_version_uuid,json=baseVersionUuid,proto3" json:"base_version_uuid,omitempty"`
	ReviewStatus    ProjectVersionReviewStatus  `protobuf:"varint,11,opt,name=review_status,json=reviewStatus,proto3,enum=nem.ProjectVersionReviewStatus" json:"review_status,omitempty"`
	Reviews         []*ProjectVersionReview     `protobuf:"bytes,12,rep,name=reviews,proto3" json:"reviews,omitempty"`
	Deployments     []*ProjectVersionDeployment `protobuf:"bytes,13,rep,name=deployments,proto3" json:"deployments,omitempty"`
	Status          ProjectVersionStatus        `protobuf:"varint,14,opt,name=status,proto3,enum=nem.ProjectVersionStatus" json:"status,omitempty"`
	CreatedAt       *timestamppb.Timestamp      `protobuf:"bytes,15,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt       *timestamppb.Timestamp      `protobuf:"bytes,16,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CreatedByUuid   string                      `protobuf:"bytes,17,opt,name=created_by_uuid,json=createdByUuid,proto3" json:"created_by_uuid,omitempty"`
	UpdatedByUuid   string                      `protobuf:"bytes,18,opt,name=updated_by_uuid,json=updatedByUuid,proto3" json:"updated_by_uuid,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ProjectVersion) Reset() {
	*x = ProjectVersion{}
	mi := &file_project_version_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProjectVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectVersion) ProtoMessage() {}

func (x *ProjectVersion) ProtoReflect() protoreflect.Message {
	mi := &file_project_version_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectVersion.ProtoReflect.Descriptor instead.
func (*ProjectVersion) Descriptor() ([]byte, []int) {
	return file_project_version_proto_rawDescGZIP(), []int{0}
}

func (x *ProjectVersion) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *ProjectVersion) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *ProjectVersion) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *ProjectVersion) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProjectVersion) GetProjectUuid() string {
	if x != nil {
		return x.ProjectUuid
	}
	return ""
}

func (x *ProjectVersion) GetEntities() []*Entity {
	if x != nil {
		return x.Entities
	}
	return nil
}

func (x *ProjectVersion) GetRelationships() []*Relationship {
	if x != nil {
		return x.Relationships
	}
	return nil
}

func (x *ProjectVersion) GetEnums() []*Enum {
	if x != nil {
		return x.Enums
	}
	return nil
}

func (x *ProjectVersion) GetServices() []*Service {
	if x != nil {
		return x.Services
	}
	return nil
}

func (x *ProjectVersion) GetBaseVersionUuid() string {
	if x != nil {
		return x.BaseVersionUuid
	}
	return ""
}

func (x *ProjectVersion) GetReviewStatus() ProjectVersionReviewStatus {
	if x != nil {
		return x.ReviewStatus
	}
	return ProjectVersionReviewStatus_PROJECT_VERSION_REVIEW_STATUS_INVALID
}

func (x *ProjectVersion) GetReviews() []*ProjectVersionReview {
	if x != nil {
		return x.Reviews
	}
	return nil
}

func (x *ProjectVersion) GetDeployments() []*ProjectVersionDeployment {
	if x != nil {
		return x.Deployments
	}
	return nil
}

func (x *ProjectVersion) GetStatus() ProjectVersionStatus {
	if x != nil {
		return x.Status
	}
	return ProjectVersionStatus_PROJECT_VERSION_STATUS_INVALID
}

func (x *ProjectVersion) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ProjectVersion) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *ProjectVersion) GetCreatedByUuid() string {
	if x != nil {
		return x.CreatedByUuid
	}
	return ""
}

func (x *ProjectVersion) GetUpdatedByUuid() string {
	if x != nil {
		return x.UpdatedByUuid
	}
	return ""
}

var File_project_version_proto protoreflect.FileDescriptor

const file_project_version_proto_rawDesc = "" +
	"\n" +
	"\x15project_version.proto\x12\x03nem\x1a\fentity.proto\x1a\n" +
	"enum.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a project_version_deployment.proto\x1a\x1cproject_version_review.proto\x1a\x12relationship.proto\x1a\rservice.proto\"\xb1\x06\n" +
	"\x0eProjectVersion\x12\x12\n" +
	"\x04uuid\x18\x01 \x01(\tR\x04uuid\x12\x18\n" +
	"\aversion\x18\x02 \x01(\x03R\aversion\x12\x1e\n" +
	"\n" +
	"identifier\x18\x03 \x01(\tR\n" +
	"identifier\x12 \n" +
	"\vdescription\x18\x04 \x01(\tR\vdescription\x12!\n" +
	"\fproject_uuid\x18\x05 \x01(\tR\vprojectUuid\x12'\n" +
	"\bentities\x18\x06 \x03(\v2\v.nem.EntityR\bentities\x127\n" +
	"\rrelationships\x18\a \x03(\v2\x11.nem.RelationshipR\rrelationships\x12\x1f\n" +
	"\x05enums\x18\b \x03(\v2\t.nem.EnumR\x05enums\x12(\n" +
	"\bservices\x18\t \x03(\v2\f.nem.ServiceR\bservices\x12*\n" +
	"\x11base_version_uuid\x18\n" +
	" \x01(\tR\x0fbaseVersionUuid\x12D\n" +
	"\rreview_status\x18\v \x01(\x0e2\x1f.nem.ProjectVersionReviewStatusR\freviewStatus\x123\n" +
	"\areviews\x18\f \x03(\v2\x19.nem.ProjectVersionReviewR\areviews\x12?\n" +
	"\vdeployments\x18\r \x03(\v2\x1d.nem.ProjectVersionDeploymentR\vdeployments\x121\n" +
	"\x06status\x18\x0e \x01(\x0e2\x19.nem.ProjectVersionStatusR\x06status\x129\n" +
	"\n" +
	"created_at\x18\x0f \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\x10 \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\x12&\n" +
	"\x0fcreated_by_uuid\x18\x11 \x01(\tR\rcreatedByUuid\x12&\n" +
	"\x0fupdated_by_uuid\x18\x12 \x01(\tR\rupdatedByUuid*\xa2\x02\n" +
	"\x1aProjectVersionReviewStatus\x12)\n" +
	"%PROJECT_VERSION_REVIEW_STATUS_INVALID\x10\x00\x12'\n" +
	"#PROJECT_VERSION_REVIEW_STATUS_DRAFT\x10\x01\x12+\n" +
	"'PROJECT_VERSION_REVIEW_STATUS_IN_REVIEW\x10\x02\x12*\n" +
	"&PROJECT_VERSION_REVIEW_STATUS_APPROVED\x10\x03\x12*\n" +
	"&PROJECT_VERSION_REVIEW_STATUS_REJECTED\x10\x04\x12+\n" +
	"'PROJECT_VERSION_REVIEW_STATUS_PUBLISHED\x10\x05*\x82\x01\n" +
	"\x14ProjectVersionStatus\x12\"\n" +
	"\x1ePROJECT_VERSION_STATUS_INVALID\x10\x00\x12!\n" +
	"\x1dPROJECT_VERSION_STATUS_ACTIVE\x10\x01\x12#\n" +
	"\x1fPROJECT_VERSION_STATUS_DISABLED\x10\x02B5\n" +
	"\x14github.com/nuzur/nemB\x0eProjectVersionP\x01Z\vnem/idl/genb\x06proto3"

var (
	file_project_version_proto_rawDescOnce sync.Once
	file_project_version_proto_rawDescData []byte
)

func file_project_version_proto_rawDescGZIP() []byte {
	file_project_version_proto_rawDescOnce.Do(func() {
		file_project_version_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_project_version_proto_rawDesc), len(file_project_version_proto_rawDesc)))
	})
	return file_project_version_proto_rawDescData
}

var file_project_version_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_project_version_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_project_version_proto_goTypes = []any{
	(ProjectVersionReviewStatus)(0),  // 0: nem.ProjectVersionReviewStatus
	(ProjectVersionStatus)(0),        // 1: nem.ProjectVersionStatus
	(*ProjectVersion)(nil),           // 2: nem.ProjectVersion
	(*Entity)(nil),                   // 3: nem.Entity
	(*Relationship)(nil),             // 4: nem.Relationship
	(*Enum)(nil),                     // 5: nem.Enum
	(*Service)(nil),                  // 6: nem.Service
	(*ProjectVersionReview)(nil),     // 7: nem.ProjectVersionReview
	(*ProjectVersionDeployment)(nil), // 8: nem.ProjectVersionDeployment
	(*timestamppb.Timestamp)(nil),    // 9: google.protobuf.Timestamp
}
var file_project_version_proto_depIdxs = []int32{
	3,  // 0: nem.ProjectVersion.entities:type_name -> nem.Entity
	4,  // 1: nem.ProjectVersion.relationships:type_name -> nem.Relationship
	5,  // 2: nem.ProjectVersion.enums:type_name -> nem.Enum
	6,  // 3: nem.ProjectVersion.services:type_name -> nem.Service
	0,  // 4: nem.ProjectVersion.review_status:type_name -> nem.ProjectVersionReviewStatus
	7,  // 5: nem.ProjectVersion.reviews:type_name -> nem.ProjectVersionReview
	8,  // 6: nem.ProjectVersion.deployments:type_name -> nem.ProjectVersionDeployment
	1,  // 7: nem.ProjectVersion.status:type_name -> nem.ProjectVersionStatus
	9,  // 8: nem.ProjectVersion.created_at:type_name -> google.protobuf.Timestamp
	9,  // 9: nem.ProjectVersion.updated_at:type_name -> google.protobuf.Timestamp
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_project_version_proto_init() }
func file_project_version_proto_init() {
	if File_project_version_proto != nil {
		return
	}
	file_entity_proto_init()
	file_enum_proto_init()
	file_project_version_deployment_proto_init()
	file_project_version_review_proto_init()
	file_relationship_proto_init()
	file_service_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_project_version_proto_rawDesc), len(file_project_version_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_version_proto_goTypes,
		DependencyIndexes: file_project_version_proto_depIdxs,
		EnumInfos:         file_project_version_proto_enumTypes,
		MessageInfos:      file_project_version_proto_msgTypes,
	}.Build()
	File_project_version_proto = out.File
	file_project_version_proto_goTypes = nil
	file_project_version_proto_depIdxs = nil
}
