package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/field_type_email_config"
	pb "github.com/nuzur/nem/idl/gen"
)

func FieldTypeEmailConfigToProto(e main_entity.FieldTypeEmailConfig) *pb.FieldTypeEmailConfig {
	return &pb.FieldTypeEmailConfig{
		AllowDomains:   e.AllowDomains,
		ExcludeDomains: e.ExcludeDomains,
	}
}

func FieldTypeEmailConfigSliceToProto(es []main_entity.FieldTypeEmailConfig) []*pb.FieldTypeEmailConfig {
	res := []*pb.FieldTypeEmailConfig{}
	for _, e := range es {
		res = append(res, FieldTypeEmailConfigToProto(e))
	}
	return res
}

func FieldTypeEmailConfigFromProto(m *pb.FieldTypeEmailConfig) main_entity.FieldTypeEmailConfig {
	if m == nil {
		return main_entity.FieldTypeEmailConfig{}
	}
	return main_entity.FieldTypeEmailConfig{
		AllowDomains:   m.GetAllowDomains(),
		ExcludeDomains: m.GetExcludeDomains(),
	}
}

func FieldTypeEmailConfigSliceFromProto(es []*pb.FieldTypeEmailConfig) []main_entity.FieldTypeEmailConfig {
	if es == nil {
		return []main_entity.FieldTypeEmailConfig{}
	}
	res := []main_entity.FieldTypeEmailConfig{}
	for _, e := range es {
		res = append(res, FieldTypeEmailConfigFromProto(e))
	}
	return res
}
