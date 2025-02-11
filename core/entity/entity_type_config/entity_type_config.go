package entity_type_config

import (
	"encoding/json"

	"nem/core/entity/entity_type_dependent_config"
	"nem/core/entity/entity_type_standalone_config"

	"fmt"

	"nem/core/entity/types"

	"math/rand"

	"time"
)

type EntityTypeConfig struct {
	Standalone entity_type_standalone_config.EntityTypeStandaloneConfig `json:"standalone"`
	Dependent  entity_type_dependent_config.EntityTypeDependentConfig   `json:"dependent"`
}

func (e EntityTypeConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e EntityTypeConfig) EntityIdentifier() string {
	return "entity_type_config"
}

func (e EntityTypeConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["standalone"] = types.SingleDependantEntityFieldType
	res["dependent"] = types.SingleDependantEntityFieldType
	return res
}

func (e EntityTypeConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e EntityTypeConfig) IsDependant() bool {
	return true
}

func EntityTypeConfigFromJSON(data json.RawMessage) EntityTypeConfig {
	entity := EntityTypeConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func EntityTypeConfigSliceFromJSON(data json.RawMessage) []EntityTypeConfig {
	entity := []EntityTypeConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e EntityTypeConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error EntityTypeConfigToJSON", err)
	}
	return res
}

func EntityTypeConfigToJSON(e EntityTypeConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error EntityTypeConfigToJSON", err)
	}
	return res
}

func EntityTypeConfigSliceToJSON(e []EntityTypeConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error EntityTypeConfigSliceToJSON", err)
	}
	return res
}

func NewEntityTypeConfigWithRandomValues() EntityTypeConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return EntityTypeConfig{
		Standalone: entity_type_standalone_config.NewEntityTypeStandaloneConfigWithRandomValues(),
		Dependent:  entity_type_dependent_config.NewEntityTypeDependentConfigWithRandomValues(),
	}
}

func NewEntityTypeConfigSliceWithRandomValues(n int) []EntityTypeConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []EntityTypeConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewEntityTypeConfigWithRandomValues())
	}
	return res
}
