package types

import (
	main_entity "nem/core/entity/user"

	"github.com/gofrs/uuid"
	"go.uber.org/zap/zapcore"
)

type FetchUserByUUIDRequest struct {
	UUID uuid.UUID
}

func (r FetchUserByUUIDRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	e.AddString("uuid", r.UUID.String())

	return nil
}

type FetchUserByUUIDResponse struct {
	Results []main_entity.User
}
