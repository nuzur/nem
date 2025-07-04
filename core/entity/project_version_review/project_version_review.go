package project_version_review

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"time"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"
)

type ProjectVersionReview struct {
	UUID      uuid.UUID `json:"uuid"`
	UserUUID  uuid.UUID `json:"user_uuid"`
	Comment   *string   `json:"comment"`
	Response  Response  `json:"response"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (e ProjectVersionReview) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e ProjectVersionReview) EntityIdentifier() string {
	return "project_version_review"
}

func (e ProjectVersionReview) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["uuid"] = types.UUIDFieldType
	res["user_uuid"] = types.UUIDFieldType
	res["comment"] = types.StringFieldType
	res["response"] = types.SingleEnumFieldType
	res["created_at"] = types.TimestampFieldType
	res["updated_at"] = types.TimestampFieldType
	return res
}

func (e ProjectVersionReview) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e ProjectVersionReview) IsDependant() bool {
	return true
}

func ProjectVersionReviewFromJSON(data json.RawMessage) ProjectVersionReview {
	entity := ProjectVersionReview{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func ProjectVersionReviewSliceFromJSON(data json.RawMessage) []ProjectVersionReview {
	entity := []ProjectVersionReview{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e ProjectVersionReview) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ProjectVersionReviewToJSON", err)
	}
	return res
}

func ProjectVersionReviewToJSON(e ProjectVersionReview) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ProjectVersionReviewToJSON", err)
	}
	return res
}

func ProjectVersionReviewSliceToJSON(e []ProjectVersionReview) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ProjectVersionReviewSliceToJSON", err)
	}
	return res
}

func NewProjectVersionReviewWithRandomValues() ProjectVersionReview {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return ProjectVersionReview{
		UUID:      randomvalues.GetRandomUUIDValue(),
		UserUUID:  randomvalues.GetRandomUUIDValue(),
		Comment:   randomvalues.GetRandomStringValuePtr(),
		Response:  randomvalues.GetRandomOptionValue[Response](2),
		CreatedAt: randomvalues.GetRandomTimeValue(),
		UpdatedAt: randomvalues.GetRandomTimeValue(),
	}
}

func NewProjectVersionReviewSliceWithRandomValues(n int) []ProjectVersionReview {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []ProjectVersionReview{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewProjectVersionReviewWithRandomValues())
	}
	return res
}
