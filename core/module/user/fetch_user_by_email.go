package user

import (
	"context"

	"github.com/nuzur/nem/core/module/user/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchUserByEmail(
	ctx context.Context,
	req types.FetchUserByEmailRequest,
	opts ...Option,
) (types.FetchUserByEmailResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchUserByEmail(
			ctx,
			nemdb.FetchUserByEmailParams{
				Email:  req.Email,
				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_by_email",
				Message:          "error in FetchUserByEmail no order",
				EntityIdentifier: "user",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserByEmailResponse{}, err
		}
		return types.FetchUserByEmailResponse{
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
			ActionIdentifier: "fetch_user_by_email",
			Message:          "error in FetchUserByEmail - order by field not found",
			EntityIdentifier: "user",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchUserByEmailResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_user_by_email",
				Message:          "error in FetchUserByEmail - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "user",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchUserByEmailResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserByEmailOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchUserByEmailOrderedByCreatedAtASCParams{
					Email:  req.Email,
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_by_email",
					Message:          "error in FetchUserByEmail - ordered asc",
					EntityIdentifier: "user",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserByEmailResponse{}, err
			}
			return types.FetchUserByEmailResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserByEmailOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchUserByEmailOrderedByCreatedAtDESCParams{
					Email:  req.Email,
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_by_email",
					Message:          "error in FetchUserByEmail - ordered desc",
					EntityIdentifier: "user",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserByEmailResponse{}, err
			}
			return types.FetchUserByEmailResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchUserByEmailOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchUserByEmailOrderedByUpdatedAtASCParams{
					Email:  req.Email,
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_by_email",
					Message:          "error in FetchUserByEmail - ordered asc",
					EntityIdentifier: "user",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserByEmailResponse{}, err
			}
			return types.FetchUserByEmailResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchUserByEmailOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchUserByEmailOrderedByUpdatedAtDESCParams{
					Email:  req.Email,
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_user_by_email",
					Message:          "error in FetchUserByEmail - ordered desc",
					EntityIdentifier: "user",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchUserByEmailResponse{}, err
			}
			return types.FetchUserByEmailResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_user_by_email_invalid",
		Message:          "error in FetchUserByEmail - invalid request",
		EntityIdentifier: "user",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchUserByEmailResponse{}, err

}
