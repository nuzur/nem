package user_project_version

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"time"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"
)

type UserProjectVersion struct {
	UUID               uuid.UUID `json:"uuid"`
	Version            int64     `json:"version"`
	ProjectVersionUUID uuid.UUID `json:"project_version_uuid"`
	UserUUID           uuid.UUID `json:"user_uuid"`
	Data               *string   `json:"data"`
	Status             Status    `json:"status"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	CreatedByUUID      uuid.UUID `json:"created_by_uuid"`
	UpdatedByUUID      uuid.UUID `json:"updated_by_uuid"`
}

func (e UserProjectVersion) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e UserProjectVersion) EntityIdentifier() string {
	return "user_project_version"
}

func (e UserProjectVersion) PrimaryKeyIdentifier() string {
	return "uuid"
}

func (e UserProjectVersion) PrimaryKeyValue() string {
	return e.UUID.String()
}

func (e UserProjectVersion) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["uuid"] = types.UUIDFieldType
	res["version"] = types.IntFieldType
	res["project_version_uuid"] = types.UUIDFieldType
	res["user_uuid"] = types.UUIDFieldType
	res["data"] = types.RawJSONFieldType
	res["status"] = types.SingleEnumFieldType
	res["created_at"] = types.TimestampFieldType
	res["updated_at"] = types.TimestampFieldType
	res["created_by_uuid"] = types.UUIDFieldType
	res["updated_by_uuid"] = types.UUIDFieldType
	return res
}

func (e UserProjectVersion) DependantFieldIdentifierToTypeMap() map[string]map[string]types.FieldType {
	res := make(map[string]map[string]types.FieldType)

	return res
}

func (e UserProjectVersion) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e UserProjectVersion) IsDependant() bool {
	return false
}

func UserProjectVersionFromJSON(data json.RawMessage) UserProjectVersion {
	entity := UserProjectVersion{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func UserProjectVersionSliceFromJSON(data json.RawMessage) []UserProjectVersion {
	entity := []UserProjectVersion{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e UserProjectVersion) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error UserProjectVersionToJSON", err)
	}
	return res
}

func UserProjectVersionToJSON(e UserProjectVersion) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error UserProjectVersionToJSON", err)
	}
	return res
}

func UserProjectVersionSliceToJSON(e []UserProjectVersion) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error UserProjectVersionSliceToJSON", err)
	}
	return res
}

func NewUserProjectVersionWithRandomValues() UserProjectVersion {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return UserProjectVersion{
		UUID:               randomvalues.GetRandomUUIDValue(),
		Version:            randomvalues.GetRandomIntValue(),
		ProjectVersionUUID: randomvalues.GetRandomUUIDValue(),
		UserUUID:           randomvalues.GetRandomUUIDValue(),
		Data:               randomvalues.GetRandomRawJSONValuePtr(),
		Status:             randomvalues.GetRandomOptionValue[Status](2),
		CreatedAt:          randomvalues.GetRandomTimeValue(),
		UpdatedAt:          randomvalues.GetRandomTimeValue(),
		CreatedByUUID:      randomvalues.GetRandomUUIDValue(),
		UpdatedByUUID:      randomvalues.GetRandomUUIDValue(),
	}
}

func NewUserProjectVersionSliceWithRandomValues(n int) []UserProjectVersion {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []UserProjectVersion{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewUserProjectVersionWithRandomValues())
	}
	return res
}
