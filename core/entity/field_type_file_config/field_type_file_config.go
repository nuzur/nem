package field_type_file_config

import (
	"encoding/json"

	"github.com/nuzur/nem/core/entity/file_storage_config"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"

	"time"
)

type FieldTypeFileConfig struct {
	StorageType       StorageType                           `json:"storage_type"`
	StorageConfig     file_storage_config.FileStorageConfig `json:"storage_config"`
	AllowedExtensions []string                              `json:"allowed_extensions"`
	MaxSize           int64                                 `json:"max_size"`
	AllowMultiple     bool                                  `json:"allow_multiple"`
	MaxFiles          int64                                 `json:"max_files"`
}

func (e FieldTypeFileConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e FieldTypeFileConfig) EntityIdentifier() string {
	return "field_type_file_config"
}

func (e FieldTypeFileConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["storage_type"] = types.SingleEnumFieldType
	res["storage_config"] = types.SingleDependantEntityFieldType
	res["allowed_extensions"] = types.ArrayFieldType
	res["max_size"] = types.IntFieldType
	res["allow_multiple"] = types.BooleanFieldType
	res["max_files"] = types.IntFieldType
	return res
}

func (e FieldTypeFileConfig) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "storage_type")
	res = append(res, "storage_config")
	res = append(res, "allowed_extensions")
	res = append(res, "max_size")
	res = append(res, "allow_multiple")
	res = append(res, "max_files")

	return res
}

func (e FieldTypeFileConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["allowed_extensions"] = types.StringFieldType
	return res
}

func (e FieldTypeFileConfig) IsDependant() bool {
	return true
}

func FieldTypeFileConfigFromJSON(data json.RawMessage) FieldTypeFileConfig {
	entity := FieldTypeFileConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func FieldTypeFileConfigSliceFromJSON(data json.RawMessage) []FieldTypeFileConfig {
	entity := []FieldTypeFileConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e FieldTypeFileConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeFileConfigToJSON", err)
	}
	return res
}

func FieldTypeFileConfigToJSON(e FieldTypeFileConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeFileConfigToJSON", err)
	}
	return res
}

func FieldTypeFileConfigSliceToJSON(e []FieldTypeFileConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FieldTypeFileConfigSliceToJSON", err)
	}
	return res
}

func NewFieldTypeFileConfigWithRandomValues() FieldTypeFileConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return FieldTypeFileConfig{
		StorageType:       randomvalues.GetRandomOptionValue[StorageType](2),
		StorageConfig:     file_storage_config.NewFileStorageConfigWithRandomValues(),
		AllowedExtensions: []string{},
		MaxSize:           randomvalues.GetRandomIntValue(),
		AllowMultiple:     randomvalues.GetRandomBoolValue(),
		MaxFiles:          randomvalues.GetRandomIntValue(),
	}
}

func NewFieldTypeFileConfigSliceWithRandomValues(n int) []FieldTypeFileConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []FieldTypeFileConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewFieldTypeFileConfigWithRandomValues())
	}
	return res
}
