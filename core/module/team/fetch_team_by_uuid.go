package team

import (
	"context"

	"nem/core/module/team/types"

	"nem/monitoring"
)

func (m *module) FetchTeamByUUID(
	ctx context.Context,
	req types.FetchTeamByUUIDRequest,
	opts ...Option,
) (types.FetchTeamByUUIDResponse, error) {

	models, err := m.repository.Queries.FetchTeamByUUID(
		ctx,
		req.UUID.String(),
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "fetch_team_by_uuid",
			Message:          "error in FetchTeamByUUID",
			EntityIdentifier: "team",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchTeamByUUIDResponse{}, err
	}
	return types.FetchTeamByUUIDResponse{
		Results: mapModelsToEntities(models),
	}, nil

}
