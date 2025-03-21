package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/store"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func StoreToProto(e main_entity.Store) *pb.Store {
	return &pb.Store{
		Uuid:          e.UUID.String(),
		Identifier:    e.Identifier,
		Description:   StringPtrToString(e.Description),
		Status:        pb.StoreStatus(e.Status),
		CreatedAt:     timestamppb.New(e.CreatedAt),
		UpdatedAt:     timestamppb.New(e.UpdatedAt),
		CreatedByUuid: e.CreatedByUUID.String(),
		UpdatedByUuid: e.UpdatedByUUID.String(),
	}
}

func StoreSliceToProto(es []main_entity.Store) []*pb.Store {
	res := []*pb.Store{}
	for _, e := range es {
		res = append(res, StoreToProto(e))
	}
	return res
}

func StoreFromProto(m *pb.Store) main_entity.Store {
	if m == nil {
		return main_entity.Store{}
	}
	return main_entity.Store{
		UUID:          uuid.FromStringOrNil(m.GetUuid()),
		Identifier:    m.GetIdentifier(),
		Description:   &m.Description,
		Status:        main_entity.Status(m.GetStatus()),
		CreatedAt:     m.GetCreatedAt().AsTime(),
		UpdatedAt:     m.GetUpdatedAt().AsTime(),
		CreatedByUUID: uuid.FromStringOrNil(m.GetCreatedByUuid()),
		UpdatedByUUID: uuid.FromStringOrNil(m.GetUpdatedByUuid()),
	}
}

func StoreSliceFromProto(es []*pb.Store) []main_entity.Store {
	if es == nil {
		return []main_entity.Store{}
	}
	res := []main_entity.Store{}
	for _, e := range es {
		res = append(res, StoreFromProto(e))
	}
	return res
}
