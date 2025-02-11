package types

import (
	"github.com/gofrs/uuid"
	main_entity "github.com/nuzur/nem/core/entity/extension_execution"
)

type UpsertRequest struct {
	ExtensionExecution main_entity.ExtensionExecution
}

type UpsertResponse struct {
	UUID uuid.UUID
}
