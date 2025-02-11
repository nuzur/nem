package field_type_float_config

import (
	"encoding/json"

	"fmt"

	"nem/core/entity/types"

	"nem/core/randomvalues"

	"math/rand"

	"time"
)

type FieldTypeFloatConfig struct {
	MinValue          float64   `json:"min_value"`
	MinValueInclusive bool      `json:"min_value_inclusive"`
	MaxValue          float64   `json:"max_value"`
	MaxValueInclusive bool      `json:"max_value_inclusive"`
	AllowNegatives    bool      `json:"allow_negatives"`
	NumberOfDecimals  int64     `json:"number_of_decimals"`
	Separator         Separator `json:"separator"`
	EnableLimits      bool      `json:"enable_limits"`
}

func (e FieldTypeFloatConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e FieldTypeFloatConfig) EntityIdentifier() string {
	return "field_type_float_config"
}

func (e FieldTypeFloatConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["min_value"] = types.FloatFieldType
	res["min_value_inclusive"] = types.BooleanFieldType
	res["max_value"] = types.FloatFieldType
	res["max_value_inclusive"] = types.BooleanFieldType
	res["allow_negatives"] = types.BooleanFieldType
	res["number_of_decimals"] = types.IntFieldType
	res["separator"] = types.SingleEnumFieldType
	res["enable_limits"] = types.BooleanFieldType
	return res
}

func (e FieldTypeFloatConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e FieldTypeFloatConfig) IsDependant() bool {
	return true
}

func FieldTypeFloatConfigFromJSON(data json.RawMessage) FieldTypeFloatConfig {
	entity := FieldTypeFloatConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func FieldTypeFloatConfigSliceFromJSON(data json.RawMessage) []FieldTypeFloatConfig {
	entity := []FieldTypeFloatConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e FieldTypeFloatConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeFloatConfigToJSON", err)
	}
	return res
}

func FieldTypeFloatConfigToJSON(e FieldTypeFloatConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeFloatConfigToJSON", err)
	}
	return res
}

func FieldTypeFloatConfigSliceToJSON(e []FieldTypeFloatConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeFloatConfigSliceToJSON", err)
	}
	return res
}

func NewFieldTypeFloatConfigWithRandomValues() FieldTypeFloatConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return FieldTypeFloatConfig{
		MinValue:          randomvalues.GetRandomFloatValue(),
		MinValueInclusive: randomvalues.GetRandomBoolValue(),
		MaxValue:          randomvalues.GetRandomFloatValue(),
		MaxValueInclusive: randomvalues.GetRandomBoolValue(),
		AllowNegatives:    randomvalues.GetRandomBoolValue(),
		NumberOfDecimals:  randomvalues.GetRandomIntValue(),
		Separator:         randomvalues.GetRandomOptionValue[Separator](2),
		EnableLimits:      randomvalues.GetRandomBoolValue(),
	}
}

func NewFieldTypeFloatConfigSliceWithRandomValues(n int) []FieldTypeFloatConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []FieldTypeFloatConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewFieldTypeFloatConfigWithRandomValues())
	}
	return res
}
