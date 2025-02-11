package field_type_integer_config

import (
	"encoding/json"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"

	"time"
)

type FieldTypeIntegerConfig struct {
	Size              Size  `json:"size"`
	MinValue          int64 `json:"min_value"`
	MinValueInclusive bool  `json:"min_value_inclusive"`
	MaxValue          int64 `json:"max_value"`
	MaxValueInclusive bool  `json:"max_value_inclusive"`
	AllowNegatives    bool  `json:"allow_negatives"`
	EnableLimits      bool  `json:"enable_limits"`
}

func (e FieldTypeIntegerConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e FieldTypeIntegerConfig) EntityIdentifier() string {
	return "field_type_integer_config"
}

func (e FieldTypeIntegerConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["size"] = types.SingleEnumFieldType
	res["min_value"] = types.IntFieldType
	res["min_value_inclusive"] = types.BooleanFieldType
	res["max_value"] = types.IntFieldType
	res["max_value_inclusive"] = types.BooleanFieldType
	res["allow_negatives"] = types.BooleanFieldType
	res["enable_limits"] = types.BooleanFieldType
	return res
}

func (e FieldTypeIntegerConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e FieldTypeIntegerConfig) IsDependant() bool {
	return true
}

func FieldTypeIntegerConfigFromJSON(data json.RawMessage) FieldTypeIntegerConfig {
	entity := FieldTypeIntegerConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func FieldTypeIntegerConfigSliceFromJSON(data json.RawMessage) []FieldTypeIntegerConfig {
	entity := []FieldTypeIntegerConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e FieldTypeIntegerConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeIntegerConfigToJSON", err)
	}
	return res
}

func FieldTypeIntegerConfigToJSON(e FieldTypeIntegerConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeIntegerConfigToJSON", err)
	}
	return res
}

func FieldTypeIntegerConfigSliceToJSON(e []FieldTypeIntegerConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeIntegerConfigSliceToJSON", err)
	}
	return res
}

func NewFieldTypeIntegerConfigWithRandomValues() FieldTypeIntegerConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return FieldTypeIntegerConfig{
		Size:              randomvalues.GetRandomOptionValue[Size](6),
		MinValue:          randomvalues.GetRandomIntValue(),
		MinValueInclusive: randomvalues.GetRandomBoolValue(),
		MaxValue:          randomvalues.GetRandomIntValue(),
		MaxValueInclusive: randomvalues.GetRandomBoolValue(),
		AllowNegatives:    randomvalues.GetRandomBoolValue(),
		EnableLimits:      randomvalues.GetRandomBoolValue(),
	}
}

func NewFieldTypeIntegerConfigSliceWithRandomValues(n int) []FieldTypeIntegerConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []FieldTypeIntegerConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewFieldTypeIntegerConfigWithRandomValues())
	}
	return res
}
