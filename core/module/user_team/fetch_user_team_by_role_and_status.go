package user_team

import (
	"context"

	"github.com/nuzur/nem/core/module/user_team/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchUserTeamByRoleAndStatus(
	ctx context.Context,
	req types.FetchUserTeamByRoleAndStatusRequest,
	opts ...Option,
) (types.FetchUserTeamByRoleAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchUserTeamByRoleAndStatus(
			ctx,
			nemdb.FetchUserTeamByRoleAndStatusParams{

				Role: req.Role.ToInt64(),

				Status: req.Status.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_team_by_role_and_status",
				Message:          "error in FetchUserTeamByRoleAndStatus no order",
				EntityIdentifier: "user_team",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserTeamByRoleAndStatusResponse{}, err
		}
		return types.FetchUserTeamByRoleAndStatusResponse{
			Results: mapModelsToEntities(models),
		}, nil
	}

	timeFields := []string{
		"created_at", "updated_at",
	}

	orderByFieldFound := false
	for _, tf := range timeFields {
		if tf == req.OrderBy {
			orderByFieldFound = true
			break
		}
	}

	if !orderByFieldFound {
		err := errors.New("order by field not found")
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "fetch_user_team_by_role_and_status",
			Message:          "error in FetchUserTeamByRoleAndStatus - order by field not found",
			EntityIdentifier: "user_team",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchUserTeamByRoleAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_team_by_role_and_status",
				Message:          "error in FetchUserTeamByRoleAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "user_team",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserTeamByRoleAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserTeamByRoleAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchUserTeamByRoleAndStatusOrderedByCreatedAtASCParams{

					Role: req.Role.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_team_by_role_and_status",
					Message:          "error in FetchUserTeamByRoleAndStatus - ordered asc",
					EntityIdentifier: "user_team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserTeamByRoleAndStatusResponse{}, err
			}
			return types.FetchUserTeamByRoleAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserTeamByRoleAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchUserTeamByRoleAndStatusOrderedByCreatedAtDESCParams{

					Role: req.Role.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_team_by_role_and_status",
					Message:          "error in FetchUserTeamByRoleAndStatus - ordered desc",
					EntityIdentifier: "user_team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserTeamByRoleAndStatusResponse{}, err
			}
			return types.FetchUserTeamByRoleAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserTeamByRoleAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchUserTeamByRoleAndStatusOrderedByUpdatedAtASCParams{

					Role: req.Role.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_team_by_role_and_status",
					Message:          "error in FetchUserTeamByRoleAndStatus - ordered asc",
					EntityIdentifier: "user_team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserTeamByRoleAndStatusResponse{}, err
			}
			return types.FetchUserTeamByRoleAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserTeamByRoleAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchUserTeamByRoleAndStatusOrderedByUpdatedAtDESCParams{

					Role: req.Role.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_team_by_role_and_status",
					Message:          "error in FetchUserTeamByRoleAndStatus - ordered desc",
					EntityIdentifier: "user_team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserTeamByRoleAndStatusResponse{}, err
			}
			return types.FetchUserTeamByRoleAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_user_team_by_role_and_status_invalid",
		Message:          "error in FetchUserTeamByRoleAndStatus - invalid request",
		EntityIdentifier: "user_team",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchUserTeamByRoleAndStatusResponse{}, err

}
