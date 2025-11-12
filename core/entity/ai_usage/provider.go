package ai_usage

import (
	"encoding/json"
	"fmt"
)

//go:generate go run github.com/dmarkham/enumer -type=Provider -json
type Provider int64

const (
	PROVIDER_INVALID Provider = iota
	PROVIDER_CLAUDE
)

func (e Provider) ToInt64() int64 {
	return int64(e)
}

func ProviderFromString(in string) Provider {
	switch in {
	case "claude":
		return PROVIDER_CLAUDE
	}
	return PROVIDER_INVALID
}

func ProviderFromPointerString(in *string) Provider {
	if in == nil {
		return PROVIDER_INVALID
	}
	return ProviderFromString(*in)
}

func (e Provider) String() string {
	switch e {
	case PROVIDER_CLAUDE:
		return "claude"
	}

	return "invalid"
}

func (e Provider) StringPtr() *string {
	val := e.String()
	return &val
}

func ProviderSliceToJSON(in []Provider) json.RawMessage {
	res := make([]int64, len(in))
	for i, e := range in {
		res[i] = int64(e)
	}
	jr, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("error marshalling Provider slice to json: %v", err)
		return json.RawMessage{}
	}
	return jr
}

func JSONToProviderSlice(in json.RawMessage) []Provider {
	res := []int64{}
	err := json.Unmarshal(in, &res)
	if err != nil {
		fmt.Printf("error unmarshing Provider slice to int slice: %v", err)
		return nil
	}
	if len(res) == 0 {
		return nil
	}
	finalRes := []Provider{}
	for _, r := range res {
		finalRes = append(finalRes, Provider(r))
	}
	return finalRes
}
