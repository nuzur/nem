package project_version_deployment

import (
	"encoding/json"

	"github.com/gofrs/uuid"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"

	"time"
)

type ProjectVersionDeployment struct {
	EnviromentUUID uuid.UUID `json:"enviroment_uuid"`
	Status         Status    `json:"status"`
}

func (e ProjectVersionDeployment) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e ProjectVersionDeployment) EntityIdentifier() string {
	return "project_version_deployment"
}

func (e ProjectVersionDeployment) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["enviroment_uuid"] = types.UUIDFieldType
	res["status"] = types.SingleEnumFieldType
	return res
}

func (e ProjectVersionDeployment) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e ProjectVersionDeployment) IsDependant() bool {
	return true
}

func ProjectVersionDeploymentFromJSON(data json.RawMessage) ProjectVersionDeployment {
	entity := ProjectVersionDeployment{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func ProjectVersionDeploymentSliceFromJSON(data json.RawMessage) []ProjectVersionDeployment {
	entity := []ProjectVersionDeployment{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e ProjectVersionDeployment) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ProjectVersionDeploymentToJSON", err)
	}
	return res
}

func ProjectVersionDeploymentToJSON(e ProjectVersionDeployment) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ProjectVersionDeploymentToJSON", err)
	}
	return res
}

func ProjectVersionDeploymentSliceToJSON(e []ProjectVersionDeployment) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ProjectVersionDeploymentSliceToJSON", err)
	}
	return res
}

func NewProjectVersionDeploymentWithRandomValues() ProjectVersionDeployment {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return ProjectVersionDeployment{
		EnviromentUUID: randomvalues.GetRandomUUIDValue(),
		Status:         randomvalues.GetRandomOptionValue[Status](5),
	}
}

func NewProjectVersionDeploymentSliceWithRandomValues(n int) []ProjectVersionDeployment {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []ProjectVersionDeployment{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewProjectVersionDeploymentWithRandomValues())
	}
	return res
}
