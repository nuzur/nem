package field_type_url_config

import (
	"encoding/json"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"

	"time"
)

type FieldTypeURLConfig struct {
	AllowDomains      []string `json:"allow_domains"`
	ExcludeDomains    []string `json:"exclude_domains"`
	HTTPSRequired     bool     `json:"https_required"`
	AllowedExtensions []string `json:"allowed_extensions"`
}

func (e FieldTypeURLConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e FieldTypeURLConfig) EntityIdentifier() string {
	return "field_type_url_config"
}

func (e FieldTypeURLConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["allow_domains"] = types.ArrayFieldType
	res["exclude_domains"] = types.ArrayFieldType
	res["https_required"] = types.BooleanFieldType
	res["allowed_extensions"] = types.ArrayFieldType
	return res
}

func (e FieldTypeURLConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["allow_domains"] = types.StringFieldType
	res["exclude_domains"] = types.StringFieldType
	res["allowed_extensions"] = types.StringFieldType
	return res
}

func (e FieldTypeURLConfig) IsDependant() bool {
	return true
}

func FieldTypeURLConfigFromJSON(data json.RawMessage) FieldTypeURLConfig {
	entity := FieldTypeURLConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func FieldTypeURLConfigSliceFromJSON(data json.RawMessage) []FieldTypeURLConfig {
	entity := []FieldTypeURLConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e FieldTypeURLConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeURLConfigToJSON", err)
	}
	return res
}

func FieldTypeURLConfigToJSON(e FieldTypeURLConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeURLConfigToJSON", err)
	}
	return res
}

func FieldTypeURLConfigSliceToJSON(e []FieldTypeURLConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeURLConfigSliceToJSON", err)
	}
	return res
}

func NewFieldTypeURLConfigWithRandomValues() FieldTypeURLConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return FieldTypeURLConfig{
		AllowDomains:      []string{},
		ExcludeDomains:    []string{},
		HTTPSRequired:     randomvalues.GetRandomBoolValue(),
		AllowedExtensions: []string{},
	}
}

func NewFieldTypeURLConfigSliceWithRandomValues(n int) []FieldTypeURLConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []FieldTypeURLConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewFieldTypeURLConfigWithRandomValues())
	}
	return res
}
