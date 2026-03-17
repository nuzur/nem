package types

import (
	main_entity "github.com/nuzur/nem/core/entity/extension"

	"go.uber.org/zap/zapcore"
)

type FetchExtensionByIdentifierAndExtensionTypeAndProAndStatusRequest struct {
	Identifier    string
	ExtensionType main_entity.ExtensionType
	Pro           bool
	Status        main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchExtensionByIdentifierAndExtensionTypeAndProAndStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchExtensionByIdentifierAndExtensionTypeAndProAndStatusResponse struct {
	Results []main_entity.Extension
}
