package types

import (
	"github.com/gofrs/uuid"
	main_entity "nem/core/entity/user_team"
)

type UpsertRequest struct {
	UserTeam main_entity.UserTeam
}

type UpsertResponse struct {
	UUID uuid.UUID
}
