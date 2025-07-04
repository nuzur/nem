package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/user_team"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func UserTeamToProto(e main_entity.UserTeam) *pb.UserTeam {
	return &pb.UserTeam{
		Uuid:                    e.UUID.String(),
		UserUuid:                e.UserUUID.String(),
		UserEmail:               StringPtrToString(e.UserEmail),
		TeamUuid:                e.TeamUUID.String(),
		Role:                    pb.UserTeamRole(e.Role),
		ReviewRequiredStructure: e.ReviewRequiredStructure,
		ReviewRequiredData:      e.ReviewRequiredData,
		Status:                  pb.UserTeamStatus(e.Status),
		CreatedAt:               timestamppb.New(e.CreatedAt),
		UpdatedAt:               timestamppb.New(e.UpdatedAt),
		CreatedByUuid:           e.CreatedByUUID.String(),
		UpdatedByUuid:           e.UpdatedByUUID.String(),
	}
}

func UserTeamSliceToProto(es []main_entity.UserTeam) []*pb.UserTeam {
	res := []*pb.UserTeam{}
	for _, e := range es {
		res = append(res, UserTeamToProto(e))
	}
	return res
}

func UserTeamFromProto(m *pb.UserTeam) main_entity.UserTeam {
	if m == nil {
		return main_entity.UserTeam{}
	}
	return main_entity.UserTeam{
		UUID:                    uuid.FromStringOrNil(m.GetUuid()),
		UserUUID:                uuid.FromStringOrNil(m.GetUserUuid()),
		UserEmail:               &m.UserEmail,
		TeamUUID:                uuid.FromStringOrNil(m.GetTeamUuid()),
		Role:                    main_entity.Role(m.GetRole()),
		ReviewRequiredStructure: m.GetReviewRequiredStructure(),
		ReviewRequiredData:      m.GetReviewRequiredData(),
		Status:                  main_entity.Status(m.GetStatus()),
		CreatedAt:               m.GetCreatedAt().AsTime(),
		UpdatedAt:               m.GetUpdatedAt().AsTime(),
		CreatedByUUID:           uuid.FromStringOrNil(m.GetCreatedByUuid()),
		UpdatedByUUID:           uuid.FromStringOrNil(m.GetUpdatedByUuid()),
	}
}

func UserTeamSliceFromProto(es []*pb.UserTeam) []main_entity.UserTeam {
	if es == nil {
		return []main_entity.UserTeam{}
	}
	res := []main_entity.UserTeam{}
	for _, e := range es {
		res = append(res, UserTeamFromProto(e))
	}
	return res
}
