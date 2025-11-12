package ai_usage

import (
	"context"

	"github.com/nuzur/nem/core/module/ai_usage/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchAiUsageByContextAndProviderAndStatus(
	ctx context.Context,
	req types.FetchAiUsageByContextAndProviderAndStatusRequest,
	opts ...Option,
) (types.FetchAiUsageByContextAndProviderAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchAiUsageByContextAndProviderAndStatus(
			ctx,
			nemdb.FetchAiUsageByContextAndProviderAndStatusParams{

				Context: req.Context.ToInt64(),

				Provider: req.Provider.ToInt64(),

				Status: req.Status.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_ai_usage_by_context_and_provider_and_status",
				Message:          "error in FetchAiUsageByContextAndProviderAndStatus no order",
				EntityIdentifier: "ai_usage",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchAiUsageByContextAndProviderAndStatusResponse{}, err
		}
		return types.FetchAiUsageByContextAndProviderAndStatusResponse{
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
			ActionIdentifier: "fetch_ai_usage_by_context_and_provider_and_status",
			Message:          "error in FetchAiUsageByContextAndProviderAndStatus - order by field not found",
			EntityIdentifier: "ai_usage",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchAiUsageByContextAndProviderAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_ai_usage_by_context_and_provider_and_status",
				Message:          "error in FetchAiUsageByContextAndProviderAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "ai_usage",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchAiUsageByContextAndProviderAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchAiUsageByContextAndProviderAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchAiUsageByContextAndProviderAndStatusOrderedByCreatedAtASCParams{

					Context: req.Context.ToInt64(),

					Provider: req.Provider.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_ai_usage_by_context_and_provider_and_status",
					Message:          "error in FetchAiUsageByContextAndProviderAndStatus - ordered asc",
					EntityIdentifier: "ai_usage",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchAiUsageByContextAndProviderAndStatusResponse{}, err
			}
			return types.FetchAiUsageByContextAndProviderAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchAiUsageByContextAndProviderAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchAiUsageByContextAndProviderAndStatusOrderedByCreatedAtDESCParams{

					Context: req.Context.ToInt64(),

					Provider: req.Provider.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_ai_usage_by_context_and_provider_and_status",
					Message:          "error in FetchAiUsageByContextAndProviderAndStatus - ordered desc",
					EntityIdentifier: "ai_usage",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchAiUsageByContextAndProviderAndStatusResponse{}, err
			}
			return types.FetchAiUsageByContextAndProviderAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchAiUsageByContextAndProviderAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchAiUsageByContextAndProviderAndStatusOrderedByUpdatedAtASCParams{

					Context: req.Context.ToInt64(),

					Provider: req.Provider.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_ai_usage_by_context_and_provider_and_status",
					Message:          "error in FetchAiUsageByContextAndProviderAndStatus - ordered asc",
					EntityIdentifier: "ai_usage",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchAiUsageByContextAndProviderAndStatusResponse{}, err
			}
			return types.FetchAiUsageByContextAndProviderAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchAiUsageByContextAndProviderAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchAiUsageByContextAndProviderAndStatusOrderedByUpdatedAtDESCParams{

					Context: req.Context.ToInt64(),

					Provider: req.Provider.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_ai_usage_by_context_and_provider_and_status",
					Message:          "error in FetchAiUsageByContextAndProviderAndStatus - ordered desc",
					EntityIdentifier: "ai_usage",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchAiUsageByContextAndProviderAndStatusResponse{}, err
			}
			return types.FetchAiUsageByContextAndProviderAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_ai_usage_by_context_and_provider_and_status_invalid",
		Message:          "error in FetchAiUsageByContextAndProviderAndStatus - invalid request",
		EntityIdentifier: "ai_usage",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchAiUsageByContextAndProviderAndStatusResponse{}, err

}
