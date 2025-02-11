package types

import (
	main_entity "github.com/nuzur/nem/core/entity/extension_version"

	"go.uber.org/zap/zapcore"
)

type FetchExtensionVersionByVersionAndStatusRequest struct {
	Version int64
	Status  main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchExtensionVersionByVersionAndStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchExtensionVersionByVersionAndStatusResponse struct {
	Results []main_entity.ExtensionVersion
}
