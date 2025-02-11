package mapper

import (
	main_entity "nem/core/entity/extension_execution"
	pb "nem/idl/gen"

	"github.com/gofrs/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ExtensionExecutionToProto(e main_entity.ExtensionExecution) *pb.ExtensionExecution {
	return &pb.ExtensionExecution{
		Uuid:                 e.UUID.String(),
		ExtensionUuid:        e.ExtensionUUID.String(),
		ExtensionVersionUuid: e.ExtensionVersionUUID.String(),
		ProjectExtensionUuid: e.ProjectExtensionUUID.String(),
		ProjectUuid:          e.ProjectUUID.String(),
		ProjectVersionUuid:   e.ProjectVersionUUID.String(),
		ExecutedByUuid:       e.ExecutedByUUID.String(),
		Metadata:             e.Metadata,
		Status:               pb.ExtensionExecutionStatus(e.Status),
		StatusMsg:            e.StatusMsg,
		CreatedAt:            timestamppb.New(e.CreatedAt),
		UpdatedAt:            timestamppb.New(e.UpdatedAt),
	}
}

func ExtensionExecutionSliceToProto(es []main_entity.ExtensionExecution) []*pb.ExtensionExecution {
	res := []*pb.ExtensionExecution{}
	for _, e := range es {
		res = append(res, ExtensionExecutionToProto(e))
	}
	return res
}

func ExtensionExecutionFromProto(m *pb.ExtensionExecution) main_entity.ExtensionExecution {
	if m == nil {
		return main_entity.ExtensionExecution{}
	}
	return main_entity.ExtensionExecution{
		UUID:                 uuid.FromStringOrNil(m.GetUuid()),
		ExtensionUUID:        uuid.FromStringOrNil(m.GetExtensionUuid()),
		ExtensionVersionUUID: uuid.FromStringOrNil(m.GetExtensionVersionUuid()),
		ProjectExtensionUUID: uuid.FromStringOrNil(m.GetProjectExtensionUuid()),
		ProjectUUID:          uuid.FromStringOrNil(m.GetProjectUuid()),
		ProjectVersionUUID:   uuid.FromStringOrNil(m.GetProjectVersionUuid()),
		ExecutedByUUID:       uuid.FromStringOrNil(m.GetExecutedByUuid()),
		Metadata:             m.GetMetadata(),
		Status:               main_entity.Status(m.GetStatus()),
		StatusMsg:            m.GetStatusMsg(),
		CreatedAt:            m.GetCreatedAt().AsTime(),
		UpdatedAt:            m.GetUpdatedAt().AsTime(),
	}
}

func ExtensionExecutionSliceFromProto(es []*pb.ExtensionExecution) []main_entity.ExtensionExecution {
	if es == nil {
		return []main_entity.ExtensionExecution{}
	}
	res := []main_entity.ExtensionExecution{}
	for _, e := range es {
		res = append(res, ExtensionExecutionFromProto(e))
	}
	return res
}
