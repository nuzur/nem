package change_request

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"github.com/nuzur/nem/core/entity/change_request_data_change"
	"github.com/nuzur/nem/core/entity/change_request_review"
	"time"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"
)

type ChangeRequest struct {
	UUID               uuid.UUID                                            `json:"uuid"`
	Version            int64                                                `json:"version"`
	Title              string                                               `json:"title"`
	Description        *string                                              `json:"description"`
	ProjectUUID        uuid.UUID                                            `json:"project_uuid"`
	ProjectVersionUUID uuid.UUID                                            `json:"project_version_uuid"`
	ChangeType         ChangeType                                           `json:"change_type"`
	DataChanges        []change_request_data_change.ChangeRequestDataChange `json:"data_changes"`
	VersionChanges     *string                                              `json:"version_changes"`
	Reviews            []change_request_review.ChangeRequestReview          `json:"reviews"`
	ReviewStatus       ReviewStatus                                         `json:"review_status"`
	OwnerUUID          uuid.UUID                                            `json:"owner_uuid"`
	Status             Status                                               `json:"status"`
	CreatedAt          time.Time                                            `json:"created_at"`
	UpdatedAt          time.Time                                            `json:"updated_at"`
	CreatedByUUID      uuid.UUID                                            `json:"created_by_uuid"`
	UpdatedByUUID      uuid.UUID                                            `json:"updated_by_uuid"`
}

func (e ChangeRequest) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e ChangeRequest) EntityIdentifier() string {
	return "change_request"
}

func (e ChangeRequest) PrimaryKeyIdentifier() string {
	return "uuid"
}

func (e ChangeRequest) PrimaryKeyValue() string {
	return e.UUID.String()
}

func (e ChangeRequest) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["uuid"] = types.UUIDFieldType
	res["version"] = types.IntFieldType
	res["title"] = types.StringFieldType
	res["description"] = types.StringFieldType
	res["project_uuid"] = types.UUIDFieldType
	res["project_version_uuid"] = types.UUIDFieldType
	res["change_type"] = types.SingleEnumFieldType
	res["data_changes"] = types.MultiDependantEntityFieldType
	res["version_changes"] = types.RawJSONFieldType
	res["reviews"] = types.MultiDependantEntityFieldType
	res["review_status"] = types.SingleEnumFieldType
	res["owner_uuid"] = types.UUIDFieldType
	res["status"] = types.SingleEnumFieldType
	res["created_at"] = types.TimestampFieldType
	res["updated_at"] = types.TimestampFieldType
	res["created_by_uuid"] = types.UUIDFieldType
	res["updated_by_uuid"] = types.UUIDFieldType
	return res
}

func (e ChangeRequest) DependantFieldIdentifierToTypeMap() map[string]map[string]types.FieldType {
	res := make(map[string]map[string]types.FieldType)

	res["data_changes"] = change_request_data_change.ChangeRequestDataChange{}.FieldIdentfierToTypeMap()
	res["reviews"] = change_request_review.ChangeRequestReview{}.FieldIdentfierToTypeMap()
	return res
}

func (e ChangeRequest) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e ChangeRequest) IsDependant() bool {
	return false
}

func ChangeRequestFromJSON(data json.RawMessage) ChangeRequest {
	entity := ChangeRequest{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func ChangeRequestSliceFromJSON(data json.RawMessage) []ChangeRequest {
	entity := []ChangeRequest{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e ChangeRequest) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ChangeRequestToJSON", err)
	}
	return res
}

func ChangeRequestToJSON(e ChangeRequest) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ChangeRequestToJSON", err)
	}
	return res
}

func ChangeRequestSliceToJSON(e []ChangeRequest) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error ChangeRequestSliceToJSON", err)
	}
	return res
}

func NewChangeRequestWithRandomValues() ChangeRequest {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return ChangeRequest{
		UUID:               randomvalues.GetRandomUUIDValue(),
		Version:            randomvalues.GetRandomIntValue(),
		Title:              randomvalues.GetRandomStringValue(),
		Description:        randomvalues.GetRandomStringValuePtr(),
		ProjectUUID:        randomvalues.GetRandomUUIDValue(),
		ProjectVersionUUID: randomvalues.GetRandomUUIDValue(),
		ChangeType:         randomvalues.GetRandomOptionValue[ChangeType](2),
		DataChanges:        change_request_data_change.NewChangeRequestDataChangeSliceWithRandomValues(rand.Intn(10)),
		VersionChanges:     randomvalues.GetRandomRawJSONValuePtr(),
		Reviews:            change_request_review.NewChangeRequestReviewSliceWithRandomValues(rand.Intn(10)),
		ReviewStatus:       randomvalues.GetRandomOptionValue[ReviewStatus](6),
		OwnerUUID:          randomvalues.GetRandomUUIDValue(),
		Status:             randomvalues.GetRandomOptionValue[Status](2),
		CreatedAt:          randomvalues.GetRandomTimeValue(),
		UpdatedAt:          randomvalues.GetRandomTimeValue(),
		CreatedByUUID:      randomvalues.GetRandomUUIDValue(),
		UpdatedByUUID:      randomvalues.GetRandomUUIDValue(),
	}
}

func NewChangeRequestSliceWithRandomValues(n int) []ChangeRequest {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []ChangeRequest{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewChangeRequestWithRandomValues())
	}
	return res
}
