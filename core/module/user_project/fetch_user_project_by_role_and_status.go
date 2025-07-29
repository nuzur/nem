package user_project

import (
	"context"

	"github.com/nuzur/nem/core/module/user_project/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchUserProjectByRoleAndStatus(
	ctx context.Context,
	req types.FetchUserProjectByRoleAndStatusRequest,
	opts ...Option,
) (types.FetchUserProjectByRoleAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchUserProjectByRoleAndStatus(
			ctx,
			nemdb.FetchUserProjectByRoleAndStatusParams{

				Role: req.Role.ToInt64(),

				Status: req.Status.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_project_by_role_and_status",
				Message:          "error in FetchUserProjectByRoleAndStatus no order",
				EntityIdentifier: "user_project",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserProjectByRoleAndStatusResponse{}, err
		}
		return types.FetchUserProjectByRoleAndStatusResponse{
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
			ActionIdentifier: "fetch_user_project_by_role_and_status",
			Message:          "error in FetchUserProjectByRoleAndStatus - order by field not found",
			EntityIdentifier: "user_project",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchUserProjectByRoleAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_project_by_role_and_status",
				Message:          "error in FetchUserProjectByRoleAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "user_project",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserProjectByRoleAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserProjectByRoleAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchUserProjectByRoleAndStatusOrderedByCreatedAtASCParams{

					Role: req.Role.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_project_by_role_and_status",
					Message:          "error in FetchUserProjectByRoleAndStatus - ordered asc",
					EntityIdentifier: "user_project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserProjectByRoleAndStatusResponse{}, err
			}
			return types.FetchUserProjectByRoleAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserProjectByRoleAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchUserProjectByRoleAndStatusOrderedByCreatedAtDESCParams{

					Role: req.Role.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_project_by_role_and_status",
					Message:          "error in FetchUserProjectByRoleAndStatus - ordered desc",
					EntityIdentifier: "user_project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserProjectByRoleAndStatusResponse{}, err
			}
			return types.FetchUserProjectByRoleAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserProjectByRoleAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchUserProjectByRoleAndStatusOrderedByUpdatedAtASCParams{

					Role: req.Role.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_project_by_role_and_status",
					Message:          "error in FetchUserProjectByRoleAndStatus - ordered asc",
					EntityIdentifier: "user_project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserProjectByRoleAndStatusResponse{}, err
			}
			return types.FetchUserProjectByRoleAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserProjectByRoleAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchUserProjectByRoleAndStatusOrderedByUpdatedAtDESCParams{

					Role: req.Role.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_project_by_role_and_status",
					Message:          "error in FetchUserProjectByRoleAndStatus - ordered desc",
					EntityIdentifier: "user_project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserProjectByRoleAndStatusResponse{}, err
			}
			return types.FetchUserProjectByRoleAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_user_project_by_role_and_status_invalid",
		Message:          "error in FetchUserProjectByRoleAndStatus - invalid request",
		EntityIdentifier: "user_project",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchUserProjectByRoleAndStatusResponse{}, err

}
