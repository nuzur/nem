package types

import (
	main_entity "github.com/nuzur/nem/core/entity/change_request"

	"go.uber.org/zap/zapcore"
)

type FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndStatusRequest struct {
	Version      int64
	ChangeType   main_entity.ChangeType
	ReviewStatus main_entity.ReviewStatus
	Status       main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndStatusResponse struct {
	Results []main_entity.ChangeRequest
}
