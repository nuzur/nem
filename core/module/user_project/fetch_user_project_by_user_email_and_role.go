package user_project

import (
	"context"

	"github.com/nuzur/nem/core/module/user_project/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"

	"github.com/nuzur/nem/core/entity/mapper"
)

func (m *module) FetchUserProjectByUserEmailAndRole(
	ctx context.Context,
	req types.FetchUserProjectByUserEmailAndRoleRequest,
	opts ...Option,
) (types.FetchUserProjectByUserEmailAndRoleResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchUserProjectByUserEmailAndRole(
			ctx,
			nemdb.FetchUserProjectByUserEmailAndRoleParams{

				UserEmail: mapper.StringPtrToSqlNullString(req.UserEmail),

				Role: req.Role.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_project_by_user_email_and_role",
				Message:          "error in FetchUserProjectByUserEmailAndRole no order",
				EntityIdentifier: "user_project",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserProjectByUserEmailAndRoleResponse{}, err
		}
		return types.FetchUserProjectByUserEmailAndRoleResponse{
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
			ActionIdentifier: "fetch_user_project_by_user_email_and_role",
			Message:          "error in FetchUserProjectByUserEmailAndRole - order by field not found",
			EntityIdentifier: "user_project",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchUserProjectByUserEmailAndRoleResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_project_by_user_email_and_role",
				Message:          "error in FetchUserProjectByUserEmailAndRole - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "user_project",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserProjectByUserEmailAndRoleResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserProjectByUserEmailAndRoleOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchUserProjectByUserEmailAndRoleOrderedByCreatedAtASCParams{

					UserEmail: mapper.StringPtrToSqlNullString(req.UserEmail),

					Role: req.Role.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_project_by_user_email_and_role",
					Message:          "error in FetchUserProjectByUserEmailAndRole - ordered asc",
					EntityIdentifier: "user_project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserProjectByUserEmailAndRoleResponse{}, err
			}
			return types.FetchUserProjectByUserEmailAndRoleResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserProjectByUserEmailAndRoleOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchUserProjectByUserEmailAndRoleOrderedByCreatedAtDESCParams{

					UserEmail: mapper.StringPtrToSqlNullString(req.UserEmail),

					Role: req.Role.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_project_by_user_email_and_role",
					Message:          "error in FetchUserProjectByUserEmailAndRole - ordered desc",
					EntityIdentifier: "user_project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserProjectByUserEmailAndRoleResponse{}, err
			}
			return types.FetchUserProjectByUserEmailAndRoleResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserProjectByUserEmailAndRoleOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchUserProjectByUserEmailAndRoleOrderedByUpdatedAtASCParams{

					UserEmail: mapper.StringPtrToSqlNullString(req.UserEmail),

					Role: req.Role.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_project_by_user_email_and_role",
					Message:          "error in FetchUserProjectByUserEmailAndRole - ordered asc",
					EntityIdentifier: "user_project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserProjectByUserEmailAndRoleResponse{}, err
			}
			return types.FetchUserProjectByUserEmailAndRoleResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserProjectByUserEmailAndRoleOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchUserProjectByUserEmailAndRoleOrderedByUpdatedAtDESCParams{

					UserEmail: mapper.StringPtrToSqlNullString(req.UserEmail),

					Role: req.Role.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_project_by_user_email_and_role",
					Message:          "error in FetchUserProjectByUserEmailAndRole - ordered desc",
					EntityIdentifier: "user_project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserProjectByUserEmailAndRoleResponse{}, err
			}
			return types.FetchUserProjectByUserEmailAndRoleResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_user_project_by_user_email_and_role_invalid",
		Message:          "error in FetchUserProjectByUserEmailAndRole - invalid request",
		EntityIdentifier: "user_project",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchUserProjectByUserEmailAndRoleResponse{}, err

}
