package change_request_review

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"time"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"
)

type ChangeRequestReview struct {
	UUID      uuid.UUID `json:"uuid"`
	UserUUID  uuid.UUID `json:"user_uuid"`
	Comment   *string   `json:"comment"`
	Response  Response  `json:"response"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (e ChangeRequestReview) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e ChangeRequestReview) EntityIdentifier() string {
	return "change_request_review"
}

func (e ChangeRequestReview) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["uuid"] = types.UUIDFieldType
	res["user_uuid"] = types.UUIDFieldType
	res["comment"] = types.StringFieldType
	res["response"] = types.SingleEnumFieldType
	res["created_at"] = types.TimestampFieldType
	res["updated_at"] = types.TimestampFieldType
	return res
}

func (e ChangeRequestReview) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e ChangeRequestReview) IsDependant() bool {
	return true
}

func ChangeRequestReviewFromJSON(data json.RawMessage) ChangeRequestReview {
	entity := ChangeRequestReview{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func ChangeRequestReviewSliceFromJSON(data json.RawMessage) []ChangeRequestReview {
	entity := []ChangeRequestReview{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e ChangeRequestReview) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ChangeRequestReviewToJSON", err)
	}
	return res
}

func ChangeRequestReviewToJSON(e ChangeRequestReview) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ChangeRequestReviewToJSON", err)
	}
	return res
}

func ChangeRequestReviewSliceToJSON(e []ChangeRequestReview) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ChangeRequestReviewSliceToJSON", err)
	}
	return res
}

func NewChangeRequestReviewWithRandomValues() ChangeRequestReview {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return ChangeRequestReview{
		UUID:      randomvalues.GetRandomUUIDValue(),
		UserUUID:  randomvalues.GetRandomUUIDValue(),
		Comment:   randomvalues.GetRandomStringValuePtr(),
		Response:  randomvalues.GetRandomOptionValue[Response](2),
		CreatedAt: randomvalues.GetRandomTimeValue(),
		UpdatedAt: randomvalues.GetRandomTimeValue(),
	}
}

func NewChangeRequestReviewSliceWithRandomValues(n int) []ChangeRequestReview {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []ChangeRequestReview{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewChangeRequestReviewWithRandomValues())
	}
	return res
}
