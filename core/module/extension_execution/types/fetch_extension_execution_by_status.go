package types

import (
	main_entity "github.com/nuzur/nem/core/entity/extension_execution"

	"go.uber.org/zap/zapcore"
)

type FetchExtensionExecutionByStatusRequest struct {
	Status main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchExtensionExecutionByStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchExtensionExecutionByStatusResponse struct {
	Results []main_entity.ExtensionExecution
}
