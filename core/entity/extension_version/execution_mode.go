package extension_version

import (
	"encoding/json"
	"fmt"
)

//go:generate go run github.com/dmarkham/enumer -type=ExecutionMode -json
type ExecutionMode int64

const (
	EXECUTION_MODE_INVALID ExecutionMode = iota
	EXECUTION_MODE_MANUAL
	EXECUTION_MODE_AUTOMATIC
)

func (e ExecutionMode) ToInt64() int64 {
	return int64(e)
}

func ExecutionModeFromString(in string) ExecutionMode {
	switch in {
	case "manual":
		return EXECUTION_MODE_MANUAL
	case "automatic":
		return EXECUTION_MODE_AUTOMATIC
	}
	return EXECUTION_MODE_INVALID
}

func ExecutionModeFromPointerString(in *string) ExecutionMode {
	if in == nil {
		return EXECUTION_MODE_INVALID
	}
	return ExecutionModeFromString(*in)
}

func (e ExecutionMode) String() string {
	switch e {
	case EXECUTION_MODE_MANUAL:
		return "manual"
	case EXECUTION_MODE_AUTOMATIC:
		return "automatic"
	}

	return "invalid"
}

func (e ExecutionMode) StringPtr() *string {
	val := e.String()
	return &val
}

func ExecutionModeSliceToJSON(in []ExecutionMode) json.RawMessage {
	res := make([]int64, len(in))
	for i, e := range in {
		res[i] = int64(e)
	}
	jr, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("error marshalling ExecutionMode slice to json: %v", err)
		return json.RawMessage{}
	}
	return jr
}

func JSONToExecutionModeSlice(in json.RawMessage) []ExecutionMode {
	res := []int64{}
	err := json.Unmarshal(in, &res)
	if err != nil {
		fmt.Printf("error unmarshing ExecutionMode slice to int slice: %v", err)
		return nil
	}
	if len(res) == 0 {
		return nil
	}
	finalRes := []ExecutionMode{}
	for _, r := range res {
		finalRes = append(finalRes, ExecutionMode(r))
	}
	return finalRes
}
