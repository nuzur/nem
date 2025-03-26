package data_change_type_config_create

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"github.com/nuzur/nem/core/entity/data_change_field_value"
	"time"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"
)

type DataChangeTypeConfigCreate struct {
	EntityUUID uuid.UUID                                      `json:"entity_uuid"`
	Keys       []data_change_field_value.DataChangeFieldValue `json:"keys"`
	CreatedAt  time.Time                                      `json:"created_at"`
}

func (e DataChangeTypeConfigCreate) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e DataChangeTypeConfigCreate) EntityIdentifier() string {
	return "data_change_type_config_create"
}

func (e DataChangeTypeConfigCreate) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["entity_uuid"] = types.UUIDFieldType
	res["keys"] = types.MultiDependantEntityFieldType
	res["created_at"] = types.TimestampFieldType
	return res
}

func (e DataChangeTypeConfigCreate) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e DataChangeTypeConfigCreate) IsDependant() bool {
	return true
}

func DataChangeTypeConfigCreateFromJSON(data json.RawMessage) DataChangeTypeConfigCreate {
	entity := DataChangeTypeConfigCreate{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func DataChangeTypeConfigCreateSliceFromJSON(data json.RawMessage) []DataChangeTypeConfigCreate {
	entity := []DataChangeTypeConfigCreate{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e DataChangeTypeConfigCreate) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error DataChangeTypeConfigCreateToJSON", err)
	}
	return res
}

func DataChangeTypeConfigCreateToJSON(e DataChangeTypeConfigCreate) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error DataChangeTypeConfigCreateToJSON", err)
	}
	return res
}

func DataChangeTypeConfigCreateSliceToJSON(e []DataChangeTypeConfigCreate) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error DataChangeTypeConfigCreateSliceToJSON", err)
	}
	return res
}

func NewDataChangeTypeConfigCreateWithRandomValues() DataChangeTypeConfigCreate {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return DataChangeTypeConfigCreate{
		EntityUUID: randomvalues.GetRandomUUIDValue(),
		Keys:       data_change_field_value.NewDataChangeFieldValueSliceWithRandomValues(rand.Intn(10)),
		CreatedAt:  randomvalues.GetRandomTimeValue(),
	}
}

func NewDataChangeTypeConfigCreateSliceWithRandomValues(n int) []DataChangeTypeConfigCreate {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []DataChangeTypeConfigCreate{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewDataChangeTypeConfigCreateWithRandomValues())
	}
	return res
}
