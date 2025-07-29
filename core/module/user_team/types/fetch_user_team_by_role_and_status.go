package types

import (
	main_entity "github.com/nuzur/nem/core/entity/user_team"

	"go.uber.org/zap/zapcore"
)

type FetchUserTeamByRoleAndStatusRequest struct {
	Role   main_entity.Role
	Status main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchUserTeamByRoleAndStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchUserTeamByRoleAndStatusResponse struct {
	Results []main_entity.UserTeam
}
