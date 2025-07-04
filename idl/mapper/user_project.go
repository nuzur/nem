package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/user_project"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func UserProjectToProto(e main_entity.UserProject) *pb.UserProject {
	return &pb.UserProject{
		Uuid:                    e.UUID.String(),
		UserUuid:                e.UserUUID.String(),
		UserEmail:               StringPtrToString(e.UserEmail),
		ProjectUuid:             e.ProjectUUID.String(),
		Role:                    pb.UserProjectRole(e.Role),
		ReviewRequiredStructure: e.ReviewRequiredStructure,
		ReviewRequiredData:      e.ReviewRequiredData,
		Status:                  pb.UserProjectStatus(e.Status),
		CreatedAt:               timestamppb.New(e.CreatedAt),
		UpdatedAt:               timestamppb.New(e.UpdatedAt),
		CreatedByUuid:           e.CreatedByUUID.String(),
		UpdatedByUuid:           e.UpdatedByUUID.String(),
	}
}

func UserProjectSliceToProto(es []main_entity.UserProject) []*pb.UserProject {
	res := []*pb.UserProject{}
	for _, e := range es {
		res = append(res, UserProjectToProto(e))
	}
	return res
}

func UserProjectFromProto(m *pb.UserProject) main_entity.UserProject {
	if m == nil {
		return main_entity.UserProject{}
	}
	return main_entity.UserProject{
		UUID:                    uuid.FromStringOrNil(m.GetUuid()),
		UserUUID:                uuid.FromStringOrNil(m.GetUserUuid()),
		UserEmail:               &m.UserEmail,
		ProjectUUID:             uuid.FromStringOrNil(m.GetProjectUuid()),
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

func UserProjectSliceFromProto(es []*pb.UserProject) []main_entity.UserProject {
	if es == nil {
		return []main_entity.UserProject{}
	}
	res := []main_entity.UserProject{}
	for _, e := range es {
		res = append(res, UserProjectFromProto(e))
	}
	return res
}
