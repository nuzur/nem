package project

import (
	"context"

	"nem/core/module/project/types"

	"errors"
	nemdb "nem/core/repository/gen"
	"nem/monitoring"
)

func (m *module) FetchProjectByVersionAndName(
	ctx context.Context,
	req types.FetchProjectByVersionAndNameRequest,
	opts ...Option,
) (types.FetchProjectByVersionAndNameResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchProjectByVersionAndName(
			ctx,
			nemdb.FetchProjectByVersionAndNameParams{
				Version: req.Version,
				Name:    req.Name,
				Offset:  req.Offset,
				Limit:   req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_project_by_version_and_name",
				Message:          "error in FetchProjectByVersionAndName no order",
				EntityIdentifier: "project",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchProjectByVersionAndNameResponse{}, err
		}
		return types.FetchProjectByVersionAndNameResponse{
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
			ActionIdentifier: "fetch_project_by_version_and_name",
			Message:          "error in FetchProjectByVersionAndName - order by field not found",
			EntityIdentifier: "project",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchProjectByVersionAndNameResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_project_by_version_and_name",
				Message:          "error in FetchProjectByVersionAndName - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "project",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchProjectByVersionAndNameResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchProjectByVersionAndNameOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchProjectByVersionAndNameOrderedByCreatedAtASCParams{
					Version: req.Version,
					Name:    req.Name,
					Offset:  req.Offset,
					Limit:   req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_by_version_and_name",
					Message:          "error in FetchProjectByVersionAndName - ordered asc",
					EntityIdentifier: "project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectByVersionAndNameResponse{}, err
			}
			return types.FetchProjectByVersionAndNameResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchProjectByVersionAndNameOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchProjectByVersionAndNameOrderedByCreatedAtDESCParams{
					Version: req.Version,
					Name:    req.Name,
					Offset:  req.Offset,
					Limit:   req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_by_version_and_name",
					Message:          "error in FetchProjectByVersionAndName - ordered desc",
					EntityIdentifier: "project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectByVersionAndNameResponse{}, err
			}
			return types.FetchProjectByVersionAndNameResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchProjectByVersionAndNameOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchProjectByVersionAndNameOrderedByUpdatedAtASCParams{
					Version: req.Version,
					Name:    req.Name,
					Offset:  req.Offset,
					Limit:   req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_by_version_and_name",
					Message:          "error in FetchProjectByVersionAndName - ordered asc",
					EntityIdentifier: "project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectByVersionAndNameResponse{}, err
			}
			return types.FetchProjectByVersionAndNameResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchProjectByVersionAndNameOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchProjectByVersionAndNameOrderedByUpdatedAtDESCParams{
					Version: req.Version,
					Name:    req.Name,
					Offset:  req.Offset,
					Limit:   req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_by_version_and_name",
					Message:          "error in FetchProjectByVersionAndName - ordered desc",
					EntityIdentifier: "project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectByVersionAndNameResponse{}, err
			}
			return types.FetchProjectByVersionAndNameResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_project_by_version_and_name_invalid",
		Message:          "error in FetchProjectByVersionAndName - invalid request",
		EntityIdentifier: "project",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchProjectByVersionAndNameResponse{}, err

}
