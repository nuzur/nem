package object_store

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"nem/core/entity/object_store_type_config"
	"time"

	"fmt"

	"nem/core/entity/types"

	"nem/core/randomvalues"

	"math/rand"
)

type ObjectStore struct {
	UUID          uuid.UUID                                      `json:"uuid"`
	Identifier    string                                         `json:"identifier"`
	Type          Type                                           `json:"type"`
	TypeConfig    object_store_type_config.ObjectStoreTypeConfig `json:"type_config"`
	Status        Status                                         `json:"status"`
	CreatedAt     time.Time                                      `json:"created_at"`
	UpdatedAt     time.Time                                      `json:"updated_at"`
	CreatedByUUID uuid.UUID                                      `json:"created_by_uuid"`
	UpdatedByUUID uuid.UUID                                      `json:"updated_by_uuid"`
}

func (e ObjectStore) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e ObjectStore) EntityIdentifier() string {
	return "object_store"
}

func (e ObjectStore) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["uuid"] = types.UUIDFieldType
	res["identifier"] = types.StringFieldType
	res["type"] = types.SingleEnumFieldType
	res["type_config"] = types.SingleDependantEntityFieldType
	res["status"] = types.SingleEnumFieldType
	res["created_at"] = types.TimestampFieldType
	res["updated_at"] = types.TimestampFieldType
	res["created_by_uuid"] = types.UUIDFieldType
	res["updated_by_uuid"] = types.UUIDFieldType
	return res
}

func (e ObjectStore) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e ObjectStore) IsDependant() bool {
	return true
}

func ObjectStoreFromJSON(data json.RawMessage) ObjectStore {
	entity := ObjectStore{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func ObjectStoreSliceFromJSON(data json.RawMessage) []ObjectStore {
	entity := []ObjectStore{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e ObjectStore) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ObjectStoreToJSON", err)
	}
	return res
}

func ObjectStoreToJSON(e ObjectStore) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ObjectStoreToJSON", err)
	}
	return res
}

func ObjectStoreSliceToJSON(e []ObjectStore) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ObjectStoreSliceToJSON", err)
	}
	return res
}

func NewObjectStoreWithRandomValues() ObjectStore {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return ObjectStore{
		UUID:          randomvalues.GetRandomUUIDValue(),
		Identifier:    randomvalues.GetRandomStringValue(),
		Type:          randomvalues.GetRandomOptionValue[Type](1),
		TypeConfig:    object_store_type_config.NewObjectStoreTypeConfigWithRandomValues(),
		Status:        randomvalues.GetRandomOptionValue[Status](2),
		CreatedAt:     randomvalues.GetRandomTimeValue(),
		UpdatedAt:     randomvalues.GetRandomTimeValue(),
		CreatedByUUID: randomvalues.GetRandomUUIDValue(),
		UpdatedByUUID: randomvalues.GetRandomUUIDValue(),
	}
}

func NewObjectStoreSliceWithRandomValues(n int) []ObjectStore {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []ObjectStore{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewObjectStoreWithRandomValues())
	}
	return res
}
