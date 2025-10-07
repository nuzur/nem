package user_connection_local_config

import (
	"encoding/json"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"

	"time"
)

type UserConnectionLocalConfig struct {
	IpAddress *string `json:"ip_address"`
	DbType    DbType  `json:"db_type"`
}

func (e UserConnectionLocalConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e UserConnectionLocalConfig) EntityIdentifier() string {
	return "user_connection_local_config"
}

func (e UserConnectionLocalConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["ip_address"] = types.StringFieldType
	res["db_type"] = types.SingleEnumFieldType
	return res
}

func (e UserConnectionLocalConfig) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "ip_address")
	res = append(res, "db_type")

	return res
}

func (e UserConnectionLocalConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e UserConnectionLocalConfig) IsDependant() bool {
	return true
}

func UserConnectionLocalConfigFromJSON(data json.RawMessage) UserConnectionLocalConfig {
	entity := UserConnectionLocalConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func UserConnectionLocalConfigSliceFromJSON(data json.RawMessage) []UserConnectionLocalConfig {
	entity := []UserConnectionLocalConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e UserConnectionLocalConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error UserConnectionLocalConfigToJSON", err)
	}
	return res
}

func UserConnectionLocalConfigToJSON(e UserConnectionLocalConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error UserConnectionLocalConfigToJSON", err)
	}
	return res
}

func UserConnectionLocalConfigSliceToJSON(e []UserConnectionLocalConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error UserConnectionLocalConfigSliceToJSON", err)
	}
	return res
}

func NewUserConnectionLocalConfigWithRandomValues() UserConnectionLocalConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return UserConnectionLocalConfig{
		IpAddress: randomvalues.GetRandomStringValuePtr(),
		DbType:    randomvalues.GetRandomOptionValue[DbType](2),
	}
}

func NewUserConnectionLocalConfigSliceWithRandomValues(n int) []UserConnectionLocalConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []UserConnectionLocalConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewUserConnectionLocalConfigWithRandomValues())
	}
	return res
}
