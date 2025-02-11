package user

import (
	"context"

	"github.com/nuzur/nem/core/module/user/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchUserByIdentifierAndEmailAndStatus(
	ctx context.Context,
	req types.FetchUserByIdentifierAndEmailAndStatusRequest,
	opts ...Option,
) (types.FetchUserByIdentifierAndEmailAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchUserByIdentifierAndEmailAndStatus(
			ctx,
			nemdb.FetchUserByIdentifierAndEmailAndStatusParams{
				Identifier: req.Identifier,
				Email:      req.Email,
				Status:     req.Status.ToInt64(),
				Offset:     req.Offset,
				Limit:      req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_by_identifier_and_email_and_status",
				Message:          "error in FetchUserByIdentifierAndEmailAndStatus no order",
				EntityIdentifier: "user",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserByIdentifierAndEmailAndStatusResponse{}, err
		}
		return types.FetchUserByIdentifierAndEmailAndStatusResponse{
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
			ActionIdentifier: "fetch_user_by_identifier_and_email_and_status",
			Message:          "error in FetchUserByIdentifierAndEmailAndStatus - order by field not found",
			EntityIdentifier: "user",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchUserByIdentifierAndEmailAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_by_identifier_and_email_and_status",
				Message:          "error in FetchUserByIdentifierAndEmailAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "user",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserByIdentifierAndEmailAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserByIdentifierAndEmailAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchUserByIdentifierAndEmailAndStatusOrderedByCreatedAtASCParams{
					Identifier: req.Identifier,
					Email:      req.Email,
					Status:     req.Status.ToInt64(),
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_by_identifier_and_email_and_status",
					Message:          "error in FetchUserByIdentifierAndEmailAndStatus - ordered asc",
					EntityIdentifier: "user",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserByIdentifierAndEmailAndStatusResponse{}, err
			}
			return types.FetchUserByIdentifierAndEmailAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserByIdentifierAndEmailAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchUserByIdentifierAndEmailAndStatusOrderedByCreatedAtDESCParams{
					Identifier: req.Identifier,
					Email:      req.Email,
					Status:     req.Status.ToInt64(),
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_by_identifier_and_email_and_status",
					Message:          "error in FetchUserByIdentifierAndEmailAndStatus - ordered desc",
					EntityIdentifier: "user",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserByIdentifierAndEmailAndStatusResponse{}, err
			}
			return types.FetchUserByIdentifierAndEmailAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserByIdentifierAndEmailAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchUserByIdentifierAndEmailAndStatusOrderedByUpdatedAtASCParams{
					Identifier: req.Identifier,
					Email:      req.Email,
					Status:     req.Status.ToInt64(),
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_by_identifier_and_email_and_status",
					Message:          "error in FetchUserByIdentifierAndEmailAndStatus - ordered asc",
					EntityIdentifier: "user",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserByIdentifierAndEmailAndStatusResponse{}, err
			}
			return types.FetchUserByIdentifierAndEmailAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserByIdentifierAndEmailAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchUserByIdentifierAndEmailAndStatusOrderedByUpdatedAtDESCParams{
					Identifier: req.Identifier,
					Email:      req.Email,
					Status:     req.Status.ToInt64(),
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_by_identifier_and_email_and_status",
					Message:          "error in FetchUserByIdentifierAndEmailAndStatus - ordered desc",
					EntityIdentifier: "user",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserByIdentifierAndEmailAndStatusResponse{}, err
			}
			return types.FetchUserByIdentifierAndEmailAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_user_by_identifier_and_email_and_status_invalid",
		Message:          "error in FetchUserByIdentifierAndEmailAndStatus - invalid request",
		EntityIdentifier: "user",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchUserByIdentifierAndEmailAndStatusResponse{}, err

}
