package enum_value

import (
	"encoding/json"

	"fmt"

	"nem/core/entity/types"

	"nem/core/randomvalues"

	"math/rand"

	"time"
)

type EnumValue struct {
	Identifier   string `json:"identifier"`
	Display      string `json:"display"`
	NumericValue int64  `json:"numeric_value"`
}

func (e EnumValue) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e EnumValue) EntityIdentifier() string {
	return "enum_value"
}

func (e EnumValue) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["identifier"] = types.StringFieldType
	res["display"] = types.StringFieldType
	res["numeric_value"] = types.IntFieldType
	return res
}

func (e EnumValue) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e EnumValue) IsDependant() bool {
	return true
}

func EnumValueFromJSON(data json.RawMessage) EnumValue {
	entity := EnumValue{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func EnumValueSliceFromJSON(data json.RawMessage) []EnumValue {
	entity := []EnumValue{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e EnumValue) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error EnumValueToJSON", err)
	}
	return res
}

func EnumValueToJSON(e EnumValue) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error EnumValueToJSON", err)
	}
	return res
}

func EnumValueSliceToJSON(e []EnumValue) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error EnumValueSliceToJSON", err)
	}
	return res
}

func NewEnumValueWithRandomValues() EnumValue {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return EnumValue{
		Identifier:   randomvalues.GetRandomStringValue(),
		Display:      randomvalues.GetRandomStringValue(),
		NumericValue: randomvalues.GetRandomIntValue(),
	}
}

func NewEnumValueSliceWithRandomValues(n int) []EnumValue {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []EnumValue{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewEnumValueWithRandomValues())
	}
	return res
}
