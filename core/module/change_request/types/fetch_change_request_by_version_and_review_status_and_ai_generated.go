package types

import (
	main_entity "github.com/nuzur/nem/core/entity/change_request"

	"go.uber.org/zap/zapcore"
)

type FetchChangeRequestByVersionAndReviewStatusAndAiGeneratedRequest struct {
	Version      int64
	ReviewStatus main_entity.ReviewStatus
	AiGenerated  bool

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchChangeRequestByVersionAndReviewStatusAndAiGeneratedRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchChangeRequestByVersionAndReviewStatusAndAiGeneratedResponse struct {
	Results []main_entity.ChangeRequest
}
