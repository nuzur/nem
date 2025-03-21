package user_connection

import (
	"context"

	"github.com/nuzur/nem/core/module/user_connection/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchUserConnectionByStatus(
	ctx context.Context,
	req types.FetchUserConnectionByStatusRequest,
	opts ...Option,
) (types.FetchUserConnectionByStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchUserConnectionByStatus(
			ctx,
			nemdb.FetchUserConnectionByStatusParams{

				Status: req.Status.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_connection_by_status",
				Message:          "error in FetchUserConnectionByStatus no order",
				EntityIdentifier: "user_connection",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserConnectionByStatusResponse{}, err
		}
		return types.FetchUserConnectionByStatusResponse{
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
			ActionIdentifier: "fetch_user_connection_by_status",
			Message:          "error in FetchUserConnectionByStatus - order by field not found",
			EntityIdentifier: "user_connection",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchUserConnectionByStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_connection_by_status",
				Message:          "error in FetchUserConnectionByStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "user_connection",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserConnectionByStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserConnectionByStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchUserConnectionByStatusOrderedByCreatedAtASCParams{

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_connection_by_status",
					Message:          "error in FetchUserConnectionByStatus - ordered asc",
					EntityIdentifier: "user_connection",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserConnectionByStatusResponse{}, err
			}
			return types.FetchUserConnectionByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserConnectionByStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchUserConnectionByStatusOrderedByCreatedAtDESCParams{

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_connection_by_status",
					Message:          "error in FetchUserConnectionByStatus - ordered desc",
					EntityIdentifier: "user_connection",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserConnectionByStatusResponse{}, err
			}
			return types.FetchUserConnectionByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserConnectionByStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchUserConnectionByStatusOrderedByUpdatedAtASCParams{

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_connection_by_status",
					Message:          "error in FetchUserConnectionByStatus - ordered asc",
					EntityIdentifier: "user_connection",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserConnectionByStatusResponse{}, err
			}
			return types.FetchUserConnectionByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserConnectionByStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchUserConnectionByStatusOrderedByUpdatedAtDESCParams{

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_connection_by_status",
					Message:          "error in FetchUserConnectionByStatus - ordered desc",
					EntityIdentifier: "user_connection",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserConnectionByStatusResponse{}, err
			}
			return types.FetchUserConnectionByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_user_connection_by_status_invalid",
		Message:          "error in FetchUserConnectionByStatus - invalid request",
		EntityIdentifier: "user_connection",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchUserConnectionByStatusResponse{}, err

}
