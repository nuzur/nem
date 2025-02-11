package types

import (
	main_entity "github.com/nuzur/nem/core/entity/extension"

	"go.uber.org/zap/zapcore"
)

type FetchExtensionByVersionAndRepositoryAndPublicAndStatusRequest struct {
	Version    int64
	Repository string
	Public     bool
	Status     main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchExtensionByVersionAndRepositoryAndPublicAndStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchExtensionByVersionAndRepositoryAndPublicAndStatusResponse struct {
	Results []main_entity.Extension
}
