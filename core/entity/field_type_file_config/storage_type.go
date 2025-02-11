package field_type_file_config

import (
	"encoding/json"
	"fmt"
)

//go:generate go run github.com/dmarkham/enumer -type=StorageType -json
type StorageType int64

const (
	STORAGE_TYPE_INVALID StorageType = iota
	STORAGE_TYPE_BINARY
	STORAGE_TYPE_OBJECT_STORE
)

func (e StorageType) ToInt64() int64 {
	return int64(e)
}

func StorageTypeFromString(in string) StorageType {
	switch in {
	case "binary":
		return STORAGE_TYPE_BINARY
	case "object_store":
		return STORAGE_TYPE_OBJECT_STORE
	}
	return STORAGE_TYPE_INVALID
}

func StorageTypeFromPointerString(in *string) StorageType {
	if in == nil {
		return STORAGE_TYPE_INVALID
	}
	return StorageTypeFromString(*in)
}

func (e StorageType) String() string {
	switch e {
	case STORAGE_TYPE_BINARY:
		return "binary"
	case STORAGE_TYPE_OBJECT_STORE:
		return "object_store"
	}

	return "invalid"
}

func (e StorageType) StringPtr() *string {
	val := e.String()
	return &val
}

func StorageTypeSliceToJSON(in []StorageType) json.RawMessage {
	res := make([]int64, len(in))
	for i, e := range in {
		res[i] = int64(e)
	}
	jr, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("error marshalling StorageType slice to json: %v", err)
		return json.RawMessage{}
	}
	return jr
}

func JSONToStorageTypeSlice(in json.RawMessage) []StorageType {
	res := []int64{}
	err := json.Unmarshal(in, &res)
	if err != nil {
		fmt.Printf("error unmarshing StorageType slice to int slice: %v", err)
		return nil
	}
	if len(res) == 0 {
		return nil
	}
	finalRes := []StorageType{}
	for _, r := range res {
		finalRes = append(finalRes, StorageType(r))
	}
	return finalRes
}
