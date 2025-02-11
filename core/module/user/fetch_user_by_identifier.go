package user

import (
	"context"

	"github.com/nuzur/nem/core/module/user/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchUserByIdentifier(
	ctx context.Context,
	req types.FetchUserByIdentifierRequest,
	opts ...Option,
) (types.FetchUserByIdentifierResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchUserByIdentifier(
			ctx,
			nemdb.FetchUserByIdentifierParams{
				Identifier: req.Identifier,
				Offset:     req.Offset,
				Limit:      req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_by_identifier",
				Message:          "error in FetchUserByIdentifier no order",
				EntityIdentifier: "user",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserByIdentifierResponse{}, err
		}
		return types.FetchUserByIdentifierResponse{
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
			ActionIdentifier: "fetch_user_by_identifier",
			Message:          "error in FetchUserByIdentifier - order by field not found",
			EntityIdentifier: "user",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchUserByIdentifierResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_by_identifier",
				Message:          "error in FetchUserByIdentifier - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "user",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserByIdentifierResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserByIdentifierOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchUserByIdentifierOrderedByCreatedAtASCParams{
					Identifier: req.Identifier,
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_by_identifier",
					Message:          "error in FetchUserByIdentifier - ordered asc",
					EntityIdentifier: "user",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserByIdentifierResponse{}, err
			}
			return types.FetchUserByIdentifierResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserByIdentifierOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchUserByIdentifierOrderedByCreatedAtDESCParams{
					Identifier: req.Identifier,
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_by_identifier",
					Message:          "error in FetchUserByIdentifier - ordered desc",
					EntityIdentifier: "user",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserByIdentifierResponse{}, err
			}
			return types.FetchUserByIdentifierResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserByIdentifierOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchUserByIdentifierOrderedByUpdatedAtASCParams{
					Identifier: req.Identifier,
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_by_identifier",
					Message:          "error in FetchUserByIdentifier - ordered asc",
					EntityIdentifier: "user",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserByIdentifierResponse{}, err
			}
			return types.FetchUserByIdentifierResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserByIdentifierOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchUserByIdentifierOrderedByUpdatedAtDESCParams{
					Identifier: req.Identifier,
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_by_identifier",
					Message:          "error in FetchUserByIdentifier - ordered desc",
					EntityIdentifier: "user",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserByIdentifierResponse{}, err
			}
			return types.FetchUserByIdentifierResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_user_by_identifier_invalid",
		Message:          "error in FetchUserByIdentifier - invalid request",
		EntityIdentifier: "user",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchUserByIdentifierResponse{}, err

}
