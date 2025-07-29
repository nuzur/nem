package types

import (
	main_entity "github.com/nuzur/nem/core/entity/user_project"

	"go.uber.org/zap/zapcore"
)

type FetchUserProjectByUserEmailAndRoleAndStatusRequest struct {
	UserEmail *string
	Role      main_entity.Role
	Status    main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchUserProjectByUserEmailAndRoleAndStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchUserProjectByUserEmailAndRoleAndStatusResponse struct {
	Results []main_entity.UserProject
}
