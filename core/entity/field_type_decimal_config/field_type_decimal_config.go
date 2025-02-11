package field_type_decimal_config

import (
	"encoding/json"

	"fmt"

	"nem/core/entity/types"

	"nem/core/randomvalues"

	"math/rand"

	"time"
)

type FieldTypeDecimalConfig struct {
	MinValue          float64   `json:"min_value"`
	MinValueInclusive bool      `json:"min_value_inclusive"`
	MaxValue          float64   `json:"max_value"`
	MaxValueInclusive bool      `json:"max_value_inclusive"`
	AllowNegatives    bool      `json:"allow_negatives"`
	NumberOfDecimals  int64     `json:"number_of_decimals"`
	Separator         Separator `json:"separator"`
	IsCurrency        bool      `json:"is_currency"`
	CurrencyCode      string    `json:"currency_code"`
	EnableLimits      bool      `json:"enable_limits"`
}

func (e FieldTypeDecimalConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e FieldTypeDecimalConfig) EntityIdentifier() string {
	return "field_type_decimal_config"
}

func (e FieldTypeDecimalConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["min_value"] = types.FloatFieldType
	res["min_value_inclusive"] = types.BooleanFieldType
	res["max_value"] = types.FloatFieldType
	res["max_value_inclusive"] = types.BooleanFieldType
	res["allow_negatives"] = types.BooleanFieldType
	res["number_of_decimals"] = types.IntFieldType
	res["separator"] = types.SingleEnumFieldType
	res["is_currency"] = types.BooleanFieldType
	res["currency_code"] = types.StringFieldType
	res["enable_limits"] = types.BooleanFieldType
	return res
}

func (e FieldTypeDecimalConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e FieldTypeDecimalConfig) IsDependant() bool {
	return true
}

func FieldTypeDecimalConfigFromJSON(data json.RawMessage) FieldTypeDecimalConfig {
	entity := FieldTypeDecimalConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func FieldTypeDecimalConfigSliceFromJSON(data json.RawMessage) []FieldTypeDecimalConfig {
	entity := []FieldTypeDecimalConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e FieldTypeDecimalConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeDecimalConfigToJSON", err)
	}
	return res
}

func FieldTypeDecimalConfigToJSON(e FieldTypeDecimalConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeDecimalConfigToJSON", err)
	}
	return res
}

func FieldTypeDecimalConfigSliceToJSON(e []FieldTypeDecimalConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeDecimalConfigSliceToJSON", err)
	}
	return res
}

func NewFieldTypeDecimalConfigWithRandomValues() FieldTypeDecimalConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return FieldTypeDecimalConfig{
		MinValue:          randomvalues.GetRandomFloatValue(),
		MinValueInclusive: randomvalues.GetRandomBoolValue(),
		MaxValue:          randomvalues.GetRandomFloatValue(),
		MaxValueInclusive: randomvalues.GetRandomBoolValue(),
		AllowNegatives:    randomvalues.GetRandomBoolValue(),
		NumberOfDecimals:  randomvalues.GetRandomIntValue(),
		Separator:         randomvalues.GetRandomOptionValue[Separator](2),
		IsCurrency:        randomvalues.GetRandomBoolValue(),
		CurrencyCode:      randomvalues.GetRandomStringValue(),
		EnableLimits:      randomvalues.GetRandomBoolValue(),
	}
}

func NewFieldTypeDecimalConfigSliceWithRandomValues(n int) []FieldTypeDecimalConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []FieldTypeDecimalConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewFieldTypeDecimalConfigWithRandomValues())
	}
	return res
}
