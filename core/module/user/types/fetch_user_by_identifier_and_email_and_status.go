package types

import (
	main_entity "nem/core/entity/user"

	"go.uber.org/zap/zapcore"
)

type FetchUserByIdentifierAndEmailAndStatusRequest struct {
	Identifier string
	Email      string
	Status     main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchUserByIdentifierAndEmailAndStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchUserByIdentifierAndEmailAndStatusResponse struct {
	Results []main_entity.User
}
