package types

import (
	"github.com/gofrs/uuid"
	main_entity "nem/core/entity/user_connection"
)

type UpsertRequest struct {
	UserConnection main_entity.UserConnection
}

type UpsertResponse struct {
	UUID uuid.UUID
}
