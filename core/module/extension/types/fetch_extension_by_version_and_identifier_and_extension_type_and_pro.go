package types

import (
	main_entity "github.com/nuzur/nem/core/entity/extension"

	"go.uber.org/zap/zapcore"
)

type FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProRequest struct {
	Version       int64
	Identifier    string
	ExtensionType main_entity.ExtensionType
	Pro           bool

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProResponse struct {
	Results []main_entity.Extension
}
