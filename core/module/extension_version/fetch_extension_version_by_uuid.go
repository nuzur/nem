package extension_version

import (
	"context"

	"nem/core/module/extension_version/types"

	"nem/monitoring"
)

func (m *module) FetchExtensionVersionByUUID(
	ctx context.Context,
	req types.FetchExtensionVersionByUUIDRequest,
	opts ...Option,
) (types.FetchExtensionVersionByUUIDResponse, error) {

	models, err := m.repository.Queries.FetchExtensionVersionByUUID(
		ctx,
		req.UUID.String(),
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "fetch_extension_version_by_uuid",
			Message:          "error in FetchExtensionVersionByUUID",
			EntityIdentifier: "extension_version",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionVersionByUUIDResponse{}, err
	}
	return types.FetchExtensionVersionByUUIDResponse{
		Results: mapModelsToEntities(models),
	}, nil

}
