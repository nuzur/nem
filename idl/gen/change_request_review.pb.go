// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: change_request_review.proto

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

type ChangeRequestReviewStatus int32

const (
	ChangeRequestReviewStatus_CHANGE_REQUEST_REVIEW_STATUS_INVALID   ChangeRequestReviewStatus = 0
	ChangeRequestReviewStatus_CHANGE_REQUEST_REVIEW_STATUS_DRAFT     ChangeRequestReviewStatus = 1
	ChangeRequestReviewStatus_CHANGE_REQUEST_REVIEW_STATUS_IN_REVIEW ChangeRequestReviewStatus = 2
	ChangeRequestReviewStatus_CHANGE_REQUEST_REVIEW_STATUS_APPROVED  ChangeRequestReviewStatus = 3
	ChangeRequestReviewStatus_CHANGE_REQUEST_REVIEW_STATUS_REJECTED  ChangeRequestReviewStatus = 4
	ChangeRequestReviewStatus_CHANGE_REQUEST_REVIEW_STATUS_PUBLISHED ChangeRequestReviewStatus = 5
)

// Enum value maps for ChangeRequestReviewStatus.
var (
	ChangeRequestReviewStatus_name = map[int32]string{
		0: "CHANGE_REQUEST_REVIEW_STATUS_INVALID",
		1: "CHANGE_REQUEST_REVIEW_STATUS_DRAFT",
		2: "CHANGE_REQUEST_REVIEW_STATUS_IN_REVIEW",
		3: "CHANGE_REQUEST_REVIEW_STATUS_APPROVED",
		4: "CHANGE_REQUEST_REVIEW_STATUS_REJECTED",
		5: "CHANGE_REQUEST_REVIEW_STATUS_PUBLISHED",
	}
	ChangeRequestReviewStatus_value = map[string]int32{
		"CHANGE_REQUEST_REVIEW_STATUS_INVALID":   0,
		"CHANGE_REQUEST_REVIEW_STATUS_DRAFT":     1,
		"CHANGE_REQUEST_REVIEW_STATUS_IN_REVIEW": 2,
		"CHANGE_REQUEST_REVIEW_STATUS_APPROVED":  3,
		"CHANGE_REQUEST_REVIEW_STATUS_REJECTED":  4,
		"CHANGE_REQUEST_REVIEW_STATUS_PUBLISHED": 5,
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
	return file_change_request_review_proto_enumTypes[0].Descriptor()
}

func (ChangeRequestReviewStatus) Type() protoreflect.EnumType {
	return &file_change_request_review_proto_enumTypes[0]
}

func (x ChangeRequestReviewStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChangeRequestReviewStatus.Descriptor instead.
func (ChangeRequestReviewStatus) EnumDescriptor() ([]byte, []int) {
	return file_change_request_review_proto_rawDescGZIP(), []int{0}
}

type ChangeRequestReview struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	Uuid          string                    `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	UserUuid      string                    `protobuf:"bytes,2,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	Comment       string                    `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	Status        ChangeRequestReviewStatus `protobuf:"varint,4,opt,name=status,proto3,enum=nem.ChangeRequestReviewStatus" json:"status,omitempty"`
	CreatedAt     *timestamppb.Timestamp    `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp    `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChangeRequestReview) Reset() {
	*x = ChangeRequestReview{}
	mi := &file_change_request_review_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChangeRequestReview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeRequestReview) ProtoMessage() {}

func (x *ChangeRequestReview) ProtoReflect() protoreflect.Message {
	mi := &file_change_request_review_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeRequestReview.ProtoReflect.Descriptor instead.
func (*ChangeRequestReview) Descriptor() ([]byte, []int) {
	return file_change_request_review_proto_rawDescGZIP(), []int{0}
}

func (x *ChangeRequestReview) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *ChangeRequestReview) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *ChangeRequestReview) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *ChangeRequestReview) GetStatus() ChangeRequestReviewStatus {
	if x != nil {
		return x.Status
	}
	return ChangeRequestReviewStatus_CHANGE_REQUEST_REVIEW_STATUS_INVALID
}

func (x *ChangeRequestReview) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ChangeRequestReview) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_change_request_review_proto protoreflect.FileDescriptor

const file_change_request_review_proto_rawDesc = "" +
	"\n" +
	"\x1bchange_request_review.proto\x12\x03nem\x1a\x1fgoogle/protobuf/timestamp.proto\"\x8e\x02\n" +
	"\x13ChangeRequestReview\x12\x12\n" +
	"\x04uuid\x18\x01 \x01(\tR\x04uuid\x12\x1b\n" +
	"\tuser_uuid\x18\x02 \x01(\tR\buserUuid\x12\x18\n" +
	"\acomment\x18\x03 \x01(\tR\acomment\x126\n" +
	"\x06status\x18\x04 \x01(\x0e2\x1e.nem.ChangeRequestReviewStatusR\x06status\x129\n" +
	"\n" +
	"created_at\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt*\x9b\x02\n" +
	"\x19ChangeRequestReviewStatus\x12(\n" +
	"$CHANGE_REQUEST_REVIEW_STATUS_INVALID\x10\x00\x12&\n" +
	"\"CHANGE_REQUEST_REVIEW_STATUS_DRAFT\x10\x01\x12*\n" +
	"&CHANGE_REQUEST_REVIEW_STATUS_IN_REVIEW\x10\x02\x12)\n" +
	"%CHANGE_REQUEST_REVIEW_STATUS_APPROVED\x10\x03\x12)\n" +
	"%CHANGE_REQUEST_REVIEW_STATUS_REJECTED\x10\x04\x12*\n" +
	"&CHANGE_REQUEST_REVIEW_STATUS_PUBLISHED\x10\x05B:\n" +
	"\x14github.com/nuzur/nemB\x13ChangeRequestReviewP\x01Z\vnem/idl/genb\x06proto3"

var (
	file_change_request_review_proto_rawDescOnce sync.Once
	file_change_request_review_proto_rawDescData []byte
)

func file_change_request_review_proto_rawDescGZIP() []byte {
	file_change_request_review_proto_rawDescOnce.Do(func() {
		file_change_request_review_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_change_request_review_proto_rawDesc), len(file_change_request_review_proto_rawDesc)))
	})
	return file_change_request_review_proto_rawDescData
}

var file_change_request_review_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_change_request_review_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_change_request_review_proto_goTypes = []any{
	(ChangeRequestReviewStatus)(0), // 0: nem.ChangeRequestReviewStatus
	(*ChangeRequestReview)(nil),    // 1: nem.ChangeRequestReview
	(*timestamppb.Timestamp)(nil),  // 2: google.protobuf.Timestamp
}
var file_change_request_review_proto_depIdxs = []int32{
	0, // 0: nem.ChangeRequestReview.status:type_name -> nem.ChangeRequestReviewStatus
	2, // 1: nem.ChangeRequestReview.created_at:type_name -> google.protobuf.Timestamp
	2, // 2: nem.ChangeRequestReview.updated_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_change_request_review_proto_init() }
func file_change_request_review_proto_init() {
	if File_change_request_review_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_change_request_review_proto_rawDesc), len(file_change_request_review_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_change_request_review_proto_goTypes,
		DependencyIndexes: file_change_request_review_proto_depIdxs,
		EnumInfos:         file_change_request_review_proto_enumTypes,
		MessageInfos:      file_change_request_review_proto_msgTypes,
	}.Build()
	File_change_request_review_proto = out.File
	file_change_request_review_proto_goTypes = nil
	file_change_request_review_proto_depIdxs = nil
}
