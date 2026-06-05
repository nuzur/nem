package local_agent

import (
	"context"

	"github.com/nuzur/nem/core/module/local_agent/types"

	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchLocalAgentByUUID(
	ctx context.Context,
	req types.FetchLocalAgentByUUIDRequest,
	opts ...Option,
) (types.FetchLocalAgentByUUIDResponse, error) {

	models, err := m.repository.Queries.FetchLocalAgentByUUID(
		ctx,
		req.UUID.String(),
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "fetch_local_agent_by_uuid",
			Message:          "error in FetchLocalAgentByUUID",
			EntityIdentifier: "local_agent",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchLocalAgentByUUIDResponse{}, err
	}
	return types.FetchLocalAgentByUUIDResponse{
		Results: mapModelsToEntities(models),
	}, nil

}
