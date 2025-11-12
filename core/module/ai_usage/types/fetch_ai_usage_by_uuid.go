package types

import (
	main_entity "github.com/nuzur/nem/core/entity/ai_usage"

	"github.com/gofrs/uuid"
	"go.uber.org/zap/zapcore"
)

type FetchAiUsageByUUIDRequest struct {
	UUID uuid.UUID
}

func (r FetchAiUsageByUUIDRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	e.AddString("uuid", r.UUID.String())

	return nil
}

type FetchAiUsageByUUIDResponse struct {
	Results []main_entity.AiUsage
}
