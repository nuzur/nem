// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: field_type_float_config.proto

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

type FieldTypeFloatConfigSeparator int32

const (
	FieldTypeFloatConfigSeparator_FIELD_TYPE_FLOAT_CONFIG_SEPARATOR_INVALID FieldTypeFloatConfigSeparator = 0
	FieldTypeFloatConfigSeparator_FIELD_TYPE_FLOAT_CONFIG_SEPARATOR_POINT   FieldTypeFloatConfigSeparator = 1
	FieldTypeFloatConfigSeparator_FIELD_TYPE_FLOAT_CONFIG_SEPARATOR_COMMA   FieldTypeFloatConfigSeparator = 2
)

// Enum value maps for FieldTypeFloatConfigSeparator.
var (
	FieldTypeFloatConfigSeparator_name = map[int32]string{
		0: "FIELD_TYPE_FLOAT_CONFIG_SEPARATOR_INVALID",
		1: "FIELD_TYPE_FLOAT_CONFIG_SEPARATOR_POINT",
		2: "FIELD_TYPE_FLOAT_CONFIG_SEPARATOR_COMMA",
	}
	FieldTypeFloatConfigSeparator_value = map[string]int32{
		"FIELD_TYPE_FLOAT_CONFIG_SEPARATOR_INVALID": 0,
		"FIELD_TYPE_FLOAT_CONFIG_SEPARATOR_POINT":   1,
		"FIELD_TYPE_FLOAT_CONFIG_SEPARATOR_COMMA":   2,
	}
)

func (x FieldTypeFloatConfigSeparator) Enum() *FieldTypeFloatConfigSeparator {
	p := new(FieldTypeFloatConfigSeparator)
	*p = x
	return p
}

func (x FieldTypeFloatConfigSeparator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FieldTypeFloatConfigSeparator) Descriptor() protoreflect.EnumDescriptor {
	return file_field_type_float_config_proto_enumTypes[0].Descriptor()
}

func (FieldTypeFloatConfigSeparator) Type() protoreflect.EnumType {
	return &file_field_type_float_config_proto_enumTypes[0]
}

func (x FieldTypeFloatConfigSeparator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FieldTypeFloatConfigSeparator.Descriptor instead.
func (FieldTypeFloatConfigSeparator) EnumDescriptor() ([]byte, []int) {
	return file_field_type_float_config_proto_rawDescGZIP(), []int{0}
}

type FieldTypeFloatConfig struct {
	state             protoimpl.MessageState        `protogen:"open.v1"`
	MinValue          float64                       `protobuf:"fixed64,1,opt,name=min_value,json=minValue,proto3" json:"min_value,omitempty"`
	MinValueInclusive bool                          `protobuf:"varint,2,opt,name=min_value_inclusive,json=minValueInclusive,proto3" json:"min_value_inclusive,omitempty"`
	MaxValue          float64                       `protobuf:"fixed64,3,opt,name=max_value,json=maxValue,proto3" json:"max_value,omitempty"`
	MaxValueInclusive bool                          `protobuf:"varint,4,opt,name=max_value_inclusive,json=maxValueInclusive,proto3" json:"max_value_inclusive,omitempty"`
	AllowNegatives    bool                          `protobuf:"varint,5,opt,name=allow_negatives,json=allowNegatives,proto3" json:"allow_negatives,omitempty"`
	NumberOfDecimals  int64                         `protobuf:"varint,6,opt,name=number_of_decimals,json=numberOfDecimals,proto3" json:"number_of_decimals,omitempty"`
	Separator         FieldTypeFloatConfigSeparator `protobuf:"varint,7,opt,name=separator,proto3,enum=nem.FieldTypeFloatConfigSeparator" json:"separator,omitempty"`
	EnableLimits      bool                          `protobuf:"varint,8,opt,name=enable_limits,json=enableLimits,proto3" json:"enable_limits,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *FieldTypeFloatConfig) Reset() {
	*x = FieldTypeFloatConfig{}
	mi := &file_field_type_float_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FieldTypeFloatConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldTypeFloatConfig) ProtoMessage() {}

func (x *FieldTypeFloatConfig) ProtoReflect() protoreflect.Message {
	mi := &file_field_type_float_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldTypeFloatConfig.ProtoReflect.Descriptor instead.
func (*FieldTypeFloatConfig) Descriptor() ([]byte, []int) {
	return file_field_type_float_config_proto_rawDescGZIP(), []int{0}
}

func (x *FieldTypeFloatConfig) GetMinValue() float64 {
	if x != nil {
		return x.MinValue
	}
	return 0
}

func (x *FieldTypeFloatConfig) GetMinValueInclusive() bool {
	if x != nil {
		return x.MinValueInclusive
	}
	return false
}

func (x *FieldTypeFloatConfig) GetMaxValue() float64 {
	if x != nil {
		return x.MaxValue
	}
	return 0
}

func (x *FieldTypeFloatConfig) GetMaxValueInclusive() bool {
	if x != nil {
		return x.MaxValueInclusive
	}
	return false
}

func (x *FieldTypeFloatConfig) GetAllowNegatives() bool {
	if x != nil {
		return x.AllowNegatives
	}
	return false
}

func (x *FieldTypeFloatConfig) GetNumberOfDecimals() int64 {
	if x != nil {
		return x.NumberOfDecimals
	}
	return 0
}

func (x *FieldTypeFloatConfig) GetSeparator() FieldTypeFloatConfigSeparator {
	if x != nil {
		return x.Separator
	}
	return FieldTypeFloatConfigSeparator_FIELD_TYPE_FLOAT_CONFIG_SEPARATOR_INVALID
}

func (x *FieldTypeFloatConfig) GetEnableLimits() bool {
	if x != nil {
		return x.EnableLimits
	}
	return false
}

var File_field_type_float_config_proto protoreflect.FileDescriptor

const file_field_type_float_config_proto_rawDesc = "" +
	"\n" +
	"\x1dfield_type_float_config.proto\x12\x03nem\"\xee\x02\n" +
	"\x14FieldTypeFloatConfig\x12\x1b\n" +
	"\tmin_value\x18\x01 \x01(\x01R\bminValue\x12.\n" +
	"\x13min_value_inclusive\x18\x02 \x01(\bR\x11minValueInclusive\x12\x1b\n" +
	"\tmax_value\x18\x03 \x01(\x01R\bmaxValue\x12.\n" +
	"\x13max_value_inclusive\x18\x04 \x01(\bR\x11maxValueInclusive\x12'\n" +
	"\x0fallow_negatives\x18\x05 \x01(\bR\x0eallowNegatives\x12,\n" +
	"\x12number_of_decimals\x18\x06 \x01(\x03R\x10numberOfDecimals\x12@\n" +
	"\tseparator\x18\a \x01(\x0e2\".nem.FieldTypeFloatConfigSeparatorR\tseparator\x12#\n" +
	"\renable_limits\x18\b \x01(\bR\fenableLimits*\xa8\x01\n" +
	"\x1dFieldTypeFloatConfigSeparator\x12-\n" +
	")FIELD_TYPE_FLOAT_CONFIG_SEPARATOR_INVALID\x10\x00\x12+\n" +
	"'FIELD_TYPE_FLOAT_CONFIG_SEPARATOR_POINT\x10\x01\x12+\n" +
	"'FIELD_TYPE_FLOAT_CONFIG_SEPARATOR_COMMA\x10\x02B;\n" +
	"\x14github.com/nuzur/nemB\x14FieldTypeFloatConfigP\x01Z\vnem/idl/genb\x06proto3"

var (
	file_field_type_float_config_proto_rawDescOnce sync.Once
	file_field_type_float_config_proto_rawDescData []byte
)

func file_field_type_float_config_proto_rawDescGZIP() []byte {
	file_field_type_float_config_proto_rawDescOnce.Do(func() {
		file_field_type_float_config_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_field_type_float_config_proto_rawDesc), len(file_field_type_float_config_proto_rawDesc)))
	})
	return file_field_type_float_config_proto_rawDescData
}

var file_field_type_float_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_field_type_float_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_field_type_float_config_proto_goTypes = []any{
	(FieldTypeFloatConfigSeparator)(0), // 0: nem.FieldTypeFloatConfigSeparator
	(*FieldTypeFloatConfig)(nil),       // 1: nem.FieldTypeFloatConfig
}
var file_field_type_float_config_proto_depIdxs = []int32{
	0, // 0: nem.FieldTypeFloatConfig.separator:type_name -> nem.FieldTypeFloatConfigSeparator
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_field_type_float_config_proto_init() }
func file_field_type_float_config_proto_init() {
	if File_field_type_float_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_field_type_float_config_proto_rawDesc), len(file_field_type_float_config_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_field_type_float_config_proto_goTypes,
		DependencyIndexes: file_field_type_float_config_proto_depIdxs,
		EnumInfos:         file_field_type_float_config_proto_enumTypes,
		MessageInfos:      file_field_type_float_config_proto_msgTypes,
	}.Build()
	File_field_type_float_config_proto = out.File
	file_field_type_float_config_proto_goTypes = nil
	file_field_type_float_config_proto_depIdxs = nil
}
