package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/review_config"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ReviewConfigToProto(e main_entity.ReviewConfig) *pb.ReviewConfig {
	return &pb.ReviewConfig{
		Uuid:            e.UUID.String(),
		ReviewUserRoles: ReviewConfigReviewUserRoleSliceToProto(e.ReviewUserRoles),
		ReviewUserUuids: UUIDSliceToStringSlice(e.ReviewUserUUIDs),
		MinReviews:      int64(e.MinReviews),
		Status:          pb.ReviewConfigStatus(e.Status),
		CreatedAt:       timestamppb.New(e.CreatedAt),
		UpdatedAt:       timestamppb.New(e.UpdatedAt),
		CreatedByUuid:   e.CreatedByUUID.String(),
		UpdatedByUuid:   e.UpdatedByUUID.String(),
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
		UUID:            uuid.FromStringOrNil(m.GetUuid()),
		ReviewUserRoles: ReviewConfigReviewUserRoleSliceFromProto(m.GetReviewUserRoles()),
		ReviewUserUUIDs: StringSliceToUUIDSlice(m.GetReviewUserUuids()),
		MinReviews:      int64(m.GetMinReviews()),
		Status:          main_entity.Status(m.GetStatus()),
		CreatedAt:       m.GetCreatedAt().AsTime(),
		UpdatedAt:       m.GetUpdatedAt().AsTime(),
		CreatedByUUID:   uuid.FromStringOrNil(m.GetCreatedByUuid()),
		UpdatedByUUID:   uuid.FromStringOrNil(m.GetUpdatedByUuid()),
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

func ReviewConfigReviewUserRoleSliceToProto(s []main_entity.ReviewUserRole) []pb.ReviewConfigReviewUserRole {
	res := []pb.ReviewConfigReviewUserRole{}
	for _, e := range s {
		res = append(res, pb.ReviewConfigReviewUserRole(e))
	}
	return res
}
func ReviewConfigReviewUserRoleSliceFromProto(s []pb.ReviewConfigReviewUserRole) []main_entity.ReviewUserRole {
	res := []main_entity.ReviewUserRole{}
	for _, e := range s {
		res = append(res, main_entity.ReviewUserRole(e))
	}
	return res
}
