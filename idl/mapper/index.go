package mapper

import (
	main_entity "nem/core/entity/index"
	pb "nem/idl/gen"

	"github.com/gofrs/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func IndexToProto(e main_entity.Index) *pb.Index {
	return &pb.Index{
		Uuid:          e.UUID.String(),
		Identifier:    e.Identifier,
		Type:          pb.IndexType(e.Type),
		Fields:        IndexFieldSliceToProto(e.Fields),
		Status:        pb.IndexStatus(e.Status),
		CreatedAt:     timestamppb.New(e.CreatedAt),
		UpdatedAt:     timestamppb.New(e.UpdatedAt),
		CreatedByUuid: e.CreatedByUUID.String(),
		UpdatedByUuid: e.UpdatedByUUID.String(),
	}
}

func IndexSliceToProto(es []main_entity.Index) []*pb.Index {
	res := []*pb.Index{}
	for _, e := range es {
		res = append(res, IndexToProto(e))
	}
	return res
}

func IndexFromProto(m *pb.Index) main_entity.Index {
	if m == nil {
		return main_entity.Index{}
	}
	return main_entity.Index{
		UUID:          uuid.FromStringOrNil(m.GetUuid()),
		Identifier:    m.GetIdentifier(),
		Type:          main_entity.Type(m.GetType()),
		Fields:        IndexFieldSliceFromProto(m.GetFields()),
		Status:        main_entity.Status(m.GetStatus()),
		CreatedAt:     m.GetCreatedAt().AsTime(),
		UpdatedAt:     m.GetUpdatedAt().AsTime(),
		CreatedByUUID: uuid.FromStringOrNil(m.GetCreatedByUuid()),
		UpdatedByUUID: uuid.FromStringOrNil(m.GetUpdatedByUuid()),
	}
}

func IndexSliceFromProto(es []*pb.Index) []main_entity.Index {
	if es == nil {
		return []main_entity.Index{}
	}
	res := []main_entity.Index{}
	for _, e := range es {
		res = append(res, IndexFromProto(e))
	}
	return res
}
