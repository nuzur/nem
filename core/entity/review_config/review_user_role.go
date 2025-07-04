package review_config

import (
	"encoding/json"
	"fmt"
)

//go:generate go run github.com/dmarkham/enumer -type=ReviewUserRole -json
type ReviewUserRole int64

const (
	REVIEW_USER_ROLE_INVALID ReviewUserRole = iota
	REVIEW_USER_ROLE_ADMIN
	REVIEW_USER_ROLE_DEVELOPER
	REVIEW_USER_ROLE_DATA_MANAGER
	REVIEW_USER_ROLE_DATA_ANALYST
	REVIEW_USER_ROLE_VIEWER
)

func (e ReviewUserRole) ToInt64() int64 {
	return int64(e)
}

func ReviewUserRoleFromString(in string) ReviewUserRole {
	switch in {
	case "admin":
		return REVIEW_USER_ROLE_ADMIN
	case "developer":
		return REVIEW_USER_ROLE_DEVELOPER
	case "data_manager":
		return REVIEW_USER_ROLE_DATA_MANAGER
	case "data_analyst":
		return REVIEW_USER_ROLE_DATA_ANALYST
	case "viewer":
		return REVIEW_USER_ROLE_VIEWER
	}
	return REVIEW_USER_ROLE_INVALID
}

func ReviewUserRoleFromPointerString(in *string) ReviewUserRole {
	if in == nil {
		return REVIEW_USER_ROLE_INVALID
	}
	return ReviewUserRoleFromString(*in)
}

func (e ReviewUserRole) String() string {
	switch e {
	case REVIEW_USER_ROLE_ADMIN:
		return "admin"
	case REVIEW_USER_ROLE_DEVELOPER:
		return "developer"
	case REVIEW_USER_ROLE_DATA_MANAGER:
		return "data_manager"
	case REVIEW_USER_ROLE_DATA_ANALYST:
		return "data_analyst"
	case REVIEW_USER_ROLE_VIEWER:
		return "viewer"
	}

	return "invalid"
}

func (e ReviewUserRole) StringPtr() *string {
	val := e.String()
	return &val
}

func ReviewUserRoleSliceToJSON(in []ReviewUserRole) json.RawMessage {
	res := make([]int64, len(in))
	for i, e := range in {
		res[i] = int64(e)
	}
	jr, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("error marshalling ReviewUserRole slice to json: %v", err)
		return json.RawMessage{}
	}
	return jr
}

func JSONToReviewUserRoleSlice(in json.RawMessage) []ReviewUserRole {
	res := []int64{}
	err := json.Unmarshal(in, &res)
	if err != nil {
		fmt.Printf("error unmarshing ReviewUserRole slice to int slice: %v", err)
		return nil
	}
	if len(res) == 0 {
		return nil
	}
	finalRes := []ReviewUserRole{}
	for _, r := range res {
		finalRes = append(finalRes, ReviewUserRole(r))
	}
	return finalRes
}
