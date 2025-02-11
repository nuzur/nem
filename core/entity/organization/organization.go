package organization

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"github.com/nuzur/nem/core/entity/membership"
	"time"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"
)

type Organization struct {
	UUID          uuid.UUID               `json:"uuid"`
	Version       int64                   `json:"version"`
	Name          string                  `json:"name"`
	Domains       []string                `json:"domains"`
	AdminUUIDs    []uuid.UUID             `json:"admin_uuids"`
	Memberships   []membership.Membership `json:"memberships"`
	Status        Status                  `json:"status"`
	CreatedAt     time.Time               `json:"created_at"`
	UpdatedAt     time.Time               `json:"updated_at"`
	CreatedByUUID uuid.UUID               `json:"created_by_uuid"`
	UpdatedByUUID uuid.UUID               `json:"updated_by_uuid"`
}

func (e Organization) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e Organization) EntityIdentifier() string {
	return "organization"
}

func (e Organization) PrimaryKeyIdentifier() string {
	return "uuid"
}

func (e Organization) PrimaryKeyValue() string {
	return e.UUID.String()
}

func (e Organization) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["uuid"] = types.UUIDFieldType
	res["version"] = types.IntFieldType
	res["name"] = types.StringFieldType
	res["domains"] = types.ArrayFieldType
	res["admin_uuids"] = types.ArrayFieldType
	res["memberships"] = types.MultiDependantEntityFieldType
	res["status"] = types.SingleEnumFieldType
	res["created_at"] = types.TimestampFieldType
	res["updated_at"] = types.TimestampFieldType
	res["created_by_uuid"] = types.UUIDFieldType
	res["updated_by_uuid"] = types.UUIDFieldType
	return res
}

func (e Organization) DependantFieldIdentifierToTypeMap() map[string]map[string]types.FieldType {
	res := make(map[string]map[string]types.FieldType)

	res["memberships"] = membership.Membership{}.FieldIdentfierToTypeMap()
	return res
}

func (e Organization) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["domains"] = types.StringFieldType
	res["admin_uuids"] = types.UUIDFieldType
	return res
}

func (e Organization) IsDependant() bool {
	return false
}

func OrganizationFromJSON(data json.RawMessage) Organization {
	entity := Organization{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func OrganizationSliceFromJSON(data json.RawMessage) []Organization {
	entity := []Organization{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e Organization) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error OrganizationToJSON", err)
	}
	return res
}

func OrganizationToJSON(e Organization) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error OrganizationToJSON", err)
	}
	return res
}

func OrganizationSliceToJSON(e []Organization) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error OrganizationSliceToJSON", err)
	}
	return res
}

func NewOrganizationWithRandomValues() Organization {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return Organization{
		UUID:          randomvalues.GetRandomUUIDValue(),
		Version:       randomvalues.GetRandomIntValue(),
		Name:          randomvalues.GetRandomStringValue(),
		Domains:       []string{},
		AdminUUIDs:    []uuid.UUID{},
		Memberships:   membership.NewMembershipSliceWithRandomValues(rand.Intn(10)),
		Status:        randomvalues.GetRandomOptionValue[Status](2),
		CreatedAt:     randomvalues.GetRandomTimeValue(),
		UpdatedAt:     randomvalues.GetRandomTimeValue(),
		CreatedByUUID: randomvalues.GetRandomUUIDValue(),
		UpdatedByUUID: randomvalues.GetRandomUUIDValue(),
	}
}

func NewOrganizationSliceWithRandomValues(n int) []Organization {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []Organization{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewOrganizationWithRandomValues())
	}
	return res
}
