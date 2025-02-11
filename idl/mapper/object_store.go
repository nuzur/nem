package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/object_store"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ObjectStoreToProto(e main_entity.ObjectStore) *pb.ObjectStore {
	return &pb.ObjectStore{
		Uuid:          e.UUID.String(),
		Identifier:    e.Identifier,
		Type:          pb.ObjectStoreType(e.Type),
		TypeConfig:    ObjectStoreTypeConfigToProto(e.TypeConfig),
		Status:        pb.ObjectStoreStatus(e.Status),
		CreatedAt:     timestamppb.New(e.CreatedAt),
		UpdatedAt:     timestamppb.New(e.UpdatedAt),
		CreatedByUuid: e.CreatedByUUID.String(),
		UpdatedByUuid: e.UpdatedByUUID.String(),
	}
}

func ObjectStoreSliceToProto(es []main_entity.ObjectStore) []*pb.ObjectStore {
	res := []*pb.ObjectStore{}
	for _, e := range es {
		res = append(res, ObjectStoreToProto(e))
	}
	return res
}

func ObjectStoreFromProto(m *pb.ObjectStore) main_entity.ObjectStore {
	if m == nil {
		return main_entity.ObjectStore{}
	}
	return main_entity.ObjectStore{
		UUID:          uuid.FromStringOrNil(m.GetUuid()),
		Identifier:    m.GetIdentifier(),
		Type:          main_entity.Type(m.GetType()),
		TypeConfig:    ObjectStoreTypeConfigFromProto(m.GetTypeConfig()),
		Status:        main_entity.Status(m.GetStatus()),
		CreatedAt:     m.GetCreatedAt().AsTime(),
		UpdatedAt:     m.GetUpdatedAt().AsTime(),
		CreatedByUUID: uuid.FromStringOrNil(m.GetCreatedByUuid()),
		UpdatedByUUID: uuid.FromStringOrNil(m.GetUpdatedByUuid()),
	}
}

func ObjectStoreSliceFromProto(es []*pb.ObjectStore) []main_entity.ObjectStore {
	if es == nil {
		return []main_entity.ObjectStore{}
	}
	res := []main_entity.ObjectStore{}
	for _, e := range es {
		res = append(res, ObjectStoreFromProto(e))
	}
	return res
}
