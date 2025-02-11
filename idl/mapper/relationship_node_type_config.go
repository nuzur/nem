package mapper

import (
	main_entity "nem/core/entity/relationship_node_type_config"
	pb "nem/idl/gen"
)

func RelationshipNodeTypeConfigToProto(e main_entity.RelationshipNodeTypeConfig) *pb.RelationshipNodeTypeConfig {
	return &pb.RelationshipNodeTypeConfig{
		Entity:  RelationshipNodeTypeEntityConfigToProto(e.Entity),
		Service: RelationshipNodeTypeServiceConfigToProto(e.Service),
	}
}

func RelationshipNodeTypeConfigSliceToProto(es []main_entity.RelationshipNodeTypeConfig) []*pb.RelationshipNodeTypeConfig {
	res := []*pb.RelationshipNodeTypeConfig{}
	for _, e := range es {
		res = append(res, RelationshipNodeTypeConfigToProto(e))
	}
	return res
}

func RelationshipNodeTypeConfigFromProto(m *pb.RelationshipNodeTypeConfig) main_entity.RelationshipNodeTypeConfig {
	if m == nil {
		return main_entity.RelationshipNodeTypeConfig{}
	}
	return main_entity.RelationshipNodeTypeConfig{
		Entity:  RelationshipNodeTypeEntityConfigFromProto(m.GetEntity()),
		Service: RelationshipNodeTypeServiceConfigFromProto(m.GetService()),
	}
}

func RelationshipNodeTypeConfigSliceFromProto(es []*pb.RelationshipNodeTypeConfig) []main_entity.RelationshipNodeTypeConfig {
	if es == nil {
		return []main_entity.RelationshipNodeTypeConfig{}
	}
	res := []main_entity.RelationshipNodeTypeConfig{}
	for _, e := range es {
		res = append(res, RelationshipNodeTypeConfigFromProto(e))
	}
	return res
}
