package types

import (
	main_entity "nem/core/entity/extension"

	"go.uber.org/zap/zapcore"
)

type FetchExtensionByVersionAndExtensionTypeAndPublicRequest struct {
	Version       int64
	ExtensionType main_entity.ExtensionType
	Public        bool

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchExtensionByVersionAndExtensionTypeAndPublicRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchExtensionByVersionAndExtensionTypeAndPublicResponse struct {
	Results []main_entity.Extension
}
