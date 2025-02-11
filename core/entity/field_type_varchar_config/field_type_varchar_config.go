package field_type_varchar_config

import (
	"encoding/json"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"

	"time"
)

type FieldTypeVarcharConfig struct {
	MinSize         int64  `json:"min_size"`
	MaxSize         int64  `json:"max_size"`
	RegexValidation string `json:"regex_validation"`
}

func (e FieldTypeVarcharConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e FieldTypeVarcharConfig) EntityIdentifier() string {
	return "field_type_varchar_config"
}

func (e FieldTypeVarcharConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["min_size"] = types.IntFieldType
	res["max_size"] = types.IntFieldType
	res["regex_validation"] = types.StringFieldType
	return res
}

func (e FieldTypeVarcharConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e FieldTypeVarcharConfig) IsDependant() bool {
	return true
}

func FieldTypeVarcharConfigFromJSON(data json.RawMessage) FieldTypeVarcharConfig {
	entity := FieldTypeVarcharConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func FieldTypeVarcharConfigSliceFromJSON(data json.RawMessage) []FieldTypeVarcharConfig {
	entity := []FieldTypeVarcharConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e FieldTypeVarcharConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeVarcharConfigToJSON", err)
	}
	return res
}

func FieldTypeVarcharConfigToJSON(e FieldTypeVarcharConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeVarcharConfigToJSON", err)
	}
	return res
}

func FieldTypeVarcharConfigSliceToJSON(e []FieldTypeVarcharConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeVarcharConfigSliceToJSON", err)
	}
	return res
}

func NewFieldTypeVarcharConfigWithRandomValues() FieldTypeVarcharConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return FieldTypeVarcharConfig{
		MinSize:         randomvalues.GetRandomIntValue(),
		MaxSize:         randomvalues.GetRandomIntValue(),
		RegexValidation: randomvalues.GetRandomStringValue(),
	}
}

func NewFieldTypeVarcharConfigSliceWithRandomValues(n int) []FieldTypeVarcharConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []FieldTypeVarcharConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewFieldTypeVarcharConfigWithRandomValues())
	}
	return res
}
