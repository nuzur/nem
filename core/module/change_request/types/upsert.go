package types

import (
	"github.com/gofrs/uuid"
	main_entity "github.com/nuzur/nem/core/entity/change_request"
)

type UpsertRequest struct {
	ChangeRequest main_entity.ChangeRequest
}

type UpsertResponse struct {
	UUID uuid.UUID
}
