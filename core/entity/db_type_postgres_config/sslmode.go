package db_type_postgres_config

import (
	"encoding/json"
	"fmt"
)

//go:generate go run github.com/dmarkham/enumer -type=Sslmode -json
type Sslmode int64

const (
	SSLMODE_INVALID Sslmode = iota
	SSLMODE_DISABLE
	SSLMODE_ALLOW
	SSLMODE_PREFER
	SSLMODE_REQUIRE
	SSLMODE_VERIFY_CA
	SSLMODE_VERIFY_FULL
)

func (e Sslmode) ToInt64() int64 {
	return int64(e)
}

func SslmodeFromString(in string) Sslmode {
	switch in {
	case "disable":
		return SSLMODE_DISABLE
	case "allow":
		return SSLMODE_ALLOW
	case "prefer":
		return SSLMODE_PREFER
	case "require":
		return SSLMODE_REQUIRE
	case "verify_ca":
		return SSLMODE_VERIFY_CA
	case "verify_full":
		return SSLMODE_VERIFY_FULL
	}
	return SSLMODE_INVALID
}

func SslmodeFromPointerString(in *string) Sslmode {
	if in == nil {
		return SSLMODE_INVALID
	}
	return SslmodeFromString(*in)
}

func (e Sslmode) String() string {
	switch e {
	case SSLMODE_DISABLE:
		return "disable"
	case SSLMODE_ALLOW:
		return "allow"
	case SSLMODE_PREFER:
		return "prefer"
	case SSLMODE_REQUIRE:
		return "require"
	case SSLMODE_VERIFY_CA:
		return "verify_ca"
	case SSLMODE_VERIFY_FULL:
		return "verify_full"
	}

	return "invalid"
}

func (e Sslmode) StringPtr() *string {
	val := e.String()
	return &val
}

func SslmodeSliceToJSON(in []Sslmode) json.RawMessage {
	res := make([]int64, len(in))
	for i, e := range in {
		res[i] = int64(e)
	}
	jr, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("error marshalling Sslmode slice to json: %v", err)
		return json.RawMessage{}
	}
	return jr
}

func JSONToSslmodeSlice(in json.RawMessage) []Sslmode {
	res := []int64{}
	err := json.Unmarshal(in, &res)
	if err != nil {
		fmt.Printf("error unmarshing Sslmode slice to int slice: %v", err)
		return nil
	}
	if len(res) == 0 {
		return nil
	}
	finalRes := []Sslmode{}
	for _, r := range res {
		finalRes = append(finalRes, Sslmode(r))
	}
	return finalRes
}
