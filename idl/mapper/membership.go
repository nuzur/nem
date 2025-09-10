package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/membership"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func MembershipToProto(e main_entity.Membership) *pb.Membership {
	return &pb.Membership{
		Uuid:            e.UUID.String(),
		OwnerUuid:       e.OwnerUUID.String(),
		Type:            pb.MembershipType(e.Type),
		StartDate:       timestamppb.New(e.StartDate),
		BillingMetadata: e.BillingMetadata,
		Status:          pb.MembershipStatus(e.Status),
		CreatedAt:       timestamppb.New(e.CreatedAt),
		UpdatedAt:       timestamppb.New(e.UpdatedAt),
		CreatedByUuid:   e.CreatedByUUID.String(),
		UpdatedByUuid:   e.UpdatedByUUID.String(),
	}
}

func MembershipSliceToProto(es []main_entity.Membership) []*pb.Membership {
	res := []*pb.Membership{}
	for _, e := range es {
		res = append(res, MembershipToProto(e))
	}
	return res
}

func MembershipFromProto(m *pb.Membership) main_entity.Membership {
	if m == nil {
		return main_entity.Membership{}
	}
	return main_entity.Membership{
		UUID:            uuid.FromStringOrNil(m.GetUuid()),
		OwnerUUID:       uuid.FromStringOrNil(m.GetOwnerUuid()),
		Type:            main_entity.Type(m.GetType()),
		StartDate:       m.GetStartDate().AsTime(),
		BillingMetadata: m.GetBillingMetadata(),
		Status:          main_entity.Status(m.GetStatus()),
		CreatedAt:       m.GetCreatedAt().AsTime(),
		UpdatedAt:       m.GetUpdatedAt().AsTime(),
		CreatedByUUID:   uuid.FromStringOrNil(m.GetCreatedByUuid()),
		UpdatedByUUID:   uuid.FromStringOrNil(m.GetUpdatedByUuid()),
	}
}

func MembershipSliceFromProto(es []*pb.Membership) []main_entity.Membership {
	if es == nil {
		return []main_entity.Membership{}
	}
	res := []main_entity.Membership{}
	for _, e := range es {
		res = append(res, MembershipFromProto(e))
	}
	return res
}
