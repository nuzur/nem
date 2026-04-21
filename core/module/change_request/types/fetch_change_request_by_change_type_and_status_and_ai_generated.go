package types

import (
	main_entity "github.com/nuzur/nem/core/entity/change_request"

	"go.uber.org/zap/zapcore"
)

type FetchChangeRequestByChangeTypeAndStatusAndAiGeneratedRequest struct {
	ChangeType  main_entity.ChangeType
	Status      main_entity.Status
	AiGenerated bool

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchChangeRequestByChangeTypeAndStatusAndAiGeneratedRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchChangeRequestByChangeTypeAndStatusAndAiGeneratedResponse struct {
	Results []main_entity.ChangeRequest
}
