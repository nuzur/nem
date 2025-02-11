package types

import (
	"github.com/gofrs/uuid"
	main_entity "nem/core/entity/extension"
)

type UpsertRequest struct {
	Extension main_entity.Extension
}

type UpsertResponse struct {
	UUID uuid.UUID
}
