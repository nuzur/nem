package project

import (
	"context"

	"github.com/nuzur/nem/core/module/project/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchProjectByVersionAndNameAndStatus(
	ctx context.Context,
	req types.FetchProjectByVersionAndNameAndStatusRequest,
	opts ...Option,
) (types.FetchProjectByVersionAndNameAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchProjectByVersionAndNameAndStatus(
			ctx,
			nemdb.FetchProjectByVersionAndNameAndStatusParams{
				Version: req.Version,
				Name:    req.Name,
				Status:  req.Status.ToInt64(),
				Offset:  req.Offset,
				Limit:   req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_project_by_version_and_name_and_status",
				Message:          "error in FetchProjectByVersionAndNameAndStatus no order",
				EntityIdentifier: "project",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchProjectByVersionAndNameAndStatusResponse{}, err
		}
		return types.FetchProjectByVersionAndNameAndStatusResponse{
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
			ActionIdentifier: "fetch_project_by_version_and_name_and_status",
			Message:          "error in FetchProjectByVersionAndNameAndStatus - order by field not found",
			EntityIdentifier: "project",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchProjectByVersionAndNameAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_project_by_version_and_name_and_status",
				Message:          "error in FetchProjectByVersionAndNameAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "project",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchProjectByVersionAndNameAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchProjectByVersionAndNameAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchProjectByVersionAndNameAndStatusOrderedByCreatedAtASCParams{
					Version: req.Version,
					Name:    req.Name,
					Status:  req.Status.ToInt64(),
					Offset:  req.Offset,
					Limit:   req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_by_version_and_name_and_status",
					Message:          "error in FetchProjectByVersionAndNameAndStatus - ordered asc",
					EntityIdentifier: "project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectByVersionAndNameAndStatusResponse{}, err
			}
			return types.FetchProjectByVersionAndNameAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchProjectByVersionAndNameAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchProjectByVersionAndNameAndStatusOrderedByCreatedAtDESCParams{
					Version: req.Version,
					Name:    req.Name,
					Status:  req.Status.ToInt64(),
					Offset:  req.Offset,
					Limit:   req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_by_version_and_name_and_status",
					Message:          "error in FetchProjectByVersionAndNameAndStatus - ordered desc",
					EntityIdentifier: "project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectByVersionAndNameAndStatusResponse{}, err
			}
			return types.FetchProjectByVersionAndNameAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchProjectByVersionAndNameAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchProjectByVersionAndNameAndStatusOrderedByUpdatedAtASCParams{
					Version: req.Version,
					Name:    req.Name,
					Status:  req.Status.ToInt64(),
					Offset:  req.Offset,
					Limit:   req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_by_version_and_name_and_status",
					Message:          "error in FetchProjectByVersionAndNameAndStatus - ordered asc",
					EntityIdentifier: "project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectByVersionAndNameAndStatusResponse{}, err
			}
			return types.FetchProjectByVersionAndNameAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchProjectByVersionAndNameAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchProjectByVersionAndNameAndStatusOrderedByUpdatedAtDESCParams{
					Version: req.Version,
					Name:    req.Name,
					Status:  req.Status.ToInt64(),
					Offset:  req.Offset,
					Limit:   req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_by_version_and_name_and_status",
					Message:          "error in FetchProjectByVersionAndNameAndStatus - ordered desc",
					EntityIdentifier: "project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectByVersionAndNameAndStatusResponse{}, err
			}
			return types.FetchProjectByVersionAndNameAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_project_by_version_and_name_and_status_invalid",
		Message:          "error in FetchProjectByVersionAndNameAndStatus - invalid request",
		EntityIdentifier: "project",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchProjectByVersionAndNameAndStatusResponse{}, err

}
