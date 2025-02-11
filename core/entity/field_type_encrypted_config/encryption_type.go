package field_type_encrypted_config

import (
	"encoding/json"
	"fmt"
)

//go:generate go run github.com/dmarkham/enumer -type=EncryptionType -json
type EncryptionType int64

const (
	ENCRYPTION_TYPE_INVALID EncryptionType = iota
	ENCRYPTION_TYPE_BCRYPT
	ENCRYPTION_TYPE_SHA512
)

func (e EncryptionType) ToInt64() int64 {
	return int64(e)
}

func EncryptionTypeFromString(in string) EncryptionType {
	switch in {
	case "bcrypt":
		return ENCRYPTION_TYPE_BCRYPT
	case "sha512":
		return ENCRYPTION_TYPE_SHA512
	}
	return ENCRYPTION_TYPE_INVALID
}

func EncryptionTypeFromPointerString(in *string) EncryptionType {
	if in == nil {
		return ENCRYPTION_TYPE_INVALID
	}
	return EncryptionTypeFromString(*in)
}

func (e EncryptionType) String() string {
	switch e {
	case ENCRYPTION_TYPE_BCRYPT:
		return "bcrypt"
	case ENCRYPTION_TYPE_SHA512:
		return "sha512"
	}

	return "invalid"
}

func (e EncryptionType) StringPtr() *string {
	val := e.String()
	return &val
}

func EncryptionTypeSliceToJSON(in []EncryptionType) json.RawMessage {
	res := make([]int64, len(in))
	for i, e := range in {
		res[i] = int64(e)
	}
	jr, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("error marshalling EncryptionType slice to json: %v", err)
		return json.RawMessage{}
	}
	return jr
}

func JSONToEncryptionTypeSlice(in json.RawMessage) []EncryptionType {
	res := []int64{}
	err := json.Unmarshal(in, &res)
	if err != nil {
		fmt.Printf("error unmarshing EncryptionType slice to int slice: %v", err)
		return nil
	}
	if len(res) == 0 {
		return nil
	}
	finalRes := []EncryptionType{}
	for _, r := range res {
		finalRes = append(finalRes, EncryptionType(r))
	}
	return finalRes
}
