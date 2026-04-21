package types

import (
	main_entity "github.com/nuzur/nem/core/entity/change_request"

	"go.uber.org/zap/zapcore"
)

type FetchChangeRequestByChangeTypeAndAiGeneratedRequest struct {
	ChangeType  main_entity.ChangeType
	AiGenerated bool

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchChangeRequestByChangeTypeAndAiGeneratedRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchChangeRequestByChangeTypeAndAiGeneratedResponse struct {
	Results []main_entity.ChangeRequest
}
