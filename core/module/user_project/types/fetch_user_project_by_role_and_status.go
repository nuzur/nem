package types

import (
	main_entity "github.com/nuzur/nem/core/entity/user_project"

	"go.uber.org/zap/zapcore"
)

type FetchUserProjectByRoleAndStatusRequest struct {
	Role   main_entity.Role
	Status main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchUserProjectByRoleAndStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchUserProjectByRoleAndStatusResponse struct {
	Results []main_entity.UserProject
}
