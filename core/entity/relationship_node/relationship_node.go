package relationship_node

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"nem/core/entity/relationship_node_type_config"
	"time"

	"fmt"

	"nem/core/entity/types"

	"nem/core/randomvalues"

	"math/rand"
)

type RelationshipNode struct {
	UUID          uuid.UUID                                                `json:"uuid"`
	Type          Type                                                     `json:"type"`
	TypeConfig    relationship_node_type_config.RelationshipNodeTypeConfig `json:"type_config"`
	CreatedAt     time.Time                                                `json:"created_at"`
	UpdatedAt     time.Time                                                `json:"updated_at"`
	CreatedByUUID uuid.UUID                                                `json:"created_by_uuid"`
	UpdatedByUUID uuid.UUID                                                `json:"updated_by_uuid"`
}

func (e RelationshipNode) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e RelationshipNode) EntityIdentifier() string {
	return "relationship_node"
}

func (e RelationshipNode) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["uuid"] = types.UUIDFieldType
	res["type"] = types.SingleEnumFieldType
	res["type_config"] = types.SingleDependantEntityFieldType
	res["created_at"] = types.TimestampFieldType
	res["updated_at"] = types.TimestampFieldType
	res["created_by_uuid"] = types.UUIDFieldType
	res["updated_by_uuid"] = types.UUIDFieldType
	return res
}

func (e RelationshipNode) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e RelationshipNode) IsDependant() bool {
	return true
}

func RelationshipNodeFromJSON(data json.RawMessage) RelationshipNode {
	entity := RelationshipNode{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func RelationshipNodeSliceFromJSON(data json.RawMessage) []RelationshipNode {
	entity := []RelationshipNode{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e RelationshipNode) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error RelationshipNodeToJSON", err)
	}
	return res
}

func RelationshipNodeToJSON(e RelationshipNode) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error RelationshipNodeToJSON", err)
	}
	return res
}

func RelationshipNodeSliceToJSON(e []RelationshipNode) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error RelationshipNodeSliceToJSON", err)
	}
	return res
}

func NewRelationshipNodeWithRandomValues() RelationshipNode {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return RelationshipNode{
		UUID:          randomvalues.GetRandomUUIDValue(),
		Type:          randomvalues.GetRandomOptionValue[Type](2),
		TypeConfig:    relationship_node_type_config.NewRelationshipNodeTypeConfigWithRandomValues(),
		CreatedAt:     randomvalues.GetRandomTimeValue(),
		UpdatedAt:     randomvalues.GetRandomTimeValue(),
		CreatedByUUID: randomvalues.GetRandomUUIDValue(),
		UpdatedByUUID: randomvalues.GetRandomUUIDValue(),
	}
}

func NewRelationshipNodeSliceWithRandomValues(n int) []RelationshipNode {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []RelationshipNode{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewRelationshipNodeWithRandomValues())
	}
	return res
}
