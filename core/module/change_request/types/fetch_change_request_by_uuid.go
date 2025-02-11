package types

import (
	main_entity "github.com/nuzur/nem/core/entity/change_request"

	"github.com/gofrs/uuid"
	"go.uber.org/zap/zapcore"
)

type FetchChangeRequestByUUIDRequest struct {
	UUID uuid.UUID
}

func (r FetchChangeRequestByUUIDRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	e.AddString("uuid", r.UUID.String())

	return nil
}

type FetchChangeRequestByUUIDResponse struct {
	Results []main_entity.ChangeRequest
}
