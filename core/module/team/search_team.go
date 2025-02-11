package team

import (
	"context"
	"fmt"

	"nem/core/module/team/types"
	nemdb "nem/core/repository/gen"
	"nem/monitoring"
)

func (m *module) Search(
	ctx context.Context,
	query string,
	limit int32,
	offset int32,
) (types.SearchTeamResponse, error) {
	request := nemdb.SearchTeamParams{
		Name:   fmt.Sprintf("%%%s%%", query),
		Offset: offset,
		Limit:  limit,
	}
	models, err := m.repository.Queries.SearchTeam(
		ctx,
		request,
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "search_team",
			Message:          "error calling repository for SearchTeam",
			EntityIdentifier: "team",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             request,
			Error:            err,
		})
		return types.SearchTeamResponse{}, err
	}

	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "search_team",
		Message:          "successfully handled SearchTeam",
		EntityIdentifier: "team",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeSuccess,
		Data:             request,
	})
	return types.SearchTeamResponse{
		Results: mapModelsToEntities(models),
	}, nil
}
