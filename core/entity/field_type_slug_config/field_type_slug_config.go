package field_type_slug_config

import (
	"encoding/json"

	"github.com/gofrs/uuid"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"

	"time"
)

type FieldTypeSlugConfig struct {
	MinSize           int64       `json:"min_size"`
	MaxSize           int64       `json:"max_size"`
	RegexValidation   *string     `json:"regex_validation"`
	BasedOnFieldUUIDs []uuid.UUID `json:"based_on_field_uuids"`
}

func (e FieldTypeSlugConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e FieldTypeSlugConfig) EntityIdentifier() string {
	return "field_type_slug_config"
}

func (e FieldTypeSlugConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["min_size"] = types.IntFieldType
	res["max_size"] = types.IntFieldType
	res["regex_validation"] = types.StringFieldType
	res["based_on_field_uuids"] = types.ArrayFieldType
	return res
}

func (e FieldTypeSlugConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["based_on_field_uuids"] = types.UUIDFieldType
	return res
}

func (e FieldTypeSlugConfig) IsDependant() bool {
	return true
}

func FieldTypeSlugConfigFromJSON(data json.RawMessage) FieldTypeSlugConfig {
	entity := FieldTypeSlugConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func FieldTypeSlugConfigSliceFromJSON(data json.RawMessage) []FieldTypeSlugConfig {
	entity := []FieldTypeSlugConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e FieldTypeSlugConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeSlugConfigToJSON", err)
	}
	return res
}

func FieldTypeSlugConfigToJSON(e FieldTypeSlugConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeSlugConfigToJSON", err)
	}
	return res
}

func FieldTypeSlugConfigSliceToJSON(e []FieldTypeSlugConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeSlugConfigSliceToJSON", err)
	}
	return res
}

func NewFieldTypeSlugConfigWithRandomValues() FieldTypeSlugConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return FieldTypeSlugConfig{
		MinSize:           randomvalues.GetRandomIntValue(),
		MaxSize:           randomvalues.GetRandomIntValue(),
		RegexValidation:   randomvalues.GetRandomStringValuePtr(),
		BasedOnFieldUUIDs: []uuid.UUID{},
	}
}

func NewFieldTypeSlugConfigSliceWithRandomValues(n int) []FieldTypeSlugConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []FieldTypeSlugConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewFieldTypeSlugConfigWithRandomValues())
	}
	return res
}
