package user_connection_type_config

import (
	"encoding/json"

	"nem/core/entity/user_connection_local_config"
	"nem/core/entity/user_connection_remote_config"

	"fmt"

	"nem/core/entity/types"

	"math/rand"

	"time"
)

type UserConnectionTypeConfig struct {
	Local  user_connection_local_config.UserConnectionLocalConfig   `json:"local"`
	Remote user_connection_remote_config.UserConnectionRemoteConfig `json:"remote"`
}

func (e UserConnectionTypeConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e UserConnectionTypeConfig) EntityIdentifier() string {
	return "user_connection_type_config"
}

func (e UserConnectionTypeConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["local"] = types.SingleDependantEntityFieldType
	res["remote"] = types.SingleDependantEntityFieldType
	return res
}

func (e UserConnectionTypeConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e UserConnectionTypeConfig) IsDependant() bool {
	return true
}

func UserConnectionTypeConfigFromJSON(data json.RawMessage) UserConnectionTypeConfig {
	entity := UserConnectionTypeConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func UserConnectionTypeConfigSliceFromJSON(data json.RawMessage) []UserConnectionTypeConfig {
	entity := []UserConnectionTypeConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e UserConnectionTypeConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error UserConnectionTypeConfigToJSON", err)
	}
	return res
}

func UserConnectionTypeConfigToJSON(e UserConnectionTypeConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error UserConnectionTypeConfigToJSON", err)
	}
	return res
}

func UserConnectionTypeConfigSliceToJSON(e []UserConnectionTypeConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error UserConnectionTypeConfigSliceToJSON", err)
	}
	return res
}

func NewUserConnectionTypeConfigWithRandomValues() UserConnectionTypeConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return UserConnectionTypeConfig{
		Local:  user_connection_local_config.NewUserConnectionLocalConfigWithRandomValues(),
		Remote: user_connection_remote_config.NewUserConnectionRemoteConfigWithRandomValues(),
	}
}

func NewUserConnectionTypeConfigSliceWithRandomValues(n int) []UserConnectionTypeConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []UserConnectionTypeConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewUserConnectionTypeConfigWithRandomValues())
	}
	return res
}
