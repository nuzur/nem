package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/ai_usage"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func AiUsageToProto(e main_entity.AiUsage) *pb.AiUsage {
	return &pb.AiUsage{
		Uuid:               e.UUID.String(),
		UserUuid:           e.UserUUID.String(),
		ProjectUuid:        e.ProjectUUID.String(),
		ProjectVersionUuid: e.ProjectVersionUUID.String(),
		UserPrompt:         e.UserPrompt,
		Step:               e.Step,
		Context:            pb.AiUsageContext(e.Context),
		Provider:           pb.AiUsageProvider(e.Provider),
		Tokens:             int64(e.Tokens),
		Status:             pb.AiUsageStatus(e.Status),
		CreatedAt:          timestamppb.New(e.CreatedAt),
		UpdatedAt:          timestamppb.New(e.UpdatedAt),
		CreatedByUuid:      e.CreatedByUUID.String(),
		UpdatedByUuid:      e.UpdatedByUUID.String(),
	}
}

func AiUsageSliceToProto(es []main_entity.AiUsage) []*pb.AiUsage {
	res := []*pb.AiUsage{}
	for _, e := range es {
		res = append(res, AiUsageToProto(e))
	}
	return res
}

func AiUsageFromProto(m *pb.AiUsage) main_entity.AiUsage {
	if m == nil {
		return main_entity.AiUsage{}
	}
	return main_entity.AiUsage{
		UUID:               uuid.FromStringOrNil(m.GetUuid()),
		UserUUID:           uuid.FromStringOrNil(m.GetUserUuid()),
		ProjectUUID:        uuid.FromStringOrNil(m.GetProjectUuid()),
		ProjectVersionUUID: uuid.FromStringOrNil(m.GetProjectVersionUuid()),
		UserPrompt:         m.GetUserPrompt(),
		Step:               m.GetStep(),
		Context:            main_entity.Context(m.GetContext()),
		Provider:           main_entity.Provider(m.GetProvider()),
		Tokens:             int64(m.GetTokens()),
		Status:             main_entity.Status(m.GetStatus()),
		CreatedAt:          m.GetCreatedAt().AsTime(),
		UpdatedAt:          m.GetUpdatedAt().AsTime(),
		CreatedByUUID:      uuid.FromStringOrNil(m.GetCreatedByUuid()),
		UpdatedByUUID:      uuid.FromStringOrNil(m.GetUpdatedByUuid()),
	}
}

func AiUsageSliceFromProto(es []*pb.AiUsage) []main_entity.AiUsage {
	if es == nil {
		return []main_entity.AiUsage{}
	}
	res := []main_entity.AiUsage{}
	for _, e := range es {
		res = append(res, AiUsageFromProto(e))
	}
	return res
}
