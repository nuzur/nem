package user_connection

import (
	"context"

	"github.com/nuzur/nem/core/module/user_connection/types"

	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchUserConnectionByUUID(
	ctx context.Context,
	req types.FetchUserConnectionByUUIDRequest,
	opts ...Option,
) (types.FetchUserConnectionByUUIDResponse, error) {

	models, err := m.repository.Queries.FetchUserConnectionByUUID(
		ctx,
		req.UUID.String(),
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "fetch_user_connection_by_uuid",
			Message:          "error in FetchUserConnectionByUUID",
			EntityIdentifier: "user_connection",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchUserConnectionByUUIDResponse{}, err
	}
	return types.FetchUserConnectionByUUIDResponse{
		Results: mapModelsToEntities(models),
	}, nil

}
