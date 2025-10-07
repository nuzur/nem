package db_type_mysql_config

import (
	"encoding/json"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"

	"time"
)

type DbTypeMysqlConfig struct {
	Params *string `json:"params"`
}

func (e DbTypeMysqlConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e DbTypeMysqlConfig) EntityIdentifier() string {
	return "db_type_mysql_config"
}

func (e DbTypeMysqlConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["params"] = types.StringFieldType
	return res
}

func (e DbTypeMysqlConfig) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "params")

	return res
}

func (e DbTypeMysqlConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e DbTypeMysqlConfig) IsDependant() bool {
	return true
}

func DbTypeMysqlConfigFromJSON(data json.RawMessage) DbTypeMysqlConfig {
	entity := DbTypeMysqlConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func DbTypeMysqlConfigSliceFromJSON(data json.RawMessage) []DbTypeMysqlConfig {
	entity := []DbTypeMysqlConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e DbTypeMysqlConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error DbTypeMysqlConfigToJSON", err)
	}
	return res
}

func DbTypeMysqlConfigToJSON(e DbTypeMysqlConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error DbTypeMysqlConfigToJSON", err)
	}
	return res
}

func DbTypeMysqlConfigSliceToJSON(e []DbTypeMysqlConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error DbTypeMysqlConfigSliceToJSON", err)
	}
	return res
}

func NewDbTypeMysqlConfigWithRandomValues() DbTypeMysqlConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return DbTypeMysqlConfig{
		Params: randomvalues.GetRandomStringValuePtr(),
	}
}

func NewDbTypeMysqlConfigSliceWithRandomValues(n int) []DbTypeMysqlConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []DbTypeMysqlConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewDbTypeMysqlConfigWithRandomValues())
	}
	return res
}
