package types

import (
	"github.com/gofrs/uuid"
	main_entity "nem/core/entity/extension_version"
)

type UpsertRequest struct {
	ExtensionVersion main_entity.ExtensionVersion
}

type UpsertResponse struct {
	UUID uuid.UUID
}
