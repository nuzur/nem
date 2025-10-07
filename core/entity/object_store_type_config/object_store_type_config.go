package object_store_type_config

import (
	"encoding/json"

	"github.com/nuzur/nem/core/entity/object_store_s3_type_config"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"math/rand"

	"time"
)

type ObjectStoreTypeConfig struct {
	S3 object_store_s3_type_config.ObjectStoreS3TypeConfig `json:"s3"`
}

func (e ObjectStoreTypeConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e ObjectStoreTypeConfig) EntityIdentifier() string {
	return "object_store_type_config"
}

func (e ObjectStoreTypeConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["s3"] = types.SingleDependantEntityFieldType
	return res
}

func (e ObjectStoreTypeConfig) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "s3")

	return res
}

func (e ObjectStoreTypeConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e ObjectStoreTypeConfig) IsDependant() bool {
	return true
}

func ObjectStoreTypeConfigFromJSON(data json.RawMessage) ObjectStoreTypeConfig {
	entity := ObjectStoreTypeConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func ObjectStoreTypeConfigSliceFromJSON(data json.RawMessage) []ObjectStoreTypeConfig {
	entity := []ObjectStoreTypeConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e ObjectStoreTypeConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ObjectStoreTypeConfigToJSON", err)
	}
	return res
}

func ObjectStoreTypeConfigToJSON(e ObjectStoreTypeConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ObjectStoreTypeConfigToJSON", err)
	}
	return res
}

func ObjectStoreTypeConfigSliceToJSON(e []ObjectStoreTypeConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ObjectStoreTypeConfigSliceToJSON", err)
	}
	return res
}

func NewObjectStoreTypeConfigWithRandomValues() ObjectStoreTypeConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return ObjectStoreTypeConfig{
		S3: object_store_s3_type_config.NewObjectStoreS3TypeConfigWithRandomValues(),
	}
}

func NewObjectStoreTypeConfigSliceWithRandomValues(n int) []ObjectStoreTypeConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []ObjectStoreTypeConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewObjectStoreTypeConfigWithRandomValues())
	}
	return res
}
