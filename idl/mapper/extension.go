package mapper

import (
	main_entity "nem/core/entity/extension"
	pb "nem/idl/gen"

	"github.com/gofrs/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ExtensionToProto(e main_entity.Extension) *pb.Extension {
	return &pb.Extension{
		Uuid:              e.UUID.String(),
		Version:           int64(e.Version),
		Identifier:        e.Identifier,
		DisplayName:       e.DisplayName,
		DisplayAuthorName: e.DisplayAuthorName,
		Description:       e.Description,
		Url:               e.URL,
		Verfied:           e.Verfied,
		Repository:        e.Repository,
		ExtensionType:     pb.ExtensionExtensionType(e.ExtensionType),
		Tags:              e.Tags,
		Public:            e.Public,
		Visibility:        VisibilitySliceToProto(e.Visibility),
		Status:            pb.ExtensionStatus(e.Status),
		OwnerUuid:         e.OwnerUUID.String(),
		CreatedAt:         timestamppb.New(e.CreatedAt),
		UpdatedAt:         timestamppb.New(e.UpdatedAt),
		CreatedByUuid:     e.CreatedByUUID.String(),
		UpdatedByUuid:     e.UpdatedByUUID.String(),
	}
}

func ExtensionSliceToProto(es []main_entity.Extension) []*pb.Extension {
	res := []*pb.Extension{}
	for _, e := range es {
		res = append(res, ExtensionToProto(e))
	}
	return res
}

func ExtensionFromProto(m *pb.Extension) main_entity.Extension {
	if m == nil {
		return main_entity.Extension{}
	}
	return main_entity.Extension{
		UUID:              uuid.FromStringOrNil(m.GetUuid()),
		Version:           int64(m.GetVersion()),
		Identifier:        m.GetIdentifier(),
		DisplayName:       m.GetDisplayName(),
		DisplayAuthorName: m.GetDisplayAuthorName(),
		Description:       m.GetDescription(),
		URL:               m.GetUrl(),
		Verfied:           m.GetVerfied(),
		Repository:        m.GetRepository(),
		ExtensionType:     main_entity.ExtensionType(m.GetExtensionType()),
		Tags:              m.GetTags(),
		Public:            m.GetPublic(),
		Visibility:        VisibilitySliceFromProto(m.GetVisibility()),
		Status:            main_entity.Status(m.GetStatus()),
		OwnerUUID:         uuid.FromStringOrNil(m.GetOwnerUuid()),
		CreatedAt:         m.GetCreatedAt().AsTime(),
		UpdatedAt:         m.GetUpdatedAt().AsTime(),
		CreatedByUUID:     uuid.FromStringOrNil(m.GetCreatedByUuid()),
		UpdatedByUUID:     uuid.FromStringOrNil(m.GetUpdatedByUuid()),
	}
}

func ExtensionSliceFromProto(es []*pb.Extension) []main_entity.Extension {
	if es == nil {
		return []main_entity.Extension{}
	}
	res := []main_entity.Extension{}
	for _, e := range es {
		res = append(res, ExtensionFromProto(e))
	}
	return res
}
