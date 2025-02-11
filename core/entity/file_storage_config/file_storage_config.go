package file_storage_config

import (
	"encoding/json"

	"nem/core/entity/file_object_storage_config"

	"fmt"

	"nem/core/entity/types"

	"math/rand"

	"time"
)

type FileStorageConfig struct {
	ObjectStore file_object_storage_config.FileObjectStorageConfig `json:"object_store"`
}

func (e FileStorageConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e FileStorageConfig) EntityIdentifier() string {
	return "file_storage_config"
}

func (e FileStorageConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["object_store"] = types.SingleDependantEntityFieldType
	return res
}

func (e FileStorageConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e FileStorageConfig) IsDependant() bool {
	return true
}

func FileStorageConfigFromJSON(data json.RawMessage) FileStorageConfig {
	entity := FileStorageConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func FileStorageConfigSliceFromJSON(data json.RawMessage) []FileStorageConfig {
	entity := []FileStorageConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e FileStorageConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FileStorageConfigToJSON", err)
	}
	return res
}

func FileStorageConfigToJSON(e FileStorageConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FileStorageConfigToJSON", err)
	}
	return res
}

func FileStorageConfigSliceToJSON(e []FileStorageConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FileStorageConfigSliceToJSON", err)
	}
	return res
}

func NewFileStorageConfigWithRandomValues() FileStorageConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return FileStorageConfig{
		ObjectStore: file_object_storage_config.NewFileObjectStorageConfigWithRandomValues(),
	}
}

func NewFileStorageConfigSliceWithRandomValues(n int) []FileStorageConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []FileStorageConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewFileStorageConfigWithRandomValues())
	}
	return res
}
