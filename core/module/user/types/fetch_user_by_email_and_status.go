package types

import (
	main_entity "nem/core/entity/user"

	"go.uber.org/zap/zapcore"
)

type FetchUserByEmailAndStatusRequest struct {
	Email  string
	Status main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchUserByEmailAndStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchUserByEmailAndStatusResponse struct {
	Results []main_entity.User
}
