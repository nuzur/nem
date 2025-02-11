package field_type_decimal_config

import (
	"encoding/json"
	"fmt"
)

//go:generate go run github.com/dmarkham/enumer -type=Separator -json
type Separator int64

const (
	SEPARATOR_INVALID Separator = iota
	SEPARATOR_POINT
	SEPARATOR_COMMA
)

func (e Separator) ToInt64() int64 {
	return int64(e)
}

func SeparatorFromString(in string) Separator {
	switch in {
	case "point":
		return SEPARATOR_POINT
	case "comma":
		return SEPARATOR_COMMA
	}
	return SEPARATOR_INVALID
}

func SeparatorFromPointerString(in *string) Separator {
	if in == nil {
		return SEPARATOR_INVALID
	}
	return SeparatorFromString(*in)
}

func (e Separator) String() string {
	switch e {
	case SEPARATOR_POINT:
		return "point"
	case SEPARATOR_COMMA:
		return "comma"
	}

	return "invalid"
}

func (e Separator) StringPtr() *string {
	val := e.String()
	return &val
}

func SeparatorSliceToJSON(in []Separator) json.RawMessage {
	res := make([]int64, len(in))
	for i, e := range in {
		res[i] = int64(e)
	}
	jr, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("error marshalling Separator slice to json: %v", err)
		return json.RawMessage{}
	}
	return jr
}

func JSONToSeparatorSlice(in json.RawMessage) []Separator {
	res := []int64{}
	err := json.Unmarshal(in, &res)
	if err != nil {
		fmt.Printf("error unmarshing Separator slice to int slice: %v", err)
		return nil
	}
	if len(res) == 0 {
		return nil
	}
	finalRes := []Separator{}
	for _, r := range res {
		finalRes = append(finalRes, Separator(r))
	}
	return finalRes
}
