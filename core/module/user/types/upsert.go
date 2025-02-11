package types

import (
	"github.com/gofrs/uuid"
	main_entity "nem/core/entity/user"
)

type UpsertRequest struct {
	User main_entity.User
}

type UpsertResponse struct {
	UUID uuid.UUID
}
