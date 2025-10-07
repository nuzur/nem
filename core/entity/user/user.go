package user

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"time"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"
)

type User struct {
	UUID        uuid.UUID `json:"uuid"`
	Identifier  string    `json:"identifier"`
	Name        *string   `json:"name"`
	LastName    *string   `json:"last_name"`
	Email       string    `json:"email"`
	UserType    UserType  `json:"user_type"`
	CountryIos2 *string   `json:"country_ios2"`
	Locale      *string   `json:"locale"`
	Metadata    *string   `json:"metadata"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (e User) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e User) EntityIdentifier() string {
	return "user"
}

func (e User) PrimaryKeyIdentifier() string {
	return "uuid"
}

func (e User) PrimaryKeyValue() string {
	return e.UUID.String()
}

func (e User) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["uuid"] = types.UUIDFieldType
	res["identifier"] = types.StringFieldType
	res["name"] = types.StringFieldType
	res["last_name"] = types.StringFieldType
	res["email"] = types.StringFieldType
	res["user_type"] = types.SingleEnumFieldType
	res["country_ios2"] = types.StringFieldType
	res["locale"] = types.StringFieldType
	res["metadata"] = types.StringFieldType
	res["status"] = types.SingleEnumFieldType
	res["created_at"] = types.TimestampFieldType
	res["updated_at"] = types.TimestampFieldType
	return res
}

func (e User) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "uuid")
	res = append(res, "identifier")
	res = append(res, "name")
	res = append(res, "last_name")
	res = append(res, "email")
	res = append(res, "user_type")
	res = append(res, "country_ios2")
	res = append(res, "locale")
	res = append(res, "metadata")
	res = append(res, "status")
	res = append(res, "created_at")
	res = append(res, "updated_at")

	return res
}

func (e User) DependantFieldIdentifierToTypeMap() map[string]map[string]types.FieldType {
	res := make(map[string]map[string]types.FieldType)

	return res
}

func (e User) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e User) IsDependant() bool {
	return false
}

func UserFromJSON(data json.RawMessage) User {
	entity := User{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func UserSliceFromJSON(data json.RawMessage) []User {
	entity := []User{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e User) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error UserToJSON", err)
	}
	return res
}

func UserToJSON(e User) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error UserToJSON", err)
	}
	return res
}

func UserSliceToJSON(e []User) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error UserSliceToJSON", err)
	}
	return res
}

func NewUserWithRandomValues() User {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return User{
		UUID:        randomvalues.GetRandomUUIDValue(),
		Identifier:  randomvalues.GetRandomStringValue(),
		Name:        randomvalues.GetRandomStringValuePtr(),
		LastName:    randomvalues.GetRandomStringValuePtr(),
		Email:       randomvalues.GetRandomStringValue(),
		UserType:    randomvalues.GetRandomOptionValue[UserType](2),
		CountryIos2: randomvalues.GetRandomStringValuePtr(),
		Locale:      randomvalues.GetRandomStringValuePtr(),
		Metadata:    randomvalues.GetRandomStringValuePtr(),
		Status:      randomvalues.GetRandomOptionValue[Status](2),
		CreatedAt:   randomvalues.GetRandomTimeValue(),
		UpdatedAt:   randomvalues.GetRandomTimeValue(),
	}
}

func NewUserSliceWithRandomValues(n int) []User {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []User{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewUserWithRandomValues())
	}
	return res
}
