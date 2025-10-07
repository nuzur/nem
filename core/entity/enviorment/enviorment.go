package enviorment

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"time"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"
)

type Enviorment struct {
	UUID          uuid.UUID `json:"uuid"`
	Identifier    string    `json:"identifier"`
	Critical      bool      `json:"critical"`
	Status        Status    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	CreatedByUUID uuid.UUID `json:"created_by_uuid"`
	UpdatedByUUID uuid.UUID `json:"updated_by_uuid"`
}

func (e Enviorment) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e Enviorment) EntityIdentifier() string {
	return "enviorment"
}

func (e Enviorment) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["uuid"] = types.UUIDFieldType
	res["identifier"] = types.StringFieldType
	res["critical"] = types.BooleanFieldType
	res["status"] = types.SingleEnumFieldType
	res["created_at"] = types.TimestampFieldType
	res["updated_at"] = types.TimestampFieldType
	res["created_by_uuid"] = types.UUIDFieldType
	res["updated_by_uuid"] = types.UUIDFieldType
	return res
}

func (e Enviorment) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "uuid")
	res = append(res, "identifier")
	res = append(res, "critical")
	res = append(res, "status")
	res = append(res, "created_at")
	res = append(res, "updated_at")
	res = append(res, "created_by_uuid")
	res = append(res, "updated_by_uuid")

	return res
}

func (e Enviorment) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e Enviorment) IsDependant() bool {
	return true
}

func EnviormentFromJSON(data json.RawMessage) Enviorment {
	entity := Enviorment{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func EnviormentSliceFromJSON(data json.RawMessage) []Enviorment {
	entity := []Enviorment{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e Enviorment) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error EnviormentToJSON", err)
	}
	return res
}

func EnviormentToJSON(e Enviorment) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error EnviormentToJSON", err)
	}
	return res
}

func EnviormentSliceToJSON(e []Enviorment) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error EnviormentSliceToJSON", err)
	}
	return res
}

func NewEnviormentWithRandomValues() Enviorment {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return Enviorment{
		UUID:          randomvalues.GetRandomUUIDValue(),
		Identifier:    randomvalues.GetRandomStringValue(),
		Critical:      randomvalues.GetRandomBoolValue(),
		Status:        randomvalues.GetRandomOptionValue[Status](2),
		CreatedAt:     randomvalues.GetRandomTimeValue(),
		UpdatedAt:     randomvalues.GetRandomTimeValue(),
		CreatedByUUID: randomvalues.GetRandomUUIDValue(),
		UpdatedByUUID: randomvalues.GetRandomUUIDValue(),
	}
}

func NewEnviormentSliceWithRandomValues(n int) []Enviorment {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []Enviorment{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewEnviormentWithRandomValues())
	}
	return res
}
