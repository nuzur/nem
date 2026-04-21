package types

import (
	main_entity "github.com/nuzur/nem/core/entity/change_request"

	"go.uber.org/zap/zapcore"
)

type FetchChangeRequestByVersionAndChangeTypeAndStatusAndAiGeneratedRequest struct {
	Version     int64
	ChangeType  main_entity.ChangeType
	Status      main_entity.Status
	AiGenerated bool

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchChangeRequestByVersionAndChangeTypeAndStatusAndAiGeneratedRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchChangeRequestByVersionAndChangeTypeAndStatusAndAiGeneratedResponse struct {
	Results []main_entity.ChangeRequest
}
