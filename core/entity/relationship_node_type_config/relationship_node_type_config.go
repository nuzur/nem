package relationship_node_type_config

import (
	"encoding/json"

	"github.com/nuzur/nem/core/entity/relationship_node_type_entity_config"
	"github.com/nuzur/nem/core/entity/relationship_node_type_service_config"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"math/rand"

	"time"
)

type RelationshipNodeTypeConfig struct {
	Entity  relationship_node_type_entity_config.RelationshipNodeTypeEntityConfig   `json:"entity"`
	Service relationship_node_type_service_config.RelationshipNodeTypeServiceConfig `json:"service"`
}

func (e RelationshipNodeTypeConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e RelationshipNodeTypeConfig) EntityIdentifier() string {
	return "relationship_node_type_config"
}

func (e RelationshipNodeTypeConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["entity"] = types.SingleDependantEntityFieldType
	res["service"] = types.SingleDependantEntityFieldType
	return res
}

func (e RelationshipNodeTypeConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e RelationshipNodeTypeConfig) IsDependant() bool {
	return true
}

func RelationshipNodeTypeConfigFromJSON(data json.RawMessage) RelationshipNodeTypeConfig {
	entity := RelationshipNodeTypeConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func RelationshipNodeTypeConfigSliceFromJSON(data json.RawMessage) []RelationshipNodeTypeConfig {
	entity := []RelationshipNodeTypeConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e RelationshipNodeTypeConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error RelationshipNodeTypeConfigToJSON", err)
	}
	return res
}

func RelationshipNodeTypeConfigToJSON(e RelationshipNodeTypeConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error RelationshipNodeTypeConfigToJSON", err)
	}
	return res
}

func RelationshipNodeTypeConfigSliceToJSON(e []RelationshipNodeTypeConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error RelationshipNodeTypeConfigSliceToJSON", err)
	}
	return res
}

func NewRelationshipNodeTypeConfigWithRandomValues() RelationshipNodeTypeConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return RelationshipNodeTypeConfig{
		Entity:  relationship_node_type_entity_config.NewRelationshipNodeTypeEntityConfigWithRandomValues(),
		Service: relationship_node_type_service_config.NewRelationshipNodeTypeServiceConfigWithRandomValues(),
	}
}

func NewRelationshipNodeTypeConfigSliceWithRandomValues(n int) []RelationshipNodeTypeConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []RelationshipNodeTypeConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewRelationshipNodeTypeConfigWithRandomValues())
	}
	return res
}
