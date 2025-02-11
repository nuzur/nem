package entity_version_config

import (
	"encoding/json"

	"github.com/nuzur/nem/core/entity/entity_version_type_config"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"

	"time"
)

type EntityVersionConfig struct {
	Type      Type                                               `json:"type"`
	Generator Generator                                          `json:"generator"`
	Config    entity_version_type_config.EntityVersionTypeConfig `json:"config"`
}

func (e EntityVersionConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e EntityVersionConfig) EntityIdentifier() string {
	return "entity_version_config"
}

func (e EntityVersionConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["type"] = types.SingleEnumFieldType
	res["generator"] = types.SingleEnumFieldType
	res["config"] = types.SingleDependantEntityFieldType
	return res
}

func (e EntityVersionConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e EntityVersionConfig) IsDependant() bool {
	return true
}

func EntityVersionConfigFromJSON(data json.RawMessage) EntityVersionConfig {
	entity := EntityVersionConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func EntityVersionConfigSliceFromJSON(data json.RawMessage) []EntityVersionConfig {
	entity := []EntityVersionConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e EntityVersionConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error EntityVersionConfigToJSON", err)
	}
	return res
}

func EntityVersionConfigToJSON(e EntityVersionConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error EntityVersionConfigToJSON", err)
	}
	return res
}

func EntityVersionConfigSliceToJSON(e []EntityVersionConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error EntityVersionConfigSliceToJSON", err)
	}
	return res
}

func NewEntityVersionConfigWithRandomValues() EntityVersionConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return EntityVersionConfig{
		Type:      randomvalues.GetRandomOptionValue[Type](2),
		Generator: randomvalues.GetRandomOptionValue[Generator](1),
		Config:    entity_version_type_config.NewEntityVersionTypeConfigWithRandomValues(),
	}
}

func NewEntityVersionConfigSliceWithRandomValues(n int) []EntityVersionConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []EntityVersionConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewEntityVersionConfigWithRandomValues())
	}
	return res
}
