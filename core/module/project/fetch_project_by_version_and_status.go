package project

import (
	"context"

	"github.com/nuzur/nem/core/module/project/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchProjectByVersionAndStatus(
	ctx context.Context,
	req types.FetchProjectByVersionAndStatusRequest,
	opts ...Option,
) (types.FetchProjectByVersionAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchProjectByVersionAndStatus(
			ctx,
			nemdb.FetchProjectByVersionAndStatusParams{
				Version: req.Version,
				Status:  req.Status.ToInt64(),
				Offset:  req.Offset,
				Limit:   req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_project_by_version_and_status",
				Message:          "error in FetchProjectByVersionAndStatus no order",
				EntityIdentifier: "project",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchProjectByVersionAndStatusResponse{}, err
		}
		return types.FetchProjectByVersionAndStatusResponse{
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
			ActionIdentifier: "fetch_project_by_version_and_status",
			Message:          "error in FetchProjectByVersionAndStatus - order by field not found",
			EntityIdentifier: "project",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchProjectByVersionAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_project_by_version_and_status",
				Message:          "error in FetchProjectByVersionAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "project",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchProjectByVersionAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchProjectByVersionAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchProjectByVersionAndStatusOrderedByCreatedAtASCParams{
					Version: req.Version,
					Status:  req.Status.ToInt64(),
					Offset:  req.Offset,
					Limit:   req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_by_version_and_status",
					Message:          "error in FetchProjectByVersionAndStatus - ordered asc",
					EntityIdentifier: "project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectByVersionAndStatusResponse{}, err
			}
			return types.FetchProjectByVersionAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchProjectByVersionAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchProjectByVersionAndStatusOrderedByCreatedAtDESCParams{
					Version: req.Version,
					Status:  req.Status.ToInt64(),
					Offset:  req.Offset,
					Limit:   req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_by_version_and_status",
					Message:          "error in FetchProjectByVersionAndStatus - ordered desc",
					EntityIdentifier: "project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectByVersionAndStatusResponse{}, err
			}
			return types.FetchProjectByVersionAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchProjectByVersionAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchProjectByVersionAndStatusOrderedByUpdatedAtASCParams{
					Version: req.Version,
					Status:  req.Status.ToInt64(),
					Offset:  req.Offset,
					Limit:   req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_by_version_and_status",
					Message:          "error in FetchProjectByVersionAndStatus - ordered asc",
					EntityIdentifier: "project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectByVersionAndStatusResponse{}, err
			}
			return types.FetchProjectByVersionAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchProjectByVersionAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchProjectByVersionAndStatusOrderedByUpdatedAtDESCParams{
					Version: req.Version,
					Status:  req.Status.ToInt64(),
					Offset:  req.Offset,
					Limit:   req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_by_version_and_status",
					Message:          "error in FetchProjectByVersionAndStatus - ordered desc",
					EntityIdentifier: "project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectByVersionAndStatusResponse{}, err
			}
			return types.FetchProjectByVersionAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_project_by_version_and_status_invalid",
		Message:          "error in FetchProjectByVersionAndStatus - invalid request",
		EntityIdentifier: "project",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchProjectByVersionAndStatusResponse{}, err

}
