package extension

import (
	"encoding/json"
	"fmt"
)

//go:generate go run github.com/dmarkham/enumer -type=ExtensionType -json
type ExtensionType int64

const (
	EXTENSION_TYPE_INVALID ExtensionType = iota
	EXTENSION_TYPE_IMPORTER
	EXTENSION_TYPE_SYNCHRONIZER
	EXTENSION_TYPE_GENERATOR
)

func (e ExtensionType) ToInt64() int64 {
	return int64(e)
}

func ExtensionTypeFromString(in string) ExtensionType {
	switch in {
	case "importer":
		return EXTENSION_TYPE_IMPORTER
	case "synchronizer":
		return EXTENSION_TYPE_SYNCHRONIZER
	case "generator":
		return EXTENSION_TYPE_GENERATOR
	}
	return EXTENSION_TYPE_INVALID
}

func ExtensionTypeFromPointerString(in *string) ExtensionType {
	if in == nil {
		return EXTENSION_TYPE_INVALID
	}
	return ExtensionTypeFromString(*in)
}

func (e ExtensionType) String() string {
	switch e {
	case EXTENSION_TYPE_IMPORTER:
		return "importer"
	case EXTENSION_TYPE_SYNCHRONIZER:
		return "synchronizer"
	case EXTENSION_TYPE_GENERATOR:
		return "generator"
	}

	return "invalid"
}

func (e ExtensionType) StringPtr() *string {
	val := e.String()
	return &val
}

func ExtensionTypeSliceToJSON(in []ExtensionType) json.RawMessage {
	res := make([]int64, len(in))
	for i, e := range in {
		res[i] = int64(e)
	}
	jr, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("error marshalling ExtensionType slice to json: %v", err)
		return json.RawMessage{}
	}
	return jr
}

func JSONToExtensionTypeSlice(in json.RawMessage) []ExtensionType {
	res := []int64{}
	err := json.Unmarshal(in, &res)
	if err != nil {
		fmt.Printf("error unmarshing ExtensionType slice to int slice: %v", err)
		return nil
	}
	if len(res) == 0 {
		return nil
	}
	finalRes := []ExtensionType{}
	for _, r := range res {
		finalRes = append(finalRes, ExtensionType(r))
	}
	return finalRes
}
