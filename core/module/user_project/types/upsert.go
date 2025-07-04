package types

import (
	"github.com/gofrs/uuid"
	main_entity "github.com/nuzur/nem/core/entity/user_project"
)

type UpsertRequest struct {
	UserProject main_entity.UserProject
}

type UpsertResponse struct {
	UUID uuid.UUID
}
