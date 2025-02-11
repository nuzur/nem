package field_type_datetime_config

import (
	"encoding/json"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"

	"time"
)

type FieldTypeDatetimeConfig struct {
	EnforceFuture   bool   `json:"enforce_future"`
	EnforcePast     bool   `json:"enforce_past"`
	DisplayTimezone string `json:"display_timezone"`
	StorageTimezone string `json:"storage_timezone"`
}

func (e FieldTypeDatetimeConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e FieldTypeDatetimeConfig) EntityIdentifier() string {
	return "field_type_datetime_config"
}

func (e FieldTypeDatetimeConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["enforce_future"] = types.BooleanFieldType
	res["enforce_past"] = types.BooleanFieldType
	res["display_timezone"] = types.StringFieldType
	res["storage_timezone"] = types.StringFieldType
	return res
}

func (e FieldTypeDatetimeConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e FieldTypeDatetimeConfig) IsDependant() bool {
	return true
}

func FieldTypeDatetimeConfigFromJSON(data json.RawMessage) FieldTypeDatetimeConfig {
	entity := FieldTypeDatetimeConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func FieldTypeDatetimeConfigSliceFromJSON(data json.RawMessage) []FieldTypeDatetimeConfig {
	entity := []FieldTypeDatetimeConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e FieldTypeDatetimeConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeDatetimeConfigToJSON", err)
	}
	return res
}

func FieldTypeDatetimeConfigToJSON(e FieldTypeDatetimeConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeDatetimeConfigToJSON", err)
	}
	return res
}

func FieldTypeDatetimeConfigSliceToJSON(e []FieldTypeDatetimeConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeDatetimeConfigSliceToJSON", err)
	}
	return res
}

func NewFieldTypeDatetimeConfigWithRandomValues() FieldTypeDatetimeConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return FieldTypeDatetimeConfig{
		EnforceFuture:   randomvalues.GetRandomBoolValue(),
		EnforcePast:     randomvalues.GetRandomBoolValue(),
		DisplayTimezone: randomvalues.GetRandomStringValue(),
		StorageTimezone: randomvalues.GetRandomStringValue(),
	}
}

func NewFieldTypeDatetimeConfigSliceWithRandomValues(n int) []FieldTypeDatetimeConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []FieldTypeDatetimeConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewFieldTypeDatetimeConfigWithRandomValues())
	}
	return res
}
