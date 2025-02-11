package project

import (
	"context"

	"github.com/nuzur/nem/core/module/project/types"

	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchProjectByUUID(
	ctx context.Context,
	req types.FetchProjectByUUIDRequest,
	opts ...Option,
) (types.FetchProjectByUUIDResponse, error) {

	models, err := m.repository.Queries.FetchProjectByUUID(
		ctx,
		req.UUID.String(),
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "fetch_project_by_uuid",
			Message:          "error in FetchProjectByUUID",
			EntityIdentifier: "project",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchProjectByUUIDResponse{}, err
	}
	return types.FetchProjectByUUIDResponse{
		Results: mapModelsToEntities(models),
	}, nil

}
