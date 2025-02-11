package types

import (
	main_entity "nem/core/entity/extension"

	"go.uber.org/zap/zapcore"
)

type FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatusRequest struct {
	Version       int64
	Repository    string
	ExtensionType main_entity.ExtensionType
	Public        bool
	Status        main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatusResponse struct {
	Results []main_entity.Extension
}
