package user_team

import (
	"context"

	"github.com/nuzur/nem/core/module/user_team/types"

	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchUserTeamByUUID(
	ctx context.Context,
	req types.FetchUserTeamByUUIDRequest,
	opts ...Option,
) (types.FetchUserTeamByUUIDResponse, error) {

	models, err := m.repository.Queries.FetchUserTeamByUUID(
		ctx,
		req.UUID.String(),
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "fetch_user_team_by_uuid",
			Message:          "error in FetchUserTeamByUUID",
			EntityIdentifier: "user_team",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchUserTeamByUUIDResponse{}, err
	}
	return types.FetchUserTeamByUUIDResponse{
		Results: mapModelsToEntities(models),
	}, nil

}
