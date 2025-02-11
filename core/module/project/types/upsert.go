package types

import (
	"github.com/gofrs/uuid"
	main_entity "github.com/nuzur/nem/core/entity/project"
)

type UpsertRequest struct {
	Project main_entity.Project
}

type UpsertResponse struct {
	UUID uuid.UUID
}
