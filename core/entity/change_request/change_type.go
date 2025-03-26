package change_request

import (
	"encoding/json"
	"fmt"
)

//go:generate go run github.com/dmarkham/enumer -type=ChangeType -json
type ChangeType int64

const (
	CHANGE_TYPE_INVALID ChangeType = iota
	CHANGE_TYPE_PROJECT_DATA
	CHANGE_TYPE_PROJECT_VERSION
)

func (e ChangeType) ToInt64() int64 {
	return int64(e)
}

func ChangeTypeFromString(in string) ChangeType {
	switch in {
	case "project_data":
		return CHANGE_TYPE_PROJECT_DATA
	case "project_version":
		return CHANGE_TYPE_PROJECT_VERSION
	}
	return CHANGE_TYPE_INVALID
}

func ChangeTypeFromPointerString(in *string) ChangeType {
	if in == nil {
		return CHANGE_TYPE_INVALID
	}
	return ChangeTypeFromString(*in)
}

func (e ChangeType) String() string {
	switch e {
	case CHANGE_TYPE_PROJECT_DATA:
		return "project_data"
	case CHANGE_TYPE_PROJECT_VERSION:
		return "project_version"
	}

	return "invalid"
}

func (e ChangeType) StringPtr() *string {
	val := e.String()
	return &val
}

func ChangeTypeSliceToJSON(in []ChangeType) json.RawMessage {
	res := make([]int64, len(in))
	for i, e := range in {
		res[i] = int64(e)
	}
	jr, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("error marshalling ChangeType slice to json: %v", err)
		return json.RawMessage{}
	}
	return jr
}

func JSONToChangeTypeSlice(in json.RawMessage) []ChangeType {
	res := []int64{}
	err := json.Unmarshal(in, &res)
	if err != nil {
		fmt.Printf("error unmarshing ChangeType slice to int slice: %v", err)
		return nil
	}
	if len(res) == 0 {
		return nil
	}
	finalRes := []ChangeType{}
	for _, r := range res {
		finalRes = append(finalRes, ChangeType(r))
	}
	return finalRes
}
