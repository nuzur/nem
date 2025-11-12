package ai_usage

import (
	"context"

	"github.com/nuzur/nem/core/module/ai_usage/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchAiUsageByProvider(
	ctx context.Context,
	req types.FetchAiUsageByProviderRequest,
	opts ...Option,
) (types.FetchAiUsageByProviderResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchAiUsageByProvider(
			ctx,
			nemdb.FetchAiUsageByProviderParams{

				Provider: req.Provider.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_ai_usage_by_provider",
				Message:          "error in FetchAiUsageByProvider no order",
				EntityIdentifier: "ai_usage",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchAiUsageByProviderResponse{}, err
		}
		return types.FetchAiUsageByProviderResponse{
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
			ActionIdentifier: "fetch_ai_usage_by_provider",
			Message:          "error in FetchAiUsageByProvider - order by field not found",
			EntityIdentifier: "ai_usage",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchAiUsageByProviderResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_ai_usage_by_provider",
				Message:          "error in FetchAiUsageByProvider - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "ai_usage",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchAiUsageByProviderResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchAiUsageByProviderOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchAiUsageByProviderOrderedByCreatedAtASCParams{

					Provider: req.Provider.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_ai_usage_by_provider",
					Message:          "error in FetchAiUsageByProvider - ordered asc",
					EntityIdentifier: "ai_usage",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchAiUsageByProviderResponse{}, err
			}
			return types.FetchAiUsageByProviderResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchAiUsageByProviderOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchAiUsageByProviderOrderedByCreatedAtDESCParams{

					Provider: req.Provider.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_ai_usage_by_provider",
					Message:          "error in FetchAiUsageByProvider - ordered desc",
					EntityIdentifier: "ai_usage",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchAiUsageByProviderResponse{}, err
			}
			return types.FetchAiUsageByProviderResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchAiUsageByProviderOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchAiUsageByProviderOrderedByUpdatedAtASCParams{

					Provider: req.Provider.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_ai_usage_by_provider",
					Message:          "error in FetchAiUsageByProvider - ordered asc",
					EntityIdentifier: "ai_usage",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchAiUsageByProviderResponse{}, err
			}
			return types.FetchAiUsageByProviderResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchAiUsageByProviderOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchAiUsageByProviderOrderedByUpdatedAtDESCParams{

					Provider: req.Provider.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_ai_usage_by_provider",
					Message:          "error in FetchAiUsageByProvider - ordered desc",
					EntityIdentifier: "ai_usage",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchAiUsageByProviderResponse{}, err
			}
			return types.FetchAiUsageByProviderResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_ai_usage_by_provider_invalid",
		Message:          "error in FetchAiUsageByProvider - invalid request",
		EntityIdentifier: "ai_usage",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchAiUsageByProviderResponse{}, err

}
