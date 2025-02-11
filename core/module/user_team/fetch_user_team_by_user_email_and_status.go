package user_team

import (
	"context"

	"nem/core/module/user_team/types"

	"errors"
	nemdb "nem/core/repository/gen"
	"nem/monitoring"
)

func (m *module) FetchUserTeamByUserEmailAndStatus(
	ctx context.Context,
	req types.FetchUserTeamByUserEmailAndStatusRequest,
	opts ...Option,
) (types.FetchUserTeamByUserEmailAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchUserTeamByUserEmailAndStatus(
			ctx,
			nemdb.FetchUserTeamByUserEmailAndStatusParams{
				UserEmail: req.UserEmail,
				Status:    req.Status.ToInt64(),
				Offset:    req.Offset,
				Limit:     req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_team_by_user_email_and_status",
				Message:          "error in FetchUserTeamByUserEmailAndStatus no order",
				EntityIdentifier: "user_team",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserTeamByUserEmailAndStatusResponse{}, err
		}
		return types.FetchUserTeamByUserEmailAndStatusResponse{
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
			ActionIdentifier: "fetch_user_team_by_user_email_and_status",
			Message:          "error in FetchUserTeamByUserEmailAndStatus - order by field not found",
			EntityIdentifier: "user_team",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchUserTeamByUserEmailAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_team_by_user_email_and_status",
				Message:          "error in FetchUserTeamByUserEmailAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "user_team",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserTeamByUserEmailAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserTeamByUserEmailAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchUserTeamByUserEmailAndStatusOrderedByCreatedAtASCParams{
					UserEmail: req.UserEmail,
					Status:    req.Status.ToInt64(),
					Offset:    req.Offset,
					Limit:     req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_team_by_user_email_and_status",
					Message:          "error in FetchUserTeamByUserEmailAndStatus - ordered asc",
					EntityIdentifier: "user_team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserTeamByUserEmailAndStatusResponse{}, err
			}
			return types.FetchUserTeamByUserEmailAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserTeamByUserEmailAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchUserTeamByUserEmailAndStatusOrderedByCreatedAtDESCParams{
					UserEmail: req.UserEmail,
					Status:    req.Status.ToInt64(),
					Offset:    req.Offset,
					Limit:     req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_team_by_user_email_and_status",
					Message:          "error in FetchUserTeamByUserEmailAndStatus - ordered desc",
					EntityIdentifier: "user_team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserTeamByUserEmailAndStatusResponse{}, err
			}
			return types.FetchUserTeamByUserEmailAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserTeamByUserEmailAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchUserTeamByUserEmailAndStatusOrderedByUpdatedAtASCParams{
					UserEmail: req.UserEmail,
					Status:    req.Status.ToInt64(),
					Offset:    req.Offset,
					Limit:     req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_team_by_user_email_and_status",
					Message:          "error in FetchUserTeamByUserEmailAndStatus - ordered asc",
					EntityIdentifier: "user_team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserTeamByUserEmailAndStatusResponse{}, err
			}
			return types.FetchUserTeamByUserEmailAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserTeamByUserEmailAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchUserTeamByUserEmailAndStatusOrderedByUpdatedAtDESCParams{
					UserEmail: req.UserEmail,
					Status:    req.Status.ToInt64(),
					Offset:    req.Offset,
					Limit:     req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_team_by_user_email_and_status",
					Message:          "error in FetchUserTeamByUserEmailAndStatus - ordered desc",
					EntityIdentifier: "user_team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserTeamByUserEmailAndStatusResponse{}, err
			}
			return types.FetchUserTeamByUserEmailAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_user_team_by_user_email_and_status_invalid",
		Message:          "error in FetchUserTeamByUserEmailAndStatus - invalid request",
		EntityIdentifier: "user_team",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchUserTeamByUserEmailAndStatusResponse{}, err

}
