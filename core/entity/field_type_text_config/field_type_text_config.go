package field_type_text_config

import (
	"encoding/json"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"

	"time"
)

type FieldTypeTextConfig struct {
	MinSize int64 `json:"min_size"`
	MaxSize int64 `json:"max_size"`
}

func (e FieldTypeTextConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e FieldTypeTextConfig) EntityIdentifier() string {
	return "field_type_text_config"
}

func (e FieldTypeTextConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["min_size"] = types.IntFieldType
	res["max_size"] = types.IntFieldType
	return res
}

func (e FieldTypeTextConfig) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "min_size")
	res = append(res, "max_size")

	return res
}

func (e FieldTypeTextConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e FieldTypeTextConfig) IsDependant() bool {
	return true
}

func FieldTypeTextConfigFromJSON(data json.RawMessage) FieldTypeTextConfig {
	entity := FieldTypeTextConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func FieldTypeTextConfigSliceFromJSON(data json.RawMessage) []FieldTypeTextConfig {
	entity := []FieldTypeTextConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e FieldTypeTextConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeTextConfigToJSON", err)
	}
	return res
}

func FieldTypeTextConfigToJSON(e FieldTypeTextConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeTextConfigToJSON", err)
	}
	return res
}

func FieldTypeTextConfigSliceToJSON(e []FieldTypeTextConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeTextConfigSliceToJSON", err)
	}
	return res
}

func NewFieldTypeTextConfigWithRandomValues() FieldTypeTextConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return FieldTypeTextConfig{
		MinSize: randomvalues.GetRandomIntValue(),
		MaxSize: randomvalues.GetRandomIntValue(),
	}
}

func NewFieldTypeTextConfigSliceWithRandomValues(n int) []FieldTypeTextConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []FieldTypeTextConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewFieldTypeTextConfigWithRandomValues())
	}
	return res
}
