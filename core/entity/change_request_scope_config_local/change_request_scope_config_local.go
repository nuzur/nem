package change_request_scope_config_local

import (
	"encoding/json"

	"github.com/gofrs/uuid"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"

	"time"
)

type ChangeRequestScopeConfigLocal struct {
	AgentUUID           uuid.UUID `json:"agent_uuid"`
	AgentConnectionUUID uuid.UUID `json:"agent_connection_uuid"`
}

func (e ChangeRequestScopeConfigLocal) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e ChangeRequestScopeConfigLocal) EntityIdentifier() string {
	return "change_request_scope_config_local"
}

func (e ChangeRequestScopeConfigLocal) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["agent_uuid"] = types.UUIDFieldType
	res["agent_connection_uuid"] = types.UUIDFieldType
	return res
}

func (e ChangeRequestScopeConfigLocal) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "agent_uuid")
	res = append(res, "agent_connection_uuid")

	return res
}

func (e ChangeRequestScopeConfigLocal) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e ChangeRequestScopeConfigLocal) IsDependant() bool {
	return true
}

func ChangeRequestScopeConfigLocalFromJSON(data json.RawMessage) ChangeRequestScopeConfigLocal {
	entity := ChangeRequestScopeConfigLocal{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func ChangeRequestScopeConfigLocalSliceFromJSON(data json.RawMessage) []ChangeRequestScopeConfigLocal {
	entity := []ChangeRequestScopeConfigLocal{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e ChangeRequestScopeConfigLocal) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ChangeRequestScopeConfigLocalToJSON", err)
	}
	return res
}

func ChangeRequestScopeConfigLocalToJSON(e ChangeRequestScopeConfigLocal) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ChangeRequestScopeConfigLocalToJSON", err)
	}
	return res
}

func ChangeRequestScopeConfigLocalSliceToJSON(e []ChangeRequestScopeConfigLocal) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ChangeRequestScopeConfigLocalSliceToJSON", err)
	}
	return res
}

func NewChangeRequestScopeConfigLocalWithRandomValues() ChangeRequestScopeConfigLocal {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return ChangeRequestScopeConfigLocal{
		AgentUUID:           randomvalues.GetRandomUUIDValue(),
		AgentConnectionUUID: randomvalues.GetRandomUUIDValue(),
	}
}

func NewChangeRequestScopeConfigLocalSliceWithRandomValues(n int) []ChangeRequestScopeConfigLocal {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []ChangeRequestScopeConfigLocal{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewChangeRequestScopeConfigLocalWithRandomValues())
	}
	return res
}
