package user

import (
	"context"

	"nem/core/module/user/types"

	"errors"
	nemdb "nem/core/repository/gen"
	"nem/monitoring"
)

func (m *module) FetchUserByIdentifierAndEmail(
	ctx context.Context,
	req types.FetchUserByIdentifierAndEmailRequest,
	opts ...Option,
) (types.FetchUserByIdentifierAndEmailResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchUserByIdentifierAndEmail(
			ctx,
			nemdb.FetchUserByIdentifierAndEmailParams{
				Identifier: req.Identifier,
				Email:      req.Email,
				Offset:     req.Offset,
				Limit:      req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_by_identifier_and_email",
				Message:          "error in FetchUserByIdentifierAndEmail no order",
				EntityIdentifier: "user",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserByIdentifierAndEmailResponse{}, err
		}
		return types.FetchUserByIdentifierAndEmailResponse{
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
			ActionIdentifier: "fetch_user_by_identifier_and_email",
			Message:          "error in FetchUserByIdentifierAndEmail - order by field not found",
			EntityIdentifier: "user",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchUserByIdentifierAndEmailResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_by_identifier_and_email",
				Message:          "error in FetchUserByIdentifierAndEmail - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "user",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserByIdentifierAndEmailResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserByIdentifierAndEmailOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchUserByIdentifierAndEmailOrderedByCreatedAtASCParams{
					Identifier: req.Identifier,
					Email:      req.Email,
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_by_identifier_and_email",
					Message:          "error in FetchUserByIdentifierAndEmail - ordered asc",
					EntityIdentifier: "user",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserByIdentifierAndEmailResponse{}, err
			}
			return types.FetchUserByIdentifierAndEmailResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserByIdentifierAndEmailOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchUserByIdentifierAndEmailOrderedByCreatedAtDESCParams{
					Identifier: req.Identifier,
					Email:      req.Email,
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_by_identifier_and_email",
					Message:          "error in FetchUserByIdentifierAndEmail - ordered desc",
					EntityIdentifier: "user",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserByIdentifierAndEmailResponse{}, err
			}
			return types.FetchUserByIdentifierAndEmailResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserByIdentifierAndEmailOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchUserByIdentifierAndEmailOrderedByUpdatedAtASCParams{
					Identifier: req.Identifier,
					Email:      req.Email,
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_by_identifier_and_email",
					Message:          "error in FetchUserByIdentifierAndEmail - ordered asc",
					EntityIdentifier: "user",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserByIdentifierAndEmailResponse{}, err
			}
			return types.FetchUserByIdentifierAndEmailResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserByIdentifierAndEmailOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchUserByIdentifierAndEmailOrderedByUpdatedAtDESCParams{
					Identifier: req.Identifier,
					Email:      req.Email,
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_by_identifier_and_email",
					Message:          "error in FetchUserByIdentifierAndEmail - ordered desc",
					EntityIdentifier: "user",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserByIdentifierAndEmailResponse{}, err
			}
			return types.FetchUserByIdentifierAndEmailResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_user_by_identifier_and_email_invalid",
		Message:          "error in FetchUserByIdentifierAndEmail - invalid request",
		EntityIdentifier: "user",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchUserByIdentifierAndEmailResponse{}, err

}
