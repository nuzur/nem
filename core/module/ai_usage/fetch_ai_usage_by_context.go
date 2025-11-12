package ai_usage

import (
	"context"

	"github.com/nuzur/nem/core/module/ai_usage/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchAiUsageByContext(
	ctx context.Context,
	req types.FetchAiUsageByContextRequest,
	opts ...Option,
) (types.FetchAiUsageByContextResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchAiUsageByContext(
			ctx,
			nemdb.FetchAiUsageByContextParams{

				Context: req.Context.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_ai_usage_by_context",
				Message:          "error in FetchAiUsageByContext no order",
				EntityIdentifier: "ai_usage",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchAiUsageByContextResponse{}, err
		}
		return types.FetchAiUsageByContextResponse{
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
			ActionIdentifier: "fetch_ai_usage_by_context",
			Message:          "error in FetchAiUsageByContext - order by field not found",
			EntityIdentifier: "ai_usage",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchAiUsageByContextResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_ai_usage_by_context",
				Message:          "error in FetchAiUsageByContext - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "ai_usage",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchAiUsageByContextResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchAiUsageByContextOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchAiUsageByContextOrderedByCreatedAtASCParams{

					Context: req.Context.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_ai_usage_by_context",
					Message:          "error in FetchAiUsageByContext - ordered asc",
					EntityIdentifier: "ai_usage",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchAiUsageByContextResponse{}, err
			}
			return types.FetchAiUsageByContextResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchAiUsageByContextOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchAiUsageByContextOrderedByCreatedAtDESCParams{

					Context: req.Context.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_ai_usage_by_context",
					Message:          "error in FetchAiUsageByContext - ordered desc",
					EntityIdentifier: "ai_usage",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchAiUsageByContextResponse{}, err
			}
			return types.FetchAiUsageByContextResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchAiUsageByContextOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchAiUsageByContextOrderedByUpdatedAtASCParams{

					Context: req.Context.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_ai_usage_by_context",
					Message:          "error in FetchAiUsageByContext - ordered asc",
					EntityIdentifier: "ai_usage",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchAiUsageByContextResponse{}, err
			}
			return types.FetchAiUsageByContextResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchAiUsageByContextOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchAiUsageByContextOrderedByUpdatedAtDESCParams{

					Context: req.Context.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_ai_usage_by_context",
					Message:          "error in FetchAiUsageByContext - ordered desc",
					EntityIdentifier: "ai_usage",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchAiUsageByContextResponse{}, err
			}
			return types.FetchAiUsageByContextResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_ai_usage_by_context_invalid",
		Message:          "error in FetchAiUsageByContext - invalid request",
		EntityIdentifier: "ai_usage",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchAiUsageByContextResponse{}, err

}
