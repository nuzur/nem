package change_request_data_change

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"github.com/nuzur/nem/core/entity/data_change_type_config"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"

	"time"
)

type ChangeRequestDataChange struct {
	UUID                 uuid.UUID                                    `json:"uuid"`
	DataChangeType       DataChangeType                               `json:"data_change_type"`
	DataChangeTypeConfig data_change_type_config.DataChangeTypeConfig `json:"data_change_type_config"`
	Order                int64                                        `json:"order"`
}

func (e ChangeRequestDataChange) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e ChangeRequestDataChange) EntityIdentifier() string {
	return "change_request_data_change"
}

func (e ChangeRequestDataChange) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["uuid"] = types.UUIDFieldType
	res["data_change_type"] = types.SingleEnumFieldType
	res["data_change_type_config"] = types.SingleDependantEntityFieldType
	res["order"] = types.IntFieldType
	return res
}

func (e ChangeRequestDataChange) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e ChangeRequestDataChange) IsDependant() bool {
	return true
}

func ChangeRequestDataChangeFromJSON(data json.RawMessage) ChangeRequestDataChange {
	entity := ChangeRequestDataChange{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func ChangeRequestDataChangeSliceFromJSON(data json.RawMessage) []ChangeRequestDataChange {
	entity := []ChangeRequestDataChange{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e ChangeRequestDataChange) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ChangeRequestDataChangeToJSON", err)
	}
	return res
}

func ChangeRequestDataChangeToJSON(e ChangeRequestDataChange) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ChangeRequestDataChangeToJSON", err)
	}
	return res
}

func ChangeRequestDataChangeSliceToJSON(e []ChangeRequestDataChange) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ChangeRequestDataChangeSliceToJSON", err)
	}
	return res
}

func NewChangeRequestDataChangeWithRandomValues() ChangeRequestDataChange {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return ChangeRequestDataChange{
		UUID:                 randomvalues.GetRandomUUIDValue(),
		DataChangeType:       randomvalues.GetRandomOptionValue[DataChangeType](3),
		DataChangeTypeConfig: data_change_type_config.NewDataChangeTypeConfigWithRandomValues(),
		Order:                randomvalues.GetRandomIntValue(),
	}
}

func NewChangeRequestDataChangeSliceWithRandomValues(n int) []ChangeRequestDataChange {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []ChangeRequestDataChange{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewChangeRequestDataChangeWithRandomValues())
	}
	return res
}
