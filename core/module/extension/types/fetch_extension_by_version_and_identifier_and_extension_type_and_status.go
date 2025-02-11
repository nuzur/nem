package types

import (
	main_entity "nem/core/entity/extension"

	"go.uber.org/zap/zapcore"
)

type FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatusRequest struct {
	Version       int64
	Identifier    string
	ExtensionType main_entity.ExtensionType
	Status        main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatusResponse struct {
	Results []main_entity.Extension
}
