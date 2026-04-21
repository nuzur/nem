package types

import (
	main_entity "github.com/nuzur/nem/core/entity/change_request"

	"go.uber.org/zap/zapcore"
)

type FetchChangeRequestByReviewStatusAndAiGeneratedRequest struct {
	ReviewStatus main_entity.ReviewStatus
	AiGenerated  bool

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchChangeRequestByReviewStatusAndAiGeneratedRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchChangeRequestByReviewStatusAndAiGeneratedResponse struct {
	Results []main_entity.ChangeRequest
}
