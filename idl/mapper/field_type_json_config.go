package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/field_type_json_config"
	pb "github.com/nuzur/nem/idl/gen"
)

func FieldTypeJSONConfigToProto(e main_entity.FieldTypeJSONConfig) *pb.FieldTypeJSONConfig {
	return &pb.FieldTypeJSONConfig{
		EnforceSchemaValidation: e.EnforceSchemaValidation,
		JsonSchema:              StringPtrToString(e.JSONSchema),
	}
}

func FieldTypeJSONConfigSliceToProto(es []main_entity.FieldTypeJSONConfig) []*pb.FieldTypeJSONConfig {
	res := []*pb.FieldTypeJSONConfig{}
	for _, e := range es {
		res = append(res, FieldTypeJSONConfigToProto(e))
	}
	return res
}

func FieldTypeJSONConfigFromProto(m *pb.FieldTypeJSONConfig) main_entity.FieldTypeJSONConfig {
	if m == nil {
		return main_entity.FieldTypeJSONConfig{}
	}
	return main_entity.FieldTypeJSONConfig{
		EnforceSchemaValidation: m.GetEnforceSchemaValidation(),
		JSONSchema:              &m.JsonSchema,
	}
}

func FieldTypeJSONConfigSliceFromProto(es []*pb.FieldTypeJSONConfig) []main_entity.FieldTypeJSONConfig {
	if es == nil {
		return []main_entity.FieldTypeJSONConfig{}
	}
	res := []main_entity.FieldTypeJSONConfig{}
	for _, e := range es {
		res = append(res, FieldTypeJSONConfigFromProto(e))
	}
	return res
}
