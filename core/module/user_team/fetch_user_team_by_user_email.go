package user_team

import (
	"context"

	"nem/core/module/user_team/types"

	"errors"
	nemdb "nem/core/repository/gen"
	"nem/monitoring"
)

func (m *module) FetchUserTeamByUserEmail(
	ctx context.Context,
	req types.FetchUserTeamByUserEmailRequest,
	opts ...Option,
) (types.FetchUserTeamByUserEmailResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchUserTeamByUserEmail(
			ctx,
			nemdb.FetchUserTeamByUserEmailParams{
				UserEmail: req.UserEmail,
				Offset:    req.Offset,
				Limit:     req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_team_by_user_email",
				Message:          "error in FetchUserTeamByUserEmail no order",
				EntityIdentifier: "user_team",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserTeamByUserEmailResponse{}, err
		}
		return types.FetchUserTeamByUserEmailResponse{
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
			ActionIdentifier: "fetch_user_team_by_user_email",
			Message:          "error in FetchUserTeamByUserEmail - order by field not found",
			EntityIdentifier: "user_team",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchUserTeamByUserEmailResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_team_by_user_email",
				Message:          "error in FetchUserTeamByUserEmail - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "user_team",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserTeamByUserEmailResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserTeamByUserEmailOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchUserTeamByUserEmailOrderedByCreatedAtASCParams{
					UserEmail: req.UserEmail,
					Offset:    req.Offset,
					Limit:     req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_team_by_user_email",
					Message:          "error in FetchUserTeamByUserEmail - ordered asc",
					EntityIdentifier: "user_team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserTeamByUserEmailResponse{}, err
			}
			return types.FetchUserTeamByUserEmailResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserTeamByUserEmailOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchUserTeamByUserEmailOrderedByCreatedAtDESCParams{
					UserEmail: req.UserEmail,
					Offset:    req.Offset,
					Limit:     req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_team_by_user_email",
					Message:          "error in FetchUserTeamByUserEmail - ordered desc",
					EntityIdentifier: "user_team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserTeamByUserEmailResponse{}, err
			}
			return types.FetchUserTeamByUserEmailResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserTeamByUserEmailOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchUserTeamByUserEmailOrderedByUpdatedAtASCParams{
					UserEmail: req.UserEmail,
					Offset:    req.Offset,
					Limit:     req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_team_by_user_email",
					Message:          "error in FetchUserTeamByUserEmail - ordered asc",
					EntityIdentifier: "user_team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserTeamByUserEmailResponse{}, err
			}
			return types.FetchUserTeamByUserEmailResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserTeamByUserEmailOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchUserTeamByUserEmailOrderedByUpdatedAtDESCParams{
					UserEmail: req.UserEmail,
					Offset:    req.Offset,
					Limit:     req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_team_by_user_email",
					Message:          "error in FetchUserTeamByUserEmail - ordered desc",
					EntityIdentifier: "user_team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserTeamByUserEmailResponse{}, err
			}
			return types.FetchUserTeamByUserEmailResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_user_team_by_user_email_invalid",
		Message:          "error in FetchUserTeamByUserEmail - invalid request",
		EntityIdentifier: "user_team",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchUserTeamByUserEmailResponse{}, err

}
