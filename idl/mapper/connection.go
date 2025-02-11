package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/connection"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConnectionToProto(e main_entity.Connection) *pb.Connection {
	return &pb.Connection{
		Uuid:           e.UUID.String(),
		Version:        int64(e.Version),
		StoreUuid:      e.StoreUUID.String(),
		EnviormentUuid: e.EnviormentUUID.String(),
		Identifier:     e.Identifier,
		DbType:         pb.ConnectionDbType(e.DbType),
		DbTypeConfig:   DbTypeConfigToProto(e.DbTypeConfig),
		DbVersion:      e.DbVersion,
		Type:           pb.ConnectionType(e.Type),
		TypeConfig:     ConnectionTypeConfigToProto(e.TypeConfig),
		Status:         pb.ConnectionStatus(e.Status),
		CreatedAt:      timestamppb.New(e.CreatedAt),
		UpdatedAt:      timestamppb.New(e.UpdatedAt),
		CreatedByUuid:  e.CreatedByUUID.String(),
		UpdatedByUuid:  e.UpdatedByUUID.String(),
	}
}

func ConnectionSliceToProto(es []main_entity.Connection) []*pb.Connection {
	res := []*pb.Connection{}
	for _, e := range es {
		res = append(res, ConnectionToProto(e))
	}
	return res
}

func ConnectionFromProto(m *pb.Connection) main_entity.Connection {
	if m == nil {
		return main_entity.Connection{}
	}
	return main_entity.Connection{
		UUID:           uuid.FromStringOrNil(m.GetUuid()),
		Version:        int64(m.GetVersion()),
		StoreUUID:      uuid.FromStringOrNil(m.GetStoreUuid()),
		EnviormentUUID: uuid.FromStringOrNil(m.GetEnviormentUuid()),
		Identifier:     m.GetIdentifier(),
		DbType:         main_entity.DbType(m.GetDbType()),
		DbTypeConfig:   DbTypeConfigFromProto(m.GetDbTypeConfig()),
		DbVersion:      m.GetDbVersion(),
		Type:           main_entity.Type(m.GetType()),
		TypeConfig:     ConnectionTypeConfigFromProto(m.GetTypeConfig()),
		Status:         main_entity.Status(m.GetStatus()),
		CreatedAt:      m.GetCreatedAt().AsTime(),
		UpdatedAt:      m.GetUpdatedAt().AsTime(),
		CreatedByUUID:  uuid.FromStringOrNil(m.GetCreatedByUuid()),
		UpdatedByUUID:  uuid.FromStringOrNil(m.GetUpdatedByUuid()),
	}
}

func ConnectionSliceFromProto(es []*pb.Connection) []main_entity.Connection {
	if es == nil {
		return []main_entity.Connection{}
	}
	res := []main_entity.Connection{}
	for _, e := range es {
		res = append(res, ConnectionFromProto(e))
	}
	return res
}
