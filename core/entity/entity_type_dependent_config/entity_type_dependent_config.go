package entity_type_dependent_config

import (
	"encoding/json"

	"github.com/gofrs/uuid"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"math/rand"

	"time"
)

type EntityTypeDependentConfig struct {
	ServiceSourceUUIDs []uuid.UUID `json:"service_source_uuids"`
	EntitySourceUUIDs  []uuid.UUID `json:"entity_source_uuids"`
}

func (e EntityTypeDependentConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e EntityTypeDependentConfig) EntityIdentifier() string {
	return "entity_type_dependent_config"
}

func (e EntityTypeDependentConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["service_source_uuids"] = types.ArrayFieldType
	res["entity_source_uuids"] = types.ArrayFieldType
	return res
}

func (e EntityTypeDependentConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["service_source_uuids"] = types.UUIDFieldType
	res["entity_source_uuids"] = types.UUIDFieldType
	return res
}

func (e EntityTypeDependentConfig) IsDependant() bool {
	return true
}

func EntityTypeDependentConfigFromJSON(data json.RawMessage) EntityTypeDependentConfig {
	entity := EntityTypeDependentConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func EntityTypeDependentConfigSliceFromJSON(data json.RawMessage) []EntityTypeDependentConfig {
	entity := []EntityTypeDependentConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e EntityTypeDependentConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error EntityTypeDependentConfigToJSON", err)
	}
	return res
}

func EntityTypeDependentConfigToJSON(e EntityTypeDependentConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error EntityTypeDependentConfigToJSON", err)
	}
	return res
}

func EntityTypeDependentConfigSliceToJSON(e []EntityTypeDependentConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error EntityTypeDependentConfigSliceToJSON", err)
	}
	return res
}

func NewEntityTypeDependentConfigWithRandomValues() EntityTypeDependentConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return EntityTypeDependentConfig{
		ServiceSourceUUIDs: []uuid.UUID{},
		EntitySourceUUIDs:  []uuid.UUID{},
	}
}

func NewEntityTypeDependentConfigSliceWithRandomValues(n int) []EntityTypeDependentConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []EntityTypeDependentConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewEntityTypeDependentConfigWithRandomValues())
	}
	return res
}
