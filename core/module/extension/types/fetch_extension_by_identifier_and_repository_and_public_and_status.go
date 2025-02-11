package types

import (
	main_entity "github.com/nuzur/nem/core/entity/extension"

	"go.uber.org/zap/zapcore"
)

type FetchExtensionByIdentifierAndRepositoryAndPublicAndStatusRequest struct {
	Identifier string
	Repository string
	Public     bool
	Status     main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchExtensionByIdentifierAndRepositoryAndPublicAndStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchExtensionByIdentifierAndRepositoryAndPublicAndStatusResponse struct {
	Results []main_entity.Extension
}
