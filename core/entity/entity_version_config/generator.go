package entity_version_config

import (
	"encoding/json"
	"fmt"
)

//go:generate go run github.com/dmarkham/enumer -type=Generator -json
type Generator int64

const (
	GENERATOR_INVALID Generator = iota
	GENERATOR_TIMESTAMP
)

func (e Generator) ToInt64() int64 {
	return int64(e)
}

func GeneratorFromString(in string) Generator {
	switch in {
	case "timestamp":
		return GENERATOR_TIMESTAMP
	}
	return GENERATOR_INVALID
}

func GeneratorFromPointerString(in *string) Generator {
	if in == nil {
		return GENERATOR_INVALID
	}
	return GeneratorFromString(*in)
}

func (e Generator) String() string {
	switch e {
	case GENERATOR_TIMESTAMP:
		return "timestamp"
	}

	return "invalid"
}

func (e Generator) StringPtr() *string {
	val := e.String()
	return &val
}

func GeneratorSliceToJSON(in []Generator) json.RawMessage {
	res := make([]int64, len(in))
	for i, e := range in {
		res[i] = int64(e)
	}
	jr, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("error marshalling Generator slice to json: %v", err)
		return json.RawMessage{}
	}
	return jr
}

func JSONToGeneratorSlice(in json.RawMessage) []Generator {
	res := []int64{}
	err := json.Unmarshal(in, &res)
	if err != nil {
		fmt.Printf("error unmarshing Generator slice to int slice: %v", err)
		return nil
	}
	if len(res) == 0 {
		return nil
	}
	finalRes := []Generator{}
	for _, r := range res {
		finalRes = append(finalRes, Generator(r))
	}
	return finalRes
}
