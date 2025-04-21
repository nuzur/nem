package types

import (
	main_entity "github.com/nuzur/nem/core/entity/user_project_version"

	"go.uber.org/zap/zapcore"
)

type FetchUserProjectVersionByStatusRequest struct {
	Status main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchUserProjectVersionByStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchUserProjectVersionByStatusResponse struct {
	Results []main_entity.UserProjectVersion
}
