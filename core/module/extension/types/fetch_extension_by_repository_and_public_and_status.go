package types

import (
	main_entity "nem/core/entity/extension"

	"go.uber.org/zap/zapcore"
)

type FetchExtensionByRepositoryAndPublicAndStatusRequest struct {
	Repository string
	Public     bool
	Status     main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchExtensionByRepositoryAndPublicAndStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchExtensionByRepositoryAndPublicAndStatusResponse struct {
	Results []main_entity.Extension
}
