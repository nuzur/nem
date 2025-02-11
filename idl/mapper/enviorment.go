package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/enviorment"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func EnviormentToProto(e main_entity.Enviorment) *pb.Enviorment {
	return &pb.Enviorment{
		Uuid:          e.UUID.String(),
		Identifier:    e.Identifier,
		Critical:      e.Critical,
		Status:        pb.EnviormentStatus(e.Status),
		CreatedAt:     timestamppb.New(e.CreatedAt),
		UpdatedAt:     timestamppb.New(e.UpdatedAt),
		CreatedByUuid: e.CreatedByUUID.String(),
		UpdatedByUuid: e.UpdatedByUUID.String(),
	}
}

func EnviormentSliceToProto(es []main_entity.Enviorment) []*pb.Enviorment {
	res := []*pb.Enviorment{}
	for _, e := range es {
		res = append(res, EnviormentToProto(e))
	}
	return res
}

func EnviormentFromProto(m *pb.Enviorment) main_entity.Enviorment {
	if m == nil {
		return main_entity.Enviorment{}
	}
	return main_entity.Enviorment{
		UUID:          uuid.FromStringOrNil(m.GetUuid()),
		Identifier:    m.GetIdentifier(),
		Critical:      m.GetCritical(),
		Status:        main_entity.Status(m.GetStatus()),
		CreatedAt:     m.GetCreatedAt().AsTime(),
		UpdatedAt:     m.GetUpdatedAt().AsTime(),
		CreatedByUUID: uuid.FromStringOrNil(m.GetCreatedByUuid()),
		UpdatedByUUID: uuid.FromStringOrNil(m.GetUpdatedByUuid()),
	}
}

func EnviormentSliceFromProto(es []*pb.Enviorment) []main_entity.Enviorment {
	if es == nil {
		return []main_entity.Enviorment{}
	}
	res := []main_entity.Enviorment{}
	for _, e := range es {
		res = append(res, EnviormentFromProto(e))
	}
	return res
}
