package types

import (
	main_entity "github.com/nuzur/nem/core/entity/project"

	"go.uber.org/zap/zapcore"
)

type FetchProjectByVersionAndStatusRequest struct {
	Version int64
	Status  main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchProjectByVersionAndStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchProjectByVersionAndStatusResponse struct {
	Results []main_entity.Project
}
