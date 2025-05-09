package randomvalues

import (
	"github.com/gofrs/uuid"
	"golang.org/x/exp/slices"
	"math/rand"
	"time"
)

func GetRandomBoolValue() bool {
	return rand.Intn(2) != 0
}

func GetRandomTimeValue() time.Time {
	return time.Now()
}

func GetRandomFloatValue() float64 {
	return rand.Float64()
}

func GetRandomIntValue() int64 {
	return rand.Int63n(100)
}

func GetRandomOptionsValues[EnumType ~int64](max int) []EnumType {
	numberOfValues := rand.Intn(max + 1)
	res := []int{}
	for i := 0; i < numberOfValues; i++ {
		newVal := rand.Intn(max + 1)
		if !slices.Contains(res, newVal) {
			res = append(res, newVal)
		}
	}

	finalRes := []EnumType{}
	for r := range res {
		finalRes = append(finalRes, EnumType(r))
	}
	return finalRes
}

func GetRandomOptionValue[EnumType ~int64](max int) EnumType {
	return EnumType(rand.Intn(max + 1))
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GetRandomStringValue() string {
	b := make([]rune, 10)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GetRandomStringValuePtr() *string {
	str := GetRandomStringValue()
	return &str
}

func GetRandomUUIDValue() uuid.UUID {
	return uuid.Must(uuid.NewV4())
}

func GetRandomRawJSONValue() string {
	return "{}"
}

func GetRandomRawJSONValuePtr() *string {
	str := GetRandomRawJSONValue()
	return &str
}
