package types

import (
	main_entity "nem/core/entity/extension"

	"go.uber.org/zap/zapcore"
)

type FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatusRequest struct {
	Version    int64
	Identifier string
	Repository string
	Public     bool
	Status     main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatusResponse struct {
	Results []main_entity.Extension
}
