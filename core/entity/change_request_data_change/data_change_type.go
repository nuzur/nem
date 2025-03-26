package change_request_data_change

import (
	"encoding/json"
	"fmt"
)

//go:generate go run github.com/dmarkham/enumer -type=DataChangeType -json
type DataChangeType int64

const (
	DATA_CHANGE_TYPE_INVALID DataChangeType = iota
	DATA_CHANGE_TYPE_UPDATE_FIELD
	DATA_CHANGE_TYPE_CREATE_RECORD
	DATA_CHANGE_TYPE_DELETE_RECORD
)

func (e DataChangeType) ToInt64() int64 {
	return int64(e)
}

func DataChangeTypeFromString(in string) DataChangeType {
	switch in {
	case "update_field":
		return DATA_CHANGE_TYPE_UPDATE_FIELD
	case "create_record":
		return DATA_CHANGE_TYPE_CREATE_RECORD
	case "delete_record":
		return DATA_CHANGE_TYPE_DELETE_RECORD
	}
	return DATA_CHANGE_TYPE_INVALID
}

func DataChangeTypeFromPointerString(in *string) DataChangeType {
	if in == nil {
		return DATA_CHANGE_TYPE_INVALID
	}
	return DataChangeTypeFromString(*in)
}

func (e DataChangeType) String() string {
	switch e {
	case DATA_CHANGE_TYPE_UPDATE_FIELD:
		return "update_field"
	case DATA_CHANGE_TYPE_CREATE_RECORD:
		return "create_record"
	case DATA_CHANGE_TYPE_DELETE_RECORD:
		return "delete_record"
	}

	return "invalid"
}

func (e DataChangeType) StringPtr() *string {
	val := e.String()
	return &val
}

func DataChangeTypeSliceToJSON(in []DataChangeType) json.RawMessage {
	res := make([]int64, len(in))
	for i, e := range in {
		res[i] = int64(e)
	}
	jr, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("error marshalling DataChangeType slice to json: %v", err)
		return json.RawMessage{}
	}
	return jr
}

func JSONToDataChangeTypeSlice(in json.RawMessage) []DataChangeType {
	res := []int64{}
	err := json.Unmarshal(in, &res)
	if err != nil {
		fmt.Printf("error unmarshing DataChangeType slice to int slice: %v", err)
		return nil
	}
	if len(res) == 0 {
		return nil
	}
	finalRes := []DataChangeType{}
	for _, r := range res {
		finalRes = append(finalRes, DataChangeType(r))
	}
	return finalRes
}
