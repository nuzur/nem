package types

import (
	main_entity "github.com/nuzur/nem/core/entity/extension"

	"go.uber.org/zap/zapcore"
)

type FetchExtensionByVersionAndIdentifierAndRepositoryRequest struct {
	Version    int64
	Identifier string
	Repository string

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchExtensionByVersionAndIdentifierAndRepositoryRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchExtensionByVersionAndIdentifierAndRepositoryResponse struct {
	Results []main_entity.Extension
}
