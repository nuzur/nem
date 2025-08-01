package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/project"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ProjectToProto(e main_entity.Project) *pb.Project {
	return &pb.Project{
		Uuid:              e.UUID.String(),
		Version:           int64(e.Version),
		Name:              e.Name,
		Description:       StringPtrToString(e.Description),
		Tags:              e.Tags,
		Url:               StringPtrToString(e.URL),
		OwnerUuid:         e.OwnerUUID.String(),
		TeamUuid:          e.TeamUUID.String(),
		AccessType:        pb.ProjectAccessType(e.AccessType),
		ProjectExtensions: ProjectExtensionSliceToProto(e.ProjectExtensions),
		Status:            pb.ProjectStatus(e.Status),
		CreatedAt:         timestamppb.New(e.CreatedAt),
		UpdatedAt:         timestamppb.New(e.UpdatedAt),
		CreatedByUuid:     e.CreatedByUUID.String(),
		UpdatedByUuid:     e.UpdatedByUUID.String(),
	}
}

func ProjectSliceToProto(es []main_entity.Project) []*pb.Project {
	res := []*pb.Project{}
	for _, e := range es {
		res = append(res, ProjectToProto(e))
	}
	return res
}

func ProjectFromProto(m *pb.Project) main_entity.Project {
	if m == nil {
		return main_entity.Project{}
	}
	return main_entity.Project{
		UUID:              uuid.FromStringOrNil(m.GetUuid()),
		Version:           int64(m.GetVersion()),
		Name:              m.GetName(),
		Description:       &m.Description,
		Tags:              m.GetTags(),
		URL:               &m.Url,
		OwnerUUID:         uuid.FromStringOrNil(m.GetOwnerUuid()),
		TeamUUID:          uuid.FromStringOrNil(m.GetTeamUuid()),
		AccessType:        main_entity.AccessType(m.GetAccessType()),
		ProjectExtensions: ProjectExtensionSliceFromProto(m.GetProjectExtensions()),
		Status:            main_entity.Status(m.GetStatus()),
		CreatedAt:         m.GetCreatedAt().AsTime(),
		UpdatedAt:         m.GetUpdatedAt().AsTime(),
		CreatedByUUID:     uuid.FromStringOrNil(m.GetCreatedByUuid()),
		UpdatedByUUID:     uuid.FromStringOrNil(m.GetUpdatedByUuid()),
	}
}

func ProjectSliceFromProto(es []*pb.Project) []main_entity.Project {
	if es == nil {
		return []main_entity.Project{}
	}
	res := []main_entity.Project{}
	for _, e := range es {
		res = append(res, ProjectFromProto(e))
	}
	return res
}
