package types

import (
	main_entity "github.com/nuzur/nem/core/entity/change_request"

	"go.uber.org/zap/zapcore"
)

type FetchChangeRequestByChangeTypeAndReviewStatusAndStatusAndAiGeneratedRequest struct {
	ChangeType   main_entity.ChangeType
	ReviewStatus main_entity.ReviewStatus
	Status       main_entity.Status
	AiGenerated  bool

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchChangeRequestByChangeTypeAndReviewStatusAndStatusAndAiGeneratedRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchChangeRequestByChangeTypeAndReviewStatusAndStatusAndAiGeneratedResponse struct {
	Results []main_entity.ChangeRequest
}
