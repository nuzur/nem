package mapper

import (
	main_entity "nem/core/entity/field_type_slug_config"
	pb "nem/idl/gen"
)

func FieldTypeSlugConfigToProto(e main_entity.FieldTypeSlugConfig) *pb.FieldTypeSlugConfig {
	return &pb.FieldTypeSlugConfig{
		MinSize:           int64(e.MinSize),
		MaxSize:           int64(e.MaxSize),
		RegexValidation:   e.RegexValidation,
		BasedOnFieldUuids: UUIDSliceToStringSlice(e.BasedOnFieldUUIDs),
	}
}

func FieldTypeSlugConfigSliceToProto(es []main_entity.FieldTypeSlugConfig) []*pb.FieldTypeSlugConfig {
	res := []*pb.FieldTypeSlugConfig{}
	for _, e := range es {
		res = append(res, FieldTypeSlugConfigToProto(e))
	}
	return res
}

func FieldTypeSlugConfigFromProto(m *pb.FieldTypeSlugConfig) main_entity.FieldTypeSlugConfig {
	if m == nil {
		return main_entity.FieldTypeSlugConfig{}
	}
	return main_entity.FieldTypeSlugConfig{
		MinSize:           int64(m.GetMinSize()),
		MaxSize:           int64(m.GetMaxSize()),
		RegexValidation:   m.GetRegexValidation(),
		BasedOnFieldUUIDs: StringSliceToUUIDSlice(m.GetBasedOnFieldUuids()),
	}
}

func FieldTypeSlugConfigSliceFromProto(es []*pb.FieldTypeSlugConfig) []main_entity.FieldTypeSlugConfig {
	if es == nil {
		return []main_entity.FieldTypeSlugConfig{}
	}
	res := []main_entity.FieldTypeSlugConfig{}
	for _, e := range es {
		res = append(res, FieldTypeSlugConfigFromProto(e))
	}
	return res
}
