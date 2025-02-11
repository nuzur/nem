package types

import (
	main_entity "github.com/nuzur/nem/core/entity/extension"

	"go.uber.org/zap/zapcore"
)

type FetchExtensionByVersionAndExtensionTypeRequest struct {
	Version       int64
	ExtensionType main_entity.ExtensionType

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchExtensionByVersionAndExtensionTypeRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchExtensionByVersionAndExtensionTypeResponse struct {
	Results []main_entity.Extension
}
