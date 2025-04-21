package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/user_project_version"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func UserProjectVersionToProto(e main_entity.UserProjectVersion) *pb.UserProjectVersion {
	return &pb.UserProjectVersion{
		Uuid:               e.UUID.String(),
		Version:            int64(e.Version),
		ProjectVersionUuid: e.ProjectVersionUUID.String(),
		UserUuid:           e.UserUUID.String(),
		Data:               StringPtrToString(e.Data),
		Status:             pb.UserProjectVersionStatus(e.Status),
		CreatedAt:          timestamppb.New(e.CreatedAt),
		UpdatedAt:          timestamppb.New(e.UpdatedAt),
		CreatedByUuid:      e.CreatedByUUID.String(),
		UpdatedByUuid:      e.UpdatedByUUID.String(),
	}
}

func UserProjectVersionSliceToProto(es []main_entity.UserProjectVersion) []*pb.UserProjectVersion {
	res := []*pb.UserProjectVersion{}
	for _, e := range es {
		res = append(res, UserProjectVersionToProto(e))
	}
	return res
}

func UserProjectVersionFromProto(m *pb.UserProjectVersion) main_entity.UserProjectVersion {
	if m == nil {
		return main_entity.UserProjectVersion{}
	}
	return main_entity.UserProjectVersion{
		UUID:               uuid.FromStringOrNil(m.GetUuid()),
		Version:            int64(m.GetVersion()),
		ProjectVersionUUID: uuid.FromStringOrNil(m.GetProjectVersionUuid()),
		UserUUID:           uuid.FromStringOrNil(m.GetUserUuid()),
		Data:               &m.Data,
		Status:             main_entity.Status(m.GetStatus()),
		CreatedAt:          m.GetCreatedAt().AsTime(),
		UpdatedAt:          m.GetUpdatedAt().AsTime(),
		CreatedByUUID:      uuid.FromStringOrNil(m.GetCreatedByUuid()),
		UpdatedByUUID:      uuid.FromStringOrNil(m.GetUpdatedByUuid()),
	}
}

func UserProjectVersionSliceFromProto(es []*pb.UserProjectVersion) []main_entity.UserProjectVersion {
	if es == nil {
		return []main_entity.UserProjectVersion{}
	}
	res := []main_entity.UserProjectVersion{}
	for _, e := range es {
		res = append(res, UserProjectVersionFromProto(e))
	}
	return res
}
