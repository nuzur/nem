package types

import (
	main_entity "github.com/nuzur/nem/core/entity/extension"

	"go.uber.org/zap/zapcore"
)

type FetchExtensionByVersionAndIdentifierAndPublicRequest struct {
	Version    int64
	Identifier string
	Public     bool

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchExtensionByVersionAndIdentifierAndPublicRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchExtensionByVersionAndIdentifierAndPublicResponse struct {
	Results []main_entity.Extension
}
