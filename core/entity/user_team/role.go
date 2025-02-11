package user_team

import (
	"encoding/json"
	"fmt"
)

//go:generate go run github.com/dmarkham/enumer -type=Role -json
type Role int64

const (
	ROLE_INVALID Role = iota
	ROLE_ADMIN
	ROLE_DEVELOPER
	ROLE_DATA_MANAGER
	ROLE_DATA_ANALYST
	ROLE_VIEWER
)

func (e Role) ToInt64() int64 {
	return int64(e)
}

func RoleFromString(in string) Role {
	switch in {
	case "admin":
		return ROLE_ADMIN
	case "developer":
		return ROLE_DEVELOPER
	case "data_manager":
		return ROLE_DATA_MANAGER
	case "data_analyst":
		return ROLE_DATA_ANALYST
	case "viewer":
		return ROLE_VIEWER
	}
	return ROLE_INVALID
}

func RoleFromPointerString(in *string) Role {
	if in == nil {
		return ROLE_INVALID
	}
	return RoleFromString(*in)
}

func (e Role) String() string {
	switch e {
	case ROLE_ADMIN:
		return "admin"
	case ROLE_DEVELOPER:
		return "developer"
	case ROLE_DATA_MANAGER:
		return "data_manager"
	case ROLE_DATA_ANALYST:
		return "data_analyst"
	case ROLE_VIEWER:
		return "viewer"
	}

	return "invalid"
}

func (e Role) StringPtr() *string {
	val := e.String()
	return &val
}

func RoleSliceToJSON(in []Role) json.RawMessage {
	res := make([]int64, len(in))
	for i, e := range in {
		res[i] = int64(e)
	}
	jr, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("error marshalling Role slice to json: %v", err)
		return json.RawMessage{}
	}
	return jr
}

func JSONToRoleSlice(in json.RawMessage) []Role {
	res := []int64{}
	err := json.Unmarshal(in, &res)
	if err != nil {
		fmt.Printf("error unmarshing Role slice to int slice: %v", err)
		return nil
	}
	if len(res) == 0 {
		return nil
	}
	finalRes := []Role{}
	for _, r := range res {
		finalRes = append(finalRes, Role(r))
	}
	return finalRes
}
