package change_request

import (
	"context"
	"fmt"

	"github.com/nuzur/nem/core/module/change_request/types"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"

	"github.com/nuzur/nem/core/entity/mapper"
)

func (m *module) Search(
	ctx context.Context,
	query string,
	limit int32,
	offset int32,
) (types.SearchChangeRequestResponse, error) {
	request := nemdb.SearchChangeRequestParams{
		Title:       fmt.Sprintf("%%%s%%", query),
		Description: mapper.StringToSqlNullString(query),
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
