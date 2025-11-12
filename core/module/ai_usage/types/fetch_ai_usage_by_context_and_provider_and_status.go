package types

import (
	main_entity "github.com/nuzur/nem/core/entity/ai_usage"

	"go.uber.org/zap/zapcore"
)

type FetchAiUsageByContextAndProviderAndStatusRequest struct {
	Context  main_entity.Context
	Provider main_entity.Provider
	Status   main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchAiUsageByContextAndProviderAndStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchAiUsageByContextAndProviderAndStatusResponse struct {
	Results []main_entity.AiUsage
}
