package types

import (
	"github.com/gofrs/uuid"
	main_entity "github.com/nuzur/nem/core/entity/membership"
)

type UpsertRequest struct {
	Membership main_entity.Membership
}

type UpsertResponse struct {
	UUID uuid.UUID
}
