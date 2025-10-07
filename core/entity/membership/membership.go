package membership

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"time"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"
)

type Membership struct {
	UUID            uuid.UUID `json:"uuid"`
	OwnerUUID       uuid.UUID `json:"owner_uuid"`
	Type            Type      `json:"type"`
	StartDate       time.Time `json:"start_date"`
	BillingMetadata string    `json:"billing_metadata"`
	Status          Status    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	CreatedByUUID   uuid.UUID `json:"created_by_uuid"`
	UpdatedByUUID   uuid.UUID `json:"updated_by_uuid"`
}

func (e Membership) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e Membership) EntityIdentifier() string {
	return "membership"
}

func (e Membership) PrimaryKeyIdentifier() string {
	return "uuid"
}

func (e Membership) PrimaryKeyValue() string {
	return e.UUID.String()
}

func (e Membership) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["uuid"] = types.UUIDFieldType
	res["owner_uuid"] = types.UUIDFieldType
	res["type"] = types.SingleEnumFieldType
	res["start_date"] = types.TimestampFieldType
	res["billing_metadata"] = types.RawJSONFieldType
	res["status"] = types.SingleEnumFieldType
	res["created_at"] = types.TimestampFieldType
	res["updated_at"] = types.TimestampFieldType
	res["created_by_uuid"] = types.UUIDFieldType
	res["updated_by_uuid"] = types.UUIDFieldType
	return res
}

func (e Membership) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "uuid")
	res = append(res, "owner_uuid")
	res = append(res, "type")
	res = append(res, "start_date")
	res = append(res, "billing_metadata")
	res = append(res, "status")
	res = append(res, "created_at")
	res = append(res, "updated_at")
	res = append(res, "created_by_uuid")
	res = append(res, "updated_by_uuid")

	return res
}

func (e Membership) DependantFieldIdentifierToTypeMap() map[string]map[string]types.FieldType {
	res := make(map[string]map[string]types.FieldType)

	return res
}

func (e Membership) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e Membership) IsDependant() bool {
	return false
}

func MembershipFromJSON(data json.RawMessage) Membership {
	entity := Membership{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func MembershipSliceFromJSON(data json.RawMessage) []Membership {
	entity := []Membership{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e Membership) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error MembershipToJSON", err)
	}
	return res
}

func MembershipToJSON(e Membership) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error MembershipToJSON", err)
	}
	return res
}

func MembershipSliceToJSON(e []Membership) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error MembershipSliceToJSON", err)
	}
	return res
}

func NewMembershipWithRandomValues() Membership {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return Membership{
		UUID:            randomvalues.GetRandomUUIDValue(),
		OwnerUUID:       randomvalues.GetRandomUUIDValue(),
		Type:            randomvalues.GetRandomOptionValue[Type](1),
		StartDate:       randomvalues.GetRandomTimeValue(),
		BillingMetadata: randomvalues.GetRandomRawJSONValue(),
		Status:          randomvalues.GetRandomOptionValue[Status](6),
		CreatedAt:       randomvalues.GetRandomTimeValue(),
		UpdatedAt:       randomvalues.GetRandomTimeValue(),
		CreatedByUUID:   randomvalues.GetRandomUUIDValue(),
		UpdatedByUUID:   randomvalues.GetRandomUUIDValue(),
	}
}

func NewMembershipSliceWithRandomValues(n int) []Membership {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []Membership{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewMembershipWithRandomValues())
	}
	return res
}
