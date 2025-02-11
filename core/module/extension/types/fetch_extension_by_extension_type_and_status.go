package types

import (
	main_entity "nem/core/entity/extension"

	"go.uber.org/zap/zapcore"
)

type FetchExtensionByExtensionTypeAndStatusRequest struct {
	ExtensionType main_entity.ExtensionType
	Status        main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchExtensionByExtensionTypeAndStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchExtensionByExtensionTypeAndStatusResponse struct {
	Results []main_entity.Extension
}
