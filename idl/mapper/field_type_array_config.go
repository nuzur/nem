package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/field_type_array_config"
	pb "github.com/nuzur/nem/idl/gen"
)

func FieldTypeArrayConfigToProto(e main_entity.FieldTypeArrayConfig) *pb.FieldTypeArrayConfig {
	return &pb.FieldTypeArrayConfig{
		MinElements:   int64(e.MinElements),
		MaxElements:   int64(e.MaxElements),
		EnforceUnique: e.EnforceUnique,
		Type:          pb.FieldTypeArrayConfigType(e.Type),
		TypeConfig:    ArrayTypeConfigToProto(e.TypeConfig),
	}
}

func FieldTypeArrayConfigSliceToProto(es []main_entity.FieldTypeArrayConfig) []*pb.FieldTypeArrayConfig {
	res := []*pb.FieldTypeArrayConfig{}
	for _, e := range es {
		res = append(res, FieldTypeArrayConfigToProto(e))
	}
	return res
}

func FieldTypeArrayConfigFromProto(m *pb.FieldTypeArrayConfig) main_entity.FieldTypeArrayConfig {
	if m == nil {
		return main_entity.FieldTypeArrayConfig{}
	}
	return main_entity.FieldTypeArrayConfig{
		MinElements:   int64(m.GetMinElements()),
		MaxElements:   int64(m.GetMaxElements()),
		EnforceUnique: m.GetEnforceUnique(),
		Type:          main_entity.Type(m.GetType()),
		TypeConfig:    ArrayTypeConfigFromProto(m.GetTypeConfig()),
	}
}

func FieldTypeArrayConfigSliceFromProto(es []*pb.FieldTypeArrayConfig) []main_entity.FieldTypeArrayConfig {
	if es == nil {
		return []main_entity.FieldTypeArrayConfig{}
	}
	res := []main_entity.FieldTypeArrayConfig{}
	for _, e := range es {
		res = append(res, FieldTypeArrayConfigFromProto(e))
	}
	return res
}
