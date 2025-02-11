package change_request

import (
	"context"

	"nem/core/module/change_request/types"

	"nem/monitoring"
)

func (m *module) FetchChangeRequestByUUID(
	ctx context.Context,
	req types.FetchChangeRequestByUUIDRequest,
	opts ...Option,
) (types.FetchChangeRequestByUUIDResponse, error) {

	models, err := m.repository.Queries.FetchChangeRequestByUUID(
		ctx,
		req.UUID.String(),
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "fetch_change_request_by_uuid",
			Message:          "error in FetchChangeRequestByUUID",
			EntityIdentifier: "change_request",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchChangeRequestByUUIDResponse{}, err
	}
	return types.FetchChangeRequestByUUIDResponse{
		Results: mapModelsToEntities(models),
	}, nil

}
