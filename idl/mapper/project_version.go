package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/project_version"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ProjectVersionToProto(e main_entity.ProjectVersion) *pb.ProjectVersion {
	return &pb.ProjectVersion{
		Uuid:            e.UUID.String(),
		Version:         int64(e.Version),
		Identifier:      e.Identifier,
		Description:     e.Description,
		ProjectUuid:     e.ProjectUUID.String(),
		Entities:        EntitySliceToProto(e.Entities),
		Relationships:   RelationshipSliceToProto(e.Relationships),
		Enums:           EnumSliceToProto(e.Enums),
		Services:        ServiceSliceToProto(e.Services),
		BaseVersionUuid: e.BaseVersionUUID.String(),
		ReviewStatus:    pb.ProjectVersionReviewStatus(e.ReviewStatus),
		Reviews:         ProjectVersionReviewSliceToProto(e.Reviews),
		Deployments:     ProjectVersionDeploymentSliceToProto(e.Deployments),
		Status:          pb.ProjectVersionStatus(e.Status),
		CreatedAt:       timestamppb.New(e.CreatedAt),
		UpdatedAt:       timestamppb.New(e.UpdatedAt),
		CreatedByUuid:   e.CreatedByUUID.String(),
		UpdatedByUuid:   e.UpdatedByUUID.String(),
	}
}

func ProjectVersionSliceToProto(es []main_entity.ProjectVersion) []*pb.ProjectVersion {
	res := []*pb.ProjectVersion{}
	for _, e := range es {
		res = append(res, ProjectVersionToProto(e))
	}
	return res
}

func ProjectVersionFromProto(m *pb.ProjectVersion) main_entity.ProjectVersion {
	if m == nil {
		return main_entity.ProjectVersion{}
	}
	return main_entity.ProjectVersion{
		UUID:            uuid.FromStringOrNil(m.GetUuid()),
		Version:         int64(m.GetVersion()),
		Identifier:      m.GetIdentifier(),
		Description:     m.GetDescription(),
		ProjectUUID:     uuid.FromStringOrNil(m.GetProjectUuid()),
		Entities:        EntitySliceFromProto(m.GetEntities()),
		Relationships:   RelationshipSliceFromProto(m.GetRelationships()),
		Enums:           EnumSliceFromProto(m.GetEnums()),
		Services:        ServiceSliceFromProto(m.GetServices()),
		BaseVersionUUID: uuid.FromStringOrNil(m.GetBaseVersionUuid()),
		ReviewStatus:    main_entity.ReviewStatus(m.GetReviewStatus()),
		Reviews:         ProjectVersionReviewSliceFromProto(m.GetReviews()),
		Deployments:     ProjectVersionDeploymentSliceFromProto(m.GetDeployments()),
		Status:          main_entity.Status(m.GetStatus()),
		CreatedAt:       m.GetCreatedAt().AsTime(),
		UpdatedAt:       m.GetUpdatedAt().AsTime(),
		CreatedByUUID:   uuid.FromStringOrNil(m.GetCreatedByUuid()),
		UpdatedByUUID:   uuid.FromStringOrNil(m.GetUpdatedByUuid()),
	}
}

func ProjectVersionSliceFromProto(es []*pb.ProjectVersion) []main_entity.ProjectVersion {
	if es == nil {
		return []main_entity.ProjectVersion{}
	}
	res := []main_entity.ProjectVersion{}
	for _, e := range es {
		res = append(res, ProjectVersionFromProto(e))
	}
	return res
}
