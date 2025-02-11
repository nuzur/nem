package types

import (
	main_entity "github.com/nuzur/nem/core/entity/project_version"

	"go.uber.org/zap/zapcore"
)

type FetchProjectVersionByReviewStatusAndStatusRequest struct {
	ReviewStatus main_entity.ReviewStatus
	Status       main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchProjectVersionByReviewStatusAndStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchProjectVersionByReviewStatusAndStatusResponse struct {
	Results []main_entity.ProjectVersion
}
