package change_request

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

//go:generate go run github.com/dmarkham/enumer -type=Scope -json
type Scope int64

const (
	SCOPE_INVALID Scope = iota
	SCOPE_LOCAL
	SCOPE_REMOTE
)

func (e Scope) ToInt64() int64 {
	return int64(e)
}

func (e Scope) ToSqlNullInt32() sql.NullInt32 {
	return sql.NullInt32{
		Int32: int32(e),
		Valid: true,
	}
}

func ScopeFromString(in string) Scope {
	switch in {
	case "local":
		return SCOPE_LOCAL
	case "remote":
		return SCOPE_REMOTE
	}
	return SCOPE_INVALID
}

func ScopeFromPointerString(in *string) Scope {
	if in == nil {
		return SCOPE_INVALID
	}
	return ScopeFromString(*in)
}

func (e Scope) String() string {
	switch e {
	case SCOPE_LOCAL:
		return "local"
	case SCOPE_REMOTE:
		return "remote"
	}

	return "invalid"
}

func (e Scope) StringPtr() *string {
	val := e.String()
	return &val
}

func ScopeSliceToJSON(in []Scope) json.RawMessage {
	res := make([]int64, len(in))
	for i, e := range in {
		res[i] = int64(e)
	}
	jr, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("error marshalling Scope slice to json: %v", err)
		return json.RawMessage{}
	}
	return jr
}

func JSONToScopeSlice(in json.RawMessage) []Scope {
	res := []int64{}
	err := json.Unmarshal(in, &res)
	if err != nil {
		fmt.Printf("error unmarshing Scope slice to int slice: %v", err)
		return nil
	}
	if len(res) == 0 {
		return nil
	}
	finalRes := []Scope{}
	for _, r := range res {
		finalRes = append(finalRes, Scope(r))
	}
	return finalRes
}
