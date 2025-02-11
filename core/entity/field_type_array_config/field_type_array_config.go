package field_type_array_config

import (
	"encoding/json"

	"github.com/nuzur/nem/core/entity/array_type_config"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"

	"time"
)

type FieldTypeArrayConfig struct {
	MinElements   int64                             `json:"min_elements"`
	MaxElements   int64                             `json:"max_elements"`
	EnforceUnique bool                              `json:"enforce_unique"`
	Type          Type                              `json:"type"`
	TypeConfig    array_type_config.ArrayTypeConfig `json:"type_config"`
}

func (e FieldTypeArrayConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e FieldTypeArrayConfig) EntityIdentifier() string {
	return "field_type_array_config"
}

func (e FieldTypeArrayConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["min_elements"] = types.IntFieldType
	res["max_elements"] = types.IntFieldType
	res["enforce_unique"] = types.BooleanFieldType
	res["type"] = types.SingleEnumFieldType
	res["type_config"] = types.SingleDependantEntityFieldType
	return res
}

func (e FieldTypeArrayConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e FieldTypeArrayConfig) IsDependant() bool {
	return true
}

func FieldTypeArrayConfigFromJSON(data json.RawMessage) FieldTypeArrayConfig {
	entity := FieldTypeArrayConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func FieldTypeArrayConfigSliceFromJSON(data json.RawMessage) []FieldTypeArrayConfig {
	entity := []FieldTypeArrayConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e FieldTypeArrayConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeArrayConfigToJSON", err)
	}
	return res
}

func FieldTypeArrayConfigToJSON(e FieldTypeArrayConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeArrayConfigToJSON", err)
	}
	return res
}

func FieldTypeArrayConfigSliceToJSON(e []FieldTypeArrayConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeArrayConfigSliceToJSON", err)
	}
	return res
}

func NewFieldTypeArrayConfigWithRandomValues() FieldTypeArrayConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return FieldTypeArrayConfig{
		MinElements:   randomvalues.GetRandomIntValue(),
		MaxElements:   randomvalues.GetRandomIntValue(),
		EnforceUnique: randomvalues.GetRandomBoolValue(),
		Type:          randomvalues.GetRandomOptionValue[Type](15),
		TypeConfig:    array_type_config.NewArrayTypeConfigWithRandomValues(),
	}
}

func NewFieldTypeArrayConfigSliceWithRandomValues(n int) []FieldTypeArrayConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []FieldTypeArrayConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewFieldTypeArrayConfigWithRandomValues())
	}
	return res
}
