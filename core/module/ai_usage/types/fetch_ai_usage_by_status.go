package types

import (
	main_entity "github.com/nuzur/nem/core/entity/ai_usage"

	"go.uber.org/zap/zapcore"
)

type FetchAiUsageByStatusRequest struct {
	Status main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchAiUsageByStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchAiUsageByStatusResponse struct {
	Results []main_entity.AiUsage
}
