package extension_execution

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"time"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"
)

type ExtensionExecution struct {
	UUID                 uuid.UUID `json:"uuid"`
	ExtensionUUID        uuid.UUID `json:"extension_uuid"`
	ExtensionVersionUUID uuid.UUID `json:"extension_version_uuid"`
	ProjectExtensionUUID uuid.UUID `json:"project_extension_uuid"`
	ProjectUUID          uuid.UUID `json:"project_uuid"`
	ProjectVersionUUID   uuid.UUID `json:"project_version_uuid"`
	ExecutedByUUID       uuid.UUID `json:"executed_by_uuid"`
	Metadata             *string   `json:"metadata"`
	Status               Status    `json:"status"`
	StatusMsg            *string   `json:"status_msg"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

func (e ExtensionExecution) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e ExtensionExecution) EntityIdentifier() string {
	return "extension_execution"
}

func (e ExtensionExecution) PrimaryKeyIdentifier() string {
	return "uuid"
}

func (e ExtensionExecution) PrimaryKeyValue() string {
	return e.UUID.String()
}

func (e ExtensionExecution) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["uuid"] = types.UUIDFieldType
	res["extension_uuid"] = types.UUIDFieldType
	res["extension_version_uuid"] = types.UUIDFieldType
	res["project_extension_uuid"] = types.UUIDFieldType
	res["project_uuid"] = types.UUIDFieldType
	res["project_version_uuid"] = types.UUIDFieldType
	res["executed_by_uuid"] = types.UUIDFieldType
	res["metadata"] = types.RawJSONFieldType
	res["status"] = types.SingleEnumFieldType
	res["status_msg"] = types.StringFieldType
	res["created_at"] = types.TimestampFieldType
	res["updated_at"] = types.TimestampFieldType
	return res
}

func (e ExtensionExecution) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "uuid")
	res = append(res, "extension_uuid")
	res = append(res, "extension_version_uuid")
	res = append(res, "project_extension_uuid")
	res = append(res, "project_uuid")
	res = append(res, "project_version_uuid")
	res = append(res, "executed_by_uuid")
	res = append(res, "metadata")
	res = append(res, "status")
	res = append(res, "status_msg")
	res = append(res, "created_at")
	res = append(res, "updated_at")

	return res
}

func (e ExtensionExecution) DependantFieldIdentifierToTypeMap() map[string]map[string]types.FieldType {
	res := make(map[string]map[string]types.FieldType)

	return res
}

func (e ExtensionExecution) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e ExtensionExecution) IsDependant() bool {
	return false
}

func ExtensionExecutionFromJSON(data json.RawMessage) ExtensionExecution {
	entity := ExtensionExecution{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func ExtensionExecutionSliceFromJSON(data json.RawMessage) []ExtensionExecution {
	entity := []ExtensionExecution{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e ExtensionExecution) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ExtensionExecutionToJSON", err)
	}
	return res
}

func ExtensionExecutionToJSON(e ExtensionExecution) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ExtensionExecutionToJSON", err)
	}
	return res
}

func ExtensionExecutionSliceToJSON(e []ExtensionExecution) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ExtensionExecutionSliceToJSON", err)
	}
	return res
}

func NewExtensionExecutionWithRandomValues() ExtensionExecution {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return ExtensionExecution{
		UUID:                 randomvalues.GetRandomUUIDValue(),
		ExtensionUUID:        randomvalues.GetRandomUUIDValue(),
		ExtensionVersionUUID: randomvalues.GetRandomUUIDValue(),
		ProjectExtensionUUID: randomvalues.GetRandomUUIDValue(),
		ProjectUUID:          randomvalues.GetRandomUUIDValue(),
		ProjectVersionUUID:   randomvalues.GetRandomUUIDValue(),
		ExecutedByUUID:       randomvalues.GetRandomUUIDValue(),
		Metadata:             randomvalues.GetRandomRawJSONValuePtr(),
		Status:               randomvalues.GetRandomOptionValue[Status](4),
		StatusMsg:            randomvalues.GetRandomStringValuePtr(),
		CreatedAt:            randomvalues.GetRandomTimeValue(),
		UpdatedAt:            randomvalues.GetRandomTimeValue(),
	}
}

func NewExtensionExecutionSliceWithRandomValues(n int) []ExtensionExecution {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []ExtensionExecution{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewExtensionExecutionWithRandomValues())
	}
	return res
}
