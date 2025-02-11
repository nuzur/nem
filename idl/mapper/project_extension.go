package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/project_extension"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ProjectExtensionToProto(e main_entity.ProjectExtension) *pb.ProjectExtension {
	return &pb.ProjectExtension{
		Uuid:                      e.UUID.String(),
		ExtensionUuid:             e.ExtensionUUID.String(),
		ExtensionVersionUuid:      e.ExtensionVersionUUID.String(),
		ConfigurationEntityValues: e.ConfigurationEntityValues,
		Blocking:                  e.Blocking,
		Status:                    pb.ProjectExtensionStatus(e.Status),
		CreatedAt:                 timestamppb.New(e.CreatedAt),
		UpdatedAt:                 timestamppb.New(e.UpdatedAt),
		CreatedByUuid:             e.CreatedByUUID.String(),
		UpdatedByUuid:             e.UpdatedByUUID.String(),
	}
}

func ProjectExtensionSliceToProto(es []main_entity.ProjectExtension) []*pb.ProjectExtension {
	res := []*pb.ProjectExtension{}
	for _, e := range es {
		res = append(res, ProjectExtensionToProto(e))
	}
	return res
}

func ProjectExtensionFromProto(m *pb.ProjectExtension) main_entity.ProjectExtension {
	if m == nil {
		return main_entity.ProjectExtension{}
	}
	return main_entity.ProjectExtension{
		UUID:                      uuid.FromStringOrNil(m.GetUuid()),
		ExtensionUUID:             uuid.FromStringOrNil(m.GetExtensionUuid()),
		ExtensionVersionUUID:      uuid.FromStringOrNil(m.GetExtensionVersionUuid()),
		ConfigurationEntityValues: m.GetConfigurationEntityValues(),
		Blocking:                  m.GetBlocking(),
		Status:                    main_entity.Status(m.GetStatus()),
		CreatedAt:                 m.GetCreatedAt().AsTime(),
		UpdatedAt:                 m.GetUpdatedAt().AsTime(),
		CreatedByUUID:             uuid.FromStringOrNil(m.GetCreatedByUuid()),
		UpdatedByUUID:             uuid.FromStringOrNil(m.GetUpdatedByUuid()),
	}
}

func ProjectExtensionSliceFromProto(es []*pb.ProjectExtension) []main_entity.ProjectExtension {
	if es == nil {
		return []main_entity.ProjectExtension{}
	}
	res := []main_entity.ProjectExtension{}
	for _, e := range es {
		res = append(res, ProjectExtensionFromProto(e))
	}
	return res
}
