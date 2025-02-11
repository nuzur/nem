package relationship

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"nem/core/entity/relationship_node"
	"time"

	"fmt"

	"nem/core/entity/types"

	"nem/core/randomvalues"

	"math/rand"
)

type Relationship struct {
	UUID          uuid.UUID                          `json:"uuid"`
	Version       int64                              `json:"version"`
	Identifier    string                             `json:"identifier"`
	Description   string                             `json:"description"`
	Cardinality   Cardinality                        `json:"cardinality"`
	From          relationship_node.RelationshipNode `json:"from"`
	To            relationship_node.RelationshipNode `json:"to"`
	UseForeignKey bool                               `json:"use_foreign_key"`
	Status        Status                             `json:"status"`
	CreatedAt     time.Time                          `json:"created_at"`
	UpdatedAt     time.Time                          `json:"updated_at"`
	CreatedByUUID uuid.UUID                          `json:"created_by_uuid"`
	UpdatedByUUID uuid.UUID                          `json:"updated_by_uuid"`
}

func (e Relationship) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e Relationship) EntityIdentifier() string {
	return "relationship"
}

func (e Relationship) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["uuid"] = types.UUIDFieldType
	res["version"] = types.IntFieldType
	res["identifier"] = types.StringFieldType
	res["description"] = types.StringFieldType
	res["cardinality"] = types.SingleEnumFieldType
	res["from"] = types.SingleDependantEntityFieldType
	res["to"] = types.SingleDependantEntityFieldType
	res["use_foreign_key"] = types.BooleanFieldType
	res["status"] = types.SingleEnumFieldType
	res["created_at"] = types.TimestampFieldType
	res["updated_at"] = types.TimestampFieldType
	res["created_by_uuid"] = types.UUIDFieldType
	res["updated_by_uuid"] = types.UUIDFieldType
	return res
}

func (e Relationship) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e Relationship) IsDependant() bool {
	return true
}

func RelationshipFromJSON(data json.RawMessage) Relationship {
	entity := Relationship{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func RelationshipSliceFromJSON(data json.RawMessage) []Relationship {
	entity := []Relationship{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e Relationship) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error RelationshipToJSON", err)
	}
	return res
}

func RelationshipToJSON(e Relationship) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error RelationshipToJSON", err)
	}
	return res
}

func RelationshipSliceToJSON(e []Relationship) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error RelationshipSliceToJSON", err)
	}
	return res
}

func NewRelationshipWithRandomValues() Relationship {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return Relationship{
		UUID:          randomvalues.GetRandomUUIDValue(),
		Version:       randomvalues.GetRandomIntValue(),
		Identifier:    randomvalues.GetRandomStringValue(),
		Description:   randomvalues.GetRandomStringValue(),
		Cardinality:   randomvalues.GetRandomOptionValue[Cardinality](2),
		From:          relationship_node.NewRelationshipNodeWithRandomValues(),
		To:            relationship_node.NewRelationshipNodeWithRandomValues(),
		UseForeignKey: randomvalues.GetRandomBoolValue(),
		Status:        randomvalues.GetRandomOptionValue[Status](2),
		CreatedAt:     randomvalues.GetRandomTimeValue(),
		UpdatedAt:     randomvalues.GetRandomTimeValue(),
		CreatedByUUID: randomvalues.GetRandomUUIDValue(),
		UpdatedByUUID: randomvalues.GetRandomUUIDValue(),
	}
}

func NewRelationshipSliceWithRandomValues(n int) []Relationship {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []Relationship{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewRelationshipWithRandomValues())
	}
	return res
}
