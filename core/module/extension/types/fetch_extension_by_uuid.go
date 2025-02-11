package types

import (
	main_entity "nem/core/entity/extension"

	"github.com/gofrs/uuid"
	"go.uber.org/zap/zapcore"
)

type FetchExtensionByUUIDRequest struct {
	UUID uuid.UUID
}

func (r FetchExtensionByUUIDRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	e.AddString("uuid", r.UUID.String())

	return nil
}

type FetchExtensionByUUIDResponse struct {
	Results []main_entity.Extension
}
