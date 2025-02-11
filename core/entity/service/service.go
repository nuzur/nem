package service

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"github.com/nuzur/nem/core/entity/element_render"
	"time"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"
)

type Service struct {
	UUID          uuid.UUID                    `json:"uuid"`
	Version       int64                        `json:"version"`
	Identifier    string                       `json:"identifier"`
	Description   string                       `json:"description"`
	Render        element_render.ElementRender `json:"render"`
	Status        Status                       `json:"status"`
	CreatedAt     time.Time                    `json:"created_at"`
	UpdatedAt     time.Time                    `json:"updated_at"`
	CreatedByUUID uuid.UUID                    `json:"created_by_uuid"`
	UpdatedByUUID uuid.UUID                    `json:"updated_by_uuid"`
}

func (e Service) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e Service) EntityIdentifier() string {
	return "service"
}

func (e Service) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["uuid"] = types.UUIDFieldType
	res["version"] = types.IntFieldType
	res["identifier"] = types.StringFieldType
	res["description"] = types.StringFieldType
	res["render"] = types.SingleDependantEntityFieldType
	res["status"] = types.SingleEnumFieldType
	res["created_at"] = types.TimestampFieldType
	res["updated_at"] = types.TimestampFieldType
	res["created_by_uuid"] = types.UUIDFieldType
	res["updated_by_uuid"] = types.UUIDFieldType
	return res
}

func (e Service) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e Service) IsDependant() bool {
	return true
}

func ServiceFromJSON(data json.RawMessage) Service {
	entity := Service{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func ServiceSliceFromJSON(data json.RawMessage) []Service {
	entity := []Service{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e Service) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ServiceToJSON", err)
	}
	return res
}

func ServiceToJSON(e Service) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ServiceToJSON", err)
	}
	return res
}

func ServiceSliceToJSON(e []Service) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ServiceSliceToJSON", err)
	}
	return res
}

func NewServiceWithRandomValues() Service {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return Service{
		UUID:          randomvalues.GetRandomUUIDValue(),
		Version:       randomvalues.GetRandomIntValue(),
		Identifier:    randomvalues.GetRandomStringValue(),
		Description:   randomvalues.GetRandomStringValue(),
		Render:        element_render.NewElementRenderWithRandomValues(),
		Status:        randomvalues.GetRandomOptionValue[Status](2),
		CreatedAt:     randomvalues.GetRandomTimeValue(),
		UpdatedAt:     randomvalues.GetRandomTimeValue(),
		CreatedByUUID: randomvalues.GetRandomUUIDValue(),
		UpdatedByUUID: randomvalues.GetRandomUUIDValue(),
	}
}

func NewServiceSliceWithRandomValues(n int) []Service {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []Service{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewServiceWithRandomValues())
	}
	return res
}
