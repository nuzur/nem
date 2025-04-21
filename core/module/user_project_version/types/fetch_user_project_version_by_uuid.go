package types

import (
	main_entity "github.com/nuzur/nem/core/entity/user_project_version"

	"github.com/gofrs/uuid"
	"go.uber.org/zap/zapcore"
)

type FetchUserProjectVersionByUUIDRequest struct {
	UUID uuid.UUID
}

func (r FetchUserProjectVersionByUUIDRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	e.AddString("uuid", r.UUID.String())

	return nil
}

type FetchUserProjectVersionByUUIDResponse struct {
	Results []main_entity.UserProjectVersion
}
