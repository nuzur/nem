package user_team

import (
	"context"

	"github.com/nuzur/nem/core/module/user_team/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"

	"github.com/nuzur/nem/core/entity/mapper"
)

func (m *module) FetchUserTeamByUserEmailAndRoleAndStatus(
	ctx context.Context,
	req types.FetchUserTeamByUserEmailAndRoleAndStatusRequest,
	opts ...Option,
) (types.FetchUserTeamByUserEmailAndRoleAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchUserTeamByUserEmailAndRoleAndStatus(
			ctx,
			nemdb.FetchUserTeamByUserEmailAndRoleAndStatusParams{

				UserEmail: mapper.StringPtrToSqlNullString(req.UserEmail),

				Role: req.Role.ToInt64(),

				Status: req.Status.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_team_by_user_email_and_role_and_status",
				Message:          "error in FetchUserTeamByUserEmailAndRoleAndStatus no order",
				EntityIdentifier: "user_team",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserTeamByUserEmailAndRoleAndStatusResponse{}, err
		}
		return types.FetchUserTeamByUserEmailAndRoleAndStatusResponse{
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
			ActionIdentifier: "fetch_user_team_by_user_email_and_role_and_status",
			Message:          "error in FetchUserTeamByUserEmailAndRoleAndStatus - order by field not found",
			EntityIdentifier: "user_team",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchUserTeamByUserEmailAndRoleAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_team_by_user_email_and_role_and_status",
				Message:          "error in FetchUserTeamByUserEmailAndRoleAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "user_team",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserTeamByUserEmailAndRoleAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserTeamByUserEmailAndRoleAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchUserTeamByUserEmailAndRoleAndStatusOrderedByCreatedAtASCParams{

					UserEmail: mapper.StringPtrToSqlNullString(req.UserEmail),

					Role: req.Role.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_team_by_user_email_and_role_and_status",
					Message:          "error in FetchUserTeamByUserEmailAndRoleAndStatus - ordered asc",
					EntityIdentifier: "user_team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserTeamByUserEmailAndRoleAndStatusResponse{}, err
			}
			return types.FetchUserTeamByUserEmailAndRoleAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserTeamByUserEmailAndRoleAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchUserTeamByUserEmailAndRoleAndStatusOrderedByCreatedAtDESCParams{

					UserEmail: mapper.StringPtrToSqlNullString(req.UserEmail),

					Role: req.Role.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_team_by_user_email_and_role_and_status",
					Message:          "error in FetchUserTeamByUserEmailAndRoleAndStatus - ordered desc",
					EntityIdentifier: "user_team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserTeamByUserEmailAndRoleAndStatusResponse{}, err
			}
			return types.FetchUserTeamByUserEmailAndRoleAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserTeamByUserEmailAndRoleAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchUserTeamByUserEmailAndRoleAndStatusOrderedByUpdatedAtASCParams{

					UserEmail: mapper.StringPtrToSqlNullString(req.UserEmail),

					Role: req.Role.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_team_by_user_email_and_role_and_status",
					Message:          "error in FetchUserTeamByUserEmailAndRoleAndStatus - ordered asc",
					EntityIdentifier: "user_team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserTeamByUserEmailAndRoleAndStatusResponse{}, err
			}
			return types.FetchUserTeamByUserEmailAndRoleAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserTeamByUserEmailAndRoleAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchUserTeamByUserEmailAndRoleAndStatusOrderedByUpdatedAtDESCParams{

					UserEmail: mapper.StringPtrToSqlNullString(req.UserEmail),

					Role: req.Role.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_team_by_user_email_and_role_and_status",
					Message:          "error in FetchUserTeamByUserEmailAndRoleAndStatus - ordered desc",
					EntityIdentifier: "user_team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserTeamByUserEmailAndRoleAndStatusResponse{}, err
			}
			return types.FetchUserTeamByUserEmailAndRoleAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_user_team_by_user_email_and_role_and_status_invalid",
		Message:          "error in FetchUserTeamByUserEmailAndRoleAndStatus - invalid request",
		EntityIdentifier: "user_team",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchUserTeamByUserEmailAndRoleAndStatusResponse{}, err

}
