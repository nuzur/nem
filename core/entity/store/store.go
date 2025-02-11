package store

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"time"

	"fmt"

	"nem/core/entity/types"

	"nem/core/randomvalues"

	"math/rand"
)

type Store struct {
	UUID          uuid.UUID `json:"uuid"`
	Identifier    string    `json:"identifier"`
	Description   string    `json:"description"`
	Status        Status    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	CreatedByUUID uuid.UUID `json:"created_by_uuid"`
	UpdatedByUUID uuid.UUID `json:"updated_by_uuid"`
}

func (e Store) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e Store) EntityIdentifier() string {
	return "store"
}

func (e Store) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["uuid"] = types.UUIDFieldType
	res["identifier"] = types.StringFieldType
	res["description"] = types.StringFieldType
	res["status"] = types.SingleEnumFieldType
	res["created_at"] = types.TimestampFieldType
	res["updated_at"] = types.TimestampFieldType
	res["created_by_uuid"] = types.UUIDFieldType
	res["updated_by_uuid"] = types.UUIDFieldType
	return res
}

func (e Store) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e Store) IsDependant() bool {
	return true
}

func StoreFromJSON(data json.RawMessage) Store {
	entity := Store{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func StoreSliceFromJSON(data json.RawMessage) []Store {
	entity := []Store{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e Store) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error StoreToJSON", err)
	}
	return res
}

func StoreToJSON(e Store) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error StoreToJSON", err)
	}
	return res
}

func StoreSliceToJSON(e []Store) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error StoreSliceToJSON", err)
	}
	return res
}

func NewStoreWithRandomValues() Store {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return Store{
		UUID:          randomvalues.GetRandomUUIDValue(),
		Identifier:    randomvalues.GetRandomStringValue(),
		Description:   randomvalues.GetRandomStringValue(),
		Status:        randomvalues.GetRandomOptionValue[Status](2),
		CreatedAt:     randomvalues.GetRandomTimeValue(),
		UpdatedAt:     randomvalues.GetRandomTimeValue(),
		CreatedByUUID: randomvalues.GetRandomUUIDValue(),
		UpdatedByUUID: randomvalues.GetRandomUUIDValue(),
	}
}

func NewStoreSliceWithRandomValues(n int) []Store {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []Store{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewStoreWithRandomValues())
	}
	return res
}
