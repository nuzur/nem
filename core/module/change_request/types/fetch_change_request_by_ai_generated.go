package types

import (
	main_entity "github.com/nuzur/nem/core/entity/change_request"

	"go.uber.org/zap/zapcore"
)

type FetchChangeRequestByAiGeneratedRequest struct {
	AiGenerated bool

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchChangeRequestByAiGeneratedRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchChangeRequestByAiGeneratedResponse struct {
	Results []main_entity.ChangeRequest
}
