package user_team

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"time"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"
)

type UserTeam struct {
	UUID          uuid.UUID `json:"uuid"`
	UserUUID      uuid.UUID `json:"user_uuid"`
	UserEmail     *string   `json:"user_email"`
	TeamUUID      uuid.UUID `json:"team_uuid"`
	Roles         []Role    `json:"roles"`
	Status        Status    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	CreatedByUUID uuid.UUID `json:"created_by_uuid"`
	UpdatedByUUID uuid.UUID `json:"updated_by_uuid"`
}

func (e UserTeam) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e UserTeam) EntityIdentifier() string {
	return "user_team"
}

func (e UserTeam) PrimaryKeyIdentifier() string {
	return "uuid"
}

func (e UserTeam) PrimaryKeyValue() string {
	return e.UUID.String()
}

func (e UserTeam) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["uuid"] = types.UUIDFieldType
	res["user_uuid"] = types.UUIDFieldType
	res["user_email"] = types.StringFieldType
	res["team_uuid"] = types.UUIDFieldType
	res["roles"] = types.MultiEnumFieldType
	res["status"] = types.SingleEnumFieldType
	res["created_at"] = types.TimestampFieldType
	res["updated_at"] = types.TimestampFieldType
	res["created_by_uuid"] = types.UUIDFieldType
	res["updated_by_uuid"] = types.UUIDFieldType
	return res
}

func (e UserTeam) DependantFieldIdentifierToTypeMap() map[string]map[string]types.FieldType {
	res := make(map[string]map[string]types.FieldType)

	return res
}

func (e UserTeam) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e UserTeam) IsDependant() bool {
	return false
}

func UserTeamFromJSON(data json.RawMessage) UserTeam {
	entity := UserTeam{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func UserTeamSliceFromJSON(data json.RawMessage) []UserTeam {
	entity := []UserTeam{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e UserTeam) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error UserTeamToJSON", err)
	}
	return res
}

func UserTeamToJSON(e UserTeam) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error UserTeamToJSON", err)
	}
	return res
}

func UserTeamSliceToJSON(e []UserTeam) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error UserTeamSliceToJSON", err)
	}
	return res
}

func NewUserTeamWithRandomValues() UserTeam {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return UserTeam{
		UUID:          randomvalues.GetRandomUUIDValue(),
		UserUUID:      randomvalues.GetRandomUUIDValue(),
		UserEmail:     randomvalues.GetRandomStringValuePtr(),
		TeamUUID:      randomvalues.GetRandomUUIDValue(),
		Roles:         randomvalues.GetRandomOptionsValues[Role](5),
		Status:        randomvalues.GetRandomOptionValue[Status](3),
		CreatedAt:     randomvalues.GetRandomTimeValue(),
		UpdatedAt:     randomvalues.GetRandomTimeValue(),
		CreatedByUUID: randomvalues.GetRandomUUIDValue(),
		UpdatedByUUID: randomvalues.GetRandomUUIDValue(),
	}
}

func NewUserTeamSliceWithRandomValues(n int) []UserTeam {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []UserTeam{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewUserTeamWithRandomValues())
	}
	return res
}
