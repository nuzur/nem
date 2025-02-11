package review_config

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"time"

	"fmt"

	"nem/core/entity/types"

	"nem/core/randomvalues"

	"math/rand"
)

type ReviewConfig struct {
	UUID          uuid.UUID  `json:"uuid"`
	Types         []Type     `json:"types"`
	UserRoles     []UserRole `json:"user_roles"`
	MinReviews    int64      `json:"min_reviews"`
	Status        Status     `json:"status"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	CreatedByUUID uuid.UUID  `json:"created_by_uuid"`
	UpdatedByUUID uuid.UUID  `json:"updated_by_uuid"`
}

func (e ReviewConfig) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e ReviewConfig) EntityIdentifier() string {
	return "review_config"
}

func (e ReviewConfig) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["uuid"] = types.UUIDFieldType
	res["types"] = types.MultiEnumFieldType
	res["user_roles"] = types.MultiEnumFieldType
	res["min_reviews"] = types.IntFieldType
	res["status"] = types.SingleEnumFieldType
	res["created_at"] = types.TimestampFieldType
	res["updated_at"] = types.TimestampFieldType
	res["created_by_uuid"] = types.UUIDFieldType
	res["updated_by_uuid"] = types.UUIDFieldType
	return res
}

func (e ReviewConfig) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e ReviewConfig) IsDependant() bool {
	return true
}

func ReviewConfigFromJSON(data json.RawMessage) ReviewConfig {
	entity := ReviewConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func ReviewConfigSliceFromJSON(data json.RawMessage) []ReviewConfig {
	entity := []ReviewConfig{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e ReviewConfig) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ReviewConfigToJSON", err)
	}
	return res
}

func ReviewConfigToJSON(e ReviewConfig) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ReviewConfigToJSON", err)
	}
	return res
}

func ReviewConfigSliceToJSON(e []ReviewConfig) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ReviewConfigSliceToJSON", err)
	}
	return res
}

func NewReviewConfigWithRandomValues() ReviewConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return ReviewConfig{
		UUID:          randomvalues.GetRandomUUIDValue(),
		Types:         randomvalues.GetRandomOptionsValues[Type](3),
		UserRoles:     randomvalues.GetRandomOptionsValues[UserRole](5),
		MinReviews:    randomvalues.GetRandomIntValue(),
		Status:        randomvalues.GetRandomOptionValue[Status](2),
		CreatedAt:     randomvalues.GetRandomTimeValue(),
		UpdatedAt:     randomvalues.GetRandomTimeValue(),
		CreatedByUUID: randomvalues.GetRandomUUIDValue(),
		UpdatedByUUID: randomvalues.GetRandomUUIDValue(),
	}
}

func NewReviewConfigSliceWithRandomValues(n int) []ReviewConfig {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []ReviewConfig{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewReviewConfigWithRandomValues())
	}
	return res
}
