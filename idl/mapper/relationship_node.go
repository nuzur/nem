package mapper

import (
	main_entity "nem/core/entity/relationship_node"
	pb "nem/idl/gen"

	"github.com/gofrs/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func RelationshipNodeToProto(e main_entity.RelationshipNode) *pb.RelationshipNode {
	return &pb.RelationshipNode{
		Uuid:          e.UUID.String(),
		Type:          pb.RelationshipNodeType(e.Type),
		TypeConfig:    RelationshipNodeTypeConfigToProto(e.TypeConfig),
		CreatedAt:     timestamppb.New(e.CreatedAt),
		UpdatedAt:     timestamppb.New(e.UpdatedAt),
		CreatedByUuid: e.CreatedByUUID.String(),
		UpdatedByUuid: e.UpdatedByUUID.String(),
	}
}

func RelationshipNodeSliceToProto(es []main_entity.RelationshipNode) []*pb.RelationshipNode {
	res := []*pb.RelationshipNode{}
	for _, e := range es {
		res = append(res, RelationshipNodeToProto(e))
	}
	return res
}

func RelationshipNodeFromProto(m *pb.RelationshipNode) main_entity.RelationshipNode {
	if m == nil {
		return main_entity.RelationshipNode{}
	}
	return main_entity.RelationshipNode{
		UUID:          uuid.FromStringOrNil(m.GetUuid()),
		Type:          main_entity.Type(m.GetType()),
		TypeConfig:    RelationshipNodeTypeConfigFromProto(m.GetTypeConfig()),
		CreatedAt:     m.GetCreatedAt().AsTime(),
		UpdatedAt:     m.GetUpdatedAt().AsTime(),
		CreatedByUUID: uuid.FromStringOrNil(m.GetCreatedByUuid()),
		UpdatedByUUID: uuid.FromStringOrNil(m.GetUpdatedByUuid()),
	}
}

func RelationshipNodeSliceFromProto(es []*pb.RelationshipNode) []main_entity.RelationshipNode {
	if es == nil {
		return []main_entity.RelationshipNode{}
	}
	res := []main_entity.RelationshipNode{}
	for _, e := range es {
		res = append(res, RelationshipNodeFromProto(e))
	}
	return res
}
