package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/field_type_url_config"
	pb "github.com/nuzur/nem/idl/gen"
)

func FieldTypeURLConfigToProto(e main_entity.FieldTypeURLConfig) *pb.FieldTypeURLConfig {
	return &pb.FieldTypeURLConfig{
		AllowDomains:      e.AllowDomains,
		ExcludeDomains:    e.ExcludeDomains,
		HttpsRequired:     e.HTTPSRequired,
		AllowedExtensions: e.AllowedExtensions,
	}
}

func FieldTypeURLConfigSliceToProto(es []main_entity.FieldTypeURLConfig) []*pb.FieldTypeURLConfig {
	res := []*pb.FieldTypeURLConfig{}
	for _, e := range es {
		res = append(res, FieldTypeURLConfigToProto(e))
	}
	return res
}

func FieldTypeURLConfigFromProto(m *pb.FieldTypeURLConfig) main_entity.FieldTypeURLConfig {
	if m == nil {
		return main_entity.FieldTypeURLConfig{}
	}
	return main_entity.FieldTypeURLConfig{
		AllowDomains:      m.GetAllowDomains(),
		ExcludeDomains:    m.GetExcludeDomains(),
		HTTPSRequired:     m.GetHttpsRequired(),
		AllowedExtensions: m.GetAllowedExtensions(),
	}
}

func FieldTypeURLConfigSliceFromProto(es []*pb.FieldTypeURLConfig) []main_entity.FieldTypeURLConfig {
	if es == nil {
		return []main_entity.FieldTypeURLConfig{}
	}
	res := []main_entity.FieldTypeURLConfig{}
	for _, e := range es {
		res = append(res, FieldTypeURLConfigFromProto(e))
	}
	return res
}
