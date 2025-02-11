package types

import (
	main_entity "github.com/nuzur/nem/core/entity/project_version"

	"go.uber.org/zap/zapcore"
)

type FetchProjectVersionByVersionRequest struct {
	Version int64

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchProjectVersionByVersionRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchProjectVersionByVersionResponse struct {
	Results []main_entity.ProjectVersion
}
