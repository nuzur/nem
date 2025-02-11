package object_store_s3_type_config

import (
	"encoding/json"

	"fmt"

	"nem/core/entity/types"

	"nem/core/randomvalues"

	"math/rand"

	"time"
)

type ObjectStoreS3TypeConfig struct {
	Region string `json:"region"`
	Key    string `json:"key"`
	Secret string `json:"secret"`
	Bucket string `json:"bucket"`
}

func (e ObjectStoreS3TypeConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e ObjectStoreS3TypeConfig) EntityIdentifier() string {
	return "object_store_s3_type_config"
}

func (e ObjectStoreS3TypeConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["region"] = types.StringFieldType
	res["key"] = types.StringFieldType
	res["secret"] = types.StringFieldType
	res["bucket"] = types.StringFieldType
	return res
}

func (e ObjectStoreS3TypeConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e ObjectStoreS3TypeConfig) IsDependant() bool {
	return true
}

func ObjectStoreS3TypeConfigFromJSON(data json.RawMessage) ObjectStoreS3TypeConfig {
	entity := ObjectStoreS3TypeConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func ObjectStoreS3TypeConfigSliceFromJSON(data json.RawMessage) []ObjectStoreS3TypeConfig {
	entity := []ObjectStoreS3TypeConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e ObjectStoreS3TypeConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ObjectStoreS3TypeConfigToJSON", err)
	}
	return res
}

func ObjectStoreS3TypeConfigToJSON(e ObjectStoreS3TypeConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ObjectStoreS3TypeConfigToJSON", err)
	}
	return res
}

func ObjectStoreS3TypeConfigSliceToJSON(e []ObjectStoreS3TypeConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ObjectStoreS3TypeConfigSliceToJSON", err)
	}
	return res
}

func NewObjectStoreS3TypeConfigWithRandomValues() ObjectStoreS3TypeConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return ObjectStoreS3TypeConfig{
		Region: randomvalues.GetRandomStringValue(),
		Key:    randomvalues.GetRandomStringValue(),
		Secret: randomvalues.GetRandomStringValue(),
		Bucket: randomvalues.GetRandomStringValue(),
	}
}

func NewObjectStoreS3TypeConfigSliceWithRandomValues(n int) []ObjectStoreS3TypeConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []ObjectStoreS3TypeConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewObjectStoreS3TypeConfigWithRandomValues())
	}
	return res
}
