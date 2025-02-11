package field_type_enum_config

import (
	"encoding/json"

	"github.com/gofrs/uuid"

	"fmt"

	"nem/core/entity/types"

	"nem/core/randomvalues"

	"math/rand"

	"time"
)

type FieldTypeEnumConfig struct {
	EnumUUID      uuid.UUID `json:"enum_uuid"`
	AllowMultiple bool      `json:"allow_multiple"`
}

func (e FieldTypeEnumConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e FieldTypeEnumConfig) EntityIdentifier() string {
	return "field_type_enum_config"
}

func (e FieldTypeEnumConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["enum_uuid"] = types.UUIDFieldType
	res["allow_multiple"] = types.BooleanFieldType
	return res
}

func (e FieldTypeEnumConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e FieldTypeEnumConfig) IsDependant() bool {
	return true
}

func FieldTypeEnumConfigFromJSON(data json.RawMessage) FieldTypeEnumConfig {
	entity := FieldTypeEnumConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func FieldTypeEnumConfigSliceFromJSON(data json.RawMessage) []FieldTypeEnumConfig {
	entity := []FieldTypeEnumConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e FieldTypeEnumConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeEnumConfigToJSON", err)
	}
	return res
}

func FieldTypeEnumConfigToJSON(e FieldTypeEnumConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeEnumConfigToJSON", err)
	}
	return res
}

func FieldTypeEnumConfigSliceToJSON(e []FieldTypeEnumConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeEnumConfigSliceToJSON", err)
	}
	return res
}

func NewFieldTypeEnumConfigWithRandomValues() FieldTypeEnumConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return FieldTypeEnumConfig{
		EnumUUID:      randomvalues.GetRandomUUIDValue(),
		AllowMultiple: randomvalues.GetRandomBoolValue(),
	}
}

func NewFieldTypeEnumConfigSliceWithRandomValues(n int) []FieldTypeEnumConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []FieldTypeEnumConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewFieldTypeEnumConfigWithRandomValues())
	}
	return res
}
