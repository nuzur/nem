package types

import (
	main_entity "github.com/nuzur/nem/core/entity/extension"

	"go.uber.org/zap/zapcore"
)

type FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicRequest struct {
	Version       int64
	Identifier    string
	ExtensionType main_entity.ExtensionType
	Public        bool

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicResponse struct {
	Results []main_entity.Extension
}
