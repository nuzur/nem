package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/entity_version_type_field_config"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"
)

func EntityVersionTypeFieldConfigToProto(e main_entity.EntityVersionTypeFieldConfig) *pb.EntityVersionTypeFieldConfig {
	return &pb.EntityVersionTypeFieldConfig{
		VersionFieldUuid: e.VersionFieldUUID.String(),
		AppendOnly:       e.AppendOnly,
	}
}

func EntityVersionTypeFieldConfigSliceToProto(es []main_entity.EntityVersionTypeFieldConfig) []*pb.EntityVersionTypeFieldConfig {
	res := []*pb.EntityVersionTypeFieldConfig{}
	for _, e := range es {
		res = append(res, EntityVersionTypeFieldConfigToProto(e))
	}
	return res
}

func EntityVersionTypeFieldConfigFromProto(m *pb.EntityVersionTypeFieldConfig) main_entity.EntityVersionTypeFieldConfig {
	if m == nil {
		return main_entity.EntityVersionTypeFieldConfig{}
	}
	return main_entity.EntityVersionTypeFieldConfig{
		VersionFieldUUID: uuid.FromStringOrNil(m.GetVersionFieldUuid()),
		AppendOnly:       m.GetAppendOnly(),
	}
}

func EntityVersionTypeFieldConfigSliceFromProto(es []*pb.EntityVersionTypeFieldConfig) []main_entity.EntityVersionTypeFieldConfig {
	if es == nil {
		return []main_entity.EntityVersionTypeFieldConfig{}
	}
	res := []main_entity.EntityVersionTypeFieldConfig{}
	for _, e := range es {
		res = append(res, EntityVersionTypeFieldConfigFromProto(e))
	}
	return res
}
