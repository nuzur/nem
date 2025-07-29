package types

import (
	main_entity "github.com/nuzur/nem/core/entity/user_project"

	"go.uber.org/zap/zapcore"
)

type FetchUserProjectByUserEmailAndRoleRequest struct {
	UserEmail *string
	Role      main_entity.Role

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchUserProjectByUserEmailAndRoleRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchUserProjectByUserEmailAndRoleResponse struct {
	Results []main_entity.UserProject
}
