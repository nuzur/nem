package user_project

import (
	"context"

	"github.com/nuzur/nem/core/module/user_project/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"

	"github.com/nuzur/nem/core/entity/mapper"
)

func (m *module) FetchUserProjectByUserEmailAndRoleAndStatus(
	ctx context.Context,
	req types.FetchUserProjectByUserEmailAndRoleAndStatusRequest,
	opts ...Option,
) (types.FetchUserProjectByUserEmailAndRoleAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchUserProjectByUserEmailAndRoleAndStatus(
			ctx,
			nemdb.FetchUserProjectByUserEmailAndRoleAndStatusParams{

				UserEmail: mapper.StringPtrToSqlNullString(req.UserEmail),

				Role: req.Role.ToInt64(),

				Status: req.Status.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_project_by_user_email_and_role_and_status",
				Message:          "error in FetchUserProjectByUserEmailAndRoleAndStatus no order",
				EntityIdentifier: "user_project",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserProjectByUserEmailAndRoleAndStatusResponse{}, err
		}
		return types.FetchUserProjectByUserEmailAndRoleAndStatusResponse{
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
			ActionIdentifier: "fetch_user_project_by_user_email_and_role_and_status",
			Message:          "error in FetchUserProjectByUserEmailAndRoleAndStatus - order by field not found",
			EntityIdentifier: "user_project",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchUserProjectByUserEmailAndRoleAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_project_by_user_email_and_role_and_status",
				Message:          "error in FetchUserProjectByUserEmailAndRoleAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "user_project",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserProjectByUserEmailAndRoleAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserProjectByUserEmailAndRoleAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchUserProjectByUserEmailAndRoleAndStatusOrderedByCreatedAtASCParams{

					UserEmail: mapper.StringPtrToSqlNullString(req.UserEmail),

					Role: req.Role.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_project_by_user_email_and_role_and_status",
					Message:          "error in FetchUserProjectByUserEmailAndRoleAndStatus - ordered asc",
					EntityIdentifier: "user_project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserProjectByUserEmailAndRoleAndStatusResponse{}, err
			}
			return types.FetchUserProjectByUserEmailAndRoleAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserProjectByUserEmailAndRoleAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchUserProjectByUserEmailAndRoleAndStatusOrderedByCreatedAtDESCParams{

					UserEmail: mapper.StringPtrToSqlNullString(req.UserEmail),

					Role: req.Role.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_project_by_user_email_and_role_and_status",
					Message:          "error in FetchUserProjectByUserEmailAndRoleAndStatus - ordered desc",
					EntityIdentifier: "user_project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserProjectByUserEmailAndRoleAndStatusResponse{}, err
			}
			return types.FetchUserProjectByUserEmailAndRoleAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserProjectByUserEmailAndRoleAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchUserProjectByUserEmailAndRoleAndStatusOrderedByUpdatedAtASCParams{

					UserEmail: mapper.StringPtrToSqlNullString(req.UserEmail),

					Role: req.Role.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_project_by_user_email_and_role_and_status",
					Message:          "error in FetchUserProjectByUserEmailAndRoleAndStatus - ordered asc",
					EntityIdentifier: "user_project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserProjectByUserEmailAndRoleAndStatusResponse{}, err
			}
			return types.FetchUserProjectByUserEmailAndRoleAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserProjectByUserEmailAndRoleAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchUserProjectByUserEmailAndRoleAndStatusOrderedByUpdatedAtDESCParams{

					UserEmail: mapper.StringPtrToSqlNullString(req.UserEmail),

					Role: req.Role.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_project_by_user_email_and_role_and_status",
					Message:          "error in FetchUserProjectByUserEmailAndRoleAndStatus - ordered desc",
					EntityIdentifier: "user_project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserProjectByUserEmailAndRoleAndStatusResponse{}, err
			}
			return types.FetchUserProjectByUserEmailAndRoleAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_user_project_by_user_email_and_role_and_status_invalid",
		Message:          "error in FetchUserProjectByUserEmailAndRoleAndStatus - invalid request",
		EntityIdentifier: "user_project",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchUserProjectByUserEmailAndRoleAndStatusResponse{}, err

}
