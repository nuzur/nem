package types

import (
	main_entity "github.com/nuzur/nem/core/entity/user"

	"go.uber.org/zap/zapcore"
)

type FetchUserByEmailRequest struct {
	Email string

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchUserByEmailRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchUserByEmailResponse struct {
	Results []main_entity.User
}
