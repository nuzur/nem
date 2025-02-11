package connection

import (
	"encoding/json"
	"fmt"
)

//go:generate go run github.com/dmarkham/enumer -type=DbType -json
type DbType int64

const (
	DB_TYPE_INVALID DbType = iota
	DB_TYPE_MYSQL
	DB_TYPE_POSTGRES
)

func (e DbType) ToInt64() int64 {
	return int64(e)
}

func DbTypeFromString(in string) DbType {
	switch in {
	case "mysql":
		return DB_TYPE_MYSQL
	case "postgres":
		return DB_TYPE_POSTGRES
	}
	return DB_TYPE_INVALID
}

func DbTypeFromPointerString(in *string) DbType {
	if in == nil {
		return DB_TYPE_INVALID
	}
	return DbTypeFromString(*in)
}

func (e DbType) String() string {
	switch e {
	case DB_TYPE_MYSQL:
		return "mysql"
	case DB_TYPE_POSTGRES:
		return "postgres"
	}

	return "invalid"
}

func (e DbType) StringPtr() *string {
	val := e.String()
	return &val
}

func DbTypeSliceToJSON(in []DbType) json.RawMessage {
	res := make([]int64, len(in))
	for i, e := range in {
		res[i] = int64(e)
	}
	jr, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("error marshalling DbType slice to json: %v", err)
		return json.RawMessage{}
	}
	return jr
}

func JSONToDbTypeSlice(in json.RawMessage) []DbType {
	res := []int64{}
	err := json.Unmarshal(in, &res)
	if err != nil {
		fmt.Printf("error unmarshing DbType slice to int slice: %v", err)
		return nil
	}
	if len(res) == 0 {
		return nil
	}
	finalRes := []DbType{}
	for _, r := range res {
		finalRes = append(finalRes, DbType(r))
	}
	return finalRes
}
