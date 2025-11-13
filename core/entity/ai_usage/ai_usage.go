package ai_usage

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"time"

	"fmt"

	"github.com/nuzur/nem/core/entity/types"

	"github.com/nuzur/nem/core/randomvalues"

	"math/rand"
)

type AiUsage struct {
	UUID               uuid.UUID `json:"uuid"`
	UserUUID           uuid.UUID `json:"user_uuid"`
	ProjectUUID        uuid.UUID `json:"project_uuid"`
	ProjectVersionUUID uuid.UUID `json:"project_version_uuid"`
	UserPrompt         string    `json:"user_prompt"`
	Step               string    `json:"step"`
	Context            Context   `json:"context"`
	Provider           Provider  `json:"provider"`
	InputTokens        int64     `json:"input_tokens"`
	OutputTokens       int64     `json:"output_tokens"`
	Status             Status    `json:"status"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	CreatedByUUID      uuid.UUID `json:"created_by_uuid"`
	UpdatedByUUID      uuid.UUID `json:"updated_by_uuid"`
}

func (e AiUsage) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e AiUsage) EntityIdentifier() string {
	return "ai_usage"
}

func (e AiUsage) PrimaryKeyIdentifier() string {
	return "uuid"
}

func (e AiUsage) PrimaryKeyValue() string {
	return e.UUID.String()
}

func (e AiUsage) FieldIdentfierToTypeMap() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	res["uuid"] = types.UUIDFieldType
	res["user_uuid"] = types.UUIDFieldType
	res["project_uuid"] = types.UUIDFieldType
	res["project_version_uuid"] = types.UUIDFieldType
	res["user_prompt"] = types.StringFieldType
	res["step"] = types.StringFieldType
	res["context"] = types.SingleEnumFieldType
	res["provider"] = types.SingleEnumFieldType
	res["input_tokens"] = types.IntFieldType
	res["output_tokens"] = types.IntFieldType
	res["status"] = types.SingleEnumFieldType
	res["created_at"] = types.TimestampFieldType
	res["updated_at"] = types.TimestampFieldType
	res["created_by_uuid"] = types.UUIDFieldType
	res["updated_by_uuid"] = types.UUIDFieldType
	return res
}

func (e AiUsage) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "uuid")
	res = append(res, "user_uuid")
	res = append(res, "project_uuid")
	res = append(res, "project_version_uuid")
	res = append(res, "user_prompt")
	res = append(res, "step")
	res = append(res, "context")
	res = append(res, "provider")
	res = append(res, "input_tokens")
	res = append(res, "output_tokens")
	res = append(res, "status")
	res = append(res, "created_at")
	res = append(res, "updated_at")
	res = append(res, "created_by_uuid")
	res = append(res, "updated_by_uuid")

	return res
}

func (e AiUsage) DependantFieldIdentifierToTypeMap() map[string]map[string]types.FieldType {
	res := make(map[string]map[string]types.FieldType)

	return res
}

func (e AiUsage) ArrayFieldIdentifierToType() map[string]types.FieldType {
	res := make(map[string]types.FieldType)

	return res
}

func (e AiUsage) IsDependant() bool {
	return false
}

func AiUsageFromJSON(data json.RawMessage) AiUsage {
	entity := AiUsage{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func AiUsageSliceFromJSON(data json.RawMessage) []AiUsage {
	entity := []AiUsage{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e AiUsage) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error AiUsageToJSON", err)
	}
	return res
}

func AiUsageToJSON(e AiUsage) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error AiUsageToJSON", err)
	}
	return res
}

func AiUsageSliceToJSON(e []AiUsage) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error AiUsageSliceToJSON", err)
	}
	return res
}

func NewAiUsageWithRandomValues() AiUsage {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	return AiUsage{
		UUID:               randomvalues.GetRandomUUIDValue(),
		UserUUID:           randomvalues.GetRandomUUIDValue(),
		ProjectUUID:        randomvalues.GetRandomUUIDValue(),
		ProjectVersionUUID: randomvalues.GetRandomUUIDValue(),
		UserPrompt:         randomvalues.GetRandomStringValue(),
		Step:               randomvalues.GetRandomStringValue(),
		Context:            randomvalues.GetRandomOptionValue[Context](3),
		Provider:           randomvalues.GetRandomOptionValue[Provider](1),
		InputTokens:        randomvalues.GetRandomIntValue(),
		OutputTokens:       randomvalues.GetRandomIntValue(),
		Status:             randomvalues.GetRandomOptionValue[Status](2),
		CreatedAt:          randomvalues.GetRandomTimeValue(),
		UpdatedAt:          randomvalues.GetRandomTimeValue(),
		CreatedByUUID:      randomvalues.GetRandomUUIDValue(),
		UpdatedByUUID:      randomvalues.GetRandomUUIDValue(),
	}
}

func NewAiUsageSliceWithRandomValues(n int) []AiUsage {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	numRecords := n
	res := []AiUsage{}
	for i := 0; i < numRecords; i++ {
		res = append(res, NewAiUsageWithRandomValues())
	}
	return res
}
