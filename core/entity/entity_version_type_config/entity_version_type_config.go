package entity_version_type_config

import (
	"encoding/json"

	"nem/core/entity/entity_version_type_entity_config"
	"nem/core/entity/entity_version_type_field_config"

	"fmt"

	"nem/core/entity/types"

	"math/rand"

	"time"
)

type EntityVersionTypeConfig struct {
	Field  entity_version_type_field_config.EntityVersionTypeFieldConfig   `json:"field"`
	Entity entity_version_type_entity_config.EntityVersionTypeEntityConfig `json:"entity"`
}

func (e EntityVersionTypeConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e EntityVersionTypeConfig) EntityIdentifier() string {
	return "entity_version_type_config"
}

func (e EntityVersionTypeConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["field"] = types.SingleDependantEntityFieldType
	res["entity"] = types.SingleDependantEntityFieldType
	return res
}

func (e EntityVersionTypeConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e EntityVersionTypeConfig) IsDependant() bool {
	return true
}

func EntityVersionTypeConfigFromJSON(data json.RawMessage) EntityVersionTypeConfig {
	entity := EntityVersionTypeConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func EntityVersionTypeConfigSliceFromJSON(data json.RawMessage) []EntityVersionTypeConfig {
	entity := []EntityVersionTypeConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e EntityVersionTypeConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error EntityVersionTypeConfigToJSON", err)
	}
	return res
}

func EntityVersionTypeConfigToJSON(e EntityVersionTypeConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error EntityVersionTypeConfigToJSON", err)
	}
	return res
}

func EntityVersionTypeConfigSliceToJSON(e []EntityVersionTypeConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error EntityVersionTypeConfigSliceToJSON", err)
	}
	return res
}

func NewEntityVersionTypeConfigWithRandomValues() EntityVersionTypeConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return EntityVersionTypeConfig{
		Field:  entity_version_type_field_config.NewEntityVersionTypeFieldConfigWithRandomValues(),
		Entity: entity_version_type_entity_config.NewEntityVersionTypeEntityConfigWithRandomValues(),
	}
}

func NewEntityVersionTypeConfigSliceWithRandomValues(n int) []EntityVersionTypeConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []EntityVersionTypeConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewEntityVersionTypeConfigWithRandomValues())
	}
	return res
}
