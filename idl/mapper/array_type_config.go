package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/array_type_config"
	pb "github.com/nuzur/nem/idl/gen"
)

func ArrayTypeConfigToProto(e main_entity.ArrayTypeConfig) *pb.ArrayTypeConfig {
	return &pb.ArrayTypeConfig{
		Integer:   FieldTypeIntegerConfigToProto(e.Integer),
		Float:     FieldTypeFloatConfigToProto(e.Float),
		Decimal:   FieldTypeDecimalConfigToProto(e.Decimal),
		Char:      FieldTypeCharConfigToProto(e.Char),
		Varchar:   FieldTypeVarcharConfigToProto(e.Varchar),
		Encrypted: FieldTypeEncryptedConfigToProto(e.Encrypted),
		Url:       FieldTypeURLConfigToProto(e.URL),
		Email:     FieldTypeEmailConfigToProto(e.Email),
		Phone:     FieldTypePhoneConfigToProto(e.Phone),
		Date:      FieldTypeDateConfigToProto(e.Date),
		Datetime:  FieldTypeDatetimeConfigToProto(e.Datetime),
		Enum:      FieldTypeEnumConfigToProto(e.Enum),
	}
}

func ArrayTypeConfigSliceToProto(es []main_entity.ArrayTypeConfig) []*pb.ArrayTypeConfig {
	res := []*pb.ArrayTypeConfig{}
	for _, e := range es {
		res = append(res, ArrayTypeConfigToProto(e))
	}
	return res
}

func ArrayTypeConfigFromProto(m *pb.ArrayTypeConfig) main_entity.ArrayTypeConfig {
	if m == nil {
		return main_entity.ArrayTypeConfig{}
	}
	return main_entity.ArrayTypeConfig{
		Integer:   FieldTypeIntegerConfigFromProto(m.GetInteger()),
		Float:     FieldTypeFloatConfigFromProto(m.GetFloat()),
		Decimal:   FieldTypeDecimalConfigFromProto(m.GetDecimal()),
		Char:      FieldTypeCharConfigFromProto(m.GetChar()),
		Varchar:   FieldTypeVarcharConfigFromProto(m.GetVarchar()),
		Encrypted: FieldTypeEncryptedConfigFromProto(m.GetEncrypted()),
		URL:       FieldTypeURLConfigFromProto(m.GetUrl()),
		Email:     FieldTypeEmailConfigFromProto(m.GetEmail()),
		Phone:     FieldTypePhoneConfigFromProto(m.GetPhone()),
		Date:      FieldTypeDateConfigFromProto(m.GetDate()),
		Datetime:  FieldTypeDatetimeConfigFromProto(m.GetDatetime()),
		Enum:      FieldTypeEnumConfigFromProto(m.GetEnum()),
	}
}

func ArrayTypeConfigSliceFromProto(es []*pb.ArrayTypeConfig) []main_entity.ArrayTypeConfig {
	if es == nil {
		return []main_entity.ArrayTypeConfig{}
	}
	res := []main_entity.ArrayTypeConfig{}
	for _, e := range es {
		res = append(res, ArrayTypeConfigFromProto(e))
	}
	return res
}
