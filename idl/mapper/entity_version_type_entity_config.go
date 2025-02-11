package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/entity_version_type_entity_config"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"
)

func EntityVersionTypeEntityConfigToProto(e main_entity.EntityVersionTypeEntityConfig) *pb.EntityVersionTypeEntityConfig {
	return &pb.EntityVersionTypeEntityConfig{
		EntityUuid:       e.EntityUUID.String(),
		VersionFieldUuid: e.VersionFieldUUID.String(),
		AppendOnly:       e.AppendOnly,
	}
}

func EntityVersionTypeEntityConfigSliceToProto(es []main_entity.EntityVersionTypeEntityConfig) []*pb.EntityVersionTypeEntityConfig {
	res := []*pb.EntityVersionTypeEntityConfig{}
	for _, e := range es {
		res = append(res, EntityVersionTypeEntityConfigToProto(e))
	}
	return res
}

func EntityVersionTypeEntityConfigFromProto(m *pb.EntityVersionTypeEntityConfig) main_entity.EntityVersionTypeEntityConfig {
	if m == nil {
		return main_entity.EntityVersionTypeEntityConfig{}
	}
	return main_entity.EntityVersionTypeEntityConfig{
		EntityUUID:       uuid.FromStringOrNil(m.GetEntityUuid()),
		VersionFieldUUID: uuid.FromStringOrNil(m.GetVersionFieldUuid()),
		AppendOnly:       m.GetAppendOnly(),
	}
}

func EntityVersionTypeEntityConfigSliceFromProto(es []*pb.EntityVersionTypeEntityConfig) []main_entity.EntityVersionTypeEntityConfig {
	if es == nil {
		return []main_entity.EntityVersionTypeEntityConfig{}
	}
	res := []main_entity.EntityVersionTypeEntityConfig{}
	for _, e := range es {
		res = append(res, EntityVersionTypeEntityConfigFromProto(e))
	}
	return res
}
