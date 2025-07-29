package user_project

import (
	"context"

	"github.com/nuzur/nem/core/module/user_project/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchUserProjectByRole(
	ctx context.Context,
	req types.FetchUserProjectByRoleRequest,
	opts ...Option,
) (types.FetchUserProjectByRoleResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchUserProjectByRole(
			ctx,
			nemdb.FetchUserProjectByRoleParams{

				Role: req.Role.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_project_by_role",
				Message:          "error in FetchUserProjectByRole no order",
				EntityIdentifier: "user_project",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserProjectByRoleResponse{}, err
		}
		return types.FetchUserProjectByRoleResponse{
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
			ActionIdentifier: "fetch_user_project_by_role",
			Message:          "error in FetchUserProjectByRole - order by field not found",
			EntityIdentifier: "user_project",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchUserProjectByRoleResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_project_by_role",
				Message:          "error in FetchUserProjectByRole - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "user_project",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserProjectByRoleResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserProjectByRoleOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchUserProjectByRoleOrderedByCreatedAtASCParams{

					Role: req.Role.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_project_by_role",
					Message:          "error in FetchUserProjectByRole - ordered asc",
					EntityIdentifier: "user_project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserProjectByRoleResponse{}, err
			}
			return types.FetchUserProjectByRoleResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserProjectByRoleOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchUserProjectByRoleOrderedByCreatedAtDESCParams{

					Role: req.Role.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_project_by_role",
					Message:          "error in FetchUserProjectByRole - ordered desc",
					EntityIdentifier: "user_project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserProjectByRoleResponse{}, err
			}
			return types.FetchUserProjectByRoleResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserProjectByRoleOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchUserProjectByRoleOrderedByUpdatedAtASCParams{

					Role: req.Role.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_project_by_role",
					Message:          "error in FetchUserProjectByRole - ordered asc",
					EntityIdentifier: "user_project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserProjectByRoleResponse{}, err
			}
			return types.FetchUserProjectByRoleResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserProjectByRoleOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchUserProjectByRoleOrderedByUpdatedAtDESCParams{

					Role: req.Role.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_project_by_role",
					Message:          "error in FetchUserProjectByRole - ordered desc",
					EntityIdentifier: "user_project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserProjectByRoleResponse{}, err
			}
			return types.FetchUserProjectByRoleResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_user_project_by_role_invalid",
		Message:          "error in FetchUserProjectByRole - invalid request",
		EntityIdentifier: "user_project",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchUserProjectByRoleResponse{}, err

}
