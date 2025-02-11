package project_extension

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"time"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"
)

type ProjectExtension struct {
	UUID                      uuid.UUID `json:"uuid"`
	ExtensionUUID             uuid.UUID `json:"extension_uuid"`
	ExtensionVersionUUID      uuid.UUID `json:"extension_version_uuid"`
	ConfigurationEntityValues string    `json:"configuration_entity_values"`
	Blocking                  bool      `json:"blocking"`
	Status                    Status    `json:"status"`
	CreatedAt                 time.Time `json:"created_at"`
	UpdatedAt                 time.Time `json:"updated_at"`
	CreatedByUUID             uuid.UUID `json:"created_by_uuid"`
	UpdatedByUUID             uuid.UUID `json:"updated_by_uuid"`
}

func (e ProjectExtension) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e ProjectExtension) EntityIdentifier() string {
	return "project_extension"
}

func (e ProjectExtension) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["uuid"] = types.UUIDFieldType
	res["extension_uuid"] = types.UUIDFieldType
	res["extension_version_uuid"] = types.UUIDFieldType
	res["configuration_entity_values"] = types.StringFieldType
	res["blocking"] = types.BooleanFieldType
	res["status"] = types.SingleEnumFieldType
	res["created_at"] = types.TimestampFieldType
	res["updated_at"] = types.TimestampFieldType
	res["created_by_uuid"] = types.UUIDFieldType
	res["updated_by_uuid"] = types.UUIDFieldType
	return res
}

func (e ProjectExtension) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e ProjectExtension) IsDependant() bool {
	return true
}

func ProjectExtensionFromJSON(data json.RawMessage) ProjectExtension {
	entity := ProjectExtension{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func ProjectExtensionSliceFromJSON(data json.RawMessage) []ProjectExtension {
	entity := []ProjectExtension{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e ProjectExtension) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ProjectExtensionToJSON", err)
	}
	return res
}

func ProjectExtensionToJSON(e ProjectExtension) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ProjectExtensionToJSON", err)
	}
	return res
}

func ProjectExtensionSliceToJSON(e []ProjectExtension) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ProjectExtensionSliceToJSON", err)
	}
	return res
}

func NewProjectExtensionWithRandomValues() ProjectExtension {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return ProjectExtension{
		UUID:                      randomvalues.GetRandomUUIDValue(),
		ExtensionUUID:             randomvalues.GetRandomUUIDValue(),
		ExtensionVersionUUID:      randomvalues.GetRandomUUIDValue(),
		ConfigurationEntityValues: randomvalues.GetRandomStringValue(),
		Blocking:                  randomvalues.GetRandomBoolValue(),
		Status:                    randomvalues.GetRandomOptionValue[Status](2),
		CreatedAt:                 randomvalues.GetRandomTimeValue(),
		UpdatedAt:                 randomvalues.GetRandomTimeValue(),
		CreatedByUUID:             randomvalues.GetRandomUUIDValue(),
		UpdatedByUUID:             randomvalues.GetRandomUUIDValue(),
	}
}

func NewProjectExtensionSliceWithRandomValues(n int) []ProjectExtension {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []ProjectExtension{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewProjectExtensionWithRandomValues())
	}
	return res
}
