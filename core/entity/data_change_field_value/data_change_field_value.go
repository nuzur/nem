package data_change_field_value

import (
	"encoding/json"

	"github.com/gofrs/uuid"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"

	"time"
)

type DataChangeFieldValue struct {
	FieldUUID uuid.UUID `json:"field_uuid"`
	Value     string    `json:"value"`
}

func (e DataChangeFieldValue) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e DataChangeFieldValue) EntityIdentifier() string {
	return "data_change_field_value"
}

func (e DataChangeFieldValue) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["field_uuid"] = types.UUIDFieldType
	res["value"] = types.StringFieldType
	return res
}

func (e DataChangeFieldValue) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e DataChangeFieldValue) IsDependant() bool {
	return true
}

func DataChangeFieldValueFromJSON(data json.RawMessage) DataChangeFieldValue {
	entity := DataChangeFieldValue{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func DataChangeFieldValueSliceFromJSON(data json.RawMessage) []DataChangeFieldValue {
	entity := []DataChangeFieldValue{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e DataChangeFieldValue) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error DataChangeFieldValueToJSON", err)
	}
	return res
}

func DataChangeFieldValueToJSON(e DataChangeFieldValue) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error DataChangeFieldValueToJSON", err)
	}
	return res
}

func DataChangeFieldValueSliceToJSON(e []DataChangeFieldValue) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error DataChangeFieldValueSliceToJSON", err)
	}
	return res
}

func NewDataChangeFieldValueWithRandomValues() DataChangeFieldValue {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return DataChangeFieldValue{
		FieldUUID: randomvalues.GetRandomUUIDValue(),
		Value:     randomvalues.GetRandomStringValue(),
	}
}

func NewDataChangeFieldValueSliceWithRandomValues(n int) []DataChangeFieldValue {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []DataChangeFieldValue{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewDataChangeFieldValueWithRandomValues())
	}
	return res
}
