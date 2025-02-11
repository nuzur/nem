package relationship

import (
	"encoding/json"
	"fmt"
)

//go:generate go run github.com/dmarkham/enumer -type=Cardinality -json
type Cardinality int64

const (
	CARDINALITY_INVALID Cardinality = iota
	CARDINALITY_ONE_TO_ONE
	CARDINALITY_ONE_TO_MANY
)

func (e Cardinality) ToInt64() int64 {
	return int64(e)
}

func CardinalityFromString(in string) Cardinality {
	switch in {
	case "one_to_one":
		return CARDINALITY_ONE_TO_ONE
	case "one_to_many":
		return CARDINALITY_ONE_TO_MANY
	}
	return CARDINALITY_INVALID
}

func CardinalityFromPointerString(in *string) Cardinality {
	if in == nil {
		return CARDINALITY_INVALID
	}
	return CardinalityFromString(*in)
}

func (e Cardinality) String() string {
	switch e {
	case CARDINALITY_ONE_TO_ONE:
		return "one_to_one"
	case CARDINALITY_ONE_TO_MANY:
		return "one_to_many"
	}

	return "invalid"
}

func (e Cardinality) StringPtr() *string {
	val := e.String()
	return &val
}

func CardinalitySliceToJSON(in []Cardinality) json.RawMessage {
	res := make([]int64, len(in))
	for i, e := range in {
		res[i] = int64(e)
	}
	jr, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("error marshalling Cardinality slice to json: %v", err)
		return json.RawMessage{}
	}
	return jr
}

func JSONToCardinalitySlice(in json.RawMessage) []Cardinality {
	res := []int64{}
	err := json.Unmarshal(in, &res)
	if err != nil {
		fmt.Printf("error unmarshing Cardinality slice to int slice: %v", err)
		return nil
	}
	if len(res) == 0 {
		return nil
	}
	finalRes := []Cardinality{}
	for _, r := range res {
		finalRes = append(finalRes, Cardinality(r))
	}
	return finalRes
}
