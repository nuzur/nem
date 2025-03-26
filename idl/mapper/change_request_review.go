package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/change_request_review"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ChangeRequestReviewToProto(e main_entity.ChangeRequestReview) *pb.ChangeRequestReview {
	return &pb.ChangeRequestReview{
		Uuid:      e.UUID.String(),
		UserUuid:  e.UserUUID.String(),
		Comment:   StringPtrToString(e.Comment),
		Response:  pb.ChangeRequestReviewResponse(e.Response),
		CreatedAt: timestamppb.New(e.CreatedAt),
		UpdatedAt: timestamppb.New(e.UpdatedAt),
	}
}

func ChangeRequestReviewSliceToProto(es []main_entity.ChangeRequestReview) []*pb.ChangeRequestReview {
	res := []*pb.ChangeRequestReview{}
	for _, e := range es {
		res = append(res, ChangeRequestReviewToProto(e))
	}
	return res
}

func ChangeRequestReviewFromProto(m *pb.ChangeRequestReview) main_entity.ChangeRequestReview {
	if m == nil {
		return main_entity.ChangeRequestReview{}
	}
	return main_entity.ChangeRequestReview{
		UUID:      uuid.FromStringOrNil(m.GetUuid()),
		UserUUID:  uuid.FromStringOrNil(m.GetUserUuid()),
		Comment:   &m.Comment,
		Response:  main_entity.Response(m.GetResponse()),
		CreatedAt: m.GetCreatedAt().AsTime(),
		UpdatedAt: m.GetUpdatedAt().AsTime(),
	}
}

func ChangeRequestReviewSliceFromProto(es []*pb.ChangeRequestReview) []main_entity.ChangeRequestReview {
	if es == nil {
		return []main_entity.ChangeRequestReview{}
	}
	res := []main_entity.ChangeRequestReview{}
	for _, e := range es {
		res = append(res, ChangeRequestReviewFromProto(e))
	}
	return res
}
