package enum

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"github.com/nuzur/nem/core/entity/enum_value"
	"time"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"
)

type Enum struct {
	UUID          uuid.UUID              `json:"uuid"`
	Version       int64                  `json:"version"`
	Identifier    string                 `json:"identifier"`
	StaticValues  []enum_value.EnumValue `json:"static_values"`
	RemoteValues  bool                   `json:"remote_values"`
	Status        Status                 `json:"status"`
	CreatedAt     time.Time              `json:"created_at"`
	UpdatedAt     time.Time              `json:"updated_at"`
	CreatedByUUID uuid.UUID              `json:"created_by_uuid"`
	UpdatedByUUID uuid.UUID              `json:"updated_by_uuid"`
}

func (e Enum) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e Enum) EntityIdentifier() string {
	return "enum"
}

func (e Enum) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["uuid"] = types.UUIDFieldType
	res["version"] = types.IntFieldType
	res["identifier"] = types.StringFieldType
	res["static_values"] = types.MultiDependantEntityFieldType
	res["remote_values"] = types.BooleanFieldType
	res["status"] = types.SingleEnumFieldType
	res["created_at"] = types.TimestampFieldType
	res["updated_at"] = types.TimestampFieldType
	res["created_by_uuid"] = types.UUIDFieldType
	res["updated_by_uuid"] = types.UUIDFieldType
	return res
}

func (e Enum) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e Enum) IsDependant() bool {
	return true
}

func EnumFromJSON(data json.RawMessage) Enum {
	entity := Enum{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func EnumSliceFromJSON(data json.RawMessage) []Enum {
	entity := []Enum{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e Enum) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error EnumToJSON", err)
	}
	return res
}

func EnumToJSON(e Enum) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error EnumToJSON", err)
	}
	return res
}

func EnumSliceToJSON(e []Enum) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error EnumSliceToJSON", err)
	}
	return res
}

func NewEnumWithRandomValues() Enum {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return Enum{
		UUID:          randomvalues.GetRandomUUIDValue(),
		Version:       randomvalues.GetRandomIntValue(),
		Identifier:    randomvalues.GetRandomStringValue(),
		StaticValues:  enum_value.NewEnumValueSliceWithRandomValues(rand.Intn(10)),
		RemoteValues:  randomvalues.GetRandomBoolValue(),
		Status:        randomvalues.GetRandomOptionValue[Status](2),
		CreatedAt:     randomvalues.GetRandomTimeValue(),
		UpdatedAt:     randomvalues.GetRandomTimeValue(),
		CreatedByUUID: randomvalues.GetRandomUUIDValue(),
		UpdatedByUUID: randomvalues.GetRandomUUIDValue(),
	}
}

func NewEnumSliceWithRandomValues(n int) []Enum {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []Enum{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewEnumWithRandomValues())
	}
	return res
}
