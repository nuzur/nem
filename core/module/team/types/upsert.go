package types

import (
	"github.com/gofrs/uuid"
	main_entity "github.com/nuzur/nem/core/entity/team"
)

type UpsertRequest struct {
	Team main_entity.Team
}

type UpsertResponse struct {
	UUID uuid.UUID
}
