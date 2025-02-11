package types

import (
	main_entity "github.com/nuzur/nem/core/entity/extension"

	"go.uber.org/zap/zapcore"
)

type FetchExtensionByExtensionTypeAndPublicAndStatusRequest struct {
	ExtensionType main_entity.ExtensionType
	Public        bool
	Status        main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchExtensionByExtensionTypeAndPublicAndStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchExtensionByExtensionTypeAndPublicAndStatusResponse struct {
	Results []main_entity.Extension
}
