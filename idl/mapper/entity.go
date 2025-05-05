package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/entity"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func EntityToProto(e main_entity.Entity) *pb.Entity {
	return &pb.Entity{
		Uuid:                 e.UUID.String(),
		Version:              int64(e.Version),
		Identifier:           e.Identifier,
		Description:          StringPtrToString(e.Description),
		Fields:               FieldSliceToProto(e.Fields),
		Type:                 pb.EntityType(e.Type),
		TypeConfig:           EntityTypeConfigToProto(e.TypeConfig),
		Render:               ElementRenderToProto(e.Render),
		Status:               pb.EntityStatus(e.Status),
		CreatedAt:            timestamppb.New(e.CreatedAt),
		UpdatedAt:            timestamppb.New(e.UpdatedAt),
		CreatedByUuid:        e.CreatedByUUID.String(),
		UpdatedByUuid:        e.UpdatedByUUID.String(),
		DataManagementConfig: EntityDataManagementConfigToProto(e.DataManagementConfig),
	}
}

func EntitySliceToProto(es []main_entity.Entity) []*pb.Entity {
	res := []*pb.Entity{}
	for _, e := range es {
		res = append(res, EntityToProto(e))
	}
	return res
}

func EntityFromProto(m *pb.Entity) main_entity.Entity {
	if m == nil {
		return main_entity.Entity{}
	}
	return main_entity.Entity{
		UUID:                 uuid.FromStringOrNil(m.GetUuid()),
		Version:              int64(m.GetVersion()),
		Identifier:           m.GetIdentifier(),
		Description:          &m.Description,
		Fields:               FieldSliceFromProto(m.GetFields()),
		Type:                 main_entity.Type(m.GetType()),
		TypeConfig:           EntityTypeConfigFromProto(m.GetTypeConfig()),
		Render:               ElementRenderFromProto(m.GetRender()),
		Status:               main_entity.Status(m.GetStatus()),
		CreatedAt:            m.GetCreatedAt().AsTime(),
		UpdatedAt:            m.GetUpdatedAt().AsTime(),
		CreatedByUUID:        uuid.FromStringOrNil(m.GetCreatedByUuid()),
		UpdatedByUUID:        uuid.FromStringOrNil(m.GetUpdatedByUuid()),
		DataManagementConfig: EntityDataManagementConfigFromProto(m.GetDataManagementConfig()),
	}
}

func EntitySliceFromProto(es []*pb.Entity) []main_entity.Entity {
	if es == nil {
		return []main_entity.Entity{}
	}
	res := []main_entity.Entity{}
	for _, e := range es {
		res = append(res, EntityFromProto(e))
	}
	return res
}
