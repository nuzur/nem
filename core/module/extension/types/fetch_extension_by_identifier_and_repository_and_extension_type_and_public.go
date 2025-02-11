package types

import (
	main_entity "github.com/nuzur/nem/core/entity/extension"

	"go.uber.org/zap/zapcore"
)

type FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicRequest struct {
	Identifier    string
	Repository    string
	ExtensionType main_entity.ExtensionType
	Public        bool

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicResponse struct {
	Results []main_entity.Extension
}
