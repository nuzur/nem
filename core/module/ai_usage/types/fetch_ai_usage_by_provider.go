package types

import (
	main_entity "github.com/nuzur/nem/core/entity/ai_usage"

	"go.uber.org/zap/zapcore"
)

type FetchAiUsageByProviderRequest struct {
	Provider main_entity.Provider

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchAiUsageByProviderRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchAiUsageByProviderResponse struct {
	Results []main_entity.AiUsage
}
