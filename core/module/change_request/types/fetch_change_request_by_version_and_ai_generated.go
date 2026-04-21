package types

import (
	main_entity "github.com/nuzur/nem/core/entity/change_request"

	"go.uber.org/zap/zapcore"
)

type FetchChangeRequestByVersionAndAiGeneratedRequest struct {
	Version     int64
	AiGenerated bool

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchChangeRequestByVersionAndAiGeneratedRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchChangeRequestByVersionAndAiGeneratedResponse struct {
	Results []main_entity.ChangeRequest
}
