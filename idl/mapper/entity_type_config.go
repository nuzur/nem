package mapper

import (
	main_entity "nem/core/entity/entity_type_config"
	pb "nem/idl/gen"
)

func EntityTypeConfigToProto(e main_entity.EntityTypeConfig) *pb.EntityTypeConfig {
	return &pb.EntityTypeConfig{
		Standalone: EntityTypeStandaloneConfigToProto(e.Standalone),
		Dependent:  EntityTypeDependentConfigToProto(e.Dependent),
	}
}

func EntityTypeConfigSliceToProto(es []main_entity.EntityTypeConfig) []*pb.EntityTypeConfig {
	res := []*pb.EntityTypeConfig{}
	for _, e := range es {
		res = append(res, EntityTypeConfigToProto(e))
	}
	return res
}

func EntityTypeConfigFromProto(m *pb.EntityTypeConfig) main_entity.EntityTypeConfig {
	if m == nil {
		return main_entity.EntityTypeConfig{}
	}
	return main_entity.EntityTypeConfig{
		Standalone: EntityTypeStandaloneConfigFromProto(m.GetStandalone()),
		Dependent:  EntityTypeDependentConfigFromProto(m.GetDependent()),
	}
}

func EntityTypeConfigSliceFromProto(es []*pb.EntityTypeConfig) []main_entity.EntityTypeConfig {
	if es == nil {
		return []main_entity.EntityTypeConfig{}
	}
	res := []main_entity.EntityTypeConfig{}
	for _, e := range es {
		res = append(res, EntityTypeConfigFromProto(e))
	}
	return res
}
