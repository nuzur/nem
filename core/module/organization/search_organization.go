package organization

import (
	"context"
	"fmt"

	"github.com/nuzur/nem/core/module/organization/types"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) Search(
	ctx context.Context,
	query string,
	limit int32,
	offset int32,
) (types.SearchOrganizationResponse, error) {
	request := nemdb.SearchOrganizationParams{
		Name:   fmt.Sprintf("%%%s%%", query),
		Offset: offset,
		Limit:  limit,
	}
	models, err := m.repository.Queries.SearchOrganization(
		ctx,
		request,
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "search_organization",
			Message:          "error calling repository for SearchOrganization",
			EntityIdentifier: "organization",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             request,
			Error:            err,
		})
		return types.SearchOrganizationResponse{}, err
	}

	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "search_organization",
		Message:          "successfully handled SearchOrganization",
		EntityIdentifier: "organization",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeSuccess,
		Data:             request,
	})
	return types.SearchOrganizationResponse{
		Results: mapModelsToEntities(models),
	}, nil
}
