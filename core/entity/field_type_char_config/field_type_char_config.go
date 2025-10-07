package field_type_char_config

import (
	"encoding/json"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"

	"time"
)

type FieldTypeCharConfig struct {
	MinSize         int64   `json:"min_size"`
	MaxSize         int64   `json:"max_size"`
	RegexValidation *string `json:"regex_validation"`
}

func (e FieldTypeCharConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e FieldTypeCharConfig) EntityIdentifier() string {
	return "field_type_char_config"
}

func (e FieldTypeCharConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["min_size"] = types.IntFieldType
	res["max_size"] = types.IntFieldType
	res["regex_validation"] = types.StringFieldType
	return res
}

func (e FieldTypeCharConfig) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "min_size")
	res = append(res, "max_size")
	res = append(res, "regex_validation")

	return res
}

func (e FieldTypeCharConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e FieldTypeCharConfig) IsDependant() bool {
	return true
}

func FieldTypeCharConfigFromJSON(data json.RawMessage) FieldTypeCharConfig {
	entity := FieldTypeCharConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func FieldTypeCharConfigSliceFromJSON(data json.RawMessage) []FieldTypeCharConfig {
	entity := []FieldTypeCharConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e FieldTypeCharConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeCharConfigToJSON", err)
	}
	return res
}

func FieldTypeCharConfigToJSON(e FieldTypeCharConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeCharConfigToJSON", err)
	}
	return res
}

func FieldTypeCharConfigSliceToJSON(e []FieldTypeCharConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeCharConfigSliceToJSON", err)
	}
	return res
}

func NewFieldTypeCharConfigWithRandomValues() FieldTypeCharConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return FieldTypeCharConfig{
		MinSize:         randomvalues.GetRandomIntValue(),
		MaxSize:         randomvalues.GetRandomIntValue(),
		RegexValidation: randomvalues.GetRandomStringValuePtr(),
	}
}

func NewFieldTypeCharConfigSliceWithRandomValues(n int) []FieldTypeCharConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []FieldTypeCharConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewFieldTypeCharConfigWithRandomValues())
	}
	return res
}
