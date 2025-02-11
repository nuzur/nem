package types

import (
	main_entity "github.com/nuzur/nem/core/entity/extension"

	"go.uber.org/zap/zapcore"
)

type FetchExtensionByRepositoryAndExtensionTypeRequest struct {
	Repository    string
	ExtensionType main_entity.ExtensionType

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchExtensionByRepositoryAndExtensionTypeRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchExtensionByRepositoryAndExtensionTypeResponse struct {
	Results []main_entity.Extension
}
