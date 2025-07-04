package types

import (
	main_entity "github.com/nuzur/nem/core/entity/user_project"

	"github.com/gofrs/uuid"
	"go.uber.org/zap/zapcore"
)

type FetchUserProjectByUUIDRequest struct {
	UUID uuid.UUID
}

func (r FetchUserProjectByUUIDRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	e.AddString("uuid", r.UUID.String())

	return nil
}

type FetchUserProjectByUUIDResponse struct {
	Results []main_entity.UserProject
}
