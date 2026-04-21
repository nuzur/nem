package types

import (
	main_entity "github.com/nuzur/nem/core/entity/change_request"

	"go.uber.org/zap/zapcore"
)

type FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndStatusAndAiGeneratedRequest struct {
	Version      int64
	ChangeType   main_entity.ChangeType
	ReviewStatus main_entity.ReviewStatus
	Status       main_entity.Status
	AiGenerated  bool

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndStatusAndAiGeneratedRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndStatusAndAiGeneratedResponse struct {
	Results []main_entity.ChangeRequest
}
