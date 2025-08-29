package change_request_metadata_data

import (
	"encoding/json"

	"github.com/gofrs/uuid"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"

	"time"
)

type ChangeRequestMetadataData struct {
	StoreUUID      uuid.UUID `json:"store_uuid"`
	ConnectionUUID uuid.UUID `json:"connection_uuid"`
	Schema         *string   `json:"schema"`
}

func (e ChangeRequestMetadataData) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e ChangeRequestMetadataData) EntityIdentifier() string {
	return "change_request_metadata_data"
}

func (e ChangeRequestMetadataData) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["store_uuid"] = types.UUIDFieldType
	res["connection_uuid"] = types.UUIDFieldType
	res["schema"] = types.StringFieldType
	return res
}

func (e ChangeRequestMetadataData) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e ChangeRequestMetadataData) IsDependant() bool {
	return true
}

func ChangeRequestMetadataDataFromJSON(data json.RawMessage) ChangeRequestMetadataData {
	entity := ChangeRequestMetadataData{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func ChangeRequestMetadataDataSliceFromJSON(data json.RawMessage) []ChangeRequestMetadataData {
	entity := []ChangeRequestMetadataData{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e ChangeRequestMetadataData) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ChangeRequestMetadataDataToJSON", err)
	}
	return res
}

func ChangeRequestMetadataDataToJSON(e ChangeRequestMetadataData) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ChangeRequestMetadataDataToJSON", err)
	}
	return res
}

func ChangeRequestMetadataDataSliceToJSON(e []ChangeRequestMetadataData) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ChangeRequestMetadataDataSliceToJSON", err)
	}
	return res
}

func NewChangeRequestMetadataDataWithRandomValues() ChangeRequestMetadataData {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return ChangeRequestMetadataData{
		StoreUUID:      randomvalues.GetRandomUUIDValue(),
		ConnectionUUID: randomvalues.GetRandomUUIDValue(),
		Schema:         randomvalues.GetRandomStringValuePtr(),
	}
}

func NewChangeRequestMetadataDataSliceWithRandomValues(n int) []ChangeRequestMetadataData {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []ChangeRequestMetadataData{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewChangeRequestMetadataDataWithRandomValues())
	}
	return res
}
