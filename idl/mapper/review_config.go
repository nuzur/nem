package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/review_config"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ReviewConfigToProto(e main_entity.ReviewConfig) *pb.ReviewConfig {
	return &pb.ReviewConfig{
		Uuid:          e.UUID.String(),
		Types:         ReviewConfigTypeSliceToProto(e.Types),
		UserRoles:     ReviewConfigUserRoleSliceToProto(e.UserRoles),
		MinReviews:    int64(e.MinReviews),
		Status:        pb.ReviewConfigStatus(e.Status),
		CreatedAt:     timestamppb.New(e.CreatedAt),
		UpdatedAt:     timestamppb.New(e.UpdatedAt),
		CreatedByUuid: e.CreatedByUUID.String(),
		UpdatedByUuid: e.UpdatedByUUID.String(),
	}
}

func ReviewConfigSliceToProto(es []main_entity.ReviewConfig) []*pb.ReviewConfig {
	res := []*pb.ReviewConfig{}
	for _, e := range es {
		res = append(res, ReviewConfigToProto(e))
	}
	return res
}

func ReviewConfigFromProto(m *pb.ReviewConfig) main_entity.ReviewConfig {
	if m == nil {
		return main_entity.ReviewConfig{}
	}
	return main_entity.ReviewConfig{
		UUID:          uuid.FromStringOrNil(m.GetUuid()),
		Types:         ReviewConfigTypeSliceFromProto(m.GetTypes()),
		UserRoles:     ReviewConfigUserRoleSliceFromProto(m.GetUserRoles()),
		MinReviews:    int64(m.GetMinReviews()),
		Status:        main_entity.Status(m.GetStatus()),
		CreatedAt:     m.GetCreatedAt().AsTime(),
		UpdatedAt:     m.GetUpdatedAt().AsTime(),
		CreatedByUUID: uuid.FromStringOrNil(m.GetCreatedByUuid()),
		UpdatedByUUID: uuid.FromStringOrNil(m.GetUpdatedByUuid()),
	}
}

func ReviewConfigSliceFromProto(es []*pb.ReviewConfig) []main_entity.ReviewConfig {
	if es == nil {
		return []main_entity.ReviewConfig{}
	}
	res := []main_entity.ReviewConfig{}
	for _, e := range es {
		res = append(res, ReviewConfigFromProto(e))
	}
	return res
}

func ReviewConfigTypeSliceToProto(s []main_entity.Type) []pb.ReviewConfigType {
	res := []pb.ReviewConfigType{}
	for _, e := range s {
		res = append(res, pb.ReviewConfigType(e))
	}
	return res
}
func ReviewConfigTypeSliceFromProto(s []pb.ReviewConfigType) []main_entity.Type {
	res := []main_entity.Type{}
	for _, e := range s {
		res = append(res, main_entity.Type(e))
	}
	return res
}

func ReviewConfigUserRoleSliceToProto(s []main_entity.UserRole) []pb.ReviewConfigUserRole {
	res := []pb.ReviewConfigUserRole{}
	for _, e := range s {
		res = append(res, pb.ReviewConfigUserRole(e))
	}
	return res
}
func ReviewConfigUserRoleSliceFromProto(s []pb.ReviewConfigUserRole) []main_entity.UserRole {
	res := []main_entity.UserRole{}
	for _, e := range s {
		res = append(res, main_entity.UserRole(e))
	}
	return res
}
