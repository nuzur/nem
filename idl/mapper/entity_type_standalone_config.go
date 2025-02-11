package mapper

import (
	main_entity "nem/core/entity/entity_type_standalone_config"
	pb "nem/idl/gen"

	"github.com/gofrs/uuid"
)

func EntityTypeStandaloneConfigToProto(e main_entity.EntityTypeStandaloneConfig) *pb.EntityTypeStandaloneConfig {
	return &pb.EntityTypeStandaloneConfig{
		StoreUuid:     e.StoreUUID.String(),
		Versioned:     e.Versioned,
		VersionConfig: EntityVersionConfigToProto(e.VersionConfig),
		Indexes:       IndexSliceToProto(e.Indexes),
	}
}

func EntityTypeStandaloneConfigSliceToProto(es []main_entity.EntityTypeStandaloneConfig) []*pb.EntityTypeStandaloneConfig {
	res := []*pb.EntityTypeStandaloneConfig{}
	for _, e := range es {
		res = append(res, EntityTypeStandaloneConfigToProto(e))
	}
	return res
}

func EntityTypeStandaloneConfigFromProto(m *pb.EntityTypeStandaloneConfig) main_entity.EntityTypeStandaloneConfig {
	if m == nil {
		return main_entity.EntityTypeStandaloneConfig{}
	}
	return main_entity.EntityTypeStandaloneConfig{
		StoreUUID:     uuid.FromStringOrNil(m.GetStoreUuid()),
		Versioned:     m.GetVersioned(),
		VersionConfig: EntityVersionConfigFromProto(m.GetVersionConfig()),
		Indexes:       IndexSliceFromProto(m.GetIndexes()),
	}
}

func EntityTypeStandaloneConfigSliceFromProto(es []*pb.EntityTypeStandaloneConfig) []main_entity.EntityTypeStandaloneConfig {
	if es == nil {
		return []main_entity.EntityTypeStandaloneConfig{}
	}
	res := []main_entity.EntityTypeStandaloneConfig{}
	for _, e := range es {
		res = append(res, EntityTypeStandaloneConfigFromProto(e))
	}
	return res
}
