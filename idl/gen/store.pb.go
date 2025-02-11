// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: store.proto

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

type StoreStatus int32

const (
	StoreStatus_STORE_STATUS_INVALID  StoreStatus = 0
	StoreStatus_STORE_STATUS_ACTIVE   StoreStatus = 1
	StoreStatus_STORE_STATUS_DISABLED StoreStatus = 2
)

// Enum value maps for StoreStatus.
var (
	StoreStatus_name = map[int32]string{
		0: "STORE_STATUS_INVALID",
		1: "STORE_STATUS_ACTIVE",
		2: "STORE_STATUS_DISABLED",
	}
	StoreStatus_value = map[string]int32{
		"STORE_STATUS_INVALID":  0,
		"STORE_STATUS_ACTIVE":   1,
		"STORE_STATUS_DISABLED": 2,
	}
)

func (x StoreStatus) Enum() *StoreStatus {
	p := new(StoreStatus)
	*p = x
	return p
}

func (x StoreStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StoreStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_store_proto_enumTypes[0].Descriptor()
}

func (StoreStatus) Type() protoreflect.EnumType {
	return &file_store_proto_enumTypes[0]
}

func (x StoreStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StoreStatus.Descriptor instead.
func (StoreStatus) EnumDescriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{0}
}

type Store struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uuid          string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Identifier    string                 `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Status        StoreStatus            `protobuf:"varint,4,opt,name=status,proto3,enum=nem.StoreStatus" json:"status,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CreatedByUuid string                 `protobuf:"bytes,7,opt,name=created_by_uuid,json=createdByUuid,proto3" json:"created_by_uuid,omitempty"`
	UpdatedByUuid string                 `protobuf:"bytes,8,opt,name=updated_by_uuid,json=updatedByUuid,proto3" json:"updated_by_uuid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Store) Reset() {
	*x = Store{}
	mi := &file_store_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Store) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Store) ProtoMessage() {}

func (x *Store) ProtoReflect() protoreflect.Message {
	mi := &file_store_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Store.ProtoReflect.Descriptor instead.
func (*Store) Descriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{0}
}

func (x *Store) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Store) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *Store) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Store) GetStatus() StoreStatus {
	if x != nil {
		return x.Status
	}
	return StoreStatus_STORE_STATUS_INVALID
}

func (x *Store) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Store) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Store) GetCreatedByUuid() string {
	if x != nil {
		return x.CreatedByUuid
	}
	return ""
}

func (x *Store) GetUpdatedByUuid() string {
	if x != nil {
		return x.UpdatedByUuid
	}
	return ""
}

var File_store_proto protoreflect.FileDescriptor

var file_store_proto_rawDesc = string([]byte{
	0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6e,
	0x65, 0x6d, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x02, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62,
	0x79, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x55, 0x75, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x55,
	0x75, 0x69, 0x64, 0x2a, 0x5b, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13,
	0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x43, 0x54,
	0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x02,
	0x42, 0x1b, 0x0a, 0x03, 0x6e, 0x65, 0x6d, 0x42, 0x05, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x01,
	0x5a, 0x0b, 0x6e, 0x65, 0x6d, 0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_store_proto_rawDescOnce sync.Once
	file_store_proto_rawDescData []byte
)

func file_store_proto_rawDescGZIP() []byte {
	file_store_proto_rawDescOnce.Do(func() {
		file_store_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_store_proto_rawDesc), len(file_store_proto_rawDesc)))
	})
	return file_store_proto_rawDescData
}

var file_store_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_store_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_store_proto_goTypes = []any{
	(StoreStatus)(0),              // 0: nem.StoreStatus
	(*Store)(nil),                 // 1: nem.Store
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_store_proto_depIdxs = []int32{
	0, // 0: nem.Store.status:type_name -> nem.StoreStatus
	2, // 1: nem.Store.created_at:type_name -> google.protobuf.Timestamp
	2, // 2: nem.Store.updated_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_store_proto_init() }
func file_store_proto_init() {
	if File_store_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_store_proto_rawDesc), len(file_store_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_proto_goTypes,
		DependencyIndexes: file_store_proto_depIdxs,
		EnumInfos:         file_store_proto_enumTypes,
		MessageInfos:      file_store_proto_msgTypes,
	}.Build()
	File_store_proto = out.File
	file_store_proto_goTypes = nil
	file_store_proto_depIdxs = nil
}
