package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/field_type_config"
	pb "github.com/nuzur/nem/idl/gen"
)

func FieldTypeConfigToProto(e main_entity.FieldTypeConfig) *pb.FieldTypeConfig {
	return &pb.FieldTypeConfig{
		Uuid:      StringPtrToString(e.UUID),
		Integer:   FieldTypeIntegerConfigToProto(e.Integer),
		Float:     FieldTypeFloatConfigToProto(e.Float),
		Decimal:   FieldTypeDecimalConfigToProto(e.Decimal),
		Boolean:   StringPtrToString(e.Boolean),
		Char:      FieldTypeCharConfigToProto(e.Char),
		Varchar:   FieldTypeVarcharConfigToProto(e.Varchar),
		Text:      FieldTypeTextConfigToProto(e.Text),
		Encrypted: FieldTypeEncryptedConfigToProto(e.Encrypted),
		Email:     FieldTypeEmailConfigToProto(e.Email),
		Phone:     FieldTypePhoneConfigToProto(e.Phone),
		Url:       FieldTypeURLConfigToProto(e.URL),
		Location:  StringPtrToString(e.Location),
		Color:     StringPtrToString(e.Color),
		RichText:  StringPtrToString(e.RichText),
		Code:      StringPtrToString(e.Code),
		Markdown:  StringPtrToString(e.Markdown),
		File:      FieldTypeFileConfigToProto(e.File),
		Image:     FieldTypeFileConfigToProto(e.Image),
		Video:     FieldTypeFileConfigToProto(e.Video),
		Audio:     FieldTypeFileConfigToProto(e.Audio),
		Enum:      FieldTypeEnumConfigToProto(e.Enum),
		Json:      FieldTypeJSONConfigToProto(e.JSON),
		Array:     FieldTypeArrayConfigToProto(e.Array),
		Date:      FieldTypeDateConfigToProto(e.Date),
		Datetime:  FieldTypeDatetimeConfigToProto(e.Datetime),
		Time:      StringPtrToString(e.Time),
		Slug:      FieldTypeSlugConfigToProto(e.Slug),
	}
}

func FieldTypeConfigSliceToProto(es []main_entity.FieldTypeConfig) []*pb.FieldTypeConfig {
	res := []*pb.FieldTypeConfig{}
	for _, e := range es {
		res = append(res, FieldTypeConfigToProto(e))
	}
	return res
}

func FieldTypeConfigFromProto(m *pb.FieldTypeConfig) main_entity.FieldTypeConfig {
	if m == nil {
		return main_entity.FieldTypeConfig{}
	}
	return main_entity.FieldTypeConfig{
		UUID:      &m.Uuid,
		Integer:   FieldTypeIntegerConfigFromProto(m.GetInteger()),
		Float:     FieldTypeFloatConfigFromProto(m.GetFloat()),
		Decimal:   FieldTypeDecimalConfigFromProto(m.GetDecimal()),
		Boolean:   &m.Boolean,
		Char:      FieldTypeCharConfigFromProto(m.GetChar()),
		Varchar:   FieldTypeVarcharConfigFromProto(m.GetVarchar()),
		Text:      FieldTypeTextConfigFromProto(m.GetText()),
		Encrypted: FieldTypeEncryptedConfigFromProto(m.GetEncrypted()),
		Email:     FieldTypeEmailConfigFromProto(m.GetEmail()),
		Phone:     FieldTypePhoneConfigFromProto(m.GetPhone()),
		URL:       FieldTypeURLConfigFromProto(m.GetUrl()),
		Location:  &m.Location,
		Color:     &m.Color,
		RichText:  &m.RichText,
		Code:      &m.Code,
		Markdown:  &m.Markdown,
		File:      FieldTypeFileConfigFromProto(m.GetFile()),
		Image:     FieldTypeFileConfigFromProto(m.GetImage()),
		Video:     FieldTypeFileConfigFromProto(m.GetVideo()),
		Audio:     FieldTypeFileConfigFromProto(m.GetAudio()),
		Enum:      FieldTypeEnumConfigFromProto(m.GetEnum()),
		JSON:      FieldTypeJSONConfigFromProto(m.GetJson()),
		Array:     FieldTypeArrayConfigFromProto(m.GetArray()),
		Date:      FieldTypeDateConfigFromProto(m.GetDate()),
		Datetime:  FieldTypeDatetimeConfigFromProto(m.GetDatetime()),
		Time:      &m.Time,
		Slug:      FieldTypeSlugConfigFromProto(m.GetSlug()),
	}
}

func FieldTypeConfigSliceFromProto(es []*pb.FieldTypeConfig) []main_entity.FieldTypeConfig {
	if es == nil {
		return []main_entity.FieldTypeConfig{}
	}
	res := []main_entity.FieldTypeConfig{}
	for _, e := range es {
		res = append(res, FieldTypeConfigFromProto(e))
	}
	return res
}
