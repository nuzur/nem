package extension_version

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"time"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"
)

type ExtensionVersion struct {
	UUID                uuid.UUID       `json:"uuid"`
	Version             int64           `json:"version"`
	ExtensionUUID       uuid.UUID       `json:"extension_uuid"`
	DisplayVersion      *string         `json:"display_version"`
	Description         *string         `json:"description"`
	RepositoryTag       string          `json:"repository_tag"`
	ConfigurationEntity *string         `json:"configuration_entity"`
	ExecutionMode       []ExecutionMode `json:"execution_mode"`
	ReviewStatus        ReviewStatus    `json:"review_status"`
	Status              Status          `json:"status"`
	CreatedAt           time.Time       `json:"created_at"`
	UpdatedAt           time.Time       `json:"updated_at"`
	CreatedByUUID       uuid.UUID       `json:"created_by_uuid"`
	UpdatedByUUID       uuid.UUID       `json:"updated_by_uuid"`
}

func (e ExtensionVersion) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e ExtensionVersion) EntityIdentifier() string {
	return "extension_version"
}

func (e ExtensionVersion) PrimaryKeyIdentifier() string {
	return "uuid"
}

func (e ExtensionVersion) PrimaryKeyValue() string {
	return e.UUID.String()
}

func (e ExtensionVersion) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["uuid"] = types.UUIDFieldType
	res["version"] = types.IntFieldType
	res["extension_uuid"] = types.UUIDFieldType
	res["display_version"] = types.StringFieldType
	res["description"] = types.StringFieldType
	res["repository_tag"] = types.StringFieldType
	res["configuration_entity"] = types.RawJSONFieldType
	res["execution_mode"] = types.MultiEnumFieldType
	res["review_status"] = types.SingleEnumFieldType
	res["status"] = types.SingleEnumFieldType
	res["created_at"] = types.TimestampFieldType
	res["updated_at"] = types.TimestampFieldType
	res["created_by_uuid"] = types.UUIDFieldType
	res["updated_by_uuid"] = types.UUIDFieldType
	return res
}

func (e ExtensionVersion) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "uuid")
	res = append(res, "version")
	res = append(res, "extension_uuid")
	res = append(res, "display_version")
	res = append(res, "description")
	res = append(res, "repository_tag")
	res = append(res, "configuration_entity")
	res = append(res, "execution_mode")
	res = append(res, "review_status")
	res = append(res, "status")
	res = append(res, "created_at")
	res = append(res, "updated_at")
	res = append(res, "created_by_uuid")
	res = append(res, "updated_by_uuid")

	return res
}

func (e ExtensionVersion) DependantFieldIdentifierToTypeMap() map[string]map[string]types.FieldType {
	res := make(map[string]map[string]types.FieldType)

	return res
}

func (e ExtensionVersion) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e ExtensionVersion) IsDependant() bool {
	return false
}

func ExtensionVersionFromJSON(data json.RawMessage) ExtensionVersion {
	entity := ExtensionVersion{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func ExtensionVersionSliceFromJSON(data json.RawMessage) []ExtensionVersion {
	entity := []ExtensionVersion{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e ExtensionVersion) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ExtensionVersionToJSON", err)
	}
	return res
}

func ExtensionVersionToJSON(e ExtensionVersion) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ExtensionVersionToJSON", err)
	}
	return res
}

func ExtensionVersionSliceToJSON(e []ExtensionVersion) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ExtensionVersionSliceToJSON", err)
	}
	return res
}

func NewExtensionVersionWithRandomValues() ExtensionVersion {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return ExtensionVersion{
		UUID:                randomvalues.GetRandomUUIDValue(),
		Version:             randomvalues.GetRandomIntValue(),
		ExtensionUUID:       randomvalues.GetRandomUUIDValue(),
		DisplayVersion:      randomvalues.GetRandomStringValuePtr(),
		Description:         randomvalues.GetRandomStringValuePtr(),
		RepositoryTag:       randomvalues.GetRandomStringValue(),
		ConfigurationEntity: randomvalues.GetRandomRawJSONValuePtr(),
		ExecutionMode:       randomvalues.GetRandomOptionsValues[ExecutionMode](2),
		ReviewStatus:        randomvalues.GetRandomOptionValue[ReviewStatus](5),
		Status:              randomvalues.GetRandomOptionValue[Status](2),
		CreatedAt:           randomvalues.GetRandomTimeValue(),
		UpdatedAt:           randomvalues.GetRandomTimeValue(),
		CreatedByUUID:       randomvalues.GetRandomUUIDValue(),
		UpdatedByUUID:       randomvalues.GetRandomUUIDValue(),
	}
}

func NewExtensionVersionSliceWithRandomValues(n int) []ExtensionVersion {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []ExtensionVersion{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewExtensionVersionWithRandomValues())
	}
	return res
}
