package change_request

import (
	"encoding/json"
	"fmt"
)

//go:generate go run github.com/dmarkham/enumer -type=ReviewType -json
type ReviewType int64

const (
	REVIEW_TYPE_INVALID ReviewType = iota
	REVIEW_TYPE_PROJECT_DATA
	REVIEW_TYPE_PROJECT_VERSION
	REVIEW_TYPE_EXTENSION_VERSION
)

func (e ReviewType) ToInt64() int64 {
	return int64(e)
}

func ReviewTypeFromString(in string) ReviewType {
	switch in {
	case "project_data":
		return REVIEW_TYPE_PROJECT_DATA
	case "project_version":
		return REVIEW_TYPE_PROJECT_VERSION
	case "extension_version":
		return REVIEW_TYPE_EXTENSION_VERSION
	}
	return REVIEW_TYPE_INVALID
}

func ReviewTypeFromPointerString(in *string) ReviewType {
	if in == nil {
		return REVIEW_TYPE_INVALID
	}
	return ReviewTypeFromString(*in)
}

func (e ReviewType) String() string {
	switch e {
	case REVIEW_TYPE_PROJECT_DATA:
		return "project_data"
	case REVIEW_TYPE_PROJECT_VERSION:
		return "project_version"
	case REVIEW_TYPE_EXTENSION_VERSION:
		return "extension_version"
	}

	return "invalid"
}

func (e ReviewType) StringPtr() *string {
	val := e.String()
	return &val
}

func ReviewTypeSliceToJSON(in []ReviewType) json.RawMessage {
	res := make([]int64, len(in))
	for i, e := range in {
		res[i] = int64(e)
	}
	jr, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("error marshalling ReviewType slice to json: %v", err)
		return json.RawMessage{}
	}
	return jr
}

func JSONToReviewTypeSlice(in json.RawMessage) []ReviewType {
	res := []int64{}
	err := json.Unmarshal(in, &res)
	if err != nil {
		fmt.Printf("error unmarshing ReviewType slice to int slice: %v", err)
		return nil
	}
	if len(res) == 0 {
		return nil
	}
	finalRes := []ReviewType{}
	for _, r := range res {
		finalRes = append(finalRes, ReviewType(r))
	}
	return finalRes
}
