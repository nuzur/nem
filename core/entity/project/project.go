package project

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"github.com/nuzur/nem/core/entity/project_extension"
	"time"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"
)

type Project struct {
	UUID              uuid.UUID                            `json:"uuid"`
	Version           int64                                `json:"version"`
	Name              string                               `json:"name"`
	Description       string                               `json:"description"`
	Tags              []string                             `json:"tags"`
	URL               string                               `json:"url"`
	OwnerUUID         uuid.UUID                            `json:"owner_uuid"`
	TeamUUID          uuid.UUID                            `json:"team_uuid"`
	ProjectExtensions []project_extension.ProjectExtension `json:"project_extensions"`
	Status            Status                               `json:"status"`
	CreatedAt         time.Time                            `json:"created_at"`
	UpdatedAt         time.Time                            `json:"updated_at"`
	CreatedByUUID     uuid.UUID                            `json:"created_by_uuid"`
	UpdatedByUUID     uuid.UUID                            `json:"updated_by_uuid"`
}

func (e Project) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e Project) EntityIdentifier() string {
	return "project"
}

func (e Project) PrimaryKeyIdentifier() string {
	return "uuid"
}

func (e Project) PrimaryKeyValue() string {
	return e.UUID.String()
}

func (e Project) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["uuid"] = types.UUIDFieldType
	res["version"] = types.IntFieldType
	res["name"] = types.StringFieldType
	res["description"] = types.StringFieldType
	res["tags"] = types.ArrayFieldType
	res["url"] = types.StringFieldType
	res["owner_uuid"] = types.UUIDFieldType
	res["team_uuid"] = types.UUIDFieldType
	res["project_extensions"] = types.MultiDependantEntityFieldType
	res["status"] = types.SingleEnumFieldType
	res["created_at"] = types.TimestampFieldType
	res["updated_at"] = types.TimestampFieldType
	res["created_by_uuid"] = types.UUIDFieldType
	res["updated_by_uuid"] = types.UUIDFieldType
	return res
}

func (e Project) DependantFieldIdentifierToTypeMap() map[string]map[string]types.FieldType {
	res := make(map[string]map[string]types.FieldType)

	res["project_extensions"] = project_extension.ProjectExtension{}.FieldIdentfierToTypeMap()
	return res
}

func (e Project) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["tags"] = types.StringFieldType
	return res
}

func (e Project) IsDependant() bool {
	return false
}

func ProjectFromJSON(data json.RawMessage) Project {
	entity := Project{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func ProjectSliceFromJSON(data json.RawMessage) []Project {
	entity := []Project{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e Project) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ProjectToJSON", err)
	}
	return res
}

func ProjectToJSON(e Project) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ProjectToJSON", err)
	}
	return res
}

func ProjectSliceToJSON(e []Project) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ProjectSliceToJSON", err)
	}
	return res
}

func NewProjectWithRandomValues() Project {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return Project{
		UUID:              randomvalues.GetRandomUUIDValue(),
		Version:           randomvalues.GetRandomIntValue(),
		Name:              randomvalues.GetRandomStringValue(),
		Description:       randomvalues.GetRandomStringValue(),
		Tags:              []string{},
		URL:               randomvalues.GetRandomStringValue(),
		OwnerUUID:         randomvalues.GetRandomUUIDValue(),
		TeamUUID:          randomvalues.GetRandomUUIDValue(),
		ProjectExtensions: project_extension.NewProjectExtensionSliceWithRandomValues(rand.Intn(10)),
		Status:            randomvalues.GetRandomOptionValue[Status](2),
		CreatedAt:         randomvalues.GetRandomTimeValue(),
		UpdatedAt:         randomvalues.GetRandomTimeValue(),
		CreatedByUUID:     randomvalues.GetRandomUUIDValue(),
		UpdatedByUUID:     randomvalues.GetRandomUUIDValue(),
	}
}

func NewProjectSliceWithRandomValues(n int) []Project {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []Project{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewProjectWithRandomValues())
	}
	return res
}
