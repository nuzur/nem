package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/project_version_review"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ProjectVersionReviewToProto(e main_entity.ProjectVersionReview) *pb.ProjectVersionReview {
	return &pb.ProjectVersionReview{
		Uuid:      e.UUID.String(),
		UserUuid:  e.UserUUID.String(),
		Comment:   StringPtrToString(e.Comment),
		Response:  pb.ProjectVersionReviewResponse(e.Response),
		CreatedAt: timestamppb.New(e.CreatedAt),
		UpdatedAt: timestamppb.New(e.UpdatedAt),
	}
}

func ProjectVersionReviewSliceToProto(es []main_entity.ProjectVersionReview) []*pb.ProjectVersionReview {
	res := []*pb.ProjectVersionReview{}
	for _, e := range es {
		res = append(res, ProjectVersionReviewToProto(e))
	}
	return res
}

func ProjectVersionReviewFromProto(m *pb.ProjectVersionReview) main_entity.ProjectVersionReview {
	if m == nil {
		return main_entity.ProjectVersionReview{}
	}
	return main_entity.ProjectVersionReview{
		UUID:      uuid.FromStringOrNil(m.GetUuid()),
		UserUUID:  uuid.FromStringOrNil(m.GetUserUuid()),
		Comment:   &m.Comment,
		Response:  main_entity.Response(m.GetResponse()),
		CreatedAt: m.GetCreatedAt().AsTime(),
		UpdatedAt: m.GetUpdatedAt().AsTime(),
	}
}

func ProjectVersionReviewSliceFromProto(es []*pb.ProjectVersionReview) []main_entity.ProjectVersionReview {
	if es == nil {
		return []main_entity.ProjectVersionReview{}
	}
	res := []main_entity.ProjectVersionReview{}
	for _, e := range es {
		res = append(res, ProjectVersionReviewFromProto(e))
	}
	return res
}
