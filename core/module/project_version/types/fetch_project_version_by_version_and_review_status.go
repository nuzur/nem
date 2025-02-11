package types

import (
	main_entity "nem/core/entity/project_version"

	"go.uber.org/zap/zapcore"
)

type FetchProjectVersionByVersionAndReviewStatusRequest struct {
	Version      int64
	ReviewStatus main_entity.ReviewStatus

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchProjectVersionByVersionAndReviewStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchProjectVersionByVersionAndReviewStatusResponse struct {
	Results []main_entity.ProjectVersion
}
