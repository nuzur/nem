package project_version

import (
	"context"

	"github.com/nuzur/nem/core/module/project_version/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchProjectVersionByVersionAndStatus(
	ctx context.Context,
	req types.FetchProjectVersionByVersionAndStatusRequest,
	opts ...Option,
) (types.FetchProjectVersionByVersionAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchProjectVersionByVersionAndStatus(
			ctx,
			nemdb.FetchProjectVersionByVersionAndStatusParams{
				Version: req.Version,
				Status:  req.Status.ToInt64(),
				Offset:  req.Offset,
				Limit:   req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_project_version_by_version_and_status",
				Message:          "error in FetchProjectVersionByVersionAndStatus no order",
				EntityIdentifier: "project_version",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchProjectVersionByVersionAndStatusResponse{}, err
		}
		return types.FetchProjectVersionByVersionAndStatusResponse{
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
			ActionIdentifier: "fetch_project_version_by_version_and_status",
			Message:          "error in FetchProjectVersionByVersionAndStatus - order by field not found",
			EntityIdentifier: "project_version",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchProjectVersionByVersionAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_project_version_by_version_and_status",
				Message:          "error in FetchProjectVersionByVersionAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "project_version",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchProjectVersionByVersionAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchProjectVersionByVersionAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchProjectVersionByVersionAndStatusOrderedByCreatedAtASCParams{
					Version: req.Version,
					Status:  req.Status.ToInt64(),
					Offset:  req.Offset,
					Limit:   req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_version_by_version_and_status",
					Message:          "error in FetchProjectVersionByVersionAndStatus - ordered asc",
					EntityIdentifier: "project_version",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectVersionByVersionAndStatusResponse{}, err
			}
			return types.FetchProjectVersionByVersionAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchProjectVersionByVersionAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchProjectVersionByVersionAndStatusOrderedByCreatedAtDESCParams{
					Version: req.Version,
					Status:  req.Status.ToInt64(),
					Offset:  req.Offset,
					Limit:   req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_version_by_version_and_status",
					Message:          "error in FetchProjectVersionByVersionAndStatus - ordered desc",
					EntityIdentifier: "project_version",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectVersionByVersionAndStatusResponse{}, err
			}
			return types.FetchProjectVersionByVersionAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchProjectVersionByVersionAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchProjectVersionByVersionAndStatusOrderedByUpdatedAtASCParams{
					Version: req.Version,
					Status:  req.Status.ToInt64(),
					Offset:  req.Offset,
					Limit:   req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_version_by_version_and_status",
					Message:          "error in FetchProjectVersionByVersionAndStatus - ordered asc",
					EntityIdentifier: "project_version",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectVersionByVersionAndStatusResponse{}, err
			}
			return types.FetchProjectVersionByVersionAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchProjectVersionByVersionAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchProjectVersionByVersionAndStatusOrderedByUpdatedAtDESCParams{
					Version: req.Version,
					Status:  req.Status.ToInt64(),
					Offset:  req.Offset,
					Limit:   req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_version_by_version_and_status",
					Message:          "error in FetchProjectVersionByVersionAndStatus - ordered desc",
					EntityIdentifier: "project_version",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectVersionByVersionAndStatusResponse{}, err
			}
			return types.FetchProjectVersionByVersionAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_project_version_by_version_and_status_invalid",
		Message:          "error in FetchProjectVersionByVersionAndStatus - invalid request",
		EntityIdentifier: "project_version",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchProjectVersionByVersionAndStatusResponse{}, err

}
