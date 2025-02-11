package types

import (
	main_entity "github.com/nuzur/nem/core/entity/team"

	"github.com/gofrs/uuid"
	"go.uber.org/zap/zapcore"
)

type FetchTeamByUUIDRequest struct {
	UUID uuid.UUID
}

func (r FetchTeamByUUIDRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	e.AddString("uuid", r.UUID.String())

	return nil
}

type FetchTeamByUUIDResponse struct {
	Results []main_entity.Team
}
