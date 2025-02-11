package mapper

import (
	main_entity "nem/core/entity/enum"
	pb "nem/idl/gen"

	"github.com/gofrs/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func EnumToProto(e main_entity.Enum) *pb.Enum {
	return &pb.Enum{
		Uuid:          e.UUID.String(),
		Version:       int64(e.Version),
		Identifier:    e.Identifier,
		StaticValues:  EnumValueSliceToProto(e.StaticValues),
		RemoteValues:  e.RemoteValues,
		Status:        pb.EnumStatus(e.Status),
		CreatedAt:     timestamppb.New(e.CreatedAt),
		UpdatedAt:     timestamppb.New(e.UpdatedAt),
		CreatedByUuid: e.CreatedByUUID.String(),
		UpdatedByUuid: e.UpdatedByUUID.String(),
	}
}

func EnumSliceToProto(es []main_entity.Enum) []*pb.Enum {
	res := []*pb.Enum{}
	for _, e := range es {
		res = append(res, EnumToProto(e))
	}
	return res
}

func EnumFromProto(m *pb.Enum) main_entity.Enum {
	if m == nil {
		return main_entity.Enum{}
	}
	return main_entity.Enum{
		UUID:          uuid.FromStringOrNil(m.GetUuid()),
		Version:       int64(m.GetVersion()),
		Identifier:    m.GetIdentifier(),
		StaticValues:  EnumValueSliceFromProto(m.GetStaticValues()),
		RemoteValues:  m.GetRemoteValues(),
		Status:        main_entity.Status(m.GetStatus()),
		CreatedAt:     m.GetCreatedAt().AsTime(),
		UpdatedAt:     m.GetUpdatedAt().AsTime(),
		CreatedByUUID: uuid.FromStringOrNil(m.GetCreatedByUuid()),
		UpdatedByUUID: uuid.FromStringOrNil(m.GetUpdatedByUuid()),
	}
}

func EnumSliceFromProto(es []*pb.Enum) []main_entity.Enum {
	if es == nil {
		return []main_entity.Enum{}
	}
	res := []main_entity.Enum{}
	for _, e := range es {
		res = append(res, EnumFromProto(e))
	}
	return res
}
