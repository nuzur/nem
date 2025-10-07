package change_request_metadata

import (
	"encoding/json"

	"github.com/nuzur/nem/core/entity/change_request_metadata_data"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"math/rand"

	"time"
)

type ChangeRequestMetadata struct {
	Data change_request_metadata_data.ChangeRequestMetadataData `json:"data"`
}

func (e ChangeRequestMetadata) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e ChangeRequestMetadata) EntityIdentifier() string {
	return "change_request_metadata"
}

func (e ChangeRequestMetadata) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["data"] = types.SingleDependantEntityFieldType
	return res
}

func (e ChangeRequestMetadata) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "data")

	return res
}

func (e ChangeRequestMetadata) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e ChangeRequestMetadata) IsDependant() bool {
	return true
}

func ChangeRequestMetadataFromJSON(data json.RawMessage) ChangeRequestMetadata {
	entity := ChangeRequestMetadata{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func ChangeRequestMetadataSliceFromJSON(data json.RawMessage) []ChangeRequestMetadata {
	entity := []ChangeRequestMetadata{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e ChangeRequestMetadata) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ChangeRequestMetadataToJSON", err)
	}
	return res
}

func ChangeRequestMetadataToJSON(e ChangeRequestMetadata) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ChangeRequestMetadataToJSON", err)
	}
	return res
}

func ChangeRequestMetadataSliceToJSON(e []ChangeRequestMetadata) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ChangeRequestMetadataSliceToJSON", err)
	}
	return res
}

func NewChangeRequestMetadataWithRandomValues() ChangeRequestMetadata {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return ChangeRequestMetadata{
		Data: change_request_metadata_data.NewChangeRequestMetadataDataWithRandomValues(),
	}
}

func NewChangeRequestMetadataSliceWithRandomValues(n int) []ChangeRequestMetadata {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []ChangeRequestMetadata{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewChangeRequestMetadataWithRandomValues())
	}
	return res
}
