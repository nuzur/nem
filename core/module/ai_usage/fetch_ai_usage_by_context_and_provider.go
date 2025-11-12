package ai_usage

import (
	"context"

	"github.com/nuzur/nem/core/module/ai_usage/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchAiUsageByContextAndProvider(
	ctx context.Context,
	req types.FetchAiUsageByContextAndProviderRequest,
	opts ...Option,
) (types.FetchAiUsageByContextAndProviderResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchAiUsageByContextAndProvider(
			ctx,
			nemdb.FetchAiUsageByContextAndProviderParams{

				Context: req.Context.ToInt64(),

				Provider: req.Provider.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_ai_usage_by_context_and_provider",
				Message:          "error in FetchAiUsageByContextAndProvider no order",
				EntityIdentifier: "ai_usage",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchAiUsageByContextAndProviderResponse{}, err
		}
		return types.FetchAiUsageByContextAndProviderResponse{
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
			ActionIdentifier: "fetch_ai_usage_by_context_and_provider",
			Message:          "error in FetchAiUsageByContextAndProvider - order by field not found",
			EntityIdentifier: "ai_usage",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchAiUsageByContextAndProviderResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_ai_usage_by_context_and_provider",
				Message:          "error in FetchAiUsageByContextAndProvider - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "ai_usage",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchAiUsageByContextAndProviderResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchAiUsageByContextAndProviderOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchAiUsageByContextAndProviderOrderedByCreatedAtASCParams{

					Context: req.Context.ToInt64(),

					Provider: req.Provider.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_ai_usage_by_context_and_provider",
					Message:          "error in FetchAiUsageByContextAndProvider - ordered asc",
					EntityIdentifier: "ai_usage",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchAiUsageByContextAndProviderResponse{}, err
			}
			return types.FetchAiUsageByContextAndProviderResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchAiUsageByContextAndProviderOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchAiUsageByContextAndProviderOrderedByCreatedAtDESCParams{

					Context: req.Context.ToInt64(),

					Provider: req.Provider.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_ai_usage_by_context_and_provider",
					Message:          "error in FetchAiUsageByContextAndProvider - ordered desc",
					EntityIdentifier: "ai_usage",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchAiUsageByContextAndProviderResponse{}, err
			}
			return types.FetchAiUsageByContextAndProviderResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchAiUsageByContextAndProviderOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchAiUsageByContextAndProviderOrderedByUpdatedAtASCParams{

					Context: req.Context.ToInt64(),

					Provider: req.Provider.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_ai_usage_by_context_and_provider",
					Message:          "error in FetchAiUsageByContextAndProvider - ordered asc",
					EntityIdentifier: "ai_usage",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchAiUsageByContextAndProviderResponse{}, err
			}
			return types.FetchAiUsageByContextAndProviderResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchAiUsageByContextAndProviderOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchAiUsageByContextAndProviderOrderedByUpdatedAtDESCParams{

					Context: req.Context.ToInt64(),

					Provider: req.Provider.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_ai_usage_by_context_and_provider",
					Message:          "error in FetchAiUsageByContextAndProvider - ordered desc",
					EntityIdentifier: "ai_usage",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchAiUsageByContextAndProviderResponse{}, err
			}
			return types.FetchAiUsageByContextAndProviderResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_ai_usage_by_context_and_provider_invalid",
		Message:          "error in FetchAiUsageByContextAndProvider - invalid request",
		EntityIdentifier: "ai_usage",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchAiUsageByContextAndProviderResponse{}, err

}
