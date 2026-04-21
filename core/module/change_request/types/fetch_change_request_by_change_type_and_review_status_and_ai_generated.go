package types

import (
	main_entity "github.com/nuzur/nem/core/entity/change_request"

	"go.uber.org/zap/zapcore"
)

type FetchChangeRequestByChangeTypeAndReviewStatusAndAiGeneratedRequest struct {
	ChangeType   main_entity.ChangeType
	ReviewStatus main_entity.ReviewStatus
	AiGenerated  bool

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchChangeRequestByChangeTypeAndReviewStatusAndAiGeneratedRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchChangeRequestByChangeTypeAndReviewStatusAndAiGeneratedResponse struct {
	Results []main_entity.ChangeRequest
}
