package types

import (
	main_entity "nem/core/entity/project"

	"github.com/gofrs/uuid"
	"go.uber.org/zap/zapcore"
)

type FetchProjectByUUIDRequest struct {
	UUID uuid.UUID
}

func (r FetchProjectByUUIDRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	e.AddString("uuid", r.UUID.String())

	return nil
}

type FetchProjectByUUIDResponse struct {
	Results []main_entity.Project
}
