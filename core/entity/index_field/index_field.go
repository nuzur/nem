package index_field

import (
	"encoding/json"

	"github.com/gofrs/uuid"

	"fmt"

	"nem/core/entity/types"

	"nem/core/randomvalues"

	"math/rand"

	"time"
)

type IndexField struct {
	FieldUUID uuid.UUID `json:"field_uuid"`
	Priority  int64     `json:"priority"`
	Order     Order     `json:"order"`
	Length    int64     `json:"length"`
}

func (e IndexField) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e IndexField) EntityIdentifier() string {
	return "index_field"
}

func (e IndexField) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["field_uuid"] = types.UUIDFieldType
	res["priority"] = types.IntFieldType
	res["order"] = types.SingleEnumFieldType
	res["length"] = types.IntFieldType
	return res
}

func (e IndexField) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e IndexField) IsDependant() bool {
	return true
}

func IndexFieldFromJSON(data json.RawMessage) IndexField {
	entity := IndexField{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func IndexFieldSliceFromJSON(data json.RawMessage) []IndexField {
	entity := []IndexField{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e IndexField) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error IndexFieldToJSON", err)
	}
	return res
}

func IndexFieldToJSON(e IndexField) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error IndexFieldToJSON", err)
	}
	return res
}

func IndexFieldSliceToJSON(e []IndexField) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error IndexFieldSliceToJSON", err)
	}
	return res
}

func NewIndexFieldWithRandomValues() IndexField {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return IndexField{
		FieldUUID: randomvalues.GetRandomUUIDValue(),
		Priority:  randomvalues.GetRandomIntValue(),
		Order:     randomvalues.GetRandomOptionValue[Order](2),
		Length:    randomvalues.GetRandomIntValue(),
	}
}

func NewIndexFieldSliceWithRandomValues(n int) []IndexField {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []IndexField{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewIndexFieldWithRandomValues())
	}
	return res
}
