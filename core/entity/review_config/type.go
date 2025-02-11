package review_config

import (
	"encoding/json"
	"fmt"
)

//go:generate go run github.com/dmarkham/enumer -type=Type -json
type Type int64

const (
	TYPE_INVALID Type = iota
	TYPE_DEFAULT
	TYPE_PROJECT_VERSION
	TYPE_EXTENSION_VERSION
)

func (e Type) ToInt64() int64 {
	return int64(e)
}

func TypeFromString(in string) Type {
	switch in {
	case "default":
		return TYPE_DEFAULT
	case "project_version":
		return TYPE_PROJECT_VERSION
	case "extension_version":
		return TYPE_EXTENSION_VERSION
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
	case TYPE_DEFAULT:
		return "default"
	case TYPE_PROJECT_VERSION:
		return "project_version"
	case TYPE_EXTENSION_VERSION:
		return "extension_version"
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
