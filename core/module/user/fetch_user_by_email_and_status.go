package user

import (
	"context"

	"nem/core/module/user/types"

	"errors"
	nemdb "nem/core/repository/gen"
	"nem/monitoring"
)

func (m *module) FetchUserByEmailAndStatus(
	ctx context.Context,
	req types.FetchUserByEmailAndStatusRequest,
	opts ...Option,
) (types.FetchUserByEmailAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchUserByEmailAndStatus(
			ctx,
			nemdb.FetchUserByEmailAndStatusParams{
				Email:  req.Email,
				Status: req.Status.ToInt64(),
				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_by_email_and_status",
				Message:          "error in FetchUserByEmailAndStatus no order",
				EntityIdentifier: "user",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserByEmailAndStatusResponse{}, err
		}
		return types.FetchUserByEmailAndStatusResponse{
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
			ActionIdentifier: "fetch_user_by_email_and_status",
			Message:          "error in FetchUserByEmailAndStatus - order by field not found",
			EntityIdentifier: "user",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchUserByEmailAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_by_email_and_status",
				Message:          "error in FetchUserByEmailAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "user",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserByEmailAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserByEmailAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchUserByEmailAndStatusOrderedByCreatedAtASCParams{
					Email:  req.Email,
					Status: req.Status.ToInt64(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_by_email_and_status",
					Message:          "error in FetchUserByEmailAndStatus - ordered asc",
					EntityIdentifier: "user",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserByEmailAndStatusResponse{}, err
			}
			return types.FetchUserByEmailAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserByEmailAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchUserByEmailAndStatusOrderedByCreatedAtDESCParams{
					Email:  req.Email,
					Status: req.Status.ToInt64(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_by_email_and_status",
					Message:          "error in FetchUserByEmailAndStatus - ordered desc",
					EntityIdentifier: "user",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserByEmailAndStatusResponse{}, err
			}
			return types.FetchUserByEmailAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserByEmailAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchUserByEmailAndStatusOrderedByUpdatedAtASCParams{
					Email:  req.Email,
					Status: req.Status.ToInt64(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_by_email_and_status",
					Message:          "error in FetchUserByEmailAndStatus - ordered asc",
					EntityIdentifier: "user",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserByEmailAndStatusResponse{}, err
			}
			return types.FetchUserByEmailAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserByEmailAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchUserByEmailAndStatusOrderedByUpdatedAtDESCParams{
					Email:  req.Email,
					Status: req.Status.ToInt64(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_by_email_and_status",
					Message:          "error in FetchUserByEmailAndStatus - ordered desc",
					EntityIdentifier: "user",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserByEmailAndStatusResponse{}, err
			}
			return types.FetchUserByEmailAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_user_by_email_and_status_invalid",
		Message:          "error in FetchUserByEmailAndStatus - invalid request",
		EntityIdentifier: "user",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchUserByEmailAndStatusResponse{}, err

}
