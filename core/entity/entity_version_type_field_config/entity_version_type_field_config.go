package entity_version_type_field_config

import (
	"encoding/json"

	"github.com/gofrs/uuid"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"

	"time"
)

type EntityVersionTypeFieldConfig struct {
	VersionFieldUUID uuid.UUID `json:"version_field_uuid"`
	AppendOnly       bool      `json:"append_only"`
}

func (e EntityVersionTypeFieldConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e EntityVersionTypeFieldConfig) EntityIdentifier() string {
	return "entity_version_type_field_config"
}

func (e EntityVersionTypeFieldConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["version_field_uuid"] = types.UUIDFieldType
	res["append_only"] = types.BooleanFieldType
	return res
}

func (e EntityVersionTypeFieldConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e EntityVersionTypeFieldConfig) IsDependant() bool {
	return true
}

func EntityVersionTypeFieldConfigFromJSON(data json.RawMessage) EntityVersionTypeFieldConfig {
	entity := EntityVersionTypeFieldConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func EntityVersionTypeFieldConfigSliceFromJSON(data json.RawMessage) []EntityVersionTypeFieldConfig {
	entity := []EntityVersionTypeFieldConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e EntityVersionTypeFieldConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error EntityVersionTypeFieldConfigToJSON", err)
	}
	return res
}

func EntityVersionTypeFieldConfigToJSON(e EntityVersionTypeFieldConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error EntityVersionTypeFieldConfigToJSON", err)
	}
	return res
}

func EntityVersionTypeFieldConfigSliceToJSON(e []EntityVersionTypeFieldConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error EntityVersionTypeFieldConfigSliceToJSON", err)
	}
	return res
}

func NewEntityVersionTypeFieldConfigWithRandomValues() EntityVersionTypeFieldConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return EntityVersionTypeFieldConfig{
		VersionFieldUUID: randomvalues.GetRandomUUIDValue(),
		AppendOnly:       randomvalues.GetRandomBoolValue(),
	}
}

func NewEntityVersionTypeFieldConfigSliceWithRandomValues(n int) []EntityVersionTypeFieldConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []EntityVersionTypeFieldConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewEntityVersionTypeFieldConfigWithRandomValues())
	}
	return res
}
