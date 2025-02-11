package visibility

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"time"

	"fmt"

	"nem/core/entity/types"

	"nem/core/randomvalues"

	"math/rand"
)

type Visibility struct {
	UUID              uuid.UUID   `json:"uuid"`
	Description       string      `json:"description"`
	OrganizationUUIDs []uuid.UUID `json:"organization_uuids"`
	TeamUUIDs         []uuid.UUID `json:"team_uuids"`
	UserUUIDs         []uuid.UUID `json:"user_uuids"`
	Roles             []Role      `json:"roles"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
	CreatedByUUID     uuid.UUID   `json:"created_by_uuid"`
	UpdatedByUUID     uuid.UUID   `json:"updated_by_uuid"`
}

func (e Visibility) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e Visibility) EntityIdentifier() string {
	return "visibility"
}

func (e Visibility) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["uuid"] = types.UUIDFieldType
	res["description"] = types.StringFieldType
	res["organization_uuids"] = types.ArrayFieldType
	res["team_uuids"] = types.ArrayFieldType
	res["user_uuids"] = types.ArrayFieldType
	res["roles"] = types.MultiEnumFieldType
	res["created_at"] = types.TimestampFieldType
	res["updated_at"] = types.TimestampFieldType
	res["created_by_uuid"] = types.UUIDFieldType
	res["updated_by_uuid"] = types.UUIDFieldType
	return res
}

func (e Visibility) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["organization_uuids"] = types.UUIDFieldType
	res["team_uuids"] = types.UUIDFieldType
	res["user_uuids"] = types.UUIDFieldType
	return res
}

func (e Visibility) IsDependant() bool {
	return true
}

func VisibilityFromJSON(data json.RawMessage) Visibility {
	entity := Visibility{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func VisibilitySliceFromJSON(data json.RawMessage) []Visibility {
	entity := []Visibility{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e Visibility) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error VisibilityToJSON", err)
	}
	return res
}

func VisibilityToJSON(e Visibility) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error VisibilityToJSON", err)
	}
	return res
}

func VisibilitySliceToJSON(e []Visibility) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error VisibilitySliceToJSON", err)
	}
	return res
}

func NewVisibilityWithRandomValues() Visibility {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return Visibility{
		UUID:              randomvalues.GetRandomUUIDValue(),
		Description:       randomvalues.GetRandomStringValue(),
		OrganizationUUIDs: []uuid.UUID{},
		TeamUUIDs:         []uuid.UUID{},
		UserUUIDs:         []uuid.UUID{},
		Roles:             randomvalues.GetRandomOptionsValues[Role](5),
		CreatedAt:         randomvalues.GetRandomTimeValue(),
		UpdatedAt:         randomvalues.GetRandomTimeValue(),
		CreatedByUUID:     randomvalues.GetRandomUUIDValue(),
		UpdatedByUUID:     randomvalues.GetRandomUUIDValue(),
	}
}

func NewVisibilitySliceWithRandomValues(n int) []Visibility {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []Visibility{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewVisibilityWithRandomValues())
	}
	return res
}
