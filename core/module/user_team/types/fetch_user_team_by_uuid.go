package types

import (
	main_entity "github.com/nuzur/nem/core/entity/user_team"

	"github.com/gofrs/uuid"
	"go.uber.org/zap/zapcore"
)

type FetchUserTeamByUUIDRequest struct {
	UUID uuid.UUID
}

func (r FetchUserTeamByUUIDRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	e.AddString("uuid", r.UUID.String())

	return nil
}

type FetchUserTeamByUUIDResponse struct {
	Results []main_entity.UserTeam
}
