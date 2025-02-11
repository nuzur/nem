package types

import (
	main_entity "github.com/nuzur/nem/core/entity/project"

	"go.uber.org/zap/zapcore"
)

type FetchProjectByVersionAndNameAndStatusRequest struct {
	Version int64
	Name    string
	Status  main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchProjectByVersionAndNameAndStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchProjectByVersionAndNameAndStatusResponse struct {
	Results []main_entity.Project
}
