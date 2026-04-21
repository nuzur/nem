package types

import (
	main_entity "github.com/nuzur/nem/core/entity/change_request"

	"go.uber.org/zap/zapcore"
)

type FetchChangeRequestByVersionAndChangeTypeAndAiGeneratedRequest struct {
	Version     int64
	ChangeType  main_entity.ChangeType
	AiGenerated bool

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchChangeRequestByVersionAndChangeTypeAndAiGeneratedRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchChangeRequestByVersionAndChangeTypeAndAiGeneratedResponse struct {
	Results []main_entity.ChangeRequest
}
