package project

import (
	"encoding/json"
	"fmt"
)

//go:generate go run github.com/dmarkham/enumer -type=AccessType -json
type AccessType int64

const (
	ACCESS_TYPE_INVALID AccessType = iota
	ACCESS_TYPE_INHERIT
	ACCESS_TYPE_CUSTOM
)

func (e AccessType) ToInt64() int64 {
	return int64(e)
}

func AccessTypeFromString(in string) AccessType {
	switch in {
	case "inherit":
		return ACCESS_TYPE_INHERIT
	case "custom":
		return ACCESS_TYPE_CUSTOM
	}
	return ACCESS_TYPE_INVALID
}

func AccessTypeFromPointerString(in *string) AccessType {
	if in == nil {
		return ACCESS_TYPE_INVALID
	}
	return AccessTypeFromString(*in)
}

func (e AccessType) String() string {
	switch e {
	case ACCESS_TYPE_INHERIT:
		return "inherit"
	case ACCESS_TYPE_CUSTOM:
		return "custom"
	}

	return "invalid"
}

func (e AccessType) StringPtr() *string {
	val := e.String()
	return &val
}

func AccessTypeSliceToJSON(in []AccessType) json.RawMessage {
	res := make([]int64, len(in))
	for i, e := range in {
		res[i] = int64(e)
	}
	jr, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("error marshalling AccessType slice to json: %v", err)
		return json.RawMessage{}
	}
	return jr
}

func JSONToAccessTypeSlice(in json.RawMessage) []AccessType {
	res := []int64{}
	err := json.Unmarshal(in, &res)
	if err != nil {
		fmt.Printf("error unmarshing AccessType slice to int slice: %v", err)
		return nil
	}
	if len(res) == 0 {
		return nil
	}
	finalRes := []AccessType{}
	for _, r := range res {
		finalRes = append(finalRes, AccessType(r))
	}
	return finalRes
}
