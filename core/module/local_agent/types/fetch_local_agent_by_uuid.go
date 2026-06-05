package types

import (
	main_entity "github.com/nuzur/nem/core/entity/local_agent"

	"github.com/gofrs/uuid"
	"go.uber.org/zap/zapcore"
)

type FetchLocalAgentByUUIDRequest struct {
	UUID uuid.UUID
}

func (r FetchLocalAgentByUUIDRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	e.AddString("uuid", r.UUID.String())

	return nil
}

type FetchLocalAgentByUUIDResponse struct {
	Results []main_entity.LocalAgent
}
