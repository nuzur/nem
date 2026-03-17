package types

import (
	main_entity "github.com/nuzur/nem/core/entity/extension"

	"go.uber.org/zap/zapcore"
)

type FetchExtensionByVersionAndExtensionTypeAndProAndPublicRequest struct {
	Version       int64
	ExtensionType main_entity.ExtensionType
	Pro           bool
	Public        bool

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchExtensionByVersionAndExtensionTypeAndProAndPublicRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchExtensionByVersionAndExtensionTypeAndProAndPublicResponse struct {
	Results []main_entity.Extension
}
