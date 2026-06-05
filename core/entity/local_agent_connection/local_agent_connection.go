package local_agent_connection

import (
	"encoding/json"

	"github.com/gofrs/uuid"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"

	"time"
)

type LocalAgentConnection struct {
	UUID          uuid.UUID `json:"uuid"`
	Name          string    `json:"name"`
	DbType        DbType    `json:"db_type"`
	DefaultSchema string    `json:"default_schema"`
}

func (e LocalAgentConnection) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e LocalAgentConnection) EntityIdentifier() string {
	return "local_agent_connection"
}

func (e LocalAgentConnection) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["uuid"] = types.UUIDFieldType
	res["name"] = types.StringFieldType
	res["db_type"] = types.SingleEnumFieldType
	res["default_schema"] = types.StringFieldType
	return res
}

func (e LocalAgentConnection) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "uuid")
	res = append(res, "name")
	res = append(res, "db_type")
	res = append(res, "default_schema")

	return res
}

func (e LocalAgentConnection) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e LocalAgentConnection) IsDependant() bool {
	return true
}

func LocalAgentConnectionFromJSON(data json.RawMessage) LocalAgentConnection {
	entity := LocalAgentConnection{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func LocalAgentConnectionSliceFromJSON(data json.RawMessage) []LocalAgentConnection {
	entity := []LocalAgentConnection{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e LocalAgentConnection) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error LocalAgentConnectionToJSON", err)
	}
	return res
}

func LocalAgentConnectionToJSON(e LocalAgentConnection) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error LocalAgentConnectionToJSON", err)
	}
	return res
}

func LocalAgentConnectionSliceToJSON(e []LocalAgentConnection) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error LocalAgentConnectionSliceToJSON", err)
	}
	return res
}

func NewLocalAgentConnectionWithRandomValues() LocalAgentConnection {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return LocalAgentConnection{
		UUID:          randomvalues.GetRandomUUIDValue(),
		Name:          randomvalues.GetRandomStringValue(),
		DbType:        randomvalues.GetRandomOptionValue[DbType](2),
		DefaultSchema: randomvalues.GetRandomStringValue(),
	}
}

func NewLocalAgentConnectionSliceWithRandomValues(n int) []LocalAgentConnection {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []LocalAgentConnection{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewLocalAgentConnectionWithRandomValues())
	}
	return res
}
