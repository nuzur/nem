package types

import (
	main_entity "github.com/nuzur/nem/core/entity/change_request"

	"go.uber.org/zap/zapcore"
)

type FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndAiGeneratedRequest struct {
	Version      int64
	ChangeType   main_entity.ChangeType
	ReviewStatus main_entity.ReviewStatus
	AiGenerated  bool

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndAiGeneratedRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndAiGeneratedResponse struct {
	Results []main_entity.ChangeRequest
}
