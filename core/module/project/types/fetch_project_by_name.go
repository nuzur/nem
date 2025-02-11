package types

import (
	main_entity "github.com/nuzur/nem/core/entity/project"

	"go.uber.org/zap/zapcore"
)

type FetchProjectByNameRequest struct {
	Name string

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchProjectByNameRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchProjectByNameResponse struct {
	Results []main_entity.Project
}
