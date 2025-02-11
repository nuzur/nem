package types

import (
	main_entity "github.com/nuzur/nem/core/entity/extension"

	"go.uber.org/zap/zapcore"
)

type FetchExtensionByVersionAndExtensionTypeAndPublicAndStatusRequest struct {
	Version       int64
	ExtensionType main_entity.ExtensionType
	Public        bool
	Status        main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchExtensionByVersionAndExtensionTypeAndPublicAndStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchExtensionByVersionAndExtensionTypeAndPublicAndStatusResponse struct {
	Results []main_entity.Extension
}
