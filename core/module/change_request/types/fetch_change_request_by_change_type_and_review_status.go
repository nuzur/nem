package types

import (
	main_entity "github.com/nuzur/nem/core/entity/change_request"

	"go.uber.org/zap/zapcore"
)

type FetchChangeRequestByChangeTypeAndReviewStatusRequest struct {
	ChangeType   main_entity.ChangeType
	ReviewStatus main_entity.ReviewStatus

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchChangeRequestByChangeTypeAndReviewStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchChangeRequestByChangeTypeAndReviewStatusResponse struct {
	Results []main_entity.ChangeRequest
}
