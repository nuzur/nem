package types

import (
	"github.com/gofrs/uuid"
	main_entity "nem/core/entity/extension_execution"
)

type UpsertRequest struct {
	ExtensionExecution main_entity.ExtensionExecution
}

type UpsertResponse struct {
	UUID uuid.UUID
}
