package user_project

import (
	"context"

	"github.com/nuzur/nem/core/module/user_project/types"

	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchUserProjectByUUID(
	ctx context.Context,
	req types.FetchUserProjectByUUIDRequest,
	opts ...Option,
) (types.FetchUserProjectByUUIDResponse, error) {

	models, err := m.repository.Queries.FetchUserProjectByUUID(
		ctx,
		req.UUID.String(),
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "fetch_user_project_by_uuid",
			Message:          "error in FetchUserProjectByUUID",
			EntityIdentifier: "user_project",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchUserProjectByUUIDResponse{}, err
	}
	return types.FetchUserProjectByUUIDResponse{
		Results: mapModelsToEntities(models),
	}, nil

}
