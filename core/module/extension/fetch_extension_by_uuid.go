package extension

import (
	"context"

	"nem/core/module/extension/types"

	"nem/monitoring"
)

func (m *module) FetchExtensionByUUID(
	ctx context.Context,
	req types.FetchExtensionByUUIDRequest,
	opts ...Option,
) (types.FetchExtensionByUUIDResponse, error) {

	models, err := m.repository.Queries.FetchExtensionByUUID(
		ctx,
		req.UUID.String(),
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "fetch_extension_by_uuid",
			Message:          "error in FetchExtensionByUUID",
			EntityIdentifier: "extension",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByUUIDResponse{}, err
	}
	return types.FetchExtensionByUUIDResponse{
		Results: mapModelsToEntities(models),
	}, nil

}
