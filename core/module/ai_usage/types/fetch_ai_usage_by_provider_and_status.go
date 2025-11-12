package types

import (
	main_entity "github.com/nuzur/nem/core/entity/ai_usage"

	"go.uber.org/zap/zapcore"
)

type FetchAiUsageByProviderAndStatusRequest struct {
	Provider main_entity.Provider
	Status   main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchAiUsageByProviderAndStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchAiUsageByProviderAndStatusResponse struct {
	Results []main_entity.AiUsage
}
