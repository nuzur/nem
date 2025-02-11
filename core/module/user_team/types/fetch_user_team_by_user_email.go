package types

import (
	main_entity "nem/core/entity/user_team"

	"go.uber.org/zap/zapcore"
)

type FetchUserTeamByUserEmailRequest struct {
	UserEmail string

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchUserTeamByUserEmailRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchUserTeamByUserEmailResponse struct {
	Results []main_entity.UserTeam
}
