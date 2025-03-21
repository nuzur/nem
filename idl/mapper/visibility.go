package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/visibility"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func VisibilityToProto(e main_entity.Visibility) *pb.Visibility {
	return &pb.Visibility{
		Uuid:              e.UUID.String(),
		Description:       StringPtrToString(e.Description),
		OrganizationUuids: UUIDSliceToStringSlice(e.OrganizationUUIDs),
		TeamUuids:         UUIDSliceToStringSlice(e.TeamUUIDs),
		UserUuids:         UUIDSliceToStringSlice(e.UserUUIDs),
		Roles:             VisibilityRoleSliceToProto(e.Roles),
		CreatedAt:         timestamppb.New(e.CreatedAt),
		UpdatedAt:         timestamppb.New(e.UpdatedAt),
		CreatedByUuid:     e.CreatedByUUID.String(),
		UpdatedByUuid:     e.UpdatedByUUID.String(),
	}
}

func VisibilitySliceToProto(es []main_entity.Visibility) []*pb.Visibility {
	res := []*pb.Visibility{}
	for _, e := range es {
		res = append(res, VisibilityToProto(e))
	}
	return res
}

func VisibilityFromProto(m *pb.Visibility) main_entity.Visibility {
	if m == nil {
		return main_entity.Visibility{}
	}
	return main_entity.Visibility{
		UUID:              uuid.FromStringOrNil(m.GetUuid()),
		Description:       &m.Description,
		OrganizationUUIDs: StringSliceToUUIDSlice(m.GetOrganizationUuids()),
		TeamUUIDs:         StringSliceToUUIDSlice(m.GetTeamUuids()),
		UserUUIDs:         StringSliceToUUIDSlice(m.GetUserUuids()),
		Roles:             VisibilityRoleSliceFromProto(m.GetRoles()),
		CreatedAt:         m.GetCreatedAt().AsTime(),
		UpdatedAt:         m.GetUpdatedAt().AsTime(),
		CreatedByUUID:     uuid.FromStringOrNil(m.GetCreatedByUuid()),
		UpdatedByUUID:     uuid.FromStringOrNil(m.GetUpdatedByUuid()),
	}
}

func VisibilitySliceFromProto(es []*pb.Visibility) []main_entity.Visibility {
	if es == nil {
		return []main_entity.Visibility{}
	}
	res := []main_entity.Visibility{}
	for _, e := range es {
		res = append(res, VisibilityFromProto(e))
	}
	return res
}

func VisibilityRoleSliceToProto(s []main_entity.Role) []pb.VisibilityRole {
	res := []pb.VisibilityRole{}
	for _, e := range s {
		res = append(res, pb.VisibilityRole(e))
	}
	return res
}
func VisibilityRoleSliceFromProto(s []pb.VisibilityRole) []main_entity.Role {
	res := []main_entity.Role{}
	for _, e := range s {
		res = append(res, main_entity.Role(e))
	}
	return res
}
