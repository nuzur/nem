package user_connection_remote_config

import (
	"encoding/json"

	"github.com/gofrs/uuid"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"

	"time"
)

type UserConnectionRemoteConfig struct {
	TeamUUID       uuid.UUID `json:"team_uuid"`
	StoreUUID      uuid.UUID `json:"store_uuid"`
	ConnectionUUID uuid.UUID `json:"connection_uuid"`
}

func (e UserConnectionRemoteConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e UserConnectionRemoteConfig) EntityIdentifier() string {
	return "user_connection_remote_config"
}

func (e UserConnectionRemoteConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["team_uuid"] = types.UUIDFieldType
	res["store_uuid"] = types.UUIDFieldType
	res["connection_uuid"] = types.UUIDFieldType
	return res
}

func (e UserConnectionRemoteConfig) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "team_uuid")
	res = append(res, "store_uuid")
	res = append(res, "connection_uuid")

	return res
}

func (e UserConnectionRemoteConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e UserConnectionRemoteConfig) IsDependant() bool {
	return true
}

func UserConnectionRemoteConfigFromJSON(data json.RawMessage) UserConnectionRemoteConfig {
	entity := UserConnectionRemoteConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func UserConnectionRemoteConfigSliceFromJSON(data json.RawMessage) []UserConnectionRemoteConfig {
	entity := []UserConnectionRemoteConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e UserConnectionRemoteConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error UserConnectionRemoteConfigToJSON", err)
	}
	return res
}

func UserConnectionRemoteConfigToJSON(e UserConnectionRemoteConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error UserConnectionRemoteConfigToJSON", err)
	}
	return res
}

func UserConnectionRemoteConfigSliceToJSON(e []UserConnectionRemoteConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error UserConnectionRemoteConfigSliceToJSON", err)
	}
	return res
}

func NewUserConnectionRemoteConfigWithRandomValues() UserConnectionRemoteConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return UserConnectionRemoteConfig{
		TeamUUID:       randomvalues.GetRandomUUIDValue(),
		StoreUUID:      randomvalues.GetRandomUUIDValue(),
		ConnectionUUID: randomvalues.GetRandomUUIDValue(),
	}
}

func NewUserConnectionRemoteConfigSliceWithRandomValues(n int) []UserConnectionRemoteConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []UserConnectionRemoteConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewUserConnectionRemoteConfigWithRandomValues())
	}
	return res
}
