package mapper

import (
	main_entity "nem/core/entity/field_type_char_config"
	pb "nem/idl/gen"
)

func FieldTypeCharConfigToProto(e main_entity.FieldTypeCharConfig) *pb.FieldTypeCharConfig {
	return &pb.FieldTypeCharConfig{
		MinSize:         int64(e.MinSize),
		MaxSize:         int64(e.MaxSize),
		RegexValidation: e.RegexValidation,
	}
}

func FieldTypeCharConfigSliceToProto(es []main_entity.FieldTypeCharConfig) []*pb.FieldTypeCharConfig {
	res := []*pb.FieldTypeCharConfig{}
	for _, e := range es {
		res = append(res, FieldTypeCharConfigToProto(e))
	}
	return res
}

func FieldTypeCharConfigFromProto(m *pb.FieldTypeCharConfig) main_entity.FieldTypeCharConfig {
	if m == nil {
		return main_entity.FieldTypeCharConfig{}
	}
	return main_entity.FieldTypeCharConfig{
		MinSize:         int64(m.GetMinSize()),
		MaxSize:         int64(m.GetMaxSize()),
		RegexValidation: m.GetRegexValidation(),
	}
}

func FieldTypeCharConfigSliceFromProto(es []*pb.FieldTypeCharConfig) []main_entity.FieldTypeCharConfig {
	if es == nil {
		return []main_entity.FieldTypeCharConfig{}
	}
	res := []main_entity.FieldTypeCharConfig{}
	for _, e := range es {
		res = append(res, FieldTypeCharConfigFromProto(e))
	}
	return res
}
