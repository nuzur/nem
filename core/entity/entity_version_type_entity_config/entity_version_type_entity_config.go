package entity_version_type_entity_config

import (
	"encoding/json"

	"github.com/gofrs/uuid"

	"fmt"

	"nem/core/entity/types"

	"nem/core/randomvalues"

	"math/rand"

	"time"
)

type EntityVersionTypeEntityConfig struct {
	EntityUUID       uuid.UUID `json:"entity_uuid"`
	VersionFieldUUID uuid.UUID `json:"version_field_uuid"`
	AppendOnly       bool      `json:"append_only"`
}

func (e EntityVersionTypeEntityConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e EntityVersionTypeEntityConfig) EntityIdentifier() string {
	return "entity_version_type_entity_config"
}

func (e EntityVersionTypeEntityConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["entity_uuid"] = types.UUIDFieldType
	res["version_field_uuid"] = types.UUIDFieldType
	res["append_only"] = types.BooleanFieldType
	return res
}

func (e EntityVersionTypeEntityConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e EntityVersionTypeEntityConfig) IsDependant() bool {
	return true
}

func EntityVersionTypeEntityConfigFromJSON(data json.RawMessage) EntityVersionTypeEntityConfig {
	entity := EntityVersionTypeEntityConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func EntityVersionTypeEntityConfigSliceFromJSON(data json.RawMessage) []EntityVersionTypeEntityConfig {
	entity := []EntityVersionTypeEntityConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e EntityVersionTypeEntityConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error EntityVersionTypeEntityConfigToJSON", err)
	}
	return res
}

func EntityVersionTypeEntityConfigToJSON(e EntityVersionTypeEntityConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error EntityVersionTypeEntityConfigToJSON", err)
	}
	return res
}

func EntityVersionTypeEntityConfigSliceToJSON(e []EntityVersionTypeEntityConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error EntityVersionTypeEntityConfigSliceToJSON", err)
	}
	return res
}

func NewEntityVersionTypeEntityConfigWithRandomValues() EntityVersionTypeEntityConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return EntityVersionTypeEntityConfig{
		EntityUUID:       randomvalues.GetRandomUUIDValue(),
		VersionFieldUUID: randomvalues.GetRandomUUIDValue(),
		AppendOnly:       randomvalues.GetRandomBoolValue(),
	}
}

func NewEntityVersionTypeEntityConfigSliceWithRandomValues(n int) []EntityVersionTypeEntityConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []EntityVersionTypeEntityConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewEntityVersionTypeEntityConfigWithRandomValues())
	}
	return res
}
