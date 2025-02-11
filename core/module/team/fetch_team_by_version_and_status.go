package team

import (
	"context"

	"nem/core/module/team/types"

	"errors"
	nemdb "nem/core/repository/gen"
	"nem/monitoring"
)

func (m *module) FetchTeamByVersionAndStatus(
	ctx context.Context,
	req types.FetchTeamByVersionAndStatusRequest,
	opts ...Option,
) (types.FetchTeamByVersionAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchTeamByVersionAndStatus(
			ctx,
			nemdb.FetchTeamByVersionAndStatusParams{
				Version: req.Version,
				Status:  req.Status.ToInt64(),
				Offset:  req.Offset,
				Limit:   req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_team_by_version_and_status",
				Message:          "error in FetchTeamByVersionAndStatus no order",
				EntityIdentifier: "team",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchTeamByVersionAndStatusResponse{}, err
		}
		return types.FetchTeamByVersionAndStatusResponse{
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
			ActionIdentifier: "fetch_team_by_version_and_status",
			Message:          "error in FetchTeamByVersionAndStatus - order by field not found",
			EntityIdentifier: "team",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchTeamByVersionAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_team_by_version_and_status",
				Message:          "error in FetchTeamByVersionAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "team",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchTeamByVersionAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchTeamByVersionAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchTeamByVersionAndStatusOrderedByCreatedAtASCParams{
					Version: req.Version,
					Status:  req.Status.ToInt64(),
					Offset:  req.Offset,
					Limit:   req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_team_by_version_and_status",
					Message:          "error in FetchTeamByVersionAndStatus - ordered asc",
					EntityIdentifier: "team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchTeamByVersionAndStatusResponse{}, err
			}
			return types.FetchTeamByVersionAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchTeamByVersionAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchTeamByVersionAndStatusOrderedByCreatedAtDESCParams{
					Version: req.Version,
					Status:  req.Status.ToInt64(),
					Offset:  req.Offset,
					Limit:   req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_team_by_version_and_status",
					Message:          "error in FetchTeamByVersionAndStatus - ordered desc",
					EntityIdentifier: "team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchTeamByVersionAndStatusResponse{}, err
			}
			return types.FetchTeamByVersionAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchTeamByVersionAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchTeamByVersionAndStatusOrderedByUpdatedAtASCParams{
					Version: req.Version,
					Status:  req.Status.ToInt64(),
					Offset:  req.Offset,
					Limit:   req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_team_by_version_and_status",
					Message:          "error in FetchTeamByVersionAndStatus - ordered asc",
					EntityIdentifier: "team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchTeamByVersionAndStatusResponse{}, err
			}
			return types.FetchTeamByVersionAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchTeamByVersionAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchTeamByVersionAndStatusOrderedByUpdatedAtDESCParams{
					Version: req.Version,
					Status:  req.Status.ToInt64(),
					Offset:  req.Offset,
					Limit:   req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_team_by_version_and_status",
					Message:          "error in FetchTeamByVersionAndStatus - ordered desc",
					EntityIdentifier: "team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchTeamByVersionAndStatusResponse{}, err
			}
			return types.FetchTeamByVersionAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_team_by_version_and_status_invalid",
		Message:          "error in FetchTeamByVersionAndStatus - invalid request",
		EntityIdentifier: "team",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchTeamByVersionAndStatusResponse{}, err

}
