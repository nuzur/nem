package types

import (
	main_entity "github.com/nuzur/nem/core/entity/ai_usage"

	"go.uber.org/zap/zapcore"
)

type FetchAiUsageByContextRequest struct {
	Context main_entity.Context

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchAiUsageByContextRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchAiUsageByContextResponse struct {
	Results []main_entity.AiUsage
}
