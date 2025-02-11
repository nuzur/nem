package types

import (
	main_entity "nem/core/entity/user"

	"go.uber.org/zap/zapcore"
)

type FetchUserByIdentifierRequest struct {
	Identifier string

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchUserByIdentifierRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchUserByIdentifierResponse struct {
	Results []main_entity.User
}
