package types

import (
	main_entity "github.com/nuzur/nem/core/entity/extension_version"

	"github.com/gofrs/uuid"
	"go.uber.org/zap/zapcore"
)

type FetchExtensionVersionByUUIDRequest struct {
	UUID uuid.UUID
}

func (r FetchExtensionVersionByUUIDRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	e.AddString("uuid", r.UUID.String())

	return nil
}

type FetchExtensionVersionByUUIDResponse struct {
	Results []main_entity.ExtensionVersion
}
