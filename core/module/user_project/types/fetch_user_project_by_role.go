package types

import (
	main_entity "github.com/nuzur/nem/core/entity/user_project"

	"go.uber.org/zap/zapcore"
)

type FetchUserProjectByRoleRequest struct {
	Role main_entity.Role

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchUserProjectByRoleRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchUserProjectByRoleResponse struct {
	Results []main_entity.UserProject
}
