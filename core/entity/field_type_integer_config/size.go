package field_type_integer_config

import (
	"encoding/json"
	"fmt"
)

//go:generate go run github.com/dmarkham/enumer -type=Size -json
type Size int64

const (
	SIZE_INVALID Size = iota
	SIZE_ONE_BIT
	SIZE_EIGHT_BITS
	SIZE_SIXTEEN_BITS
	SIZE_TWENTY_FOUR_BITS
	SIZE_THIRTY_TWO_BITS
	SIZE_SIXTY_FOUR_BITS
)

func (e Size) ToInt64() int64 {
	return int64(e)
}

func SizeFromString(in string) Size {
	switch in {
	case "one_bit":
		return SIZE_ONE_BIT
	case "eight_bits":
		return SIZE_EIGHT_BITS
	case "sixteen_bits":
		return SIZE_SIXTEEN_BITS
	case "twenty_four_bits":
		return SIZE_TWENTY_FOUR_BITS
	case "thirty_two_bits":
		return SIZE_THIRTY_TWO_BITS
	case "sixty_four_bits":
		return SIZE_SIXTY_FOUR_BITS
	}
	return SIZE_INVALID
}

func SizeFromPointerString(in *string) Size {
	if in == nil {
		return SIZE_INVALID
	}
	return SizeFromString(*in)
}

func (e Size) String() string {
	switch e {
	case SIZE_ONE_BIT:
		return "one_bit"
	case SIZE_EIGHT_BITS:
		return "eight_bits"
	case SIZE_SIXTEEN_BITS:
		return "sixteen_bits"
	case SIZE_TWENTY_FOUR_BITS:
		return "twenty_four_bits"
	case SIZE_THIRTY_TWO_BITS:
		return "thirty_two_bits"
	case SIZE_SIXTY_FOUR_BITS:
		return "sixty_four_bits"
	}

	return "invalid"
}

func (e Size) StringPtr() *string {
	val := e.String()
	return &val
}

func SizeSliceToJSON(in []Size) json.RawMessage {
	res := make([]int64, len(in))
	for i, e := range in {
		res[i] = int64(e)
	}
	jr, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("error marshalling Size slice to json: %v", err)
		return json.RawMessage{}
	}
	return jr
}

func JSONToSizeSlice(in json.RawMessage) []Size {
	res := []int64{}
	err := json.Unmarshal(in, &res)
	if err != nil {
		fmt.Printf("error unmarshing Size slice to int slice: %v", err)
		return nil
	}
	if len(res) == 0 {
		return nil
	}
	finalRes := []Size{}
	for _, r := range res {
		finalRes = append(finalRes, Size(r))
	}
	return finalRes
}
