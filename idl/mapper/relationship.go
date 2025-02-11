package mapper

import (
	main_entity "nem/core/entity/relationship"
	pb "nem/idl/gen"

	"github.com/gofrs/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func RelationshipToProto(e main_entity.Relationship) *pb.Relationship {
	return &pb.Relationship{
		Uuid:          e.UUID.String(),
		Version:       int64(e.Version),
		Identifier:    e.Identifier,
		Description:   e.Description,
		Cardinality:   pb.RelationshipCardinality(e.Cardinality),
		From:          RelationshipNodeToProto(e.From),
		To:            RelationshipNodeToProto(e.To),
		UseForeignKey: e.UseForeignKey,
		Status:        pb.RelationshipStatus(e.Status),
		CreatedAt:     timestamppb.New(e.CreatedAt),
		UpdatedAt:     timestamppb.New(e.UpdatedAt),
		CreatedByUuid: e.CreatedByUUID.String(),
		UpdatedByUuid: e.UpdatedByUUID.String(),
	}
}

func RelationshipSliceToProto(es []main_entity.Relationship) []*pb.Relationship {
	res := []*pb.Relationship{}
	for _, e := range es {
		res = append(res, RelationshipToProto(e))
	}
	return res
}

func RelationshipFromProto(m *pb.Relationship) main_entity.Relationship {
	if m == nil {
		return main_entity.Relationship{}
	}
	return main_entity.Relationship{
		UUID:          uuid.FromStringOrNil(m.GetUuid()),
		Version:       int64(m.GetVersion()),
		Identifier:    m.GetIdentifier(),
		Description:   m.GetDescription(),
		Cardinality:   main_entity.Cardinality(m.GetCardinality()),
		From:          RelationshipNodeFromProto(m.GetFrom()),
		To:            RelationshipNodeFromProto(m.GetTo()),
		UseForeignKey: m.GetUseForeignKey(),
		Status:        main_entity.Status(m.GetStatus()),
		CreatedAt:     m.GetCreatedAt().AsTime(),
		UpdatedAt:     m.GetUpdatedAt().AsTime(),
		CreatedByUUID: uuid.FromStringOrNil(m.GetCreatedByUuid()),
		UpdatedByUUID: uuid.FromStringOrNil(m.GetUpdatedByUuid()),
	}
}

func RelationshipSliceFromProto(es []*pb.Relationship) []main_entity.Relationship {
	if es == nil {
		return []main_entity.Relationship{}
	}
	res := []main_entity.Relationship{}
	for _, e := range es {
		res = append(res, RelationshipFromProto(e))
	}
	return res
}
