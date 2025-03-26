// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: field_type_encrypted_config.proto

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

type FieldTypeEncryptedConfigEncryptionType int32

const (
	FieldTypeEncryptedConfigEncryptionType_FIELD_TYPE_ENCRYPTED_CONFIG_ENCRYPTION_TYPE_INVALID FieldTypeEncryptedConfigEncryptionType = 0
	FieldTypeEncryptedConfigEncryptionType_FIELD_TYPE_ENCRYPTED_CONFIG_ENCRYPTION_TYPE_BCRYPT  FieldTypeEncryptedConfigEncryptionType = 1
	FieldTypeEncryptedConfigEncryptionType_FIELD_TYPE_ENCRYPTED_CONFIG_ENCRYPTION_TYPE_SHA_512 FieldTypeEncryptedConfigEncryptionType = 2
)

// Enum value maps for FieldTypeEncryptedConfigEncryptionType.
var (
	FieldTypeEncryptedConfigEncryptionType_name = map[int32]string{
		0: "FIELD_TYPE_ENCRYPTED_CONFIG_ENCRYPTION_TYPE_INVALID",
		1: "FIELD_TYPE_ENCRYPTED_CONFIG_ENCRYPTION_TYPE_BCRYPT",
		2: "FIELD_TYPE_ENCRYPTED_CONFIG_ENCRYPTION_TYPE_SHA_512",
	}
	FieldTypeEncryptedConfigEncryptionType_value = map[string]int32{
		"FIELD_TYPE_ENCRYPTED_CONFIG_ENCRYPTION_TYPE_INVALID": 0,
		"FIELD_TYPE_ENCRYPTED_CONFIG_ENCRYPTION_TYPE_BCRYPT":  1,
		"FIELD_TYPE_ENCRYPTED_CONFIG_ENCRYPTION_TYPE_SHA_512": 2,
	}
)

func (x FieldTypeEncryptedConfigEncryptionType) Enum() *FieldTypeEncryptedConfigEncryptionType {
	p := new(FieldTypeEncryptedConfigEncryptionType)
	*p = x
	return p
}

func (x FieldTypeEncryptedConfigEncryptionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FieldTypeEncryptedConfigEncryptionType) Descriptor() protoreflect.EnumDescriptor {
	return file_field_type_encrypted_config_proto_enumTypes[0].Descriptor()
}

func (FieldTypeEncryptedConfigEncryptionType) Type() protoreflect.EnumType {
	return &file_field_type_encrypted_config_proto_enumTypes[0]
}

func (x FieldTypeEncryptedConfigEncryptionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FieldTypeEncryptedConfigEncryptionType.Descriptor instead.
func (FieldTypeEncryptedConfigEncryptionType) EnumDescriptor() ([]byte, []int) {
	return file_field_type_encrypted_config_proto_rawDescGZIP(), []int{0}
}

type FieldTypeEncryptedConfig struct {
	state           protoimpl.MessageState                 `protogen:"open.v1"`
	MinSize         int64                                  `protobuf:"varint,1,opt,name=min_size,json=minSize,proto3" json:"min_size,omitempty"`
	MaxSize         int64                                  `protobuf:"varint,2,opt,name=max_size,json=maxSize,proto3" json:"max_size,omitempty"`
	RegexValidation string                                 `protobuf:"bytes,3,opt,name=regex_validation,json=regexValidation,proto3" json:"regex_validation,omitempty"`
	EncryptionType  FieldTypeEncryptedConfigEncryptionType `protobuf:"varint,4,opt,name=encryption_type,json=encryptionType,proto3,enum=nem.FieldTypeEncryptedConfigEncryptionType" json:"encryption_type,omitempty"`
	UseSalt         bool                                   `protobuf:"varint,5,opt,name=use_salt,json=useSalt,proto3" json:"use_salt,omitempty"`
	SaltFieldUuids  []string                               `protobuf:"bytes,6,rep,name=salt_field_uuids,json=saltFieldUuids,proto3" json:"salt_field_uuids,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *FieldTypeEncryptedConfig) Reset() {
	*x = FieldTypeEncryptedConfig{}
	mi := &file_field_type_encrypted_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FieldTypeEncryptedConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldTypeEncryptedConfig) ProtoMessage() {}

func (x *FieldTypeEncryptedConfig) ProtoReflect() protoreflect.Message {
	mi := &file_field_type_encrypted_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldTypeEncryptedConfig.ProtoReflect.Descriptor instead.
func (*FieldTypeEncryptedConfig) Descriptor() ([]byte, []int) {
	return file_field_type_encrypted_config_proto_rawDescGZIP(), []int{0}
}

func (x *FieldTypeEncryptedConfig) GetMinSize() int64 {
	if x != nil {
		return x.MinSize
	}
	return 0
}

func (x *FieldTypeEncryptedConfig) GetMaxSize() int64 {
	if x != nil {
		return x.MaxSize
	}
	return 0
}

func (x *FieldTypeEncryptedConfig) GetRegexValidation() string {
	if x != nil {
		return x.RegexValidation
	}
	return ""
}

func (x *FieldTypeEncryptedConfig) GetEncryptionType() FieldTypeEncryptedConfigEncryptionType {
	if x != nil {
		return x.EncryptionType
	}
	return FieldTypeEncryptedConfigEncryptionType_FIELD_TYPE_ENCRYPTED_CONFIG_ENCRYPTION_TYPE_INVALID
}

func (x *FieldTypeEncryptedConfig) GetUseSalt() bool {
	if x != nil {
		return x.UseSalt
	}
	return false
}

func (x *FieldTypeEncryptedConfig) GetSaltFieldUuids() []string {
	if x != nil {
		return x.SaltFieldUuids
	}
	return nil
}

var File_field_type_encrypted_config_proto protoreflect.FileDescriptor

const file_field_type_encrypted_config_proto_rawDesc = "" +
	"\n" +
	"!field_type_encrypted_config.proto\x12\x03nem\"\x96\x02\n" +
	"\x18FieldTypeEncryptedConfig\x12\x19\n" +
	"\bmin_size\x18\x01 \x01(\x03R\aminSize\x12\x19\n" +
	"\bmax_size\x18\x02 \x01(\x03R\amaxSize\x12)\n" +
	"\x10regex_validation\x18\x03 \x01(\tR\x0fregexValidation\x12T\n" +
	"\x0fencryption_type\x18\x04 \x01(\x0e2+.nem.FieldTypeEncryptedConfigEncryptionTypeR\x0eencryptionType\x12\x19\n" +
	"\buse_salt\x18\x05 \x01(\bR\auseSalt\x12(\n" +
	"\x10salt_field_uuids\x18\x06 \x03(\tR\x0esaltFieldUuids*\xd2\x01\n" +
	"&FieldTypeEncryptedConfigEncryptionType\x127\n" +
	"3FIELD_TYPE_ENCRYPTED_CONFIG_ENCRYPTION_TYPE_INVALID\x10\x00\x126\n" +
	"2FIELD_TYPE_ENCRYPTED_CONFIG_ENCRYPTION_TYPE_BCRYPT\x10\x01\x127\n" +
	"3FIELD_TYPE_ENCRYPTED_CONFIG_ENCRYPTION_TYPE_SHA_512\x10\x02B?\n" +
	"\x14github.com/nuzur/nemB\x18FieldTypeEncryptedConfigP\x01Z\vnem/idl/genb\x06proto3"

var (
	file_field_type_encrypted_config_proto_rawDescOnce sync.Once
	file_field_type_encrypted_config_proto_rawDescData []byte
)

func file_field_type_encrypted_config_proto_rawDescGZIP() []byte {
	file_field_type_encrypted_config_proto_rawDescOnce.Do(func() {
		file_field_type_encrypted_config_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_field_type_encrypted_config_proto_rawDesc), len(file_field_type_encrypted_config_proto_rawDesc)))
	})
	return file_field_type_encrypted_config_proto_rawDescData
}

var file_field_type_encrypted_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_field_type_encrypted_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_field_type_encrypted_config_proto_goTypes = []any{
	(FieldTypeEncryptedConfigEncryptionType)(0), // 0: nem.FieldTypeEncryptedConfigEncryptionType
	(*FieldTypeEncryptedConfig)(nil),            // 1: nem.FieldTypeEncryptedConfig
}
var file_field_type_encrypted_config_proto_depIdxs = []int32{
	0, // 0: nem.FieldTypeEncryptedConfig.encryption_type:type_name -> nem.FieldTypeEncryptedConfigEncryptionType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_field_type_encrypted_config_proto_init() }
func file_field_type_encrypted_config_proto_init() {
	if File_field_type_encrypted_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_field_type_encrypted_config_proto_rawDesc), len(file_field_type_encrypted_config_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_field_type_encrypted_config_proto_goTypes,
		DependencyIndexes: file_field_type_encrypted_config_proto_depIdxs,
		EnumInfos:         file_field_type_encrypted_config_proto_enumTypes,
		MessageInfos:      file_field_type_encrypted_config_proto_msgTypes,
	}.Build()
	File_field_type_encrypted_config_proto = out.File
	file_field_type_encrypted_config_proto_goTypes = nil
	file_field_type_encrypted_config_proto_depIdxs = nil
}
