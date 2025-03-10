package mapper

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
)

func JSONToStringSlice(data json.RawMessage) []string {
	res := []string{}
	err := json.Unmarshal(data, &res)
	if err != nil {
		fmt.Printf("error unmarshaling json to string slice: %v \n", err)
		return []string{}
	}
	return res
}

func JSONToUUIDSlice(data json.RawMessage) []uuid.UUID {
	stringSlice := JSONToStringSlice(data)

	res := []uuid.UUID{}
	for _, s := range stringSlice {
		res = append(res, uuid.FromStringOrNil(s))
	}
	return res
}

func JSONToIntSlice(data json.RawMessage) []int64 {
	res := []int64{}
	err := json.Unmarshal(data, &res)
	if err != nil {
		fmt.Printf("error unmarshaling json to int slice: %v \n", err)
		return []int64{}
	}
	return res
}

func JSONToFloatSlice(data json.RawMessage) []float64 {
	res := []float64{}
	err := json.Unmarshal(data, &res)
	if err != nil {
		fmt.Printf("error unmarshaling json to float slice: %v \n", err)
		return []float64{}
	}
	return res
}

func JSONToBooleanSlice(data json.RawMessage) []bool {
	res := []bool{}
	err := json.Unmarshal(data, &res)
	if err != nil {
		fmt.Printf("error unmarshaling json to boolean slice: %v \n", err)
		return []bool{}
	}
	return res
}

func JSONToDateSlice(data json.RawMessage) []time.Time {
	res := []time.Time{}
	err := json.Unmarshal(data, &res)
	if err != nil {
		fmt.Printf("error unmarshaling json to date slice: %v \n", err)
		return []time.Time{}
	}
	return res
}

func JSONToDatetimeSlice(data json.RawMessage) []time.Time {
	res := []time.Time{}
	err := json.Unmarshal(data, &res)
	if err != nil {
		fmt.Printf("error unmarshaling json to datetime slice: %v \n", err)
		return []time.Time{}
	}
	return res
}

func SliceToJSON(slice interface{}) json.RawMessage {
	res, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("error marshaling slice to json: %v \n", err)
		return json.RawMessage{}
	}
	return res
}
