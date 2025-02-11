package project

import (
	"context"

	"github.com/nuzur/nem/core/module/project/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchProjectByVersion(
	ctx context.Context,
	req types.FetchProjectByVersionRequest,
	opts ...Option,
) (types.FetchProjectByVersionResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchProjectByVersion(
			ctx,
			nemdb.FetchProjectByVersionParams{
				Version: req.Version,
				Offset:  req.Offset,
				Limit:   req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_project_by_version",
				Message:          "error in FetchProjectByVersion no order",
				EntityIdentifier: "project",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchProjectByVersionResponse{}, err
		}
		return types.FetchProjectByVersionResponse{
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
			ActionIdentifier: "fetch_project_by_version",
			Message:          "error in FetchProjectByVersion - order by field not found",
			EntityIdentifier: "project",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchProjectByVersionResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_project_by_version",
				Message:          "error in FetchProjectByVersion - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "project",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchProjectByVersionResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchProjectByVersionOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchProjectByVersionOrderedByCreatedAtASCParams{
					Version: req.Version,
					Offset:  req.Offset,
					Limit:   req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_by_version",
					Message:          "error in FetchProjectByVersion - ordered asc",
					EntityIdentifier: "project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectByVersionResponse{}, err
			}
			return types.FetchProjectByVersionResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchProjectByVersionOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchProjectByVersionOrderedByCreatedAtDESCParams{
					Version: req.Version,
					Offset:  req.Offset,
					Limit:   req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_by_version",
					Message:          "error in FetchProjectByVersion - ordered desc",
					EntityIdentifier: "project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectByVersionResponse{}, err
			}
			return types.FetchProjectByVersionResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchProjectByVersionOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchProjectByVersionOrderedByUpdatedAtASCParams{
					Version: req.Version,
					Offset:  req.Offset,
					Limit:   req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_by_version",
					Message:          "error in FetchProjectByVersion - ordered asc",
					EntityIdentifier: "project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectByVersionResponse{}, err
			}
			return types.FetchProjectByVersionResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchProjectByVersionOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchProjectByVersionOrderedByUpdatedAtDESCParams{
					Version: req.Version,
					Offset:  req.Offset,
					Limit:   req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_by_version",
					Message:          "error in FetchProjectByVersion - ordered desc",
					EntityIdentifier: "project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectByVersionResponse{}, err
			}
			return types.FetchProjectByVersionResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_project_by_version_invalid",
		Message:          "error in FetchProjectByVersion - invalid request",
		EntityIdentifier: "project",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchProjectByVersionResponse{}, err

}
