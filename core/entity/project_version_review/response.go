package project_version_review

import (
	"encoding/json"
	"fmt"
)

//go:generate go run github.com/dmarkham/enumer -type=Response -json
type Response int64

const (
	RESPONSE_INVALID Response = iota
	RESPONSE_APPROVED
	RESPONSE_REJECTED
)

func (e Response) ToInt64() int64 {
	return int64(e)
}

func ResponseFromString(in string) Response {
	switch in {
	case "approved":
		return RESPONSE_APPROVED
	case "rejected":
		return RESPONSE_REJECTED
	}
	return RESPONSE_INVALID
}

func ResponseFromPointerString(in *string) Response {
	if in == nil {
		return RESPONSE_INVALID
	}
	return ResponseFromString(*in)
}

func (e Response) String() string {
	switch e {
	case RESPONSE_APPROVED:
		return "approved"
	case RESPONSE_REJECTED:
		return "rejected"
	}

	return "invalid"
}

func (e Response) StringPtr() *string {
	val := e.String()
	return &val
}

func ResponseSliceToJSON(in []Response) json.RawMessage {
	res := make([]int64, len(in))
	for i, e := range in {
		res[i] = int64(e)
	}
	jr, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("error marshalling Response slice to json: %v", err)
		return json.RawMessage{}
	}
	return jr
}

func JSONToResponseSlice(in json.RawMessage) []Response {
	res := []int64{}
	err := json.Unmarshal(in, &res)
	if err != nil {
		fmt.Printf("error unmarshing Response slice to int slice: %v", err)
		return nil
	}
	if len(res) == 0 {
		return nil
	}
	finalRes := []Response{}
	for _, r := range res {
		finalRes = append(finalRes, Response(r))
	}
	return finalRes
}
