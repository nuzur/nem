package user_project_version

import (
	"context"

	"github.com/nuzur/nem/core/module/user_project_version/types"

	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchUserProjectVersionByUUID(
	ctx context.Context,
	req types.FetchUserProjectVersionByUUIDRequest,
	opts ...Option,
) (types.FetchUserProjectVersionByUUIDResponse, error) {

	models, err := m.repository.Queries.FetchUserProjectVersionByUUID(
		ctx,
		req.UUID.String(),
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "fetch_user_project_version_by_uuid",
			Message:          "error in FetchUserProjectVersionByUUID",
			EntityIdentifier: "user_project_version",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchUserProjectVersionByUUIDResponse{}, err
	}
	return types.FetchUserProjectVersionByUUIDResponse{
		Results: mapModelsToEntities(models),
	}, nil

}
