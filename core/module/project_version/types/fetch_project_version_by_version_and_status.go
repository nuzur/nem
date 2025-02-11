package types

import (
	main_entity "nem/core/entity/project_version"

	"go.uber.org/zap/zapcore"
)

type FetchProjectVersionByVersionAndStatusRequest struct {
	Version int64
	Status  main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchProjectVersionByVersionAndStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchProjectVersionByVersionAndStatusResponse struct {
	Results []main_entity.ProjectVersion
}
