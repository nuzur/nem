package mapper

import (
	main_entity "nem/core/entity/entity_type_dependent_config"
	pb "nem/idl/gen"
)

func EntityTypeDependentConfigToProto(e main_entity.EntityTypeDependentConfig) *pb.EntityTypeDependentConfig {
	return &pb.EntityTypeDependentConfig{
		ServiceSourceUuids: UUIDSliceToStringSlice(e.ServiceSourceUUIDs),
		EntitySourceUuids:  UUIDSliceToStringSlice(e.EntitySourceUUIDs),
	}
}

func EntityTypeDependentConfigSliceToProto(es []main_entity.EntityTypeDependentConfig) []*pb.EntityTypeDependentConfig {
	res := []*pb.EntityTypeDependentConfig{}
	for _, e := range es {
		res = append(res, EntityTypeDependentConfigToProto(e))
	}
	return res
}

func EntityTypeDependentConfigFromProto(m *pb.EntityTypeDependentConfig) main_entity.EntityTypeDependentConfig {
	if m == nil {
		return main_entity.EntityTypeDependentConfig{}
	}
	return main_entity.EntityTypeDependentConfig{
		ServiceSourceUUIDs: StringSliceToUUIDSlice(m.GetServiceSourceUuids()),
		EntitySourceUUIDs:  StringSliceToUUIDSlice(m.GetEntitySourceUuids()),
	}
}

func EntityTypeDependentConfigSliceFromProto(es []*pb.EntityTypeDependentConfig) []main_entity.EntityTypeDependentConfig {
	if es == nil {
		return []main_entity.EntityTypeDependentConfig{}
	}
	res := []main_entity.EntityTypeDependentConfig{}
	for _, e := range es {
		res = append(res, EntityTypeDependentConfigFromProto(e))
	}
	return res
}
