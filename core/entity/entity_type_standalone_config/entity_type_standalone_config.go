package entity_type_standalone_config

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"github.com/nuzur/nem/core/entity/entity_version_config"
	"github.com/nuzur/nem/core/entity/index"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"

	"time"
)

type EntityTypeStandaloneConfig struct {
	StoreUUID     uuid.UUID                                 `json:"store_uuid"`
	Versioned     bool                                      `json:"versioned"`
	VersionConfig entity_version_config.EntityVersionConfig `json:"version_config"`
	Indexes       []index.Index                             `json:"indexes"`
}

func (e EntityTypeStandaloneConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e EntityTypeStandaloneConfig) EntityIdentifier() string {
	return "entity_type_standalone_config"
}

func (e EntityTypeStandaloneConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["store_uuid"] = types.UUIDFieldType
	res["versioned"] = types.BooleanFieldType
	res["version_config"] = types.SingleDependantEntityFieldType
	res["indexes"] = types.MultiDependantEntityFieldType
	return res
}

func (e EntityTypeStandaloneConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e EntityTypeStandaloneConfig) IsDependant() bool {
	return true
}

func EntityTypeStandaloneConfigFromJSON(data json.RawMessage) EntityTypeStandaloneConfig {
	entity := EntityTypeStandaloneConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func EntityTypeStandaloneConfigSliceFromJSON(data json.RawMessage) []EntityTypeStandaloneConfig {
	entity := []EntityTypeStandaloneConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e EntityTypeStandaloneConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error EntityTypeStandaloneConfigToJSON", err)
	}
	return res
}

func EntityTypeStandaloneConfigToJSON(e EntityTypeStandaloneConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error EntityTypeStandaloneConfigToJSON", err)
	}
	return res
}

func EntityTypeStandaloneConfigSliceToJSON(e []EntityTypeStandaloneConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error EntityTypeStandaloneConfigSliceToJSON", err)
	}
	return res
}

func NewEntityTypeStandaloneConfigWithRandomValues() EntityTypeStandaloneConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return EntityTypeStandaloneConfig{
		StoreUUID:     randomvalues.GetRandomUUIDValue(),
		Versioned:     randomvalues.GetRandomBoolValue(),
		VersionConfig: entity_version_config.NewEntityVersionConfigWithRandomValues(),
		Indexes:       index.NewIndexSliceWithRandomValues(rand.Intn(10)),
	}
}

func NewEntityTypeStandaloneConfigSliceWithRandomValues(n int) []EntityTypeStandaloneConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []EntityTypeStandaloneConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewEntityTypeStandaloneConfigWithRandomValues())
	}
	return res
}
