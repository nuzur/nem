package types

import (
	main_entity "nem/core/entity/extension"

	"go.uber.org/zap/zapcore"
)

type FetchExtensionByVersionAndRepositoryAndExtensionTypeAndStatusRequest struct {
	Version       int64
	Repository    string
	ExtensionType main_entity.ExtensionType
	Status        main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchExtensionByVersionAndRepositoryAndExtensionTypeAndStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchExtensionByVersionAndRepositoryAndExtensionTypeAndStatusResponse struct {
	Results []main_entity.Extension
}
