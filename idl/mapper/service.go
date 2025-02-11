package mapper

import (
	main_entity "nem/core/entity/service"
	pb "nem/idl/gen"

	"github.com/gofrs/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ServiceToProto(e main_entity.Service) *pb.Service {
	return &pb.Service{
		Uuid:          e.UUID.String(),
		Version:       int64(e.Version),
		Identifier:    e.Identifier,
		Description:   e.Description,
		Render:        ElementRenderToProto(e.Render),
		Status:        pb.ServiceStatus(e.Status),
		CreatedAt:     timestamppb.New(e.CreatedAt),
		UpdatedAt:     timestamppb.New(e.UpdatedAt),
		CreatedByUuid: e.CreatedByUUID.String(),
		UpdatedByUuid: e.UpdatedByUUID.String(),
	}
}

func ServiceSliceToProto(es []main_entity.Service) []*pb.Service {
	res := []*pb.Service{}
	for _, e := range es {
		res = append(res, ServiceToProto(e))
	}
	return res
}

func ServiceFromProto(m *pb.Service) main_entity.Service {
	if m == nil {
		return main_entity.Service{}
	}
	return main_entity.Service{
		UUID:          uuid.FromStringOrNil(m.GetUuid()),
		Version:       int64(m.GetVersion()),
		Identifier:    m.GetIdentifier(),
		Description:   m.GetDescription(),
		Render:        ElementRenderFromProto(m.GetRender()),
		Status:        main_entity.Status(m.GetStatus()),
		CreatedAt:     m.GetCreatedAt().AsTime(),
		UpdatedAt:     m.GetUpdatedAt().AsTime(),
		CreatedByUUID: uuid.FromStringOrNil(m.GetCreatedByUuid()),
		UpdatedByUUID: uuid.FromStringOrNil(m.GetUpdatedByUuid()),
	}
}

func ServiceSliceFromProto(es []*pb.Service) []main_entity.Service {
	if es == nil {
		return []main_entity.Service{}
	}
	res := []main_entity.Service{}
	for _, e := range es {
		res = append(res, ServiceFromProto(e))
	}
	return res
}
