package types

import (
	main_entity "nem/core/entity/extension"

	"go.uber.org/zap/zapcore"
)

type FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatusRequest struct {
	Identifier    string
	Repository    string
	ExtensionType main_entity.ExtensionType
	Status        main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatusResponse struct {
	Results []main_entity.Extension
}
