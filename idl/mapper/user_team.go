package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/user_team"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func UserTeamToProto(e main_entity.UserTeam) *pb.UserTeam {
	return &pb.UserTeam{
		Uuid:          e.UUID.String(),
		UserUuid:      e.UserUUID.String(),
		UserEmail:     e.UserEmail,
		TeamUuid:      e.TeamUUID.String(),
		Roles:         UserTeamRoleSliceToProto(e.Roles),
		Status:        pb.UserTeamStatus(e.Status),
		CreatedAt:     timestamppb.New(e.CreatedAt),
		UpdatedAt:     timestamppb.New(e.UpdatedAt),
		CreatedByUuid: e.CreatedByUUID.String(),
		UpdatedByUuid: e.UpdatedByUUID.String(),
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
		UUID:          uuid.FromStringOrNil(m.GetUuid()),
		UserUUID:      uuid.FromStringOrNil(m.GetUserUuid()),
		UserEmail:     m.GetUserEmail(),
		TeamUUID:      uuid.FromStringOrNil(m.GetTeamUuid()),
		Roles:         UserTeamRoleSliceFromProto(m.GetRoles()),
		Status:        main_entity.Status(m.GetStatus()),
		CreatedAt:     m.GetCreatedAt().AsTime(),
		UpdatedAt:     m.GetUpdatedAt().AsTime(),
		CreatedByUUID: uuid.FromStringOrNil(m.GetCreatedByUuid()),
		UpdatedByUUID: uuid.FromStringOrNil(m.GetUpdatedByUuid()),
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

func UserTeamRoleSliceToProto(s []main_entity.Role) []pb.UserTeamRole {
	res := []pb.UserTeamRole{}
	for _, e := range s {
		res = append(res, pb.UserTeamRole(e))
	}
	return res
}
func UserTeamRoleSliceFromProto(s []pb.UserTeamRole) []main_entity.Role {
	res := []main_entity.Role{}
	for _, e := range s {
		res = append(res, main_entity.Role(e))
	}
	return res
}
