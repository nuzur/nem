package field_type_phone_config

import (
	"encoding/json"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"math/rand"

	"time"
)

type FieldTypePhoneConfig struct {
	AllowCountries   []string `json:"allow_countries"`
	ExcludeCountries []string `json:"exclude_countries"`
}

func (e FieldTypePhoneConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e FieldTypePhoneConfig) EntityIdentifier() string {
	return "field_type_phone_config"
}

func (e FieldTypePhoneConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["allow_countries"] = types.ArrayFieldType
	res["exclude_countries"] = types.ArrayFieldType
	return res
}

func (e FieldTypePhoneConfig) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "allow_countries")
	res = append(res, "exclude_countries")

	return res
}

func (e FieldTypePhoneConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["allow_countries"] = types.StringFieldType
	res["exclude_countries"] = types.StringFieldType
	return res
}

func (e FieldTypePhoneConfig) IsDependant() bool {
	return true
}

func FieldTypePhoneConfigFromJSON(data json.RawMessage) FieldTypePhoneConfig {
	entity := FieldTypePhoneConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func FieldTypePhoneConfigSliceFromJSON(data json.RawMessage) []FieldTypePhoneConfig {
	entity := []FieldTypePhoneConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e FieldTypePhoneConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypePhoneConfigToJSON", err)
	}
	return res
}

func FieldTypePhoneConfigToJSON(e FieldTypePhoneConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypePhoneConfigToJSON", err)
	}
	return res
}

func FieldTypePhoneConfigSliceToJSON(e []FieldTypePhoneConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypePhoneConfigSliceToJSON", err)
	}
	return res
}

func NewFieldTypePhoneConfigWithRandomValues() FieldTypePhoneConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return FieldTypePhoneConfig{
		AllowCountries:   []string{},
		ExcludeCountries: []string{},
	}
}

func NewFieldTypePhoneConfigSliceWithRandomValues(n int) []FieldTypePhoneConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []FieldTypePhoneConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewFieldTypePhoneConfigWithRandomValues())
	}
	return res
}
