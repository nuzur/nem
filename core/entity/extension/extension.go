package extension

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"github.com/nuzur/nem/core/entity/visibility"
	"time"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"
)

type Extension struct {
	UUID              uuid.UUID               `json:"uuid"`
	Version           int64                   `json:"version"`
	Identifier        string                  `json:"identifier"`
	DisplayName       string                  `json:"display_name"`
	DisplayAuthorName string                  `json:"display_author_name"`
	Description       string                  `json:"description"`
	URL               string                  `json:"url"`
	Verfied           bool                    `json:"verfied"`
	Repository        string                  `json:"repository"`
	ExtensionType     ExtensionType           `json:"extension_type"`
	Tags              []string                `json:"tags"`
	Public            bool                    `json:"public"`
	Visibility        []visibility.Visibility `json:"visibility"`
	Status            Status                  `json:"status"`
	OwnerUUID         uuid.UUID               `json:"owner_uuid"`
	CreatedAt         time.Time               `json:"created_at"`
	UpdatedAt         time.Time               `json:"updated_at"`
	CreatedByUUID     uuid.UUID               `json:"created_by_uuid"`
	UpdatedByUUID     uuid.UUID               `json:"updated_by_uuid"`
}

func (e Extension) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e Extension) EntityIdentifier() string {
	return "extension"
}

func (e Extension) PrimaryKeyIdentifier() string {
	return "uuid"
}

func (e Extension) PrimaryKeyValue() string {
	return e.UUID.String()
}

func (e Extension) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["uuid"] = types.UUIDFieldType
	res["version"] = types.IntFieldType
	res["identifier"] = types.StringFieldType
	res["display_name"] = types.StringFieldType
	res["display_author_name"] = types.StringFieldType
	res["description"] = types.StringFieldType
	res["url"] = types.StringFieldType
	res["verfied"] = types.BooleanFieldType
	res["repository"] = types.StringFieldType
	res["extension_type"] = types.SingleEnumFieldType
	res["tags"] = types.ArrayFieldType
	res["public"] = types.BooleanFieldType
	res["visibility"] = types.MultiDependantEntityFieldType
	res["status"] = types.SingleEnumFieldType
	res["owner_uuid"] = types.UUIDFieldType
	res["created_at"] = types.TimestampFieldType
	res["updated_at"] = types.TimestampFieldType
	res["created_by_uuid"] = types.UUIDFieldType
	res["updated_by_uuid"] = types.UUIDFieldType
	return res
}

func (e Extension) DependantFieldIdentifierToTypeMap() map[string]map[string]types.FieldType {
	res := make(map[string]map[string]types.FieldType)

	res["visibility"] = visibility.Visibility{}.FieldIdentfierToTypeMap()
	return res
}

func (e Extension) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["tags"] = types.StringFieldType
	return res
}

func (e Extension) IsDependant() bool {
	return false
}

func ExtensionFromJSON(data json.RawMessage) Extension {
	entity := Extension{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func ExtensionSliceFromJSON(data json.RawMessage) []Extension {
	entity := []Extension{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e Extension) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ExtensionToJSON", err)
	}
	return res
}

func ExtensionToJSON(e Extension) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ExtensionToJSON", err)
	}
	return res
}

func ExtensionSliceToJSON(e []Extension) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ExtensionSliceToJSON", err)
	}
	return res
}

func NewExtensionWithRandomValues() Extension {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return Extension{
		UUID:              randomvalues.GetRandomUUIDValue(),
		Version:           randomvalues.GetRandomIntValue(),
		Identifier:        randomvalues.GetRandomStringValue(),
		DisplayName:       randomvalues.GetRandomStringValue(),
		DisplayAuthorName: randomvalues.GetRandomStringValue(),
		Description:       randomvalues.GetRandomStringValue(),
		URL:               randomvalues.GetRandomStringValue(),
		Verfied:           randomvalues.GetRandomBoolValue(),
		Repository:        randomvalues.GetRandomStringValue(),
		ExtensionType:     randomvalues.GetRandomOptionValue[ExtensionType](3),
		Tags:              []string{},
		Public:            randomvalues.GetRandomBoolValue(),
		Visibility:        visibility.NewVisibilitySliceWithRandomValues(rand.Intn(10)),
		Status:            randomvalues.GetRandomOptionValue[Status](2),
		OwnerUUID:         randomvalues.GetRandomUUIDValue(),
		CreatedAt:         randomvalues.GetRandomTimeValue(),
		UpdatedAt:         randomvalues.GetRandomTimeValue(),
		CreatedByUUID:     randomvalues.GetRandomUUIDValue(),
		UpdatedByUUID:     randomvalues.GetRandomUUIDValue(),
	}
}

func NewExtensionSliceWithRandomValues(n int) []Extension {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []Extension{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewExtensionWithRandomValues())
	}
	return res
}
