package index

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"github.com/nuzur/nem/core/entity/index_field"
	"time"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"
)

type Index struct {
	UUID          uuid.UUID                `json:"uuid"`
	Identifier    string                   `json:"identifier"`
	Type          Type                     `json:"type"`
	Fields        []index_field.IndexField `json:"fields"`
	Status        Status                   `json:"status"`
	CreatedAt     time.Time                `json:"created_at"`
	UpdatedAt     time.Time                `json:"updated_at"`
	CreatedByUUID uuid.UUID                `json:"created_by_uuid"`
	UpdatedByUUID uuid.UUID                `json:"updated_by_uuid"`
}

func (e Index) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e Index) EntityIdentifier() string {
	return "index"
}

func (e Index) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["uuid"] = types.UUIDFieldType
	res["identifier"] = types.StringFieldType
	res["type"] = types.SingleEnumFieldType
	res["fields"] = types.MultiDependantEntityFieldType
	res["status"] = types.SingleEnumFieldType
	res["created_at"] = types.TimestampFieldType
	res["updated_at"] = types.TimestampFieldType
	res["created_by_uuid"] = types.UUIDFieldType
	res["updated_by_uuid"] = types.UUIDFieldType
	return res
}

func (e Index) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "uuid")
	res = append(res, "identifier")
	res = append(res, "type")
	res = append(res, "fields")
	res = append(res, "status")
	res = append(res, "created_at")
	res = append(res, "updated_at")
	res = append(res, "created_by_uuid")
	res = append(res, "updated_by_uuid")

	return res
}

func (e Index) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e Index) IsDependant() bool {
	return true
}

func IndexFromJSON(data json.RawMessage) Index {
	entity := Index{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func IndexSliceFromJSON(data json.RawMessage) []Index {
	entity := []Index{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e Index) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error IndexToJSON", err)
	}
	return res
}

func IndexToJSON(e Index) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error IndexToJSON", err)
	}
	return res
}

func IndexSliceToJSON(e []Index) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error IndexSliceToJSON", err)
	}
	return res
}

func NewIndexWithRandomValues() Index {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return Index{
		UUID:          randomvalues.GetRandomUUIDValue(),
		Identifier:    randomvalues.GetRandomStringValue(),
		Type:          randomvalues.GetRandomOptionValue[Type](4),
		Fields:        index_field.NewIndexFieldSliceWithRandomValues(rand.Intn(10)),
		Status:        randomvalues.GetRandomOptionValue[Status](2),
		CreatedAt:     randomvalues.GetRandomTimeValue(),
		UpdatedAt:     randomvalues.GetRandomTimeValue(),
		CreatedByUUID: randomvalues.GetRandomUUIDValue(),
		UpdatedByUUID: randomvalues.GetRandomUUIDValue(),
	}
}

func NewIndexSliceWithRandomValues(n int) []Index {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []Index{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewIndexWithRandomValues())
	}
	return res
}
