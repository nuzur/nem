package ai_usage

import (
	"context"

	"github.com/nuzur/nem/core/module/ai_usage/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchAiUsageByProviderAndStatus(
	ctx context.Context,
	req types.FetchAiUsageByProviderAndStatusRequest,
	opts ...Option,
) (types.FetchAiUsageByProviderAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchAiUsageByProviderAndStatus(
			ctx,
			nemdb.FetchAiUsageByProviderAndStatusParams{

				Provider: req.Provider.ToInt64(),

				Status: req.Status.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_ai_usage_by_provider_and_status",
				Message:          "error in FetchAiUsageByProviderAndStatus no order",
				EntityIdentifier: "ai_usage",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchAiUsageByProviderAndStatusResponse{}, err
		}
		return types.FetchAiUsageByProviderAndStatusResponse{
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
			ActionIdentifier: "fetch_ai_usage_by_provider_and_status",
			Message:          "error in FetchAiUsageByProviderAndStatus - order by field not found",
			EntityIdentifier: "ai_usage",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchAiUsageByProviderAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_ai_usage_by_provider_and_status",
				Message:          "error in FetchAiUsageByProviderAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "ai_usage",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchAiUsageByProviderAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchAiUsageByProviderAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchAiUsageByProviderAndStatusOrderedByCreatedAtASCParams{

					Provider: req.Provider.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_ai_usage_by_provider_and_status",
					Message:          "error in FetchAiUsageByProviderAndStatus - ordered asc",
					EntityIdentifier: "ai_usage",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchAiUsageByProviderAndStatusResponse{}, err
			}
			return types.FetchAiUsageByProviderAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchAiUsageByProviderAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchAiUsageByProviderAndStatusOrderedByCreatedAtDESCParams{

					Provider: req.Provider.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_ai_usage_by_provider_and_status",
					Message:          "error in FetchAiUsageByProviderAndStatus - ordered desc",
					EntityIdentifier: "ai_usage",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchAiUsageByProviderAndStatusResponse{}, err
			}
			return types.FetchAiUsageByProviderAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchAiUsageByProviderAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchAiUsageByProviderAndStatusOrderedByUpdatedAtASCParams{

					Provider: req.Provider.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_ai_usage_by_provider_and_status",
					Message:          "error in FetchAiUsageByProviderAndStatus - ordered asc",
					EntityIdentifier: "ai_usage",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchAiUsageByProviderAndStatusResponse{}, err
			}
			return types.FetchAiUsageByProviderAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchAiUsageByProviderAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchAiUsageByProviderAndStatusOrderedByUpdatedAtDESCParams{

					Provider: req.Provider.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_ai_usage_by_provider_and_status",
					Message:          "error in FetchAiUsageByProviderAndStatus - ordered desc",
					EntityIdentifier: "ai_usage",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchAiUsageByProviderAndStatusResponse{}, err
			}
			return types.FetchAiUsageByProviderAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_ai_usage_by_provider_and_status_invalid",
		Message:          "error in FetchAiUsageByProviderAndStatus - invalid request",
		EntityIdentifier: "ai_usage",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchAiUsageByProviderAndStatusResponse{}, err

}
