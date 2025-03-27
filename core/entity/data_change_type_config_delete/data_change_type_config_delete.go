package data_change_type_config_delete

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

type DataChangeTypeConfigDelete struct {
	EntityUUID uuid.UUID                                      `json:"entity_uuid"`
	Keys       []data_change_field_value.DataChangeFieldValue `json:"keys"`
	CreatedAt  time.Time                                      `json:"created_at"`
}

func (e DataChangeTypeConfigDelete) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e DataChangeTypeConfigDelete) EntityIdentifier() string {
	return "data_change_type_config_delete"
}

func (e DataChangeTypeConfigDelete) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["entity_uuid"] = types.UUIDFieldType
	res["keys"] = types.MultiDependantEntityFieldType
	res["created_at"] = types.TimestampFieldType
	return res
}

func (e DataChangeTypeConfigDelete) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e DataChangeTypeConfigDelete) IsDependant() bool {
	return true
}

func DataChangeTypeConfigDeleteFromJSON(data json.RawMessage) DataChangeTypeConfigDelete {
	entity := DataChangeTypeConfigDelete{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func DataChangeTypeConfigDeleteSliceFromJSON(data json.RawMessage) []DataChangeTypeConfigDelete {
	entity := []DataChangeTypeConfigDelete{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e DataChangeTypeConfigDelete) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error DataChangeTypeConfigDeleteToJSON", err)
	}
	return res
}

func DataChangeTypeConfigDeleteToJSON(e DataChangeTypeConfigDelete) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error DataChangeTypeConfigDeleteToJSON", err)
	}
	return res
}

func DataChangeTypeConfigDeleteSliceToJSON(e []DataChangeTypeConfigDelete) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error DataChangeTypeConfigDeleteSliceToJSON", err)
	}
	return res
}

func NewDataChangeTypeConfigDeleteWithRandomValues() DataChangeTypeConfigDelete {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return DataChangeTypeConfigDelete{
		EntityUUID: randomvalues.GetRandomUUIDValue(),
		Keys:       data_change_field_value.NewDataChangeFieldValueSliceWithRandomValues(rand.Intn(10)),
		CreatedAt:  randomvalues.GetRandomTimeValue(),
	}
}

func NewDataChangeTypeConfigDeleteSliceWithRandomValues(n int) []DataChangeTypeConfigDelete {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []DataChangeTypeConfigDelete{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewDataChangeTypeConfigDeleteWithRandomValues())
	}
	return res
}
