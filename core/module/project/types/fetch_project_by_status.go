package types

import (
	main_entity "nem/core/entity/project"

	"go.uber.org/zap/zapcore"
)

type FetchProjectByStatusRequest struct {
	Status main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchProjectByStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchProjectByStatusResponse struct {
	Results []main_entity.Project
}
