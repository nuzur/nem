package review_config

import (
	"encoding/json"
	"fmt"
)

//go:generate go run github.com/dmarkham/enumer -type=UserRole -json
type UserRole int64

const (
	USER_ROLE_INVALID UserRole = iota
	USER_ROLE_ADMIN
	USER_ROLE_DEVELOPER
	USER_ROLE_DATA_MANAGER
	USER_ROLE_DATA_ANALYST
	USER_ROLE_VIEWER
)

func (e UserRole) ToInt64() int64 {
	return int64(e)
}

func UserRoleFromString(in string) UserRole {
	switch in {
	case "admin":
		return USER_ROLE_ADMIN
	case "developer":
		return USER_ROLE_DEVELOPER
	case "data_manager":
		return USER_ROLE_DATA_MANAGER
	case "data_analyst":
		return USER_ROLE_DATA_ANALYST
	case "viewer":
		return USER_ROLE_VIEWER
	}
	return USER_ROLE_INVALID
}

func UserRoleFromPointerString(in *string) UserRole {
	if in == nil {
		return USER_ROLE_INVALID
	}
	return UserRoleFromString(*in)
}

func (e UserRole) String() string {
	switch e {
	case USER_ROLE_ADMIN:
		return "admin"
	case USER_ROLE_DEVELOPER:
		return "developer"
	case USER_ROLE_DATA_MANAGER:
		return "data_manager"
	case USER_ROLE_DATA_ANALYST:
		return "data_analyst"
	case USER_ROLE_VIEWER:
		return "viewer"
	}

	return "invalid"
}

func (e UserRole) StringPtr() *string {
	val := e.String()
	return &val
}

func UserRoleSliceToJSON(in []UserRole) json.RawMessage {
	res := make([]int64, len(in))
	for i, e := range in {
		res[i] = int64(e)
	}
	jr, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("error marshalling UserRole slice to json: %v", err)
		return json.RawMessage{}
	}
	return jr
}

func JSONToUserRoleSlice(in json.RawMessage) []UserRole {
	res := []int64{}
	err := json.Unmarshal(in, &res)
	if err != nil {
		fmt.Printf("error unmarshing UserRole slice to int slice: %v", err)
		return nil
	}
	if len(res) == 0 {
		return nil
	}
	finalRes := []UserRole{}
	for _, r := range res {
		finalRes = append(finalRes, UserRole(r))
	}
	return finalRes
}
