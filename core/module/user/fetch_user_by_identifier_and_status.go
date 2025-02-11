package user

import (
	"context"

	"nem/core/module/user/types"

	"errors"
	nemdb "nem/core/repository/gen"
	"nem/monitoring"
)

func (m *module) FetchUserByIdentifierAndStatus(
	ctx context.Context,
	req types.FetchUserByIdentifierAndStatusRequest,
	opts ...Option,
) (types.FetchUserByIdentifierAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchUserByIdentifierAndStatus(
			ctx,
			nemdb.FetchUserByIdentifierAndStatusParams{
				Identifier: req.Identifier,
				Status:     req.Status.ToInt64(),
				Offset:     req.Offset,
				Limit:      req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_by_identifier_and_status",
				Message:          "error in FetchUserByIdentifierAndStatus no order",
				EntityIdentifier: "user",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserByIdentifierAndStatusResponse{}, err
		}
		return types.FetchUserByIdentifierAndStatusResponse{
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
			ActionIdentifier: "fetch_user_by_identifier_and_status",
			Message:          "error in FetchUserByIdentifierAndStatus - order by field not found",
			EntityIdentifier: "user",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchUserByIdentifierAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_by_identifier_and_status",
				Message:          "error in FetchUserByIdentifierAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "user",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserByIdentifierAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserByIdentifierAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchUserByIdentifierAndStatusOrderedByCreatedAtASCParams{
					Identifier: req.Identifier,
					Status:     req.Status.ToInt64(),
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_by_identifier_and_status",
					Message:          "error in FetchUserByIdentifierAndStatus - ordered asc",
					EntityIdentifier: "user",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserByIdentifierAndStatusResponse{}, err
			}
			return types.FetchUserByIdentifierAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserByIdentifierAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchUserByIdentifierAndStatusOrderedByCreatedAtDESCParams{
					Identifier: req.Identifier,
					Status:     req.Status.ToInt64(),
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_by_identifier_and_status",
					Message:          "error in FetchUserByIdentifierAndStatus - ordered desc",
					EntityIdentifier: "user",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserByIdentifierAndStatusResponse{}, err
			}
			return types.FetchUserByIdentifierAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserByIdentifierAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchUserByIdentifierAndStatusOrderedByUpdatedAtASCParams{
					Identifier: req.Identifier,
					Status:     req.Status.ToInt64(),
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_by_identifier_and_status",
					Message:          "error in FetchUserByIdentifierAndStatus - ordered asc",
					EntityIdentifier: "user",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserByIdentifierAndStatusResponse{}, err
			}
			return types.FetchUserByIdentifierAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserByIdentifierAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchUserByIdentifierAndStatusOrderedByUpdatedAtDESCParams{
					Identifier: req.Identifier,
					Status:     req.Status.ToInt64(),
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_by_identifier_and_status",
					Message:          "error in FetchUserByIdentifierAndStatus - ordered desc",
					EntityIdentifier: "user",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserByIdentifierAndStatusResponse{}, err
			}
			return types.FetchUserByIdentifierAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_user_by_identifier_and_status_invalid",
		Message:          "error in FetchUserByIdentifierAndStatus - invalid request",
		EntityIdentifier: "user",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchUserByIdentifierAndStatusResponse{}, err

}
