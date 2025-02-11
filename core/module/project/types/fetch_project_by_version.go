package types

import (
	main_entity "github.com/nuzur/nem/core/entity/project"

	"go.uber.org/zap/zapcore"
)

type FetchProjectByVersionRequest struct {
	Version int64

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchProjectByVersionRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchProjectByVersionResponse struct {
	Results []main_entity.Project
}
