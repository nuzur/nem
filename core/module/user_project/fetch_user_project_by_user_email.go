package user_project

import (
	"context"

	"github.com/nuzur/nem/core/module/user_project/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"

	"github.com/nuzur/nem/core/entity/mapper"
)

func (m *module) FetchUserProjectByUserEmail(
	ctx context.Context,
	req types.FetchUserProjectByUserEmailRequest,
	opts ...Option,
) (types.FetchUserProjectByUserEmailResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchUserProjectByUserEmail(
			ctx,
			nemdb.FetchUserProjectByUserEmailParams{

				UserEmail: mapper.StringPtrToSqlNullString(req.UserEmail),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_project_by_user_email",
				Message:          "error in FetchUserProjectByUserEmail no order",
				EntityIdentifier: "user_project",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserProjectByUserEmailResponse{}, err
		}
		return types.FetchUserProjectByUserEmailResponse{
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
			ActionIdentifier: "fetch_user_project_by_user_email",
			Message:          "error in FetchUserProjectByUserEmail - order by field not found",
			EntityIdentifier: "user_project",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchUserProjectByUserEmailResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_project_by_user_email",
				Message:          "error in FetchUserProjectByUserEmail - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "user_project",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserProjectByUserEmailResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserProjectByUserEmailOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchUserProjectByUserEmailOrderedByCreatedAtASCParams{

					UserEmail: mapper.StringPtrToSqlNullString(req.UserEmail),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_project_by_user_email",
					Message:          "error in FetchUserProjectByUserEmail - ordered asc",
					EntityIdentifier: "user_project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserProjectByUserEmailResponse{}, err
			}
			return types.FetchUserProjectByUserEmailResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserProjectByUserEmailOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchUserProjectByUserEmailOrderedByCreatedAtDESCParams{

					UserEmail: mapper.StringPtrToSqlNullString(req.UserEmail),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_project_by_user_email",
					Message:          "error in FetchUserProjectByUserEmail - ordered desc",
					EntityIdentifier: "user_project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserProjectByUserEmailResponse{}, err
			}
			return types.FetchUserProjectByUserEmailResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserProjectByUserEmailOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchUserProjectByUserEmailOrderedByUpdatedAtASCParams{

					UserEmail: mapper.StringPtrToSqlNullString(req.UserEmail),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_project_by_user_email",
					Message:          "error in FetchUserProjectByUserEmail - ordered asc",
					EntityIdentifier: "user_project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserProjectByUserEmailResponse{}, err
			}
			return types.FetchUserProjectByUserEmailResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserProjectByUserEmailOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchUserProjectByUserEmailOrderedByUpdatedAtDESCParams{

					UserEmail: mapper.StringPtrToSqlNullString(req.UserEmail),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_project_by_user_email",
					Message:          "error in FetchUserProjectByUserEmail - ordered desc",
					EntityIdentifier: "user_project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserProjectByUserEmailResponse{}, err
			}
			return types.FetchUserProjectByUserEmailResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_user_project_by_user_email_invalid",
		Message:          "error in FetchUserProjectByUserEmail - invalid request",
		EntityIdentifier: "user_project",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchUserProjectByUserEmailResponse{}, err

}
