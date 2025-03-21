package field_type_date_config

import (
	"encoding/json"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"

	"time"
)

type FieldTypeDateConfig struct {
	EnforceFuture   bool    `json:"enforce_future"`
	EnforcePast     bool    `json:"enforce_past"`
	DisplayTimezone *string `json:"display_timezone"`
	StorageTimezone *string `json:"storage_timezone"`
}

func (e FieldTypeDateConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e FieldTypeDateConfig) EntityIdentifier() string {
	return "field_type_date_config"
}

func (e FieldTypeDateConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["enforce_future"] = types.BooleanFieldType
	res["enforce_past"] = types.BooleanFieldType
	res["display_timezone"] = types.StringFieldType
	res["storage_timezone"] = types.StringFieldType
	return res
}

func (e FieldTypeDateConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e FieldTypeDateConfig) IsDependant() bool {
	return true
}

func FieldTypeDateConfigFromJSON(data json.RawMessage) FieldTypeDateConfig {
	entity := FieldTypeDateConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func FieldTypeDateConfigSliceFromJSON(data json.RawMessage) []FieldTypeDateConfig {
	entity := []FieldTypeDateConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e FieldTypeDateConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeDateConfigToJSON", err)
	}
	return res
}

func FieldTypeDateConfigToJSON(e FieldTypeDateConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeDateConfigToJSON", err)
	}
	return res
}

func FieldTypeDateConfigSliceToJSON(e []FieldTypeDateConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeDateConfigSliceToJSON", err)
	}
	return res
}

func NewFieldTypeDateConfigWithRandomValues() FieldTypeDateConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return FieldTypeDateConfig{
		EnforceFuture:   randomvalues.GetRandomBoolValue(),
		EnforcePast:     randomvalues.GetRandomBoolValue(),
		DisplayTimezone: randomvalues.GetRandomStringValuePtr(),
		StorageTimezone: randomvalues.GetRandomStringValuePtr(),
	}
}

func NewFieldTypeDateConfigSliceWithRandomValues(n int) []FieldTypeDateConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []FieldTypeDateConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewFieldTypeDateConfigWithRandomValues())
	}
	return res
}
