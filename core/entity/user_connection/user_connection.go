package user_connection

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"github.com/nuzur/nem/core/entity/user_connection_execution"
	"github.com/nuzur/nem/core/entity/user_connection_type_config"
	"time"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"
)

type UserConnection struct {
	UUID               uuid.UUID                                            `json:"uuid"`
	UserUUID           uuid.UUID                                            `json:"user_uuid"`
	ProjectUUID        uuid.UUID                                            `json:"project_uuid"`
	ProjectVersionUUID uuid.UUID                                            `json:"project_version_uuid"`
	Type               Type                                                 `json:"type"`
	TypeConfig         user_connection_type_config.UserConnectionTypeConfig `json:"type_config"`
	DbSchema           string                                               `json:"db_schema"`
	Executions         []user_connection_execution.UserConnectionExecution  `json:"executions"`
	Status             Status                                               `json:"status"`
	CreatedAt          time.Time                                            `json:"created_at"`
	UpdatedAt          time.Time                                            `json:"updated_at"`
}

func (e UserConnection) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e UserConnection) EntityIdentifier() string {
	return "user_connection"
}

func (e UserConnection) PrimaryKeyIdentifier() string {
	return "uuid"
}

func (e UserConnection) PrimaryKeyValue() string {
	return e.UUID.String()
}

func (e UserConnection) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["uuid"] = types.UUIDFieldType
	res["user_uuid"] = types.UUIDFieldType
	res["project_uuid"] = types.UUIDFieldType
	res["project_version_uuid"] = types.UUIDFieldType
	res["type"] = types.SingleEnumFieldType
	res["type_config"] = types.SingleDependantEntityFieldType
	res["db_schema"] = types.StringFieldType
	res["executions"] = types.MultiDependantEntityFieldType
	res["status"] = types.SingleEnumFieldType
	res["created_at"] = types.TimestampFieldType
	res["updated_at"] = types.TimestampFieldType
	return res
}

func (e UserConnection) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "uuid")
	res = append(res, "user_uuid")
	res = append(res, "project_uuid")
	res = append(res, "project_version_uuid")
	res = append(res, "type")
	res = append(res, "type_config")
	res = append(res, "db_schema")
	res = append(res, "executions")
	res = append(res, "status")
	res = append(res, "created_at")
	res = append(res, "updated_at")

	return res
}

func (e UserConnection) DependantFieldIdentifierToTypeMap() map[string]map[string]types.FieldType {
	res := make(map[string]map[string]types.FieldType)

	res["type_config"] = user_connection_type_config.UserConnectionTypeConfig{}.FieldIdentfierToTypeMap()
	res["executions"] = user_connection_execution.UserConnectionExecution{}.FieldIdentfierToTypeMap()
	return res
}

func (e UserConnection) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e UserConnection) IsDependant() bool {
	return false
}

func UserConnectionFromJSON(data json.RawMessage) UserConnection {
	entity := UserConnection{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func UserConnectionSliceFromJSON(data json.RawMessage) []UserConnection {
	entity := []UserConnection{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e UserConnection) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error UserConnectionToJSON", err)
	}
	return res
}

func UserConnectionToJSON(e UserConnection) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error UserConnectionToJSON", err)
	}
	return res
}

func UserConnectionSliceToJSON(e []UserConnection) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error UserConnectionSliceToJSON", err)
	}
	return res
}

func NewUserConnectionWithRandomValues() UserConnection {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return UserConnection{
		UUID:               randomvalues.GetRandomUUIDValue(),
		UserUUID:           randomvalues.GetRandomUUIDValue(),
		ProjectUUID:        randomvalues.GetRandomUUIDValue(),
		ProjectVersionUUID: randomvalues.GetRandomUUIDValue(),
		Type:               randomvalues.GetRandomOptionValue[Type](2),
		TypeConfig:         user_connection_type_config.NewUserConnectionTypeConfigWithRandomValues(),
		DbSchema:           randomvalues.GetRandomStringValue(),
		Executions:         user_connection_execution.NewUserConnectionExecutionSliceWithRandomValues(rand.Intn(10)),
		Status:             randomvalues.GetRandomOptionValue[Status](2),
		CreatedAt:          randomvalues.GetRandomTimeValue(),
		UpdatedAt:          randomvalues.GetRandomTimeValue(),
	}
}

func NewUserConnectionSliceWithRandomValues(n int) []UserConnection {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []UserConnection{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewUserConnectionWithRandomValues())
	}
	return res
}
