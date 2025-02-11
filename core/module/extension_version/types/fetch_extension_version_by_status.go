package types

import (
	main_entity "nem/core/entity/extension_version"

	"go.uber.org/zap/zapcore"
)

type FetchExtensionVersionByStatusRequest struct {
	Status main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchExtensionVersionByStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchExtensionVersionByStatusResponse struct {
	Results []main_entity.ExtensionVersion
}
