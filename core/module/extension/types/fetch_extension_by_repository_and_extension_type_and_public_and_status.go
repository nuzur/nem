package types

import (
	main_entity "github.com/nuzur/nem/core/entity/extension"

	"go.uber.org/zap/zapcore"
)

type FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatusRequest struct {
	Repository    string
	ExtensionType main_entity.ExtensionType
	Public        bool
	Status        main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatusResponse struct {
	Results []main_entity.Extension
}
