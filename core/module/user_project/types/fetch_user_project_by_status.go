package types

import (
	main_entity "github.com/nuzur/nem/core/entity/user_project"

	"go.uber.org/zap/zapcore"
)

type FetchUserProjectByStatusRequest struct {
	Status main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchUserProjectByStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchUserProjectByStatusResponse struct {
	Results []main_entity.UserProject
}
