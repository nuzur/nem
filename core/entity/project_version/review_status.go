package project_version

import (
	"encoding/json"
	"fmt"
)

//go:generate go run github.com/dmarkham/enumer -type=ReviewStatus -json
type ReviewStatus int64

const (
	REVIEW_STATUS_INVALID ReviewStatus = iota
	REVIEW_STATUS_DRAFT
	REVIEW_STATUS_IN_REVIEW
	REVIEW_STATUS_APPROVED
	REVIEW_STATUS_REJECTED
	REVIEW_STATUS_PUBLISHED
)

func (e ReviewStatus) ToInt64() int64 {
	return int64(e)
}

func ReviewStatusFromString(in string) ReviewStatus {
	switch in {
	case "draft":
		return REVIEW_STATUS_DRAFT
	case "in_review":
		return REVIEW_STATUS_IN_REVIEW
	case "approved":
		return REVIEW_STATUS_APPROVED
	case "rejected":
		return REVIEW_STATUS_REJECTED
	case "published":
		return REVIEW_STATUS_PUBLISHED
	}
	return REVIEW_STATUS_INVALID
}

func ReviewStatusFromPointerString(in *string) ReviewStatus {
	if in == nil {
		return REVIEW_STATUS_INVALID
	}
	return ReviewStatusFromString(*in)
}

func (e ReviewStatus) String() string {
	switch e {
	case REVIEW_STATUS_DRAFT:
		return "draft"
	case REVIEW_STATUS_IN_REVIEW:
		return "in_review"
	case REVIEW_STATUS_APPROVED:
		return "approved"
	case REVIEW_STATUS_REJECTED:
		return "rejected"
	case REVIEW_STATUS_PUBLISHED:
		return "published"
	}

	return "invalid"
}

func (e ReviewStatus) StringPtr() *string {
	val := e.String()
	return &val
}

func ReviewStatusSliceToJSON(in []ReviewStatus) json.RawMessage {
	res := make([]int64, len(in))
	for i, e := range in {
		res[i] = int64(e)
	}
	jr, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("error marshalling ReviewStatus slice to json: %v", err)
		return json.RawMessage{}
	}
	return jr
}

func JSONToReviewStatusSlice(in json.RawMessage) []ReviewStatus {
	res := []int64{}
	err := json.Unmarshal(in, &res)
	if err != nil {
		fmt.Printf("error unmarshing ReviewStatus slice to int slice: %v", err)
		return nil
	}
	if len(res) == 0 {
		return nil
	}
	finalRes := []ReviewStatus{}
	for _, r := range res {
		finalRes = append(finalRes, ReviewStatus(r))
	}
	return finalRes
}
