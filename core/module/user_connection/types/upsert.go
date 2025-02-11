package types

import (
	"github.com/gofrs/uuid"
	main_entity "github.com/nuzur/nem/core/entity/user_connection"
)

type UpsertRequest struct {
	UserConnection main_entity.UserConnection
}

type UpsertResponse struct {
	UUID uuid.UUID
}
