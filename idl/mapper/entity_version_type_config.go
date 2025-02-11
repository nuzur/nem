package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/entity_version_type_config"
	pb "github.com/nuzur/nem/idl/gen"
)

func EntityVersionTypeConfigToProto(e main_entity.EntityVersionTypeConfig) *pb.EntityVersionTypeConfig {
	return &pb.EntityVersionTypeConfig{
		Field:  EntityVersionTypeFieldConfigToProto(e.Field),
		Entity: EntityVersionTypeEntityConfigToProto(e.Entity),
	}
}

func EntityVersionTypeConfigSliceToProto(es []main_entity.EntityVersionTypeConfig) []*pb.EntityVersionTypeConfig {
	res := []*pb.EntityVersionTypeConfig{}
	for _, e := range es {
		res = append(res, EntityVersionTypeConfigToProto(e))
	}
	return res
}

func EntityVersionTypeConfigFromProto(m *pb.EntityVersionTypeConfig) main_entity.EntityVersionTypeConfig {
	if m == nil {
		return main_entity.EntityVersionTypeConfig{}
	}
	return main_entity.EntityVersionTypeConfig{
		Field:  EntityVersionTypeFieldConfigFromProto(m.GetField()),
		Entity: EntityVersionTypeEntityConfigFromProto(m.GetEntity()),
	}
}

func EntityVersionTypeConfigSliceFromProto(es []*pb.EntityVersionTypeConfig) []main_entity.EntityVersionTypeConfig {
	if es == nil {
		return []main_entity.EntityVersionTypeConfig{}
	}
	res := []main_entity.EntityVersionTypeConfig{}
	for _, e := range es {
		res = append(res, EntityVersionTypeConfigFromProto(e))
	}
	return res
}
