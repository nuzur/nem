package types

import (
	main_entity "nem/core/entity/user_connection"

	"github.com/gofrs/uuid"
	"go.uber.org/zap/zapcore"
)

type FetchUserConnectionByUUIDRequest struct {
	UUID uuid.UUID
}

func (r FetchUserConnectionByUUIDRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	e.AddString("uuid", r.UUID.String())

	return nil
}

type FetchUserConnectionByUUIDResponse struct {
	Results []main_entity.UserConnection
}
