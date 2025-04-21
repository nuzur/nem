package types

import (
	"github.com/gofrs/uuid"
	main_entity "github.com/nuzur/nem/core/entity/user_project_version"
)

type UpsertRequest struct {
	UserProjectVersion main_entity.UserProjectVersion
}

type UpsertResponse struct {
	UUID uuid.UUID
}
