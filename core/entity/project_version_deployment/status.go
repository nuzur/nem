package project_version_deployment

import (
	"encoding/json"
	"fmt"
)

//go:generate go run github.com/dmarkham/enumer -type=Status -json
type Status int64

const (
	STATUS_INVALID Status = iota
	STATUS_NOT_STARTED
	STATUS_STARTED
	STATUS_COMPLETED
	STATUS_FAILED
	STATUS_ROLLED_BACK
)

func (e Status) ToInt64() int64 {
	return int64(e)
}

func StatusFromString(in string) Status {
	switch in {
	case "not_started":
		return STATUS_NOT_STARTED
	case "started":
		return STATUS_STARTED
	case "completed":
		return STATUS_COMPLETED
	case "failed":
		return STATUS_FAILED
	case "rolled_back":
		return STATUS_ROLLED_BACK
	}
	return STATUS_INVALID
}

func StatusFromPointerString(in *string) Status {
	if in == nil {
		return STATUS_INVALID
	}
	return StatusFromString(*in)
}

func (e Status) String() string {
	switch e {
	case STATUS_NOT_STARTED:
		return "not_started"
	case STATUS_STARTED:
		return "started"
	case STATUS_COMPLETED:
		return "completed"
	case STATUS_FAILED:
		return "failed"
	case STATUS_ROLLED_BACK:
		return "rolled_back"
	}

	return "invalid"
}

func (e Status) StringPtr() *string {
	val := e.String()
	return &val
}

func StatusSliceToJSON(in []Status) json.RawMessage {
	res := make([]int64, len(in))
	for i, e := range in {
		res[i] = int64(e)
	}
	jr, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("error marshalling Status slice to json: %v", err)
		return json.RawMessage{}
	}
	return jr
}

func JSONToStatusSlice(in json.RawMessage) []Status {
	res := []int64{}
	err := json.Unmarshal(in, &res)
	if err != nil {
		fmt.Printf("error unmarshing Status slice to int slice: %v", err)
		return nil
	}
	if len(res) == 0 {
		return nil
	}
	finalRes := []Status{}
	for _, r := range res {
		finalRes = append(finalRes, Status(r))
	}
	return finalRes
}
