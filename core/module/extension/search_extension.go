package extension

import (
	"context"
	"fmt"

	"github.com/nuzur/nem/core/module/extension/types"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) Search(
	ctx context.Context,
	query string,
	limit int32,
	offset int32,
) (types.SearchExtensionResponse, error) {
	request := nemdb.SearchExtensionParams{
		Identifier:  fmt.Sprintf("%%%s%%", query),
		Description: fmt.Sprintf("%%%s%%", query),
		Offset:      offset,
		Limit:       limit,
	}
	models, err := m.repository.Queries.SearchExtension(
		ctx,
		request,
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "search_extension",
			Message:          "error calling repository for SearchExtension",
			EntityIdentifier: "extension",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             request,
			Error:            err,
		})
		return types.SearchExtensionResponse{}, err
	}

	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "search_extension",
		Message:          "successfully handled SearchExtension",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeSuccess,
		Data:             request,
	})
	return types.SearchExtensionResponse{
		Results: mapModelsToEntities(models),
	}, nil
}
