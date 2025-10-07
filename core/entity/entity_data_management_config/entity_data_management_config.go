package entity_data_management_config

import (
	"encoding/json"

	"github.com/gofrs/uuid"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"math/rand"

	"time"
)

type EntityDataManagementConfig struct {
	ListDisplayFields []uuid.UUID `json:"list_display_fields"`
}

func (e EntityDataManagementConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e EntityDataManagementConfig) EntityIdentifier() string {
	return "entity_data_management_config"
}

func (e EntityDataManagementConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["list_display_fields"] = types.ArrayFieldType
	return res
}

func (e EntityDataManagementConfig) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "list_display_fields")

	return res
}

func (e EntityDataManagementConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["list_display_fields"] = types.UUIDFieldType
	return res
}

func (e EntityDataManagementConfig) IsDependant() bool {
	return true
}

func EntityDataManagementConfigFromJSON(data json.RawMessage) EntityDataManagementConfig {
	entity := EntityDataManagementConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func EntityDataManagementConfigSliceFromJSON(data json.RawMessage) []EntityDataManagementConfig {
	entity := []EntityDataManagementConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e EntityDataManagementConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error EntityDataManagementConfigToJSON", err)
	}
	return res
}

func EntityDataManagementConfigToJSON(e EntityDataManagementConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error EntityDataManagementConfigToJSON", err)
	}
	return res
}

func EntityDataManagementConfigSliceToJSON(e []EntityDataManagementConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error EntityDataManagementConfigSliceToJSON", err)
	}
	return res
}

func NewEntityDataManagementConfigWithRandomValues() EntityDataManagementConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return EntityDataManagementConfig{
		ListDisplayFields: []uuid.UUID{},
	}
}

func NewEntityDataManagementConfigSliceWithRandomValues(n int) []EntityDataManagementConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []EntityDataManagementConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewEntityDataManagementConfigWithRandomValues())
	}
	return res
}
