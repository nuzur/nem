package types

import (
	main_entity "nem/core/entity/project_version"

	"github.com/gofrs/uuid"
	"go.uber.org/zap/zapcore"
)

type FetchProjectVersionByUUIDRequest struct {
	UUID uuid.UUID
}

func (r FetchProjectVersionByUUIDRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	e.AddString("uuid", r.UUID.String())

	return nil
}

type FetchProjectVersionByUUIDResponse struct {
	Results []main_entity.ProjectVersion
}
