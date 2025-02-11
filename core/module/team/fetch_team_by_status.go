package team

import (
	"context"

	"github.com/nuzur/nem/core/module/team/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchTeamByStatus(
	ctx context.Context,
	req types.FetchTeamByStatusRequest,
	opts ...Option,
) (types.FetchTeamByStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchTeamByStatus(
			ctx,
			nemdb.FetchTeamByStatusParams{
				Status: req.Status.ToInt64(),
				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_team_by_status",
				Message:          "error in FetchTeamByStatus no order",
				EntityIdentifier: "team",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchTeamByStatusResponse{}, err
		}
		return types.FetchTeamByStatusResponse{
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
			ActionIdentifier: "fetch_team_by_status",
			Message:          "error in FetchTeamByStatus - order by field not found",
			EntityIdentifier: "team",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchTeamByStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_team_by_status",
				Message:          "error in FetchTeamByStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "team",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchTeamByStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchTeamByStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchTeamByStatusOrderedByCreatedAtASCParams{
					Status: req.Status.ToInt64(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_team_by_status",
					Message:          "error in FetchTeamByStatus - ordered asc",
					EntityIdentifier: "team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchTeamByStatusResponse{}, err
			}
			return types.FetchTeamByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchTeamByStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchTeamByStatusOrderedByCreatedAtDESCParams{
					Status: req.Status.ToInt64(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_team_by_status",
					Message:          "error in FetchTeamByStatus - ordered desc",
					EntityIdentifier: "team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchTeamByStatusResponse{}, err
			}
			return types.FetchTeamByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchTeamByStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchTeamByStatusOrderedByUpdatedAtASCParams{
					Status: req.Status.ToInt64(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_team_by_status",
					Message:          "error in FetchTeamByStatus - ordered asc",
					EntityIdentifier: "team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchTeamByStatusResponse{}, err
			}
			return types.FetchTeamByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchTeamByStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchTeamByStatusOrderedByUpdatedAtDESCParams{
					Status: req.Status.ToInt64(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_team_by_status",
					Message:          "error in FetchTeamByStatus - ordered desc",
					EntityIdentifier: "team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchTeamByStatusResponse{}, err
			}
			return types.FetchTeamByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_team_by_status_invalid",
		Message:          "error in FetchTeamByStatus - invalid request",
		EntityIdentifier: "team",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchTeamByStatusResponse{}, err

}
