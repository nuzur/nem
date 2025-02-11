package user

import (
	"context"

	"nem/core/module/user/types"

	"nem/monitoring"
)

func (m *module) FetchUserByUUID(
	ctx context.Context,
	req types.FetchUserByUUIDRequest,
	opts ...Option,
) (types.FetchUserByUUIDResponse, error) {

	models, err := m.repository.Queries.FetchUserByUUID(
		ctx,
		req.UUID.String(),
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "fetch_user_by_uuid",
			Message:          "error in FetchUserByUUID",
			EntityIdentifier: "user",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchUserByUUIDResponse{}, err
	}
	return types.FetchUserByUUIDResponse{
		Results: mapModelsToEntities(models),
	}, nil

}
