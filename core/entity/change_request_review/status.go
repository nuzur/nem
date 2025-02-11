package change_request_review

import (
	"encoding/json"
	"fmt"
)

//go:generate go run github.com/dmarkham/enumer -type=Status -json
type Status int64

const (
	STATUS_INVALID Status = iota
	STATUS_DRAFT
	STATUS_IN_REVIEW
	STATUS_APPROVED
	STATUS_REJECTED
	STATUS_PUBLISHED
)

func (e Status) ToInt64() int64 {
	return int64(e)
}

func StatusFromString(in string) Status {
	switch in {
	case "draft":
		return STATUS_DRAFT
	case "in_review":
		return STATUS_IN_REVIEW
	case "approved":
		return STATUS_APPROVED
	case "rejected":
		return STATUS_REJECTED
	case "published":
		return STATUS_PUBLISHED
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
	case STATUS_DRAFT:
		return "draft"
	case STATUS_IN_REVIEW:
		return "in_review"
	case STATUS_APPROVED:
		return "approved"
	case STATUS_REJECTED:
		return "rejected"
	case STATUS_PUBLISHED:
		return "published"
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
