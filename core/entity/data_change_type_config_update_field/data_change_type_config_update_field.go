package data_change_type_config_update_field

import (
	"encoding/json"

	"github.com/gofrs/uuid"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"

	"time"
)

type DataChangeTypeConfigUpdateField struct {
	FieldUUID    uuid.UUID `json:"field_uuid"`
	CurrentValue string    `json:"current_value"`
	NewValue     string    `json:"new_value"`
}

func (e DataChangeTypeConfigUpdateField) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e DataChangeTypeConfigUpdateField) EntityIdentifier() string {
	return "data_change_type_config_update_field"
}

func (e DataChangeTypeConfigUpdateField) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["field_uuid"] = types.UUIDFieldType
	res["current_value"] = types.StringFieldType
	res["new_value"] = types.StringFieldType
	return res
}

func (e DataChangeTypeConfigUpdateField) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "field_uuid")
	res = append(res, "current_value")
	res = append(res, "new_value")

	return res
}

func (e DataChangeTypeConfigUpdateField) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e DataChangeTypeConfigUpdateField) IsDependant() bool {
	return true
}

func DataChangeTypeConfigUpdateFieldFromJSON(data json.RawMessage) DataChangeTypeConfigUpdateField {
	entity := DataChangeTypeConfigUpdateField{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func DataChangeTypeConfigUpdateFieldSliceFromJSON(data json.RawMessage) []DataChangeTypeConfigUpdateField {
	entity := []DataChangeTypeConfigUpdateField{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e DataChangeTypeConfigUpdateField) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error DataChangeTypeConfigUpdateFieldToJSON", err)
	}
	return res
}

func DataChangeTypeConfigUpdateFieldToJSON(e DataChangeTypeConfigUpdateField) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error DataChangeTypeConfigUpdateFieldToJSON", err)
	}
	return res
}

func DataChangeTypeConfigUpdateFieldSliceToJSON(e []DataChangeTypeConfigUpdateField) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error DataChangeTypeConfigUpdateFieldSliceToJSON", err)
	}
	return res
}

func NewDataChangeTypeConfigUpdateFieldWithRandomValues() DataChangeTypeConfigUpdateField {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return DataChangeTypeConfigUpdateField{
		FieldUUID:    randomvalues.GetRandomUUIDValue(),
		CurrentValue: randomvalues.GetRandomStringValue(),
		NewValue:     randomvalues.GetRandomStringValue(),
	}
}

func NewDataChangeTypeConfigUpdateFieldSliceWithRandomValues(n int) []DataChangeTypeConfigUpdateField {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []DataChangeTypeConfigUpdateField{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewDataChangeTypeConfigUpdateFieldWithRandomValues())
	}
	return res
}
