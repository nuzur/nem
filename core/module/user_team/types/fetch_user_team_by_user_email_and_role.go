package types

import (
	main_entity "github.com/nuzur/nem/core/entity/user_team"

	"go.uber.org/zap/zapcore"
)

type FetchUserTeamByUserEmailAndRoleRequest struct {
	UserEmail *string
	Role      main_entity.Role

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchUserTeamByUserEmailAndRoleRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchUserTeamByUserEmailAndRoleResponse struct {
	Results []main_entity.UserTeam
}
