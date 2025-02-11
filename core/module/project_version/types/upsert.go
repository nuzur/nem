package types

import (
	"github.com/gofrs/uuid"
	main_entity "github.com/nuzur/nem/core/entity/project_version"
)

type UpsertRequest struct {
	ProjectVersion main_entity.ProjectVersion
}

type UpsertResponse struct {
	UUID uuid.UUID
}
