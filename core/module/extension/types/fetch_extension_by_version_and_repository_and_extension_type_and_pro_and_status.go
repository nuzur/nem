package types

import (
	main_entity "github.com/nuzur/nem/core/entity/extension"

	"go.uber.org/zap/zapcore"
)

type FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatusRequest struct {
	Version       int64
	Repository    string
	ExtensionType main_entity.ExtensionType
	Pro           bool
	Status        main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatusResponse struct {
	Results []main_entity.Extension
}
