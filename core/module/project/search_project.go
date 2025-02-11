package project

import (
	"context"
	"fmt"

	"github.com/nuzur/nem/core/module/project/types"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) Search(
	ctx context.Context,
	query string,
	limit int32,
	offset int32,
) (types.SearchProjectResponse, error) {
	request := nemdb.SearchProjectParams{
		Name:        fmt.Sprintf("%%%s%%", query),
		Description: fmt.Sprintf("%%%s%%", query),
		Offset:      offset,
		Limit:       limit,
	}
	models, err := m.repository.Queries.SearchProject(
		ctx,
		request,
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "search_project",
			Message:          "error calling repository for SearchProject",
			EntityIdentifier: "project",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             request,
			Error:            err,
		})
		return types.SearchProjectResponse{}, err
	}

	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "search_project",
		Message:          "successfully handled SearchProject",
		EntityIdentifier: "project",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeSuccess,
		Data:             request,
	})
	return types.SearchProjectResponse{
		Results: mapModelsToEntities(models),
	}, nil
}
