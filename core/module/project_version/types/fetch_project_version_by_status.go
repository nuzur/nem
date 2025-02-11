package types

import (
	main_entity "nem/core/entity/project_version"

	"go.uber.org/zap/zapcore"
)

type FetchProjectVersionByStatusRequest struct {
	Status main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchProjectVersionByStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchProjectVersionByStatusResponse struct {
	Results []main_entity.ProjectVersion
}
