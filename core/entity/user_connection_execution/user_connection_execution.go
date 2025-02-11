package user_connection_execution

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"time"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"
)

type UserConnectionExecution struct {
	UUID          uuid.UUID `json:"uuid"`
	Status        Status    `json:"status"`
	ResultsPath   string    `json:"results_path"`
	StartedAt     time.Time `json:"started_at"`
	FinishedAt    time.Time `json:"finished_at"`
	EstimatedTime string    `json:"estimated_time"`
	UserMsg       string    `json:"user_msg"`
}

func (e UserConnectionExecution) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e UserConnectionExecution) EntityIdentifier() string {
	return "user_connection_execution"
}

func (e UserConnectionExecution) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["uuid"] = types.UUIDFieldType
	res["status"] = types.SingleEnumFieldType
	res["results_path"] = types.StringFieldType
	res["started_at"] = types.TimestampFieldType
	res["finished_at"] = types.TimestampFieldType
	res["estimated_time"] = types.StringFieldType
	res["user_msg"] = types.StringFieldType
	return res
}

func (e UserConnectionExecution) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e UserConnectionExecution) IsDependant() bool {
	return true
}

func UserConnectionExecutionFromJSON(data json.RawMessage) UserConnectionExecution {
	entity := UserConnectionExecution{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func UserConnectionExecutionSliceFromJSON(data json.RawMessage) []UserConnectionExecution {
	entity := []UserConnectionExecution{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e UserConnectionExecution) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error UserConnectionExecutionToJSON", err)
	}
	return res
}

func UserConnectionExecutionToJSON(e UserConnectionExecution) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error UserConnectionExecutionToJSON", err)
	}
	return res
}

func UserConnectionExecutionSliceToJSON(e []UserConnectionExecution) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error UserConnectionExecutionSliceToJSON", err)
	}
	return res
}

func NewUserConnectionExecutionWithRandomValues() UserConnectionExecution {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return UserConnectionExecution{
		UUID:          randomvalues.GetRandomUUIDValue(),
		Status:        randomvalues.GetRandomOptionValue[Status](3),
		ResultsPath:   randomvalues.GetRandomStringValue(),
		StartedAt:     randomvalues.GetRandomTimeValue(),
		FinishedAt:    randomvalues.GetRandomTimeValue(),
		EstimatedTime: randomvalues.GetRandomStringValue(),
		UserMsg:       randomvalues.GetRandomStringValue(),
	}
}

func NewUserConnectionExecutionSliceWithRandomValues(n int) []UserConnectionExecution {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []UserConnectionExecution{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewUserConnectionExecutionWithRandomValues())
	}
	return res
}
