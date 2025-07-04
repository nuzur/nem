package user_project

import (
	"context"

	"github.com/nuzur/nem/core/module/user_project/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchUserProjectByStatus(
	ctx context.Context,
	req types.FetchUserProjectByStatusRequest,
	opts ...Option,
) (types.FetchUserProjectByStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchUserProjectByStatus(
			ctx,
			nemdb.FetchUserProjectByStatusParams{

				Status: req.Status.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_project_by_status",
				Message:          "error in FetchUserProjectByStatus no order",
				EntityIdentifier: "user_project",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserProjectByStatusResponse{}, err
		}
		return types.FetchUserProjectByStatusResponse{
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
			ActionIdentifier: "fetch_user_project_by_status",
			Message:          "error in FetchUserProjectByStatus - order by field not found",
			EntityIdentifier: "user_project",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchUserProjectByStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_project_by_status",
				Message:          "error in FetchUserProjectByStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "user_project",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserProjectByStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserProjectByStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchUserProjectByStatusOrderedByCreatedAtASCParams{

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_project_by_status",
					Message:          "error in FetchUserProjectByStatus - ordered asc",
					EntityIdentifier: "user_project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserProjectByStatusResponse{}, err
			}
			return types.FetchUserProjectByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserProjectByStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchUserProjectByStatusOrderedByCreatedAtDESCParams{

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_project_by_status",
					Message:          "error in FetchUserProjectByStatus - ordered desc",
					EntityIdentifier: "user_project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserProjectByStatusResponse{}, err
			}
			return types.FetchUserProjectByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserProjectByStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchUserProjectByStatusOrderedByUpdatedAtASCParams{

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_project_by_status",
					Message:          "error in FetchUserProjectByStatus - ordered asc",
					EntityIdentifier: "user_project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserProjectByStatusResponse{}, err
			}
			return types.FetchUserProjectByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserProjectByStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchUserProjectByStatusOrderedByUpdatedAtDESCParams{

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_project_by_status",
					Message:          "error in FetchUserProjectByStatus - ordered desc",
					EntityIdentifier: "user_project",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserProjectByStatusResponse{}, err
			}
			return types.FetchUserProjectByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_user_project_by_status_invalid",
		Message:          "error in FetchUserProjectByStatus - invalid request",
		EntityIdentifier: "user_project",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchUserProjectByStatusResponse{}, err

}
