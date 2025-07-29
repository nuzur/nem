package types

import (
	main_entity "github.com/nuzur/nem/core/entity/change_request"

	"go.uber.org/zap/zapcore"
)

type FetchChangeRequestByVersionAndReviewStatusAndStatusRequest struct {
	Version      int64
	ReviewStatus main_entity.ReviewStatus
	Status       main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchChangeRequestByVersionAndReviewStatusAndStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchChangeRequestByVersionAndReviewStatusAndStatusResponse struct {
	Results []main_entity.ChangeRequest
}
