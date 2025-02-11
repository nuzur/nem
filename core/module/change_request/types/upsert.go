package types

import (
	"github.com/gofrs/uuid"
	main_entity "nem/core/entity/change_request"
)

type UpsertRequest struct {
	ChangeRequest main_entity.ChangeRequest
}

type UpsertResponse struct {
	UUID uuid.UUID
}
