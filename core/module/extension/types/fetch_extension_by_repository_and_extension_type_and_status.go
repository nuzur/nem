package types

import (
	main_entity "github.com/nuzur/nem/core/entity/extension"

	"go.uber.org/zap/zapcore"
)

type FetchExtensionByRepositoryAndExtensionTypeAndStatusRequest struct {
	Repository    string
	ExtensionType main_entity.ExtensionType
	Status        main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchExtensionByRepositoryAndExtensionTypeAndStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchExtensionByRepositoryAndExtensionTypeAndStatusResponse struct {
	Results []main_entity.Extension
}
