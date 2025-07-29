package types

import (
	main_entity "github.com/nuzur/nem/core/entity/user_team"

	"go.uber.org/zap/zapcore"
)

type FetchUserTeamByUserEmailAndRoleAndStatusRequest struct {
	UserEmail *string
	Role      main_entity.Role
	Status    main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchUserTeamByUserEmailAndRoleAndStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchUserTeamByUserEmailAndRoleAndStatusResponse struct {
	Results []main_entity.UserTeam
}
