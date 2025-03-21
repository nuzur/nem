package team

import (
	"context"

	"github.com/nuzur/nem/core/module/team/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchTeamByVersion(
	ctx context.Context,
	req types.FetchTeamByVersionRequest,
	opts ...Option,
) (types.FetchTeamByVersionResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchTeamByVersion(
			ctx,
			nemdb.FetchTeamByVersionParams{

				Version: req.Version,

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_team_by_version",
				Message:          "error in FetchTeamByVersion no order",
				EntityIdentifier: "team",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchTeamByVersionResponse{}, err
		}
		return types.FetchTeamByVersionResponse{
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
			ActionIdentifier: "fetch_team_by_version",
			Message:          "error in FetchTeamByVersion - order by field not found",
			EntityIdentifier: "team",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchTeamByVersionResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_team_by_version",
				Message:          "error in FetchTeamByVersion - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "team",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchTeamByVersionResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchTeamByVersionOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchTeamByVersionOrderedByCreatedAtASCParams{

					Version: req.Version,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_team_by_version",
					Message:          "error in FetchTeamByVersion - ordered asc",
					EntityIdentifier: "team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchTeamByVersionResponse{}, err
			}
			return types.FetchTeamByVersionResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchTeamByVersionOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchTeamByVersionOrderedByCreatedAtDESCParams{

					Version: req.Version,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_team_by_version",
					Message:          "error in FetchTeamByVersion - ordered desc",
					EntityIdentifier: "team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchTeamByVersionResponse{}, err
			}
			return types.FetchTeamByVersionResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchTeamByVersionOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchTeamByVersionOrderedByUpdatedAtASCParams{

					Version: req.Version,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_team_by_version",
					Message:          "error in FetchTeamByVersion - ordered asc",
					EntityIdentifier: "team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchTeamByVersionResponse{}, err
			}
			return types.FetchTeamByVersionResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchTeamByVersionOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchTeamByVersionOrderedByUpdatedAtDESCParams{

					Version: req.Version,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_team_by_version",
					Message:          "error in FetchTeamByVersion - ordered desc",
					EntityIdentifier: "team",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchTeamByVersionResponse{}, err
			}
			return types.FetchTeamByVersionResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_team_by_version_invalid",
		Message:          "error in FetchTeamByVersion - invalid request",
		EntityIdentifier: "team",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchTeamByVersionResponse{}, err

}
