package field

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"nem/core/entity/field_type_config"
	"time"

	"fmt"

	"nem/core/entity/types"

	"nem/core/randomvalues"

	"math/rand"
)

type Field struct {
	UUID             uuid.UUID                         `json:"uuid"`
	Version          int64                             `json:"version"`
	Identifier       string                            `json:"identifier"`
	Description      string                            `json:"description"`
	Type             Type                              `json:"type"`
	TypeConfig       field_type_config.FieldTypeConfig `json:"type_config"`
	Required         bool                              `json:"required"`
	Key              bool                              `json:"key"`
	KeyAutoIncrement bool                              `json:"key_auto_increment"`
	Unique           bool                              `json:"unique"`
	Deprecated       bool                              `json:"deprecated"`
	Generated        bool                              `json:"generated"`
	Status           Status                            `json:"status"`
	CreatedAt        time.Time                         `json:"created_at"`
	UpdatedAt        time.Time                         `json:"updated_at"`
	CreatedByUUID    uuid.UUID                         `json:"created_by_uuid"`
	UpdatedByUUID    uuid.UUID                         `json:"updated_by_uuid"`
}

func (e Field) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e Field) EntityIdentifier() string {
	return "field"
}

func (e Field) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["uuid"] = types.UUIDFieldType
	res["version"] = types.IntFieldType
	res["identifier"] = types.StringFieldType
	res["description"] = types.StringFieldType
	res["type"] = types.SingleEnumFieldType
	res["type_config"] = types.SingleDependantEntityFieldType
	res["required"] = types.BooleanFieldType
	res["key"] = types.BooleanFieldType
	res["key_auto_increment"] = types.BooleanFieldType
	res["unique"] = types.BooleanFieldType
	res["deprecated"] = types.BooleanFieldType
	res["generated"] = types.BooleanFieldType
	res["status"] = types.SingleEnumFieldType
	res["created_at"] = types.TimestampFieldType
	res["updated_at"] = types.TimestampFieldType
	res["created_by_uuid"] = types.UUIDFieldType
	res["updated_by_uuid"] = types.UUIDFieldType
	return res
}

func (e Field) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e Field) IsDependant() bool {
	return true
}

func FieldFromJSON(data json.RawMessage) Field {
	entity := Field{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func FieldSliceFromJSON(data json.RawMessage) []Field {
	entity := []Field{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e Field) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldToJSON", err)
	}
	return res
}

func FieldToJSON(e Field) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldToJSON", err)
	}
	return res
}

func FieldSliceToJSON(e []Field) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldSliceToJSON", err)
	}
	return res
}

func NewFieldWithRandomValues() Field {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return Field{
		UUID:             randomvalues.GetRandomUUIDValue(),
		Version:          randomvalues.GetRandomIntValue(),
		Identifier:       randomvalues.GetRandomStringValue(),
		Description:      randomvalues.GetRandomStringValue(),
		Type:             randomvalues.GetRandomOptionValue[Type](28),
		TypeConfig:       field_type_config.NewFieldTypeConfigWithRandomValues(),
		Required:         randomvalues.GetRandomBoolValue(),
		Key:              randomvalues.GetRandomBoolValue(),
		KeyAutoIncrement: randomvalues.GetRandomBoolValue(),
		Unique:           randomvalues.GetRandomBoolValue(),
		Deprecated:       randomvalues.GetRandomBoolValue(),
		Generated:        randomvalues.GetRandomBoolValue(),
		Status:           randomvalues.GetRandomOptionValue[Status](2),
		CreatedAt:        randomvalues.GetRandomTimeValue(),
		UpdatedAt:        randomvalues.GetRandomTimeValue(),
		CreatedByUUID:    randomvalues.GetRandomUUIDValue(),
		UpdatedByUUID:    randomvalues.GetRandomUUIDValue(),
	}
}

func NewFieldSliceWithRandomValues(n int) []Field {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []Field{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewFieldWithRandomValues())
	}
	return res
}
