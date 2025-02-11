package user_team

import (
	"context"

	"github.com/nuzur/nem/core/module/user_team/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchUserTeamByStatus(
	ctx context.Context,
	req types.FetchUserTeamByStatusRequest,
	opts ...Option,
) (types.FetchUserTeamByStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchUserTeamByStatus(
			ctx,
			nemdb.FetchUserTeamByStatusParams{
				Status: req.Status.ToInt64(),
				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_team_by_status",
				Message:          "error in FetchUserTeamByStatus no order",
				EntityIdentifier: "user_team",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserTeamByStatusResponse{}, err
		}
		return types.FetchUserTeamByStatusResponse{
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
			ActionIdentifier: "fetch_user_team_by_status",
			Message:          "error in FetchUserTeamByStatus - order by field not found",
			EntityIdentifier: "user_team",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchUserTeamByStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_team_by_status",
				Message:          "error in FetchUserTeamByStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "user_team",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserTeamByStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserTeamByStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchUserTeamByStatusOrderedByCreatedAtASCParams{
					Status: req.Status.ToInt64(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_team_by_status",
					Message:          "error in FetchUserTeamByStatus - ordered asc",
					EntityIdentifier: "user_team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserTeamByStatusResponse{}, err
			}
			return types.FetchUserTeamByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserTeamByStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchUserTeamByStatusOrderedByCreatedAtDESCParams{
					Status: req.Status.ToInt64(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_team_by_status",
					Message:          "error in FetchUserTeamByStatus - ordered desc",
					EntityIdentifier: "user_team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserTeamByStatusResponse{}, err
			}
			return types.FetchUserTeamByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserTeamByStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchUserTeamByStatusOrderedByUpdatedAtASCParams{
					Status: req.Status.ToInt64(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_team_by_status",
					Message:          "error in FetchUserTeamByStatus - ordered asc",
					EntityIdentifier: "user_team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserTeamByStatusResponse{}, err
			}
			return types.FetchUserTeamByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserTeamByStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchUserTeamByStatusOrderedByUpdatedAtDESCParams{
					Status: req.Status.ToInt64(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_team_by_status",
					Message:          "error in FetchUserTeamByStatus - ordered desc",
					EntityIdentifier: "user_team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserTeamByStatusResponse{}, err
			}
			return types.FetchUserTeamByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_user_team_by_status_invalid",
		Message:          "error in FetchUserTeamByStatus - invalid request",
		EntityIdentifier: "user_team",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchUserTeamByStatusResponse{}, err

}
