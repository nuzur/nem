package relationship_node_type_entity_config

import (
	"encoding/json"

	"github.com/gofrs/uuid"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"

	"time"
)

type RelationshipNodeTypeEntityConfig struct {
	EntityUUID      uuid.UUID   `json:"entity_uuid"`
	FieldUUIDs      []uuid.UUID `json:"field_uuids"`
	FieldsGenerated bool        `json:"fields_generated"`
}

func (e RelationshipNodeTypeEntityConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e RelationshipNodeTypeEntityConfig) EntityIdentifier() string {
	return "relationship_node_type_entity_config"
}

func (e RelationshipNodeTypeEntityConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["entity_uuid"] = types.UUIDFieldType
	res["field_uuids"] = types.ArrayFieldType
	res["fields_generated"] = types.BooleanFieldType
	return res
}

func (e RelationshipNodeTypeEntityConfig) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "entity_uuid")
	res = append(res, "field_uuids")
	res = append(res, "fields_generated")

	return res
}

func (e RelationshipNodeTypeEntityConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["field_uuids"] = types.UUIDFieldType
	return res
}

func (e RelationshipNodeTypeEntityConfig) IsDependant() bool {
	return true
}

func RelationshipNodeTypeEntityConfigFromJSON(data json.RawMessage) RelationshipNodeTypeEntityConfig {
	entity := RelationshipNodeTypeEntityConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func RelationshipNodeTypeEntityConfigSliceFromJSON(data json.RawMessage) []RelationshipNodeTypeEntityConfig {
	entity := []RelationshipNodeTypeEntityConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e RelationshipNodeTypeEntityConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error RelationshipNodeTypeEntityConfigToJSON", err)
	}
	return res
}

func RelationshipNodeTypeEntityConfigToJSON(e RelationshipNodeTypeEntityConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error RelationshipNodeTypeEntityConfigToJSON", err)
	}
	return res
}

func RelationshipNodeTypeEntityConfigSliceToJSON(e []RelationshipNodeTypeEntityConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error RelationshipNodeTypeEntityConfigSliceToJSON", err)
	}
	return res
}

func NewRelationshipNodeTypeEntityConfigWithRandomValues() RelationshipNodeTypeEntityConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return RelationshipNodeTypeEntityConfig{
		EntityUUID:      randomvalues.GetRandomUUIDValue(),
		FieldUUIDs:      []uuid.UUID{},
		FieldsGenerated: randomvalues.GetRandomBoolValue(),
	}
}

func NewRelationshipNodeTypeEntityConfigSliceWithRandomValues(n int) []RelationshipNodeTypeEntityConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []RelationshipNodeTypeEntityConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewRelationshipNodeTypeEntityConfigWithRandomValues())
	}
	return res
}
