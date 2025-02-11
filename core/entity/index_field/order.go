package index_field

import (
	"encoding/json"
	"fmt"
)

//go:generate go run github.com/dmarkham/enumer -type=Order -json
type Order int64

const (
	ORDER_INVALID Order = iota
	ORDER_ASC
	ORDER_DESC
)

func (e Order) ToInt64() int64 {
	return int64(e)
}

func OrderFromString(in string) Order {
	switch in {
	case "asc":
		return ORDER_ASC
	case "desc":
		return ORDER_DESC
	}
	return ORDER_INVALID
}

func OrderFromPointerString(in *string) Order {
	if in == nil {
		return ORDER_INVALID
	}
	return OrderFromString(*in)
}

func (e Order) String() string {
	switch e {
	case ORDER_ASC:
		return "asc"
	case ORDER_DESC:
		return "desc"
	}

	return "invalid"
}

func (e Order) StringPtr() *string {
	val := e.String()
	return &val
}

func OrderSliceToJSON(in []Order) json.RawMessage {
	res := make([]int64, len(in))
	for i, e := range in {
		res[i] = int64(e)
	}
	jr, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("error marshalling Order slice to json: %v", err)
		return json.RawMessage{}
	}
	return jr
}

func JSONToOrderSlice(in json.RawMessage) []Order {
	res := []int64{}
	err := json.Unmarshal(in, &res)
	if err != nil {
		fmt.Printf("error unmarshing Order slice to int slice: %v", err)
		return nil
	}
	if len(res) == 0 {
		return nil
	}
	finalRes := []Order{}
	for _, r := range res {
		finalRes = append(finalRes, Order(r))
	}
	return finalRes
}
