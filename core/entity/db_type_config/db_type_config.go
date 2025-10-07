package db_type_config

import (
	"encoding/json"

	"github.com/nuzur/nem/core/entity/db_type_mysql_config"
	"github.com/nuzur/nem/core/entity/db_type_postgres_config"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"math/rand"

	"time"
)

type DbTypeConfig struct {
	Postgres db_type_postgres_config.DbTypePostgresConfig `json:"postgres"`
	Mysql    db_type_mysql_config.DbTypeMysqlConfig       `json:"mysql"`
}

func (e DbTypeConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e DbTypeConfig) EntityIdentifier() string {
	return "db_type_config"
}

func (e DbTypeConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["postgres"] = types.SingleDependantEntityFieldType
	res["mysql"] = types.SingleDependantEntityFieldType
	return res
}

func (e DbTypeConfig) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "postgres")
	res = append(res, "mysql")

	return res
}

func (e DbTypeConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e DbTypeConfig) IsDependant() bool {
	return true
}

func DbTypeConfigFromJSON(data json.RawMessage) DbTypeConfig {
	entity := DbTypeConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func DbTypeConfigSliceFromJSON(data json.RawMessage) []DbTypeConfig {
	entity := []DbTypeConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e DbTypeConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error DbTypeConfigToJSON", err)
	}
	return res
}

func DbTypeConfigToJSON(e DbTypeConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error DbTypeConfigToJSON", err)
	}
	return res
}

func DbTypeConfigSliceToJSON(e []DbTypeConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error DbTypeConfigSliceToJSON", err)
	}
	return res
}

func NewDbTypeConfigWithRandomValues() DbTypeConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return DbTypeConfig{
		Postgres: db_type_postgres_config.NewDbTypePostgresConfigWithRandomValues(),
		Mysql:    db_type_mysql_config.NewDbTypeMysqlConfigWithRandomValues(),
	}
}

func NewDbTypeConfigSliceWithRandomValues(n int) []DbTypeConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []DbTypeConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewDbTypeConfigWithRandomValues())
	}
	return res
}
