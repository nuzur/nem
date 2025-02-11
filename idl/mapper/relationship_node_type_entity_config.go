package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/relationship_node_type_entity_config"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"
)

func RelationshipNodeTypeEntityConfigToProto(e main_entity.RelationshipNodeTypeEntityConfig) *pb.RelationshipNodeTypeEntityConfig {
	return &pb.RelationshipNodeTypeEntityConfig{
		EntityUuid:      e.EntityUUID.String(),
		FieldUuids:      UUIDSliceToStringSlice(e.FieldUUIDs),
		FieldsGenerated: e.FieldsGenerated,
	}
}

func RelationshipNodeTypeEntityConfigSliceToProto(es []main_entity.RelationshipNodeTypeEntityConfig) []*pb.RelationshipNodeTypeEntityConfig {
	res := []*pb.RelationshipNodeTypeEntityConfig{}
	for _, e := range es {
		res = append(res, RelationshipNodeTypeEntityConfigToProto(e))
	}
	return res
}

func RelationshipNodeTypeEntityConfigFromProto(m *pb.RelationshipNodeTypeEntityConfig) main_entity.RelationshipNodeTypeEntityConfig {
	if m == nil {
		return main_entity.RelationshipNodeTypeEntityConfig{}
	}
	return main_entity.RelationshipNodeTypeEntityConfig{
		EntityUUID:      uuid.FromStringOrNil(m.GetEntityUuid()),
		FieldUUIDs:      StringSliceToUUIDSlice(m.GetFieldUuids()),
		FieldsGenerated: m.GetFieldsGenerated(),
	}
}

func RelationshipNodeTypeEntityConfigSliceFromProto(es []*pb.RelationshipNodeTypeEntityConfig) []main_entity.RelationshipNodeTypeEntityConfig {
	if es == nil {
		return []main_entity.RelationshipNodeTypeEntityConfig{}
	}
	res := []main_entity.RelationshipNodeTypeEntityConfig{}
	for _, e := range es {
		res = append(res, RelationshipNodeTypeEntityConfigFromProto(e))
	}
	return res
}
