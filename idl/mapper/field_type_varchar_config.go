package mapper

import (
	main_entity "nem/core/entity/field_type_varchar_config"
	pb "nem/idl/gen"
)

func FieldTypeVarcharConfigToProto(e main_entity.FieldTypeVarcharConfig) *pb.FieldTypeVarcharConfig {
	return &pb.FieldTypeVarcharConfig{
		MinSize:         int64(e.MinSize),
		MaxSize:         int64(e.MaxSize),
		RegexValidation: e.RegexValidation,
	}
}

func FieldTypeVarcharConfigSliceToProto(es []main_entity.FieldTypeVarcharConfig) []*pb.FieldTypeVarcharConfig {
	res := []*pb.FieldTypeVarcharConfig{}
	for _, e := range es {
		res = append(res, FieldTypeVarcharConfigToProto(e))
	}
	return res
}

func FieldTypeVarcharConfigFromProto(m *pb.FieldTypeVarcharConfig) main_entity.FieldTypeVarcharConfig {
	if m == nil {
		return main_entity.FieldTypeVarcharConfig{}
	}
	return main_entity.FieldTypeVarcharConfig{
		MinSize:         int64(m.GetMinSize()),
		MaxSize:         int64(m.GetMaxSize()),
		RegexValidation: m.GetRegexValidation(),
	}
}

func FieldTypeVarcharConfigSliceFromProto(es []*pb.FieldTypeVarcharConfig) []main_entity.FieldTypeVarcharConfig {
	if es == nil {
		return []main_entity.FieldTypeVarcharConfig{}
	}
	res := []main_entity.FieldTypeVarcharConfig{}
	for _, e := range es {
		res = append(res, FieldTypeVarcharConfigFromProto(e))
	}
	return res
}
