package field_type_email_config

import (
	"encoding/json"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"math/rand"

	"time"
)

type FieldTypeEmailConfig struct {
	AllowDomains   []string `json:"allow_domains"`
	ExcludeDomains []string `json:"exclude_domains"`
}

func (e FieldTypeEmailConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e FieldTypeEmailConfig) EntityIdentifier() string {
	return "field_type_email_config"
}

func (e FieldTypeEmailConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["allow_domains"] = types.ArrayFieldType
	res["exclude_domains"] = types.ArrayFieldType
	return res
}

func (e FieldTypeEmailConfig) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "allow_domains")
	res = append(res, "exclude_domains")

	return res
}

func (e FieldTypeEmailConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["allow_domains"] = types.StringFieldType
	res["exclude_domains"] = types.StringFieldType
	return res
}

func (e FieldTypeEmailConfig) IsDependant() bool {
	return true
}

func FieldTypeEmailConfigFromJSON(data json.RawMessage) FieldTypeEmailConfig {
	entity := FieldTypeEmailConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func FieldTypeEmailConfigSliceFromJSON(data json.RawMessage) []FieldTypeEmailConfig {
	entity := []FieldTypeEmailConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e FieldTypeEmailConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeEmailConfigToJSON", err)
	}
	return res
}

func FieldTypeEmailConfigToJSON(e FieldTypeEmailConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeEmailConfigToJSON", err)
	}
	return res
}

func FieldTypeEmailConfigSliceToJSON(e []FieldTypeEmailConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeEmailConfigSliceToJSON", err)
	}
	return res
}

func NewFieldTypeEmailConfigWithRandomValues() FieldTypeEmailConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return FieldTypeEmailConfig{
		AllowDomains:   []string{},
		ExcludeDomains: []string{},
	}
}

func NewFieldTypeEmailConfigSliceWithRandomValues(n int) []FieldTypeEmailConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []FieldTypeEmailConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewFieldTypeEmailConfigWithRandomValues())
	}
	return res
}
