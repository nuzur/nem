package project_version

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"github.com/nuzur/nem/core/entity/entity"
	"github.com/nuzur/nem/core/entity/enum"
	"github.com/nuzur/nem/core/entity/project_version_deployment"
	"github.com/nuzur/nem/core/entity/relationship"
	"github.com/nuzur/nem/core/entity/service"
	"time"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"
)

type ProjectVersion struct {
	UUID            uuid.UUID                                             `json:"uuid"`
	Version         int64                                                 `json:"version"`
	Identifier      string                                                `json:"identifier"`
	Description     string                                                `json:"description"`
	ProjectUUID     uuid.UUID                                             `json:"project_uuid"`
	Entities        []entity.Entity                                       `json:"entities"`
	Relationships   []relationship.Relationship                           `json:"relationships"`
	Enums           []enum.Enum                                           `json:"enums"`
	Services        []service.Service                                     `json:"services"`
	BaseVersionUUID uuid.UUID                                             `json:"base_version_uuid"`
	ReviewStatus    ReviewStatus                                          `json:"review_status"`
	Deployments     []project_version_deployment.ProjectVersionDeployment `json:"deployments"`
	Status          Status                                                `json:"status"`
	CreatedAt       time.Time                                             `json:"created_at"`
	UpdatedAt       time.Time                                             `json:"updated_at"`
	CreatedByUUID   uuid.UUID                                             `json:"created_by_uuid"`
	UpdatedByUUID   uuid.UUID                                             `json:"updated_by_uuid"`
}

func (e ProjectVersion) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e ProjectVersion) EntityIdentifier() string {
	return "project_version"
}

func (e ProjectVersion) PrimaryKeyIdentifier() string {
	return "uuid"
}

func (e ProjectVersion) PrimaryKeyValue() string {
	return e.UUID.String()
}

func (e ProjectVersion) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["uuid"] = types.UUIDFieldType
	res["version"] = types.IntFieldType
	res["identifier"] = types.StringFieldType
	res["description"] = types.StringFieldType
	res["project_uuid"] = types.UUIDFieldType
	res["entities"] = types.MultiDependantEntityFieldType
	res["relationships"] = types.MultiDependantEntityFieldType
	res["enums"] = types.MultiDependantEntityFieldType
	res["services"] = types.MultiDependantEntityFieldType
	res["base_version_uuid"] = types.UUIDFieldType
	res["review_status"] = types.SingleEnumFieldType
	res["deployments"] = types.MultiDependantEntityFieldType
	res["status"] = types.SingleEnumFieldType
	res["created_at"] = types.TimestampFieldType
	res["updated_at"] = types.TimestampFieldType
	res["created_by_uuid"] = types.UUIDFieldType
	res["updated_by_uuid"] = types.UUIDFieldType
	return res
}

func (e ProjectVersion) DependantFieldIdentifierToTypeMap() map[string]map[string]types.FieldType {
	res := make(map[string]map[string]types.FieldType)

	res["entities"] = entity.Entity{}.FieldIdentfierToTypeMap()
	res["relationships"] = relationship.Relationship{}.FieldIdentfierToTypeMap()
	res["enums"] = enum.Enum{}.FieldIdentfierToTypeMap()
	res["services"] = service.Service{}.FieldIdentfierToTypeMap()
	res["deployments"] = project_version_deployment.ProjectVersionDeployment{}.FieldIdentfierToTypeMap()
	return res
}

func (e ProjectVersion) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e ProjectVersion) IsDependant() bool {
	return false
}

func ProjectVersionFromJSON(data json.RawMessage) ProjectVersion {
	entity := ProjectVersion{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func ProjectVersionSliceFromJSON(data json.RawMessage) []ProjectVersion {
	entity := []ProjectVersion{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e ProjectVersion) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ProjectVersionToJSON", err)
	}
	return res
}

func ProjectVersionToJSON(e ProjectVersion) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ProjectVersionToJSON", err)
	}
	return res
}

func ProjectVersionSliceToJSON(e []ProjectVersion) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ProjectVersionSliceToJSON", err)
	}
	return res
}

func NewProjectVersionWithRandomValues() ProjectVersion {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return ProjectVersion{
		UUID:            randomvalues.GetRandomUUIDValue(),
		Version:         randomvalues.GetRandomIntValue(),
		Identifier:      randomvalues.GetRandomStringValue(),
		Description:     randomvalues.GetRandomStringValue(),
		ProjectUUID:     randomvalues.GetRandomUUIDValue(),
		Entities:        entity.NewEntitySliceWithRandomValues(rand.Intn(10)),
		Relationships:   relationship.NewRelationshipSliceWithRandomValues(rand.Intn(10)),
		Enums:           enum.NewEnumSliceWithRandomValues(rand.Intn(10)),
		Services:        service.NewServiceSliceWithRandomValues(rand.Intn(10)),
		BaseVersionUUID: randomvalues.GetRandomUUIDValue(),
		ReviewStatus:    randomvalues.GetRandomOptionValue[ReviewStatus](5),
		Deployments:     project_version_deployment.NewProjectVersionDeploymentSliceWithRandomValues(rand.Intn(10)),
		Status:          randomvalues.GetRandomOptionValue[Status](2),
		CreatedAt:       randomvalues.GetRandomTimeValue(),
		UpdatedAt:       randomvalues.GetRandomTimeValue(),
		CreatedByUUID:   randomvalues.GetRandomUUIDValue(),
		UpdatedByUUID:   randomvalues.GetRandomUUIDValue(),
	}
}

func NewProjectVersionSliceWithRandomValues(n int) []ProjectVersion {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []ProjectVersion{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewProjectVersionWithRandomValues())
	}
	return res
}
