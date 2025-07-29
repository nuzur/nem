package user_team

import (
	"context"

	"github.com/nuzur/nem/core/module/user_team/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchUserTeamByRole(
	ctx context.Context,
	req types.FetchUserTeamByRoleRequest,
	opts ...Option,
) (types.FetchUserTeamByRoleResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchUserTeamByRole(
			ctx,
			nemdb.FetchUserTeamByRoleParams{

				Role: req.Role.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_team_by_role",
				Message:          "error in FetchUserTeamByRole no order",
				EntityIdentifier: "user_team",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserTeamByRoleResponse{}, err
		}
		return types.FetchUserTeamByRoleResponse{
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
			ActionIdentifier: "fetch_user_team_by_role",
			Message:          "error in FetchUserTeamByRole - order by field not found",
			EntityIdentifier: "user_team",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchUserTeamByRoleResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_team_by_role",
				Message:          "error in FetchUserTeamByRole - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "user_team",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserTeamByRoleResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserTeamByRoleOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchUserTeamByRoleOrderedByCreatedAtASCParams{

					Role: req.Role.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_team_by_role",
					Message:          "error in FetchUserTeamByRole - ordered asc",
					EntityIdentifier: "user_team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserTeamByRoleResponse{}, err
			}
			return types.FetchUserTeamByRoleResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserTeamByRoleOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchUserTeamByRoleOrderedByCreatedAtDESCParams{

					Role: req.Role.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_team_by_role",
					Message:          "error in FetchUserTeamByRole - ordered desc",
					EntityIdentifier: "user_team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserTeamByRoleResponse{}, err
			}
			return types.FetchUserTeamByRoleResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserTeamByRoleOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchUserTeamByRoleOrderedByUpdatedAtASCParams{

					Role: req.Role.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_team_by_role",
					Message:          "error in FetchUserTeamByRole - ordered asc",
					EntityIdentifier: "user_team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserTeamByRoleResponse{}, err
			}
			return types.FetchUserTeamByRoleResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserTeamByRoleOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchUserTeamByRoleOrderedByUpdatedAtDESCParams{

					Role: req.Role.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_team_by_role",
					Message:          "error in FetchUserTeamByRole - ordered desc",
					EntityIdentifier: "user_team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserTeamByRoleResponse{}, err
			}
			return types.FetchUserTeamByRoleResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_user_team_by_role_invalid",
		Message:          "error in FetchUserTeamByRole - invalid request",
		EntityIdentifier: "user_team",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchUserTeamByRoleResponse{}, err

}
