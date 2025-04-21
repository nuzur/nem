package user_project_version

import (
	"context"

	"github.com/nuzur/nem/core/module/user_project_version/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchUserProjectVersionByStatus(
	ctx context.Context,
	req types.FetchUserProjectVersionByStatusRequest,
	opts ...Option,
) (types.FetchUserProjectVersionByStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchUserProjectVersionByStatus(
			ctx,
			nemdb.FetchUserProjectVersionByStatusParams{

				Status: req.Status.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_project_version_by_status",
				Message:          "error in FetchUserProjectVersionByStatus no order",
				EntityIdentifier: "user_project_version",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserProjectVersionByStatusResponse{}, err
		}
		return types.FetchUserProjectVersionByStatusResponse{
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
			ActionIdentifier: "fetch_user_project_version_by_status",
			Message:          "error in FetchUserProjectVersionByStatus - order by field not found",
			EntityIdentifier: "user_project_version",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchUserProjectVersionByStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_project_version_by_status",
				Message:          "error in FetchUserProjectVersionByStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "user_project_version",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserProjectVersionByStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserProjectVersionByStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchUserProjectVersionByStatusOrderedByCreatedAtASCParams{

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_project_version_by_status",
					Message:          "error in FetchUserProjectVersionByStatus - ordered asc",
					EntityIdentifier: "user_project_version",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserProjectVersionByStatusResponse{}, err
			}
			return types.FetchUserProjectVersionByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserProjectVersionByStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchUserProjectVersionByStatusOrderedByCreatedAtDESCParams{

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_project_version_by_status",
					Message:          "error in FetchUserProjectVersionByStatus - ordered desc",
					EntityIdentifier: "user_project_version",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserProjectVersionByStatusResponse{}, err
			}
			return types.FetchUserProjectVersionByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserProjectVersionByStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchUserProjectVersionByStatusOrderedByUpdatedAtASCParams{

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_project_version_by_status",
					Message:          "error in FetchUserProjectVersionByStatus - ordered asc",
					EntityIdentifier: "user_project_version",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserProjectVersionByStatusResponse{}, err
			}
			return types.FetchUserProjectVersionByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserProjectVersionByStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchUserProjectVersionByStatusOrderedByUpdatedAtDESCParams{

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_project_version_by_status",
					Message:          "error in FetchUserProjectVersionByStatus - ordered desc",
					EntityIdentifier: "user_project_version",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserProjectVersionByStatusResponse{}, err
			}
			return types.FetchUserProjectVersionByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_user_project_version_by_status_invalid",
		Message:          "error in FetchUserProjectVersionByStatus - invalid request",
		EntityIdentifier: "user_project_version",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchUserProjectVersionByStatusResponse{}, err

}
