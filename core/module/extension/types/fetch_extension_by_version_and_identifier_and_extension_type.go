package types

import (
	main_entity "github.com/nuzur/nem/core/entity/extension"

	"go.uber.org/zap/zapcore"
)

type FetchExtensionByVersionAndIdentifierAndExtensionTypeRequest struct {
	Version       int64
	Identifier    string
	ExtensionType main_entity.ExtensionType

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchExtensionByVersionAndIdentifierAndExtensionTypeRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchExtensionByVersionAndIdentifierAndExtensionTypeResponse struct {
	Results []main_entity.Extension
}
