package organization

import (
	"context"

	"github.com/nuzur/nem/core/module/organization/types"

	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchOrganizationByUUID(
	ctx context.Context,
	req types.FetchOrganizationByUUIDRequest,
	opts ...Option,
) (types.FetchOrganizationByUUIDResponse, error) {

	models, err := m.repository.Queries.FetchOrganizationByUUID(
		ctx,
		req.UUID.String(),
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "fetch_organization_by_uuid",
			Message:          "error in FetchOrganizationByUUID",
			EntityIdentifier: "organization",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchOrganizationByUUIDResponse{}, err
	}
	return types.FetchOrganizationByUUIDResponse{
		Results: mapModelsToEntities(models),
	}, nil

}
