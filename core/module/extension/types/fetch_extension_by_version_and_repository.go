package types

import (
	main_entity "github.com/nuzur/nem/core/entity/extension"

	"go.uber.org/zap/zapcore"
)

type FetchExtensionByVersionAndRepositoryRequest struct {
	Version    int64
	Repository string

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchExtensionByVersionAndRepositoryRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchExtensionByVersionAndRepositoryResponse struct {
	Results []main_entity.Extension
}
