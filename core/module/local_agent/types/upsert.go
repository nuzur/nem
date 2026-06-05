package types

import (
	"github.com/gofrs/uuid"
	main_entity "github.com/nuzur/nem/core/entity/local_agent"
)

type UpsertRequest struct {
	LocalAgent main_entity.LocalAgent
}

type UpsertResponse struct {
	UUID uuid.UUID
}
