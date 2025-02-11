package types

import (
	main_entity "nem/core/entity/user_team"

	"go.uber.org/zap/zapcore"
)

type FetchUserTeamByUserEmailAndStatusRequest struct {
	UserEmail string
	Status    main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchUserTeamByUserEmailAndStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchUserTeamByUserEmailAndStatusResponse struct {
	Results []main_entity.UserTeam
}
