package mapper

import (
	main_entity "nem/core/entity/field_type_phone_config"
	pb "nem/idl/gen"
)

func FieldTypePhoneConfigToProto(e main_entity.FieldTypePhoneConfig) *pb.FieldTypePhoneConfig {
	return &pb.FieldTypePhoneConfig{
		AllowCountries:   e.AllowCountries,
		ExcludeCountries: e.ExcludeCountries,
	}
}

func FieldTypePhoneConfigSliceToProto(es []main_entity.FieldTypePhoneConfig) []*pb.FieldTypePhoneConfig {
	res := []*pb.FieldTypePhoneConfig{}
	for _, e := range es {
		res = append(res, FieldTypePhoneConfigToProto(e))
	}
	return res
}

func FieldTypePhoneConfigFromProto(m *pb.FieldTypePhoneConfig) main_entity.FieldTypePhoneConfig {
	if m == nil {
		return main_entity.FieldTypePhoneConfig{}
	}
	return main_entity.FieldTypePhoneConfig{
		AllowCountries:   m.GetAllowCountries(),
		ExcludeCountries: m.GetExcludeCountries(),
	}
}

func FieldTypePhoneConfigSliceFromProto(es []*pb.FieldTypePhoneConfig) []main_entity.FieldTypePhoneConfig {
	if es == nil {
		return []main_entity.FieldTypePhoneConfig{}
	}
	res := []main_entity.FieldTypePhoneConfig{}
	for _, e := range es {
		res = append(res, FieldTypePhoneConfigFromProto(e))
	}
	return res
}
