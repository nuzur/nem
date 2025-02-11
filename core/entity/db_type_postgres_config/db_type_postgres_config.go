package db_type_postgres_config

import (
	"encoding/json"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"

	"time"
)

type DbTypePostgresConfig struct {
	Database string  `json:"database"`
	Sslmode  Sslmode `json:"sslmode"`
}

func (e DbTypePostgresConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e DbTypePostgresConfig) EntityIdentifier() string {
	return "db_type_postgres_config"
}

func (e DbTypePostgresConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["database"] = types.StringFieldType
	res["sslmode"] = types.SingleEnumFieldType
	return res
}

func (e DbTypePostgresConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e DbTypePostgresConfig) IsDependant() bool {
	return true
}

func DbTypePostgresConfigFromJSON(data json.RawMessage) DbTypePostgresConfig {
	entity := DbTypePostgresConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func DbTypePostgresConfigSliceFromJSON(data json.RawMessage) []DbTypePostgresConfig {
	entity := []DbTypePostgresConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e DbTypePostgresConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error DbTypePostgresConfigToJSON", err)
	}
	return res
}

func DbTypePostgresConfigToJSON(e DbTypePostgresConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error DbTypePostgresConfigToJSON", err)
	}
	return res
}

func DbTypePostgresConfigSliceToJSON(e []DbTypePostgresConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error DbTypePostgresConfigSliceToJSON", err)
	}
	return res
}

func NewDbTypePostgresConfigWithRandomValues() DbTypePostgresConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return DbTypePostgresConfig{
		Database: randomvalues.GetRandomStringValue(),
		Sslmode:  randomvalues.GetRandomOptionValue[Sslmode](6),
	}
}

func NewDbTypePostgresConfigSliceWithRandomValues(n int) []DbTypePostgresConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []DbTypePostgresConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewDbTypePostgresConfigWithRandomValues())
	}
	return res
}
