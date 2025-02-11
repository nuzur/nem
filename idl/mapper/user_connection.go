package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/user_connection"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func UserConnectionToProto(e main_entity.UserConnection) *pb.UserConnection {
	return &pb.UserConnection{
		Uuid:               e.UUID.String(),
		UserUuid:           e.UserUUID.String(),
		ProjectUuid:        e.ProjectUUID.String(),
		ProjectVersionUuid: e.ProjectVersionUUID.String(),
		Type:               pb.UserConnectionType(e.Type),
		TypeConfig:         UserConnectionTypeConfigToProto(e.TypeConfig),
		DbSchema:           e.DbSchema,
		Executions:         UserConnectionExecutionSliceToProto(e.Executions),
		Status:             pb.UserConnectionStatus(e.Status),
		CreatedAt:          timestamppb.New(e.CreatedAt),
		UpdatedAt:          timestamppb.New(e.UpdatedAt),
	}
}

func UserConnectionSliceToProto(es []main_entity.UserConnection) []*pb.UserConnection {
	res := []*pb.UserConnection{}
	for _, e := range es {
		res = append(res, UserConnectionToProto(e))
	}
	return res
}

func UserConnectionFromProto(m *pb.UserConnection) main_entity.UserConnection {
	if m == nil {
		return main_entity.UserConnection{}
	}
	return main_entity.UserConnection{
		UUID:               uuid.FromStringOrNil(m.GetUuid()),
		UserUUID:           uuid.FromStringOrNil(m.GetUserUuid()),
		ProjectUUID:        uuid.FromStringOrNil(m.GetProjectUuid()),
		ProjectVersionUUID: uuid.FromStringOrNil(m.GetProjectVersionUuid()),
		Type:               main_entity.Type(m.GetType()),
		TypeConfig:         UserConnectionTypeConfigFromProto(m.GetTypeConfig()),
		DbSchema:           m.GetDbSchema(),
		Executions:         UserConnectionExecutionSliceFromProto(m.GetExecutions()),
		Status:             main_entity.Status(m.GetStatus()),
		CreatedAt:          m.GetCreatedAt().AsTime(),
		UpdatedAt:          m.GetUpdatedAt().AsTime(),
	}
}

func UserConnectionSliceFromProto(es []*pb.UserConnection) []main_entity.UserConnection {
	if es == nil {
		return []main_entity.UserConnection{}
	}
	res := []main_entity.UserConnection{}
	for _, e := range es {
		res = append(res, UserConnectionFromProto(e))
	}
	return res
}
