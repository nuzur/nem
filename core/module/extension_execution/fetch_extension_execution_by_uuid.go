package extension_execution

import (
	"context"

	"github.com/nuzur/nem/core/module/extension_execution/types"

	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchExtensionExecutionByUUID(
	ctx context.Context,
	req types.FetchExtensionExecutionByUUIDRequest,
	opts ...Option,
) (types.FetchExtensionExecutionByUUIDResponse, error) {

	models, err := m.repository.Queries.FetchExtensionExecutionByUUID(
		ctx,
		req.UUID.String(),
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "fetch_extension_execution_by_uuid",
			Message:          "error in FetchExtensionExecutionByUUID",
			EntityIdentifier: "extension_execution",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionExecutionByUUIDResponse{}, err
	}
	return types.FetchExtensionExecutionByUUIDResponse{
		Results: mapModelsToEntities(models),
	}, nil

}
