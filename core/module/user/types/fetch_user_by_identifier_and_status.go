package types

import (
	main_entity "github.com/nuzur/nem/core/entity/user"

	"go.uber.org/zap/zapcore"
)

type FetchUserByIdentifierAndStatusRequest struct {
	Identifier string
	Status     main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchUserByIdentifierAndStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchUserByIdentifierAndStatusResponse struct {
	Results []main_entity.User
}
