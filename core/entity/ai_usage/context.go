package ai_usage

import (
	"encoding/json"
	"fmt"
)

//go:generate go run github.com/dmarkham/enumer -type=Context -json
type Context int64

const (
	CONTEXT_INVALID Context = iota
	CONTEXT_EDITOR_ASSISTANT
	CONTEXT_DATA_MANAGER
	CONTEXT_EDITOR_LAYOUT
)

func (e Context) ToInt64() int64 {
	return int64(e)
}

func ContextFromString(in string) Context {
	switch in {
	case "editor_assistant":
		return CONTEXT_EDITOR_ASSISTANT
	case "data_manager":
		return CONTEXT_DATA_MANAGER
	case "editor_layout":
		return CONTEXT_EDITOR_LAYOUT
	}
	return CONTEXT_INVALID
}

func ContextFromPointerString(in *string) Context {
	if in == nil {
		return CONTEXT_INVALID
	}
	return ContextFromString(*in)
}

func (e Context) String() string {
	switch e {
	case CONTEXT_EDITOR_ASSISTANT:
		return "editor_assistant"
	case CONTEXT_DATA_MANAGER:
		return "data_manager"
	case CONTEXT_EDITOR_LAYOUT:
		return "editor_layout"
	}

	return "invalid"
}

func (e Context) StringPtr() *string {
	val := e.String()
	return &val
}

func ContextSliceToJSON(in []Context) json.RawMessage {
	res := make([]int64, len(in))
	for i, e := range in {
		res[i] = int64(e)
	}
	jr, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("error marshalling Context slice to json: %v", err)
		return json.RawMessage{}
	}
	return jr
}

func JSONToContextSlice(in json.RawMessage) []Context {
	res := []int64{}
	err := json.Unmarshal(in, &res)
	if err != nil {
		fmt.Printf("error unmarshing Context slice to int slice: %v", err)
		return nil
	}
	if len(res) == 0 {
		return nil
	}
	finalRes := []Context{}
	for _, r := range res {
		finalRes = append(finalRes, Context(r))
	}
	return finalRes
}
