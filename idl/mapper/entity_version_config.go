package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/entity_version_config"
	pb "github.com/nuzur/nem/idl/gen"
)

func EntityVersionConfigToProto(e main_entity.EntityVersionConfig) *pb.EntityVersionConfig {
	return &pb.EntityVersionConfig{
		Type:      pb.EntityVersionConfigType(e.Type),
		Generator: pb.EntityVersionConfigGenerator(e.Generator),
		Config:    EntityVersionTypeConfigToProto(e.Config),
	}
}

func EntityVersionConfigSliceToProto(es []main_entity.EntityVersionConfig) []*pb.EntityVersionConfig {
	res := []*pb.EntityVersionConfig{}
	for _, e := range es {
		res = append(res, EntityVersionConfigToProto(e))
	}
	return res
}

func EntityVersionConfigFromProto(m *pb.EntityVersionConfig) main_entity.EntityVersionConfig {
	if m == nil {
		return main_entity.EntityVersionConfig{}
	}
	return main_entity.EntityVersionConfig{
		Type:      main_entity.Type(m.GetType()),
		Generator: main_entity.Generator(m.GetGenerator()),
		Config:    EntityVersionTypeConfigFromProto(m.GetConfig()),
	}
}

func EntityVersionConfigSliceFromProto(es []*pb.EntityVersionConfig) []main_entity.EntityVersionConfig {
	if es == nil {
		return []main_entity.EntityVersionConfig{}
	}
	res := []main_entity.EntityVersionConfig{}
	for _, e := range es {
		res = append(res, EntityVersionConfigFromProto(e))
	}
	return res
}
