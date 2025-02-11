package types

import (
	main_entity "nem/core/entity/extension"

	"go.uber.org/zap/zapcore"
)

type FetchExtensionByPublicRequest struct {
	Public bool

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchExtensionByPublicRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchExtensionByPublicResponse struct {
	Results []main_entity.Extension
}
