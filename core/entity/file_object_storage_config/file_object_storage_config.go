package file_object_storage_config

import (
	"encoding/json"

	"github.com/gofrs/uuid"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"

	"time"
)

type FileObjectStorageConfig struct {
	ObjectStoreUUID uuid.UUID `json:"object_store_uuid"`
	Path            *string   `json:"path"`
}

func (e FileObjectStorageConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e FileObjectStorageConfig) EntityIdentifier() string {
	return "file_object_storage_config"
}

func (e FileObjectStorageConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["object_store_uuid"] = types.UUIDFieldType
	res["path"] = types.StringFieldType
	return res
}

func (e FileObjectStorageConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e FileObjectStorageConfig) IsDependant() bool {
	return true
}

func FileObjectStorageConfigFromJSON(data json.RawMessage) FileObjectStorageConfig {
	entity := FileObjectStorageConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func FileObjectStorageConfigSliceFromJSON(data json.RawMessage) []FileObjectStorageConfig {
	entity := []FileObjectStorageConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e FileObjectStorageConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FileObjectStorageConfigToJSON", err)
	}
	return res
}

func FileObjectStorageConfigToJSON(e FileObjectStorageConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FileObjectStorageConfigToJSON", err)
	}
	return res
}

func FileObjectStorageConfigSliceToJSON(e []FileObjectStorageConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error FileObjectStorageConfigSliceToJSON", err)
	}
	return res
}

func NewFileObjectStorageConfigWithRandomValues() FileObjectStorageConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return FileObjectStorageConfig{
		ObjectStoreUUID: randomvalues.GetRandomUUIDValue(),
		Path:            randomvalues.GetRandomStringValuePtr(),
	}
}

func NewFileObjectStorageConfigSliceWithRandomValues(n int) []FileObjectStorageConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []FileObjectStorageConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewFileObjectStorageConfigWithRandomValues())
	}
	return res
}
