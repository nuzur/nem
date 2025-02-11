package relationship_node_type_service_config

import (
	"encoding/json"

	"github.com/gofrs/uuid"

	"fmt"

	"nem/core/entity/types"

	"nem/core/randomvalues"

	"math/rand"

	"time"
)

type RelationshipNodeTypeServiceConfig struct {
	ServiceUUID uuid.UUID `json:"service_uuid"`
}

func (e RelationshipNodeTypeServiceConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e RelationshipNodeTypeServiceConfig) EntityIdentifier() string {
	return "relationship_node_type_service_config"
}

func (e RelationshipNodeTypeServiceConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["service_uuid"] = types.UUIDFieldType
	return res
}

func (e RelationshipNodeTypeServiceConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e RelationshipNodeTypeServiceConfig) IsDependant() bool {
	return true
}

func RelationshipNodeTypeServiceConfigFromJSON(data json.RawMessage) RelationshipNodeTypeServiceConfig {
	entity := RelationshipNodeTypeServiceConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func RelationshipNodeTypeServiceConfigSliceFromJSON(data json.RawMessage) []RelationshipNodeTypeServiceConfig {
	entity := []RelationshipNodeTypeServiceConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e RelationshipNodeTypeServiceConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error RelationshipNodeTypeServiceConfigToJSON", err)
	}
	return res
}

func RelationshipNodeTypeServiceConfigToJSON(e RelationshipNodeTypeServiceConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error RelationshipNodeTypeServiceConfigToJSON", err)
	}
	return res
}

func RelationshipNodeTypeServiceConfigSliceToJSON(e []RelationshipNodeTypeServiceConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error RelationshipNodeTypeServiceConfigSliceToJSON", err)
	}
	return res
}

func NewRelationshipNodeTypeServiceConfigWithRandomValues() RelationshipNodeTypeServiceConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return RelationshipNodeTypeServiceConfig{
		ServiceUUID: randomvalues.GetRandomUUIDValue(),
	}
}

func NewRelationshipNodeTypeServiceConfigSliceWithRandomValues(n int) []RelationshipNodeTypeServiceConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []RelationshipNodeTypeServiceConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewRelationshipNodeTypeServiceConfigWithRandomValues())
	}
	return res
}
