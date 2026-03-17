package types

import (
	main_entity "github.com/nuzur/nem/core/entity/extension"

	"go.uber.org/zap/zapcore"
)

type FetchExtensionByExtensionTypeAndProRequest struct {
	ExtensionType main_entity.ExtensionType
	Pro           bool

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchExtensionByExtensionTypeAndProRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchExtensionByExtensionTypeAndProResponse struct {
	Results []main_entity.Extension
}
