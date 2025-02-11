package user

import (
	"encoding/json"
	"fmt"
)

//go:generate go run github.com/dmarkham/enumer -type=UserType -json
type UserType int64

const (
	USER_TYPE_INVALID UserType = iota
	USER_TYPE_ADMIN
	USER_TYPE_REGULAR
)

func (e UserType) ToInt64() int64 {
	return int64(e)
}

func UserTypeFromString(in string) UserType {
	switch in {
	case "admin":
		return USER_TYPE_ADMIN
	case "regular":
		return USER_TYPE_REGULAR
	}
	return USER_TYPE_INVALID
}

func UserTypeFromPointerString(in *string) UserType {
	if in == nil {
		return USER_TYPE_INVALID
	}
	return UserTypeFromString(*in)
}

func (e UserType) String() string {
	switch e {
	case USER_TYPE_ADMIN:
		return "admin"
	case USER_TYPE_REGULAR:
		return "regular"
	}

	return "invalid"
}

func (e UserType) StringPtr() *string {
	val := e.String()
	return &val
}

func UserTypeSliceToJSON(in []UserType) json.RawMessage {
	res := make([]int64, len(in))
	for i, e := range in {
		res[i] = int64(e)
	}
	jr, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("error marshalling UserType slice to json: %v", err)
		return json.RawMessage{}
	}
	return jr
}

func JSONToUserTypeSlice(in json.RawMessage) []UserType {
	res := []int64{}
	err := json.Unmarshal(in, &res)
	if err != nil {
		fmt.Printf("error unmarshing UserType slice to int slice: %v", err)
		return nil
	}
	if len(res) == 0 {
		return nil
	}
	finalRes := []UserType{}
	for _, r := range res {
		finalRes = append(finalRes, UserType(r))
	}
	return finalRes
}
