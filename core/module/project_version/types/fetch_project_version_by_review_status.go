package types

import (
	main_entity "nem/core/entity/project_version"

	"go.uber.org/zap/zapcore"
)

type FetchProjectVersionByReviewStatusRequest struct {
	ReviewStatus main_entity.ReviewStatus

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchProjectVersionByReviewStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchProjectVersionByReviewStatusResponse struct {
	Results []main_entity.ProjectVersion
}
