package ai_usage

import (
	"context"

	"github.com/nuzur/nem/core/module/ai_usage/types"

	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchAiUsageByUUID(
	ctx context.Context,
	req types.FetchAiUsageByUUIDRequest,
	opts ...Option,
) (types.FetchAiUsageByUUIDResponse, error) {

	models, err := m.repository.Queries.FetchAiUsageByUUID(
		ctx,
		req.UUID.String(),
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "fetch_ai_usage_by_uuid",
			Message:          "error in FetchAiUsageByUUID",
			EntityIdentifier: "ai_usage",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchAiUsageByUUIDResponse{}, err
	}
	return types.FetchAiUsageByUUIDResponse{
		Results: mapModelsToEntities(models),
	}, nil

}
