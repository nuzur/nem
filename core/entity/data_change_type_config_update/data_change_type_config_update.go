package data_change_type_config_update

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"github.com/nuzur/nem/core/entity/data_change_field_value"
	"github.com/nuzur/nem/core/entity/data_change_type_config_update_field"
	"time"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"
)

type DataChangeTypeConfigUpdate struct {
	EntityUUID uuid.UUID                                                              `json:"entity_uuid"`
	Fields     []data_change_type_config_update_field.DataChangeTypeConfigUpdateField `json:"fields"`
	Keys       []data_change_field_value.DataChangeFieldValue                         `json:"keys"`
	CreatedAt  time.Time                                                              `json:"created_at"`
}

func (e DataChangeTypeConfigUpdate) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e DataChangeTypeConfigUpdate) EntityIdentifier() string {
	return "data_change_type_config_update"
}

func (e DataChangeTypeConfigUpdate) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["entity_uuid"] = types.UUIDFieldType
	res["fields"] = types.MultiDependantEntityFieldType
	res["keys"] = types.MultiDependantEntityFieldType
	res["created_at"] = types.TimestampFieldType
	return res
}

func (e DataChangeTypeConfigUpdate) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "entity_uuid")
	res = append(res, "fields")
	res = append(res, "keys")
	res = append(res, "created_at")

	return res
}

func (e DataChangeTypeConfigUpdate) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e DataChangeTypeConfigUpdate) IsDependant() bool {
	return true
}

func DataChangeTypeConfigUpdateFromJSON(data json.RawMessage) DataChangeTypeConfigUpdate {
	entity := DataChangeTypeConfigUpdate{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func DataChangeTypeConfigUpdateSliceFromJSON(data json.RawMessage) []DataChangeTypeConfigUpdate {
	entity := []DataChangeTypeConfigUpdate{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e DataChangeTypeConfigUpdate) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error DataChangeTypeConfigUpdateToJSON", err)
	}
	return res
}

func DataChangeTypeConfigUpdateToJSON(e DataChangeTypeConfigUpdate) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error DataChangeTypeConfigUpdateToJSON", err)
	}
	return res
}

func DataChangeTypeConfigUpdateSliceToJSON(e []DataChangeTypeConfigUpdate) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error DataChangeTypeConfigUpdateSliceToJSON", err)
	}
	return res
}

func NewDataChangeTypeConfigUpdateWithRandomValues() DataChangeTypeConfigUpdate {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return DataChangeTypeConfigUpdate{
		EntityUUID: randomvalues.GetRandomUUIDValue(),
		Fields:     data_change_type_config_update_field.NewDataChangeTypeConfigUpdateFieldSliceWithRandomValues(rand.Intn(10)),
		Keys:       data_change_field_value.NewDataChangeFieldValueSliceWithRandomValues(rand.Intn(10)),
		CreatedAt:  randomvalues.GetRandomTimeValue(),
	}
}

func NewDataChangeTypeConfigUpdateSliceWithRandomValues(n int) []DataChangeTypeConfigUpdate {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []DataChangeTypeConfigUpdate{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewDataChangeTypeConfigUpdateWithRandomValues())
	}
	return res
}
