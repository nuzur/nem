package types

import (
	main_entity "github.com/nuzur/nem/core/entity/user_team"

	"go.uber.org/zap/zapcore"
)

type FetchUserTeamByRoleRequest struct {
	Role main_entity.Role

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchUserTeamByRoleRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchUserTeamByRoleResponse struct {
	Results []main_entity.UserTeam
}
