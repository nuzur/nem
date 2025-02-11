package types

import (
	main_entity "nem/core/entity/user"

	"go.uber.org/zap/zapcore"
)

type FetchUserByStatusRequest struct {
	Status main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchUserByStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchUserByStatusResponse struct {
	Results []main_entity.User
}
