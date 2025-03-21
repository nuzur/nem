package entity

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"github.com/nuzur/nem/core/entity/element_render"
	"github.com/nuzur/nem/core/entity/entity_type_config"
	"github.com/nuzur/nem/core/entity/field"
	"time"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"
)

type Entity struct {
	UUID          uuid.UUID                           `json:"uuid"`
	Version       int64                               `json:"version"`
	Identifier    string                              `json:"identifier"`
	Description   *string                             `json:"description"`
	Fields        []field.Field                       `json:"fields"`
	Type          Type                                `json:"type"`
	TypeConfig    entity_type_config.EntityTypeConfig `json:"type_config"`
	Render        element_render.ElementRender        `json:"render"`
	Status        Status                              `json:"status"`
	CreatedAt     time.Time                           `json:"created_at"`
	UpdatedAt     time.Time                           `json:"updated_at"`
	CreatedByUUID uuid.UUID                           `json:"created_by_uuid"`
	UpdatedByUUID uuid.UUID                           `json:"updated_by_uuid"`
}

func (e Entity) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e Entity) EntityIdentifier() string {
	return "entity"
}

func (e Entity) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["uuid"] = types.UUIDFieldType
	res["version"] = types.IntFieldType
	res["identifier"] = types.StringFieldType
	res["description"] = types.StringFieldType
	res["fields"] = types.MultiDependantEntityFieldType
	res["type"] = types.SingleEnumFieldType
	res["type_config"] = types.SingleDependantEntityFieldType
	res["render"] = types.SingleDependantEntityFieldType
	res["status"] = types.SingleEnumFieldType
	res["created_at"] = types.TimestampFieldType
	res["updated_at"] = types.TimestampFieldType
	res["created_by_uuid"] = types.UUIDFieldType
	res["updated_by_uuid"] = types.UUIDFieldType
	return res
}

func (e Entity) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e Entity) IsDependant() bool {
	return true
}

func EntityFromJSON(data json.RawMessage) Entity {
	entity := Entity{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func EntitySliceFromJSON(data json.RawMessage) []Entity {
	entity := []Entity{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e Entity) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error EntityToJSON", err)
	}
	return res
}

func EntityToJSON(e Entity) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error EntityToJSON", err)
	}
	return res
}

func EntitySliceToJSON(e []Entity) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error EntitySliceToJSON", err)
	}
	return res
}

func NewEntityWithRandomValues() Entity {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return Entity{
		UUID:          randomvalues.GetRandomUUIDValue(),
		Version:       randomvalues.GetRandomIntValue(),
		Identifier:    randomvalues.GetRandomStringValue(),
		Description:   randomvalues.GetRandomStringValuePtr(),
		Fields:        field.NewFieldSliceWithRandomValues(rand.Intn(10)),
		Type:          randomvalues.GetRandomOptionValue[Type](2),
		TypeConfig:    entity_type_config.NewEntityTypeConfigWithRandomValues(),
		Render:        element_render.NewElementRenderWithRandomValues(),
		Status:        randomvalues.GetRandomOptionValue[Status](2),
		CreatedAt:     randomvalues.GetRandomTimeValue(),
		UpdatedAt:     randomvalues.GetRandomTimeValue(),
		CreatedByUUID: randomvalues.GetRandomUUIDValue(),
		UpdatedByUUID: randomvalues.GetRandomUUIDValue(),
	}
}

func NewEntitySliceWithRandomValues(n int) []Entity {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []Entity{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewEntityWithRandomValues())
	}
	return res
}
