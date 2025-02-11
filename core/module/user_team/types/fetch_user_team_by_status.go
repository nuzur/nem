package types

import (
	main_entity "nem/core/entity/user_team"

	"go.uber.org/zap/zapcore"
)

type FetchUserTeamByStatusRequest struct {
	Status main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchUserTeamByStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchUserTeamByStatusResponse struct {
	Results []main_entity.UserTeam
}
