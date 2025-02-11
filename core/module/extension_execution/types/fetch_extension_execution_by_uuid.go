package types

import (
	main_entity "github.com/nuzur/nem/core/entity/extension_execution"

	"github.com/gofrs/uuid"
	"go.uber.org/zap/zapcore"
)

type FetchExtensionExecutionByUUIDRequest struct {
	UUID uuid.UUID
}

func (r FetchExtensionExecutionByUUIDRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	e.AddString("uuid", r.UUID.String())

	return nil
}

type FetchExtensionExecutionByUUIDResponse struct {
	Results []main_entity.ExtensionExecution
}
