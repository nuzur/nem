// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: change_request.proto

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

type ChangeRequestChangeType int32

const (
	ChangeRequestChangeType_CHANGE_REQUEST_CHANGE_TYPE_INVALID         ChangeRequestChangeType = 0
	ChangeRequestChangeType_CHANGE_REQUEST_CHANGE_TYPE_PROJECT_DATA    ChangeRequestChangeType = 1
	ChangeRequestChangeType_CHANGE_REQUEST_CHANGE_TYPE_PROJECT_VERSION ChangeRequestChangeType = 2
)

// Enum value maps for ChangeRequestChangeType.
var (
	ChangeRequestChangeType_name = map[int32]string{
		0: "CHANGE_REQUEST_CHANGE_TYPE_INVALID",
		1: "CHANGE_REQUEST_CHANGE_TYPE_PROJECT_DATA",
		2: "CHANGE_REQUEST_CHANGE_TYPE_PROJECT_VERSION",
	}
	ChangeRequestChangeType_value = map[string]int32{
		"CHANGE_REQUEST_CHANGE_TYPE_INVALID":         0,
		"CHANGE_REQUEST_CHANGE_TYPE_PROJECT_DATA":    1,
		"CHANGE_REQUEST_CHANGE_TYPE_PROJECT_VERSION": 2,
	}
)

func (x ChangeRequestChangeType) Enum() *ChangeRequestChangeType {
	p := new(ChangeRequestChangeType)
	*p = x
	return p
}

func (x ChangeRequestChangeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChangeRequestChangeType) Descriptor() protoreflect.EnumDescriptor {
	return file_change_request_proto_enumTypes[0].Descriptor()
}

func (ChangeRequestChangeType) Type() protoreflect.EnumType {
	return &file_change_request_proto_enumTypes[0]
}

func (x ChangeRequestChangeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChangeRequestChangeType.Descriptor instead.
func (ChangeRequestChangeType) EnumDescriptor() ([]byte, []int) {
	return file_change_request_proto_rawDescGZIP(), []int{0}
}

type ChangeRequestReviewStatus int32

const (
	ChangeRequestReviewStatus_CHANGE_REQUEST_REVIEW_STATUS_INVALID   ChangeRequestReviewStatus = 0
	ChangeRequestReviewStatus_CHANGE_REQUEST_REVIEW_STATUS_DRAFT     ChangeRequestReviewStatus = 1
	ChangeRequestReviewStatus_CHANGE_REQUEST_REVIEW_STATUS_IN_REVIEW ChangeRequestReviewStatus = 2
	ChangeRequestReviewStatus_CHANGE_REQUEST_REVIEW_STATUS_APPROVED  ChangeRequestReviewStatus = 3
	ChangeRequestReviewStatus_CHANGE_REQUEST_REVIEW_STATUS_REJECTED  ChangeRequestReviewStatus = 4
	ChangeRequestReviewStatus_CHANGE_REQUEST_REVIEW_STATUS_DISCARDED ChangeRequestReviewStatus = 5
	ChangeRequestReviewStatus_CHANGE_REQUEST_REVIEW_STATUS_APPLIED   ChangeRequestReviewStatus = 6
)

// Enum value maps for ChangeRequestReviewStatus.
var (
	ChangeRequestReviewStatus_name = map[int32]string{
		0: "CHANGE_REQUEST_REVIEW_STATUS_INVALID",
		1: "CHANGE_REQUEST_REVIEW_STATUS_DRAFT",
		2: "CHANGE_REQUEST_REVIEW_STATUS_IN_REVIEW",
		3: "CHANGE_REQUEST_REVIEW_STATUS_APPROVED",
		4: "CHANGE_REQUEST_REVIEW_STATUS_REJECTED",
		5: "CHANGE_REQUEST_REVIEW_STATUS_DISCARDED",
		6: "CHANGE_REQUEST_REVIEW_STATUS_APPLIED",
	}
	ChangeRequestReviewStatus_value = map[string]int32{
		"CHANGE_REQUEST_REVIEW_STATUS_INVALID":   0,
		"CHANGE_REQUEST_REVIEW_STATUS_DRAFT":     1,
		"CHANGE_REQUEST_REVIEW_STATUS_IN_REVIEW": 2,
		"CHANGE_REQUEST_REVIEW_STATUS_APPROVED":  3,
		"CHANGE_REQUEST_REVIEW_STATUS_REJECTED":  4,
		"CHANGE_REQUEST_REVIEW_STATUS_DISCARDED": 5,
		"CHANGE_REQUEST_REVIEW_STATUS_APPLIED":   6,
	}
)

func (x ChangeRequestReviewStatus) Enum() *ChangeRequestReviewStatus {
	p := new(ChangeRequestReviewStatus)
	*p = x
	return p
}

func (x ChangeRequestReviewStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChangeRequestReviewStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_change_request_proto_enumTypes[1].Descriptor()
}

func (ChangeRequestReviewStatus) Type() protoreflect.EnumType {
	return &file_change_request_proto_enumTypes[1]
}

func (x ChangeRequestReviewStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChangeRequestReviewStatus.Descriptor instead.
func (ChangeRequestReviewStatus) EnumDescriptor() ([]byte, []int) {
	return file_change_request_proto_rawDescGZIP(), []int{1}
}

type ChangeRequestStatus int32

const (
	ChangeRequestStatus_CHANGE_REQUEST_STATUS_INVALID  ChangeRequestStatus = 0
	ChangeRequestStatus_CHANGE_REQUEST_STATUS_ACTIVE   ChangeRequestStatus = 1
	ChangeRequestStatus_CHANGE_REQUEST_STATUS_DISABLED ChangeRequestStatus = 2
)

// Enum value maps for ChangeRequestStatus.
var (
	ChangeRequestStatus_name = map[int32]string{
		0: "CHANGE_REQUEST_STATUS_INVALID",
		1: "CHANGE_REQUEST_STATUS_ACTIVE",
		2: "CHANGE_REQUEST_STATUS_DISABLED",
	}
	ChangeRequestStatus_value = map[string]int32{
		"CHANGE_REQUEST_STATUS_INVALID":  0,
		"CHANGE_REQUEST_STATUS_ACTIVE":   1,
		"CHANGE_REQUEST_STATUS_DISABLED": 2,
	}
)

func (x ChangeRequestStatus) Enum() *ChangeRequestStatus {
	p := new(ChangeRequestStatus)
	*p = x
	return p
}

func (x ChangeRequestStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChangeRequestStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_change_request_proto_enumTypes[2].Descriptor()
}

func (ChangeRequestStatus) Type() protoreflect.EnumType {
	return &file_change_request_proto_enumTypes[2]
}

func (x ChangeRequestStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChangeRequestStatus.Descriptor instead.
func (ChangeRequestStatus) EnumDescriptor() ([]byte, []int) {
	return file_change_request_proto_rawDescGZIP(), []int{2}
}

type ChangeRequest struct {
	state              protoimpl.MessageState     `protogen:"open.v1"`
	Uuid               string                     `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Version            int64                      `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	Title              string                     `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description        string                     `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	ProjectVersionUuid string                     `protobuf:"bytes,5,opt,name=project_version_uuid,json=projectVersionUuid,proto3" json:"project_version_uuid,omitempty"`
	ChangeType         ChangeRequestChangeType    `protobuf:"varint,6,opt,name=change_type,json=changeType,proto3,enum=nem.ChangeRequestChangeType" json:"change_type,omitempty"`
	DataChanges        []*ChangeRequestDataChange `protobuf:"bytes,7,rep,name=data_changes,json=dataChanges,proto3" json:"data_changes,omitempty"`
	VersionChanges     string                     `protobuf:"bytes,8,opt,name=version_changes,json=versionChanges,proto3" json:"version_changes,omitempty"`
	Reviews            []*ChangeRequestReview     `protobuf:"bytes,9,rep,name=reviews,proto3" json:"reviews,omitempty"`
	ReviewStatus       ChangeRequestReviewStatus  `protobuf:"varint,10,opt,name=review_status,json=reviewStatus,proto3,enum=nem.ChangeRequestReviewStatus" json:"review_status,omitempty"`
	OwnerUuid          string                     `protobuf:"bytes,11,opt,name=owner_uuid,json=ownerUuid,proto3" json:"owner_uuid,omitempty"`
	Status             ChangeRequestStatus        `protobuf:"varint,12,opt,name=status,proto3,enum=nem.ChangeRequestStatus" json:"status,omitempty"`
	CreatedAt          *timestamppb.Timestamp     `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt          *timestamppb.Timestamp     `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CreatedByUuid      string                     `protobuf:"bytes,15,opt,name=created_by_uuid,json=createdByUuid,proto3" json:"created_by_uuid,omitempty"`
	UpdatedByUuid      string                     `protobuf:"bytes,16,opt,name=updated_by_uuid,json=updatedByUuid,proto3" json:"updated_by_uuid,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *ChangeRequest) Reset() {
	*x = ChangeRequest{}
	mi := &file_change_request_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeRequest) ProtoMessage() {}

func (x *ChangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_change_request_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeRequest.ProtoReflect.Descriptor instead.
func (*ChangeRequest) Descriptor() ([]byte, []int) {
	return file_change_request_proto_rawDescGZIP(), []int{0}
}

func (x *ChangeRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *ChangeRequest) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *ChangeRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ChangeRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ChangeRequest) GetProjectVersionUuid() string {
	if x != nil {
		return x.ProjectVersionUuid
	}
	return ""
}

func (x *ChangeRequest) GetChangeType() ChangeRequestChangeType {
	if x != nil {
		return x.ChangeType
	}
	return ChangeRequestChangeType_CHANGE_REQUEST_CHANGE_TYPE_INVALID
}

func (x *ChangeRequest) GetDataChanges() []*ChangeRequestDataChange {
	if x != nil {
		return x.DataChanges
	}
	return nil
}

func (x *ChangeRequest) GetVersionChanges() string {
	if x != nil {
		return x.VersionChanges
	}
	return ""
}

func (x *ChangeRequest) GetReviews() []*ChangeRequestReview {
	if x != nil {
		return x.Reviews
	}
	return nil
}

func (x *ChangeRequest) GetReviewStatus() ChangeRequestReviewStatus {
	if x != nil {
		return x.ReviewStatus
	}
	return ChangeRequestReviewStatus_CHANGE_REQUEST_REVIEW_STATUS_INVALID
}

func (x *ChangeRequest) GetOwnerUuid() string {
	if x != nil {
		return x.OwnerUuid
	}
	return ""
}

func (x *ChangeRequest) GetStatus() ChangeRequestStatus {
	if x != nil {
		return x.Status
	}
	return ChangeRequestStatus_CHANGE_REQUEST_STATUS_INVALID
}

func (x *ChangeRequest) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ChangeRequest) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *ChangeRequest) GetCreatedByUuid() string {
	if x != nil {
		return x.CreatedByUuid
	}
	return ""
}

func (x *ChangeRequest) GetUpdatedByUuid() string {
	if x != nil {
		return x.UpdatedByUuid
	}
	return ""
}

var File_change_request_proto protoreflect.FileDescriptor

const file_change_request_proto_rawDesc = "" +
	"\n" +
	"\x14change_request.proto\x12\x03nem\x1a change_request_data_change.proto\x1a\x1bchange_request_review.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xe0\x05\n" +
	"\rChangeRequest\x12\x12\n" +
	"\x04uuid\x18\x01 \x01(\tR\x04uuid\x12\x18\n" +
	"\aversion\x18\x02 \x01(\x03R\aversion\x12\x14\n" +
	"\x05title\x18\x03 \x01(\tR\x05title\x12 \n" +
	"\vdescription\x18\x04 \x01(\tR\vdescription\x120\n" +
	"\x14project_version_uuid\x18\x05 \x01(\tR\x12projectVersionUuid\x12=\n" +
	"\vchange_type\x18\x06 \x01(\x0e2\x1c.nem.ChangeRequestChangeTypeR\n" +
	"changeType\x12?\n" +
	"\fdata_changes\x18\a \x03(\v2\x1c.nem.ChangeRequestDataChangeR\vdataChanges\x12'\n" +
	"\x0fversion_changes\x18\b \x01(\tR\x0eversionChanges\x122\n" +
	"\areviews\x18\t \x03(\v2\x18.nem.ChangeRequestReviewR\areviews\x12C\n" +
	"\rreview_status\x18\n" +
	" \x01(\x0e2\x1e.nem.ChangeRequestReviewStatusR\freviewStatus\x12\x1d\n" +
	"\n" +
	"owner_uuid\x18\v \x01(\tR\townerUuid\x120\n" +
	"\x06status\x18\f \x01(\x0e2\x18.nem.ChangeRequestStatusR\x06status\x129\n" +
	"\n" +
	"created_at\x18\r \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\x0e \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\x12&\n" +
	"\x0fcreated_by_uuid\x18\x0f \x01(\tR\rcreatedByUuid\x12&\n" +
	"\x0fupdated_by_uuid\x18\x10 \x01(\tR\rupdatedByUuid*\x9e\x01\n" +
	"\x17ChangeRequestChangeType\x12&\n" +
	"\"CHANGE_REQUEST_CHANGE_TYPE_INVALID\x10\x00\x12+\n" +
	"'CHANGE_REQUEST_CHANGE_TYPE_PROJECT_DATA\x10\x01\x12.\n" +
	"*CHANGE_REQUEST_CHANGE_TYPE_PROJECT_VERSION\x10\x02*\xc5\x02\n" +
	"\x19ChangeRequestReviewStatus\x12(\n" +
	"$CHANGE_REQUEST_REVIEW_STATUS_INVALID\x10\x00\x12&\n" +
	"\"CHANGE_REQUEST_REVIEW_STATUS_DRAFT\x10\x01\x12*\n" +
	"&CHANGE_REQUEST_REVIEW_STATUS_IN_REVIEW\x10\x02\x12)\n" +
	"%CHANGE_REQUEST_REVIEW_STATUS_APPROVED\x10\x03\x12)\n" +
	"%CHANGE_REQUEST_REVIEW_STATUS_REJECTED\x10\x04\x12*\n" +
	"&CHANGE_REQUEST_REVIEW_STATUS_DISCARDED\x10\x05\x12(\n" +
	"$CHANGE_REQUEST_REVIEW_STATUS_APPLIED\x10\x06*~\n" +
	"\x13ChangeRequestStatus\x12!\n" +
	"\x1dCHANGE_REQUEST_STATUS_INVALID\x10\x00\x12 \n" +
	"\x1cCHANGE_REQUEST_STATUS_ACTIVE\x10\x01\x12\"\n" +
	"\x1eCHANGE_REQUEST_STATUS_DISABLED\x10\x02B4\n" +
	"\x14github.com/nuzur/nemB\rChangeRequestP\x01Z\vnem/idl/genb\x06proto3"

var (
	file_change_request_proto_rawDescOnce sync.Once
	file_change_request_proto_rawDescData []byte
)

func file_change_request_proto_rawDescGZIP() []byte {
	file_change_request_proto_rawDescOnce.Do(func() {
		file_change_request_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_change_request_proto_rawDesc), len(file_change_request_proto_rawDesc)))
	})
	return file_change_request_proto_rawDescData
}

var file_change_request_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_change_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_change_request_proto_goTypes = []any{
	(ChangeRequestChangeType)(0),    // 0: nem.ChangeRequestChangeType
	(ChangeRequestReviewStatus)(0),  // 1: nem.ChangeRequestReviewStatus
	(ChangeRequestStatus)(0),        // 2: nem.ChangeRequestStatus
	(*ChangeRequest)(nil),           // 3: nem.ChangeRequest
	(*ChangeRequestDataChange)(nil), // 4: nem.ChangeRequestDataChange
	(*ChangeRequestReview)(nil),     // 5: nem.ChangeRequestReview
	(*timestamppb.Timestamp)(nil),   // 6: google.protobuf.Timestamp
}
var file_change_request_proto_depIdxs = []int32{
	0, // 0: nem.ChangeRequest.change_type:type_name -> nem.ChangeRequestChangeType
	4, // 1: nem.ChangeRequest.data_changes:type_name -> nem.ChangeRequestDataChange
	5, // 2: nem.ChangeRequest.reviews:type_name -> nem.ChangeRequestReview
	1, // 3: nem.ChangeRequest.review_status:type_name -> nem.ChangeRequestReviewStatus
	2, // 4: nem.ChangeRequest.status:type_name -> nem.ChangeRequestStatus
	6, // 5: nem.ChangeRequest.created_at:type_name -> google.protobuf.Timestamp
	6, // 6: nem.ChangeRequest.updated_at:type_name -> google.protobuf.Timestamp
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_change_request_proto_init() }
func file_change_request_proto_init() {
	if File_change_request_proto != nil {
		return
	}
	file_change_request_data_change_proto_init()
	file_change_request_review_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_change_request_proto_rawDesc), len(file_change_request_proto_rawDesc)),
			NumEnums:      3,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_change_request_proto_goTypes,
		DependencyIndexes: file_change_request_proto_depIdxs,
		EnumInfos:         file_change_request_proto_enumTypes,
		MessageInfos:      file_change_request_proto_msgTypes,
	}.Build()
	File_change_request_proto = out.File
	file_change_request_proto_goTypes = nil
	file_change_request_proto_depIdxs = nil
}
