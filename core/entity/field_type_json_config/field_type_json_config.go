package field_type_json_config

import (
	"encoding/json"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"

	"time"
)

type FieldTypeJSONConfig struct {
	EnforceSchemaValidation bool    `json:"enforce_schema_validation"`
	JSONSchema              *string `json:"json_schema"`
}

func (e FieldTypeJSONConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e FieldTypeJSONConfig) EntityIdentifier() string {
	return "field_type_json_config"
}

func (e FieldTypeJSONConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["enforce_schema_validation"] = types.BooleanFieldType
	res["json_schema"] = types.StringFieldType
	return res
}

func (e FieldTypeJSONConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e FieldTypeJSONConfig) IsDependant() bool {
	return true
}

func FieldTypeJSONConfigFromJSON(data json.RawMessage) FieldTypeJSONConfig {
	entity := FieldTypeJSONConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func FieldTypeJSONConfigSliceFromJSON(data json.RawMessage) []FieldTypeJSONConfig {
	entity := []FieldTypeJSONConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e FieldTypeJSONConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeJSONConfigToJSON", err)
	}
	return res
}

func FieldTypeJSONConfigToJSON(e FieldTypeJSONConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeJSONConfigToJSON", err)
	}
	return res
}

func FieldTypeJSONConfigSliceToJSON(e []FieldTypeJSONConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeJSONConfigSliceToJSON", err)
	}
	return res
}

func NewFieldTypeJSONConfigWithRandomValues() FieldTypeJSONConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return FieldTypeJSONConfig{
		EnforceSchemaValidation: randomvalues.GetRandomBoolValue(),
		JSONSchema:              randomvalues.GetRandomStringValuePtr(),
	}
}

func NewFieldTypeJSONConfigSliceWithRandomValues(n int) []FieldTypeJSONConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []FieldTypeJSONConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewFieldTypeJSONConfigWithRandomValues())
	}
	return res
}
