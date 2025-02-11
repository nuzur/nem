package types

import (
	main_entity "github.com/nuzur/nem/core/entity/project"

	"go.uber.org/zap/zapcore"
)

type FetchProjectByVersionAndNameRequest struct {
	Version int64
	Name    string

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchProjectByVersionAndNameRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchProjectByVersionAndNameResponse struct {
	Results []main_entity.Project
}
