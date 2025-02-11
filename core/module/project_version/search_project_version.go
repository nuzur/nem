package project_version

import (
	"context"
	"fmt"

	"nem/core/module/project_version/types"
	nemdb "nem/core/repository/gen"
	"nem/monitoring"
)

func (m *module) Search(
	ctx context.Context,
	query string,
	limit int32,
	offset int32,
) (types.SearchProjectVersionResponse, error) {
	request := nemdb.SearchProjectVersionParams{
		Description: fmt.Sprintf("%%%s%%", query),
		Offset:      offset,
		Limit:       limit,
	}
	models, err := m.repository.Queries.SearchProjectVersion(
		ctx,
		request,
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "search_project_version",
			Message:          "error calling repository for SearchProjectVersion",
			EntityIdentifier: "project_version",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             request,
			Error:            err,
		})
		return types.SearchProjectVersionResponse{}, err
	}

	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "search_project_version",
		Message:          "successfully handled SearchProjectVersion",
		EntityIdentifier: "project_version",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeSuccess,
		Data:             request,
	})
	return types.SearchProjectVersionResponse{
		Results: mapModelsToEntities(models),
	}, nil
}
