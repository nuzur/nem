package types

import (
	main_entity "nem/core/entity/user_connection"

	"go.uber.org/zap/zapcore"
)

type FetchUserConnectionByStatusRequest struct {
	Status main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchUserConnectionByStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchUserConnectionByStatusResponse struct {
	Results []main_entity.UserConnection
}
