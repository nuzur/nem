package field_type_array_config

import (
	"encoding/json"
	"fmt"
)

//go:generate go run github.com/dmarkham/enumer -type=Type -json
type Type int64

const (
	TYPE_INVALID Type = iota
	TYPE_UUID
	TYPE_INTEGER
	TYPE_FLOAT
	TYPE_DECIMAL
	TYPE_CHAR
	TYPE_VARCHAR
	TYPE_EMAIL
	TYPE_PHONE
	TYPE_URL
	TYPE_COLOR
	TYPE_DATE
	TYPE_DATETIME
	TYPE_ENCRYPTED
	TYPE_TIME
	TYPE_ENUM
)

func (e Type) ToInt64() int64 {
	return int64(e)
}

func TypeFromString(in string) Type {
	switch in {
	case "uuid":
		return TYPE_UUID
	case "integer":
		return TYPE_INTEGER
	case "float":
		return TYPE_FLOAT
	case "decimal":
		return TYPE_DECIMAL
	case "char":
		return TYPE_CHAR
	case "varchar":
		return TYPE_VARCHAR
	case "email":
		return TYPE_EMAIL
	case "phone":
		return TYPE_PHONE
	case "url":
		return TYPE_URL
	case "color":
		return TYPE_COLOR
	case "date":
		return TYPE_DATE
	case "datetime":
		return TYPE_DATETIME
	case "encrypted":
		return TYPE_ENCRYPTED
	case "time":
		return TYPE_TIME
	case "enum":
		return TYPE_ENUM
	}
	return TYPE_INVALID
}

func TypeFromPointerString(in *string) Type {
	if in == nil {
		return TYPE_INVALID
	}
	return TypeFromString(*in)
}

func (e Type) String() string {
	switch e {
	case TYPE_UUID:
		return "uuid"
	case TYPE_INTEGER:
		return "integer"
	case TYPE_FLOAT:
		return "float"
	case TYPE_DECIMAL:
		return "decimal"
	case TYPE_CHAR:
		return "char"
	case TYPE_VARCHAR:
		return "varchar"
	case TYPE_EMAIL:
		return "email"
	case TYPE_PHONE:
		return "phone"
	case TYPE_URL:
		return "url"
	case TYPE_COLOR:
		return "color"
	case TYPE_DATE:
		return "date"
	case TYPE_DATETIME:
		return "datetime"
	case TYPE_ENCRYPTED:
		return "encrypted"
	case TYPE_TIME:
		return "time"
	case TYPE_ENUM:
		return "enum"
	}

	return "invalid"
}

func (e Type) StringPtr() *string {
	val := e.String()
	return &val
}

func TypeSliceToJSON(in []Type) json.RawMessage {
	res := make([]int64, len(in))
	for i, e := range in {
		res[i] = int64(e)
	}
	jr, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("error marshalling Type slice to json: %v", err)
		return json.RawMessage{}
	}
	return jr
}

func JSONToTypeSlice(in json.RawMessage) []Type {
	res := []int64{}
	err := json.Unmarshal(in, &res)
	if err != nil {
		fmt.Printf("error unmarshing Type slice to int slice: %v", err)
		return nil
	}
	if len(res) == 0 {
		return nil
	}
	finalRes := []Type{}
	for _, r := range res {
		finalRes = append(finalRes, Type(r))
	}
	return finalRes
}
