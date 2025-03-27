package data_change_type_config

import (
	"encoding/json"

	"github.com/nuzur/nem/core/entity/data_change_type_config_create"
	"github.com/nuzur/nem/core/entity/data_change_type_config_delete"
	"github.com/nuzur/nem/core/entity/data_change_type_config_update"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"math/rand"

	"time"
)

type DataChangeTypeConfig struct {
	Update data_change_type_config_update.DataChangeTypeConfigUpdate `json:"update"`
	Create data_change_type_config_create.DataChangeTypeConfigCreate `json:"create"`
	Delete data_change_type_config_delete.DataChangeTypeConfigDelete `json:"delete"`
}

func (e DataChangeTypeConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e DataChangeTypeConfig) EntityIdentifier() string {
	return "data_change_type_config"
}

func (e DataChangeTypeConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["update"] = types.SingleDependantEntityFieldType
	res["create"] = types.SingleDependantEntityFieldType
	res["delete"] = types.SingleDependantEntityFieldType
	return res
}

func (e DataChangeTypeConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e DataChangeTypeConfig) IsDependant() bool {
	return true
}

func DataChangeTypeConfigFromJSON(data json.RawMessage) DataChangeTypeConfig {
	entity := DataChangeTypeConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func DataChangeTypeConfigSliceFromJSON(data json.RawMessage) []DataChangeTypeConfig {
	entity := []DataChangeTypeConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e DataChangeTypeConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error DataChangeTypeConfigToJSON", err)
	}
	return res
}

func DataChangeTypeConfigToJSON(e DataChangeTypeConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error DataChangeTypeConfigToJSON", err)
	}
	return res
}

func DataChangeTypeConfigSliceToJSON(e []DataChangeTypeConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error DataChangeTypeConfigSliceToJSON", err)
	}
	return res
}

func NewDataChangeTypeConfigWithRandomValues() DataChangeTypeConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return DataChangeTypeConfig{
		Update: data_change_type_config_update.NewDataChangeTypeConfigUpdateWithRandomValues(),
		Create: data_change_type_config_create.NewDataChangeTypeConfigCreateWithRandomValues(),
		Delete: data_change_type_config_delete.NewDataChangeTypeConfigDeleteWithRandomValues(),
	}
}

func NewDataChangeTypeConfigSliceWithRandomValues(n int) []DataChangeTypeConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []DataChangeTypeConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewDataChangeTypeConfigWithRandomValues())
	}
	return res
}
