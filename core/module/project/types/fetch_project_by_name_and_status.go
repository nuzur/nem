package types

import (
	main_entity "nem/core/entity/project"

	"go.uber.org/zap/zapcore"
)

type FetchProjectByNameAndStatusRequest struct {
	Name   string
	Status main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchProjectByNameAndStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchProjectByNameAndStatusResponse struct {
	Results []main_entity.Project
}
