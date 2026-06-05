package local_agent

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"github.com/nuzur/nem/core/entity/local_agent_connection"
	"time"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"
)

type LocalAgent struct {
	UUID          uuid.UUID                                     `json:"uuid"`
	UserUUID      uuid.UUID                                     `json:"user_uuid"`
	TokenHash     *string                                       `json:"token_hash"`
	MachineName   *string                                       `json:"machine_name"`
	Os            *string                                       `json:"os"`
	CliVersion    *string                                       `json:"cli_version"`
	Connections   []local_agent_connection.LocalAgentConnection `json:"connections"`
	Status        Status                                        `json:"status"`
	LastSeenAt    *time.Time                                    `json:"last_seen_at"`
	RevokedAt     *time.Time                                    `json:"revoked_at"`
	CreatedAt     time.Time                                     `json:"created_at"`
	UpdatedAt     time.Time                                     `json:"updated_at"`
	CreatedByUUID uuid.UUID                                     `json:"created_by_uuid"`
	UpdatedByUUID uuid.UUID                                     `json:"updated_by_uuid"`
}

func (e LocalAgent) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e LocalAgent) EntityIdentifier() string {
	return "local_agent"
}

func (e LocalAgent) PrimaryKeyIdentifier() string {
	return "uuid"
}

func (e LocalAgent) PrimaryKeyValue() string {
	return e.UUID.String()
}

func (e LocalAgent) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["uuid"] = types.UUIDFieldType
	res["user_uuid"] = types.UUIDFieldType
	res["token_hash"] = types.StringFieldType
	res["machine_name"] = types.StringFieldType
	res["os"] = types.StringFieldType
	res["cli_version"] = types.StringFieldType
	res["connections"] = types.MultiDependantEntityFieldType
	res["status"] = types.SingleEnumFieldType
	res["last_seen_at"] = types.TimestampFieldType
	res["revoked_at"] = types.TimestampFieldType
	res["created_at"] = types.TimestampFieldType
	res["updated_at"] = types.TimestampFieldType
	res["created_by_uuid"] = types.UUIDFieldType
	res["updated_by_uuid"] = types.UUIDFieldType
	return res
}

func (e LocalAgent) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "uuid")
	res = append(res, "user_uuid")
	res = append(res, "token_hash")
	res = append(res, "machine_name")
	res = append(res, "os")
	res = append(res, "cli_version")
	res = append(res, "connections")
	res = append(res, "status")
	res = append(res, "last_seen_at")
	res = append(res, "revoked_at")
	res = append(res, "created_at")
	res = append(res, "updated_at")
	res = append(res, "created_by_uuid")
	res = append(res, "updated_by_uuid")

	return res
}

func (e LocalAgent) DependantFieldIdentifierToTypeMap() map[string]map[string]types.FieldType {
	res := make(map[string]map[string]types.FieldType)

	res["connections"] = local_agent_connection.LocalAgentConnection{}.FieldIdentfierToTypeMap()
	return res
}

func (e LocalAgent) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e LocalAgent) IsDependant() bool {
	return false
}

func LocalAgentFromJSON(data json.RawMessage) LocalAgent {
	entity := LocalAgent{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func LocalAgentSliceFromJSON(data json.RawMessage) []LocalAgent {
	entity := []LocalAgent{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e LocalAgent) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error LocalAgentToJSON", err)
	}
	return res
}

func LocalAgentToJSON(e LocalAgent) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error LocalAgentToJSON", err)
	}
	return res
}

func LocalAgentSliceToJSON(e []LocalAgent) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error LocalAgentSliceToJSON", err)
	}
	return res
}

func NewLocalAgentWithRandomValues() LocalAgent {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return LocalAgent{
		UUID:          randomvalues.GetRandomUUIDValue(),
		UserUUID:      randomvalues.GetRandomUUIDValue(),
		TokenHash:     randomvalues.GetRandomStringValuePtr(),
		MachineName:   randomvalues.GetRandomStringValuePtr(),
		Os:            randomvalues.GetRandomStringValuePtr(),
		CliVersion:    randomvalues.GetRandomStringValuePtr(),
		Connections:   local_agent_connection.NewLocalAgentConnectionSliceWithRandomValues(rand.Intn(10)),
		Status:        randomvalues.GetRandomOptionValue[Status](3),
		LastSeenAt:    randomvalues.GetRandomTimeValuePtr(),
		RevokedAt:     randomvalues.GetRandomTimeValuePtr(),
		CreatedAt:     randomvalues.GetRandomTimeValue(),
		UpdatedAt:     randomvalues.GetRandomTimeValue(),
		CreatedByUUID: randomvalues.GetRandomUUIDValue(),
		UpdatedByUUID: randomvalues.GetRandomUUIDValue(),
	}
}

func NewLocalAgentSliceWithRandomValues(n int) []LocalAgent {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []LocalAgent{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewLocalAgentWithRandomValues())
	}
	return res
}
