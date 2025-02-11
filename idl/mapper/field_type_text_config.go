package mapper

import (
	main_entity "nem/core/entity/field_type_text_config"
	pb "nem/idl/gen"
)

func FieldTypeTextConfigToProto(e main_entity.FieldTypeTextConfig) *pb.FieldTypeTextConfig {
	return &pb.FieldTypeTextConfig{
		MinSize: int64(e.MinSize),
		MaxSize: int64(e.MaxSize),
	}
}

func FieldTypeTextConfigSliceToProto(es []main_entity.FieldTypeTextConfig) []*pb.FieldTypeTextConfig {
	res := []*pb.FieldTypeTextConfig{}
	for _, e := range es {
		res = append(res, FieldTypeTextConfigToProto(e))
	}
	return res
}

func FieldTypeTextConfigFromProto(m *pb.FieldTypeTextConfig) main_entity.FieldTypeTextConfig {
	if m == nil {
		return main_entity.FieldTypeTextConfig{}
	}
	return main_entity.FieldTypeTextConfig{
		MinSize: int64(m.GetMinSize()),
		MaxSize: int64(m.GetMaxSize()),
	}
}

func FieldTypeTextConfigSliceFromProto(es []*pb.FieldTypeTextConfig) []main_entity.FieldTypeTextConfig {
	if es == nil {
		return []main_entity.FieldTypeTextConfig{}
	}
	res := []main_entity.FieldTypeTextConfig{}
	for _, e := range es {
		res = append(res, FieldTypeTextConfigFromProto(e))
	}
	return res
}
