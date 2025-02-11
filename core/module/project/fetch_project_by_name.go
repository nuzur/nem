package project

import (
	"context"

	"github.com/nuzur/nem/core/module/project/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchProjectByName(
	ctx context.Context,
	req types.FetchProjectByNameRequest,
	opts ...Option,
) (types.FetchProjectByNameResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchProjectByName(
			ctx,
			nemdb.FetchProjectByNameParams{
				Name:   req.Name,
				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_project_by_name",
				Message:          "error in FetchProjectByName no order",
				EntityIdentifier: "project",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchProjectByNameResponse{}, err
		}
		return types.FetchProjectByNameResponse{
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
			ActionIdentifier: "fetch_project_by_name",
			Message:          "error in FetchProjectByName - order by field not found",
			EntityIdentifier: "project",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchProjectByNameResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_project_by_name",
				Message:          "error in FetchProjectByName - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "project",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchProjectByNameResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchProjectByNameOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchProjectByNameOrderedByCreatedAtASCParams{
					Name:   req.Name,
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_by_name",
					Message:          "error in FetchProjectByName - ordered asc",
					EntityIdentifier: "project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectByNameResponse{}, err
			}
			return types.FetchProjectByNameResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchProjectByNameOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchProjectByNameOrderedByCreatedAtDESCParams{
					Name:   req.Name,
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_by_name",
					Message:          "error in FetchProjectByName - ordered desc",
					EntityIdentifier: "project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectByNameResponse{}, err
			}
			return types.FetchProjectByNameResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchProjectByNameOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchProjectByNameOrderedByUpdatedAtASCParams{
					Name:   req.Name,
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_by_name",
					Message:          "error in FetchProjectByName - ordered asc",
					EntityIdentifier: "project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectByNameResponse{}, err
			}
			return types.FetchProjectByNameResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchProjectByNameOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchProjectByNameOrderedByUpdatedAtDESCParams{
					Name:   req.Name,
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_by_name",
					Message:          "error in FetchProjectByName - ordered desc",
					EntityIdentifier: "project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectByNameResponse{}, err
			}
			return types.FetchProjectByNameResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_project_by_name_invalid",
		Message:          "error in FetchProjectByName - invalid request",
		EntityIdentifier: "project",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchProjectByNameResponse{}, err

}
