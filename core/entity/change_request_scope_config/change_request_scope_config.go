package change_request_scope_config

import (
	"encoding/json"

	"github.com/nuzur/nem/core/entity/change_request_scope_config_local"
	"github.com/nuzur/nem/core/entity/change_request_scope_config_remote"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"math/rand"

	"time"
)

type ChangeRequestScopeConfig struct {
	Local  change_request_scope_config_local.ChangeRequestScopeConfigLocal   `json:"local"`
	Remote change_request_scope_config_remote.ChangeRequestScopeConfigRemote `json:"remote"`
}

func (e ChangeRequestScopeConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e ChangeRequestScopeConfig) EntityIdentifier() string {
	return "change_request_scope_config"
}

func (e ChangeRequestScopeConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["local"] = types.SingleDependantEntityFieldType
	res["remote"] = types.SingleDependantEntityFieldType
	return res
}

func (e ChangeRequestScopeConfig) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "local")
	res = append(res, "remote")

	return res
}

func (e ChangeRequestScopeConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e ChangeRequestScopeConfig) IsDependant() bool {
	return true
}

func ChangeRequestScopeConfigFromJSON(data json.RawMessage) ChangeRequestScopeConfig {
	entity := ChangeRequestScopeConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func ChangeRequestScopeConfigSliceFromJSON(data json.RawMessage) []ChangeRequestScopeConfig {
	entity := []ChangeRequestScopeConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e ChangeRequestScopeConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ChangeRequestScopeConfigToJSON", err)
	}
	return res
}

func ChangeRequestScopeConfigToJSON(e ChangeRequestScopeConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ChangeRequestScopeConfigToJSON", err)
	}
	return res
}

func ChangeRequestScopeConfigSliceToJSON(e []ChangeRequestScopeConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ChangeRequestScopeConfigSliceToJSON", err)
	}
	return res
}

func NewChangeRequestScopeConfigWithRandomValues() ChangeRequestScopeConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return ChangeRequestScopeConfig{
		Local:  change_request_scope_config_local.NewChangeRequestScopeConfigLocalWithRandomValues(),
		Remote: change_request_scope_config_remote.NewChangeRequestScopeConfigRemoteWithRandomValues(),
	}
}

func NewChangeRequestScopeConfigSliceWithRandomValues(n int) []ChangeRequestScopeConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []ChangeRequestScopeConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewChangeRequestScopeConfigWithRandomValues())
	}
	return res
}
