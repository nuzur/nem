package types

import (
	main_entity "nem/core/entity/extension"

	"go.uber.org/zap/zapcore"
)

type FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusRequest struct {
	Version       int64
	Identifier    string
	Repository    string
	ExtensionType main_entity.ExtensionType
	Public        bool
	Status        main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusResponse struct {
	Results []main_entity.Extension
}
