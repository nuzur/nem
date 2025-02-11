package project_version

import (
	"context"

	"github.com/nuzur/nem/core/module/project_version/types"

	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchProjectVersionByUUID(
	ctx context.Context,
	req types.FetchProjectVersionByUUIDRequest,
	opts ...Option,
) (types.FetchProjectVersionByUUIDResponse, error) {

	models, err := m.repository.Queries.FetchProjectVersionByUUID(
		ctx,
		req.UUID.String(),
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "fetch_project_version_by_uuid",
			Message:          "error in FetchProjectVersionByUUID",
			EntityIdentifier: "project_version",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchProjectVersionByUUIDResponse{}, err
	}
	return types.FetchProjectVersionByUUIDResponse{
		Results: mapModelsToEntities(models),
	}, nil

}
