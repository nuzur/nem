package user_project

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"time"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"
)

type UserProject struct {
	UUID                    uuid.UUID `json:"uuid"`
	UserUUID                uuid.UUID `json:"user_uuid"`
	UserEmail               *string   `json:"user_email"`
	ProjectUUID             uuid.UUID `json:"project_uuid"`
	Role                    Role      `json:"role"`
	ReviewRequiredStructure bool      `json:"review_required_structure"`
	ReviewRequiredData      bool      `json:"review_required_data"`
	Status                  Status    `json:"status"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
	CreatedByUUID           uuid.UUID `json:"created_by_uuid"`
	UpdatedByUUID           uuid.UUID `json:"updated_by_uuid"`
}

func (e UserProject) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e UserProject) EntityIdentifier() string {
	return "user_project"
}

func (e UserProject) PrimaryKeyIdentifier() string {
	return "uuid"
}

func (e UserProject) PrimaryKeyValue() string {
	return e.UUID.String()
}

func (e UserProject) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["uuid"] = types.UUIDFieldType
	res["user_uuid"] = types.UUIDFieldType
	res["user_email"] = types.StringFieldType
	res["project_uuid"] = types.UUIDFieldType
	res["role"] = types.SingleEnumFieldType
	res["review_required_structure"] = types.BooleanFieldType
	res["review_required_data"] = types.BooleanFieldType
	res["status"] = types.SingleEnumFieldType
	res["created_at"] = types.TimestampFieldType
	res["updated_at"] = types.TimestampFieldType
	res["created_by_uuid"] = types.UUIDFieldType
	res["updated_by_uuid"] = types.UUIDFieldType
	return res
}

func (e UserProject) DependantFieldIdentifierToTypeMap() map[string]map[string]types.FieldType {
	res := make(map[string]map[string]types.FieldType)

	return res
}

func (e UserProject) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e UserProject) IsDependant() bool {
	return false
}

func UserProjectFromJSON(data json.RawMessage) UserProject {
	entity := UserProject{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func UserProjectSliceFromJSON(data json.RawMessage) []UserProject {
	entity := []UserProject{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e UserProject) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error UserProjectToJSON", err)
	}
	return res
}

func UserProjectToJSON(e UserProject) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error UserProjectToJSON", err)
	}
	return res
}

func UserProjectSliceToJSON(e []UserProject) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error UserProjectSliceToJSON", err)
	}
	return res
}

func NewUserProjectWithRandomValues() UserProject {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return UserProject{
		UUID:                    randomvalues.GetRandomUUIDValue(),
		UserUUID:                randomvalues.GetRandomUUIDValue(),
		UserEmail:               randomvalues.GetRandomStringValuePtr(),
		ProjectUUID:             randomvalues.GetRandomUUIDValue(),
		Role:                    randomvalues.GetRandomOptionValue[Role](5),
		ReviewRequiredStructure: randomvalues.GetRandomBoolValue(),
		ReviewRequiredData:      randomvalues.GetRandomBoolValue(),
		Status:                  randomvalues.GetRandomOptionValue[Status](2),
		CreatedAt:               randomvalues.GetRandomTimeValue(),
		UpdatedAt:               randomvalues.GetRandomTimeValue(),
		CreatedByUUID:           randomvalues.GetRandomUUIDValue(),
		UpdatedByUUID:           randomvalues.GetRandomUUIDValue(),
	}
}

func NewUserProjectSliceWithRandomValues(n int) []UserProject {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []UserProject{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewUserProjectWithRandomValues())
	}
	return res
}
