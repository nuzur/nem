package types

import (
	"github.com/gofrs/uuid"
	main_entity "nem/core/entity/project"
)

type UpsertRequest struct {
	Project main_entity.Project
}

type UpsertResponse struct {
	UUID uuid.UUID
}
