package connection

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"github.com/nuzur/nem/core/entity/connection_type_config"
	"github.com/nuzur/nem/core/entity/db_type_config"
	"time"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"
)

type Connection struct {
	UUID           uuid.UUID                                   `json:"uuid"`
	Version        int64                                       `json:"version"`
	StoreUUID      uuid.UUID                                   `json:"store_uuid"`
	EnviormentUUID uuid.UUID                                   `json:"enviorment_uuid"`
	Identifier     string                                      `json:"identifier"`
	DbType         DbType                                      `json:"db_type"`
	DbTypeConfig   db_type_config.DbTypeConfig                 `json:"db_type_config"`
	DbVersion      string                                      `json:"db_version"`
	Type           Type                                        `json:"type"`
	TypeConfig     connection_type_config.ConnectionTypeConfig `json:"type_config"`
	Status         Status                                      `json:"status"`
	CreatedAt      time.Time                                   `json:"created_at"`
	UpdatedAt      time.Time                                   `json:"updated_at"`
	CreatedByUUID  uuid.UUID                                   `json:"created_by_uuid"`
	UpdatedByUUID  uuid.UUID                                   `json:"updated_by_uuid"`
}

func (e Connection) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e Connection) EntityIdentifier() string {
	return "connection"
}

func (e Connection) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["uuid"] = types.UUIDFieldType
	res["version"] = types.IntFieldType
	res["store_uuid"] = types.UUIDFieldType
	res["enviorment_uuid"] = types.UUIDFieldType
	res["identifier"] = types.StringFieldType
	res["db_type"] = types.SingleEnumFieldType
	res["db_type_config"] = types.SingleDependantEntityFieldType
	res["db_version"] = types.StringFieldType
	res["type"] = types.SingleEnumFieldType
	res["type_config"] = types.SingleDependantEntityFieldType
	res["status"] = types.SingleEnumFieldType
	res["created_at"] = types.TimestampFieldType
	res["updated_at"] = types.TimestampFieldType
	res["created_by_uuid"] = types.UUIDFieldType
	res["updated_by_uuid"] = types.UUIDFieldType
	return res
}

func (e Connection) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e Connection) IsDependant() bool {
	return true
}

func ConnectionFromJSON(data json.RawMessage) Connection {
	entity := Connection{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func ConnectionSliceFromJSON(data json.RawMessage) []Connection {
	entity := []Connection{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e Connection) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ConnectionToJSON", err)
	}
	return res
}

func ConnectionToJSON(e Connection) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ConnectionToJSON", err)
	}
	return res
}

func ConnectionSliceToJSON(e []Connection) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ConnectionSliceToJSON", err)
	}
	return res
}

func NewConnectionWithRandomValues() Connection {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return Connection{
		UUID:           randomvalues.GetRandomUUIDValue(),
		Version:        randomvalues.GetRandomIntValue(),
		StoreUUID:      randomvalues.GetRandomUUIDValue(),
		EnviormentUUID: randomvalues.GetRandomUUIDValue(),
		Identifier:     randomvalues.GetRandomStringValue(),
		DbType:         randomvalues.GetRandomOptionValue[DbType](2),
		DbTypeConfig:   db_type_config.NewDbTypeConfigWithRandomValues(),
		DbVersion:      randomvalues.GetRandomStringValue(),
		Type:           randomvalues.GetRandomOptionValue[Type](2),
		TypeConfig:     connection_type_config.NewConnectionTypeConfigWithRandomValues(),
		Status:         randomvalues.GetRandomOptionValue[Status](2),
		CreatedAt:      randomvalues.GetRandomTimeValue(),
		UpdatedAt:      randomvalues.GetRandomTimeValue(),
		CreatedByUUID:  randomvalues.GetRandomUUIDValue(),
		UpdatedByUUID:  randomvalues.GetRandomUUIDValue(),
	}
}

func NewConnectionSliceWithRandomValues(n int) []Connection {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []Connection{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewConnectionWithRandomValues())
	}
	return res
}
