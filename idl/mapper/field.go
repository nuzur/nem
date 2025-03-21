package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/field"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func FieldToProto(e main_entity.Field) *pb.Field {
	return &pb.Field{
		Uuid:             e.UUID.String(),
		Version:          int64(e.Version),
		Identifier:       e.Identifier,
		Description:      StringPtrToString(e.Description),
		Type:             pb.FieldType(e.Type),
		TypeConfig:       FieldTypeConfigToProto(e.TypeConfig),
		Required:         e.Required,
		Key:              e.Key,
		KeyAutoIncrement: e.KeyAutoIncrement,
		Unique:           e.Unique,
		Deprecated:       e.Deprecated,
		Generated:        e.Generated,
		Status:           pb.FieldStatus(e.Status),
		CreatedAt:        timestamppb.New(e.CreatedAt),
		UpdatedAt:        timestamppb.New(e.UpdatedAt),
		CreatedByUuid:    e.CreatedByUUID.String(),
		UpdatedByUuid:    e.UpdatedByUUID.String(),
	}
}

func FieldSliceToProto(es []main_entity.Field) []*pb.Field {
	res := []*pb.Field{}
	for _, e := range es {
		res = append(res, FieldToProto(e))
	}
	return res
}

func FieldFromProto(m *pb.Field) main_entity.Field {
	if m == nil {
		return main_entity.Field{}
	}
	return main_entity.Field{
		UUID:             uuid.FromStringOrNil(m.GetUuid()),
		Version:          int64(m.GetVersion()),
		Identifier:       m.GetIdentifier(),
		Description:      &m.Description,
		Type:             main_entity.Type(m.GetType()),
		TypeConfig:       FieldTypeConfigFromProto(m.GetTypeConfig()),
		Required:         m.GetRequired(),
		Key:              m.GetKey(),
		KeyAutoIncrement: m.GetKeyAutoIncrement(),
		Unique:           m.GetUnique(),
		Deprecated:       m.GetDeprecated(),
		Generated:        m.GetGenerated(),
		Status:           main_entity.Status(m.GetStatus()),
		CreatedAt:        m.GetCreatedAt().AsTime(),
		UpdatedAt:        m.GetUpdatedAt().AsTime(),
		CreatedByUUID:    uuid.FromStringOrNil(m.GetCreatedByUuid()),
		UpdatedByUUID:    uuid.FromStringOrNil(m.GetUpdatedByUuid()),
	}
}

func FieldSliceFromProto(es []*pb.Field) []main_entity.Field {
	if es == nil {
		return []main_entity.Field{}
	}
	res := []main_entity.Field{}
	for _, e := range es {
		res = append(res, FieldFromProto(e))
	}
	return res
}
