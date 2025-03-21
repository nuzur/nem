package types

import (
	main_entity "github.com/nuzur/nem/core/entity/extension"

	"go.uber.org/zap/zapcore"
)

type FetchExtensionByIdentifierAndRepositoryAndStatusRequest struct {
	Identifier string
	Repository string
	Status     main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchExtensionByIdentifierAndRepositoryAndStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchExtensionByIdentifierAndRepositoryAndStatusResponse struct {
	Results []main_entity.Extension
}
