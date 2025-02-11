package types

import (
	"github.com/gofrs/uuid"
	main_entity "github.com/nuzur/nem/core/entity/user_team"
)

type UpsertRequest struct {
	UserTeam main_entity.UserTeam
}

type UpsertResponse struct {
	UUID uuid.UUID
}
