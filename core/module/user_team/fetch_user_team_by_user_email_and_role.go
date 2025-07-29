package user_team

import (
	"context"

	"github.com/nuzur/nem/core/module/user_team/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"

	"github.com/nuzur/nem/core/entity/mapper"
)

func (m *module) FetchUserTeamByUserEmailAndRole(
	ctx context.Context,
	req types.FetchUserTeamByUserEmailAndRoleRequest,
	opts ...Option,
) (types.FetchUserTeamByUserEmailAndRoleResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchUserTeamByUserEmailAndRole(
			ctx,
			nemdb.FetchUserTeamByUserEmailAndRoleParams{

				UserEmail: mapper.StringPtrToSqlNullString(req.UserEmail),

				Role: req.Role.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_team_by_user_email_and_role",
				Message:          "error in FetchUserTeamByUserEmailAndRole no order",
				EntityIdentifier: "user_team",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserTeamByUserEmailAndRoleResponse{}, err
		}
		return types.FetchUserTeamByUserEmailAndRoleResponse{
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
			ActionIdentifier: "fetch_user_team_by_user_email_and_role",
			Message:          "error in FetchUserTeamByUserEmailAndRole - order by field not found",
			EntityIdentifier: "user_team",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchUserTeamByUserEmailAndRoleResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_team_by_user_email_and_role",
				Message:          "error in FetchUserTeamByUserEmailAndRole - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "user_team",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserTeamByUserEmailAndRoleResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserTeamByUserEmailAndRoleOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchUserTeamByUserEmailAndRoleOrderedByCreatedAtASCParams{

					UserEmail: mapper.StringPtrToSqlNullString(req.UserEmail),

					Role: req.Role.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_team_by_user_email_and_role",
					Message:          "error in FetchUserTeamByUserEmailAndRole - ordered asc",
					EntityIdentifier: "user_team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserTeamByUserEmailAndRoleResponse{}, err
			}
			return types.FetchUserTeamByUserEmailAndRoleResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserTeamByUserEmailAndRoleOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchUserTeamByUserEmailAndRoleOrderedByCreatedAtDESCParams{

					UserEmail: mapper.StringPtrToSqlNullString(req.UserEmail),

					Role: req.Role.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_team_by_user_email_and_role",
					Message:          "error in FetchUserTeamByUserEmailAndRole - ordered desc",
					EntityIdentifier: "user_team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserTeamByUserEmailAndRoleResponse{}, err
			}
			return types.FetchUserTeamByUserEmailAndRoleResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserTeamByUserEmailAndRoleOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchUserTeamByUserEmailAndRoleOrderedByUpdatedAtASCParams{

					UserEmail: mapper.StringPtrToSqlNullString(req.UserEmail),

					Role: req.Role.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_team_by_user_email_and_role",
					Message:          "error in FetchUserTeamByUserEmailAndRole - ordered asc",
					EntityIdentifier: "user_team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserTeamByUserEmailAndRoleResponse{}, err
			}
			return types.FetchUserTeamByUserEmailAndRoleResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserTeamByUserEmailAndRoleOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchUserTeamByUserEmailAndRoleOrderedByUpdatedAtDESCParams{

					UserEmail: mapper.StringPtrToSqlNullString(req.UserEmail),

					Role: req.Role.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_team_by_user_email_and_role",
					Message:          "error in FetchUserTeamByUserEmailAndRole - ordered desc",
					EntityIdentifier: "user_team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserTeamByUserEmailAndRoleResponse{}, err
			}
			return types.FetchUserTeamByUserEmailAndRoleResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_user_team_by_user_email_and_role_invalid",
		Message:          "error in FetchUserTeamByUserEmailAndRole - invalid request",
		EntityIdentifier: "user_team",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchUserTeamByUserEmailAndRoleResponse{}, err

}
