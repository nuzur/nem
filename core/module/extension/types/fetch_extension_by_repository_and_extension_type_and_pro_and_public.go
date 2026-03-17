package types

import (
	main_entity "github.com/nuzur/nem/core/entity/extension"

	"go.uber.org/zap/zapcore"
)

type FetchExtensionByRepositoryAndExtensionTypeAndProAndPublicRequest struct {
	Repository    string
	ExtensionType main_entity.ExtensionType
	Pro           bool
	Public        bool

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchExtensionByRepositoryAndExtensionTypeAndProAndPublicRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchExtensionByRepositoryAndExtensionTypeAndProAndPublicResponse struct {
	Results []main_entity.Extension
}
