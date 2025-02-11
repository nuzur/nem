package types

import (
	main_entity "github.com/nuzur/nem/core/entity/user"

	"go.uber.org/zap/zapcore"
)

type FetchUserByIdentifierAndEmailRequest struct {
	Identifier string
	Email      string

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchUserByIdentifierAndEmailRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchUserByIdentifierAndEmailResponse struct {
	Results []main_entity.User
}
