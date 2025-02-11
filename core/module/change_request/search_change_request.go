package change_request

import (
	"context"
	"fmt"

	"nem/core/module/change_request/types"
	nemdb "nem/core/repository/gen"
	"nem/monitoring"
)

func (m *module) Search(
	ctx context.Context,
	query string,
	limit int32,
	offset int32,
) (types.SearchChangeRequestResponse, error) {
	request := nemdb.SearchChangeRequestParams{
		Title:       fmt.Sprintf("%%%s%%", query),
		Description: fmt.Sprintf("%%%s%%", query),
		Offset:      offset,
		Limit:       limit,
	}
	models, err := m.repository.Queries.SearchChangeRequest(
		ctx,
		request,
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "search_change_request",
			Message:          "error calling repository for SearchChangeRequest",
			EntityIdentifier: "change_request",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             request,
			Error:            err,
		})
		return types.SearchChangeRequestResponse{}, err
	}

	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "search_change_request",
		Message:          "successfully handled SearchChangeRequest",
		EntityIdentifier: "change_request",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeSuccess,
		Data:             request,
	})
	return types.SearchChangeRequestResponse{
		Results: mapModelsToEntities(models),
	}, nil
}
