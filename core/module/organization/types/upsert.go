package types

import (
	"github.com/gofrs/uuid"
	main_entity "github.com/nuzur/nem/core/entity/organization"
)

type UpsertRequest struct {
	Organization main_entity.Organization
}

type UpsertResponse struct {
	UUID uuid.UUID
}
