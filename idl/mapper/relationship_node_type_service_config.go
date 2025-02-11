package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/relationship_node_type_service_config"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"
)

func RelationshipNodeTypeServiceConfigToProto(e main_entity.RelationshipNodeTypeServiceConfig) *pb.RelationshipNodeTypeServiceConfig {
	return &pb.RelationshipNodeTypeServiceConfig{
		ServiceUuid: e.ServiceUUID.String(),
	}
}

func RelationshipNodeTypeServiceConfigSliceToProto(es []main_entity.RelationshipNodeTypeServiceConfig) []*pb.RelationshipNodeTypeServiceConfig {
	res := []*pb.RelationshipNodeTypeServiceConfig{}
	for _, e := range es {
		res = append(res, RelationshipNodeTypeServiceConfigToProto(e))
	}
	return res
}

func RelationshipNodeTypeServiceConfigFromProto(m *pb.RelationshipNodeTypeServiceConfig) main_entity.RelationshipNodeTypeServiceConfig {
	if m == nil {
		return main_entity.RelationshipNodeTypeServiceConfig{}
	}
	return main_entity.RelationshipNodeTypeServiceConfig{
		ServiceUUID: uuid.FromStringOrNil(m.GetServiceUuid()),
	}
}

func RelationshipNodeTypeServiceConfigSliceFromProto(es []*pb.RelationshipNodeTypeServiceConfig) []main_entity.RelationshipNodeTypeServiceConfig {
	if es == nil {
		return []main_entity.RelationshipNodeTypeServiceConfig{}
	}
	res := []main_entity.RelationshipNodeTypeServiceConfig{}
	for _, e := range es {
		res = append(res, RelationshipNodeTypeServiceConfigFromProto(e))
	}
	return res
}
