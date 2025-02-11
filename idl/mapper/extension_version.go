package mapper

import (
	main_entity "nem/core/entity/extension_version"
	pb "nem/idl/gen"

	"github.com/gofrs/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ExtensionVersionToProto(e main_entity.ExtensionVersion) *pb.ExtensionVersion {
	return &pb.ExtensionVersion{
		Uuid:                e.UUID.String(),
		Version:             int64(e.Version),
		ExtensionUuid:       e.ExtensionUUID.String(),
		DisplayVersion:      e.DisplayVersion,
		Description:         e.Description,
		RepositoryTag:       e.RepositoryTag,
		ConfigurationEntity: e.ConfigurationEntity,
		ExecutionMode:       ExtensionVersionExecutionModeSliceToProto(e.ExecutionMode),
		ReviewStatus:        pb.ExtensionVersionReviewStatus(e.ReviewStatus),
		Status:              pb.ExtensionVersionStatus(e.Status),
		CreatedAt:           timestamppb.New(e.CreatedAt),
		UpdatedAt:           timestamppb.New(e.UpdatedAt),
		CreatedByUuid:       e.CreatedByUUID.String(),
		UpdatedByUuid:       e.UpdatedByUUID.String(),
	}
}

func ExtensionVersionSliceToProto(es []main_entity.ExtensionVersion) []*pb.ExtensionVersion {
	res := []*pb.ExtensionVersion{}
	for _, e := range es {
		res = append(res, ExtensionVersionToProto(e))
	}
	return res
}

func ExtensionVersionFromProto(m *pb.ExtensionVersion) main_entity.ExtensionVersion {
	if m == nil {
		return main_entity.ExtensionVersion{}
	}
	return main_entity.ExtensionVersion{
		UUID:                uuid.FromStringOrNil(m.GetUuid()),
		Version:             int64(m.GetVersion()),
		ExtensionUUID:       uuid.FromStringOrNil(m.GetExtensionUuid()),
		DisplayVersion:      m.GetDisplayVersion(),
		Description:         m.GetDescription(),
		RepositoryTag:       m.GetRepositoryTag(),
		ConfigurationEntity: m.GetConfigurationEntity(),
		ExecutionMode:       ExtensionVersionExecutionModeSliceFromProto(m.GetExecutionMode()),
		ReviewStatus:        main_entity.ReviewStatus(m.GetReviewStatus()),
		Status:              main_entity.Status(m.GetStatus()),
		CreatedAt:           m.GetCreatedAt().AsTime(),
		UpdatedAt:           m.GetUpdatedAt().AsTime(),
		CreatedByUUID:       uuid.FromStringOrNil(m.GetCreatedByUuid()),
		UpdatedByUUID:       uuid.FromStringOrNil(m.GetUpdatedByUuid()),
	}
}

func ExtensionVersionSliceFromProto(es []*pb.ExtensionVersion) []main_entity.ExtensionVersion {
	if es == nil {
		return []main_entity.ExtensionVersion{}
	}
	res := []main_entity.ExtensionVersion{}
	for _, e := range es {
		res = append(res, ExtensionVersionFromProto(e))
	}
	return res
}

func ExtensionVersionExecutionModeSliceToProto(s []main_entity.ExecutionMode) []pb.ExtensionVersionExecutionMode {
	res := []pb.ExtensionVersionExecutionMode{}
	for _, e := range s {
		res = append(res, pb.ExtensionVersionExecutionMode(e))
	}
	return res
}
func ExtensionVersionExecutionModeSliceFromProto(s []pb.ExtensionVersionExecutionMode) []main_entity.ExecutionMode {
	res := []main_entity.ExecutionMode{}
	for _, e := range s {
		res = append(res, main_entity.ExecutionMode(e))
	}
	return res
}
