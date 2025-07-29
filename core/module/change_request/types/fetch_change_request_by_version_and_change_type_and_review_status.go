package types

import (
	main_entity "github.com/nuzur/nem/core/entity/change_request"

	"go.uber.org/zap/zapcore"
)

type FetchChangeRequestByVersionAndChangeTypeAndReviewStatusRequest struct {
	Version      int64
	ChangeType   main_entity.ChangeType
	ReviewStatus main_entity.ReviewStatus

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchChangeRequestByVersionAndChangeTypeAndReviewStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchChangeRequestByVersionAndChangeTypeAndReviewStatusResponse struct {
	Results []main_entity.ChangeRequest
}
