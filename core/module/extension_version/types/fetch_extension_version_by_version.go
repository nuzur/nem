package types

import (
	main_entity "github.com/nuzur/nem/core/entity/extension_version"

	"go.uber.org/zap/zapcore"
)

type FetchExtensionVersionByVersionRequest struct {
	Version int64

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchExtensionVersionByVersionRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchExtensionVersionByVersionResponse struct {
	Results []main_entity.ExtensionVersion
}
