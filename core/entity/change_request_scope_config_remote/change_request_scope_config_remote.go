package change_request_scope_config_remote

import (
	"encoding/json"

	"github.com/gofrs/uuid"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"

	"time"
)

type ChangeRequestScopeConfigRemote struct {
	TeamUUID       uuid.UUID `json:"team_uuid"`
	StoreUUID      uuid.UUID `json:"store_uuid"`
	ConnectionUUID uuid.UUID `json:"connection_uuid"`
}

func (e ChangeRequestScopeConfigRemote) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e ChangeRequestScopeConfigRemote) EntityIdentifier() string {
	return "change_request_scope_config_remote"
}

func (e ChangeRequestScopeConfigRemote) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["team_uuid"] = types.UUIDFieldType
	res["store_uuid"] = types.UUIDFieldType
	res["connection_uuid"] = types.UUIDFieldType
	return res
}

func (e ChangeRequestScopeConfigRemote) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "team_uuid")
	res = append(res, "store_uuid")
	res = append(res, "connection_uuid")

	return res
}

func (e ChangeRequestScopeConfigRemote) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e ChangeRequestScopeConfigRemote) IsDependant() bool {
	return true
}

func ChangeRequestScopeConfigRemoteFromJSON(data json.RawMessage) ChangeRequestScopeConfigRemote {
	entity := ChangeRequestScopeConfigRemote{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func ChangeRequestScopeConfigRemoteSliceFromJSON(data json.RawMessage) []ChangeRequestScopeConfigRemote {
	entity := []ChangeRequestScopeConfigRemote{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e ChangeRequestScopeConfigRemote) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ChangeRequestScopeConfigRemoteToJSON", err)
	}
	return res
}

func ChangeRequestScopeConfigRemoteToJSON(e ChangeRequestScopeConfigRemote) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ChangeRequestScopeConfigRemoteToJSON", err)
	}
	return res
}

func ChangeRequestScopeConfigRemoteSliceToJSON(e []ChangeRequestScopeConfigRemote) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ChangeRequestScopeConfigRemoteSliceToJSON", err)
	}
	return res
}

func NewChangeRequestScopeConfigRemoteWithRandomValues() ChangeRequestScopeConfigRemote {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return ChangeRequestScopeConfigRemote{
		TeamUUID:       randomvalues.GetRandomUUIDValue(),
		StoreUUID:      randomvalues.GetRandomUUIDValue(),
		ConnectionUUID: randomvalues.GetRandomUUIDValue(),
	}
}

func NewChangeRequestScopeConfigRemoteSliceWithRandomValues(n int) []ChangeRequestScopeConfigRemote {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []ChangeRequestScopeConfigRemote{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewChangeRequestScopeConfigRemoteWithRandomValues())
	}
	return res
}
