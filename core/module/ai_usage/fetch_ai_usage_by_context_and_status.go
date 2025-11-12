package ai_usage

import (
	"context"

	"github.com/nuzur/nem/core/module/ai_usage/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchAiUsageByContextAndStatus(
	ctx context.Context,
	req types.FetchAiUsageByContextAndStatusRequest,
	opts ...Option,
) (types.FetchAiUsageByContextAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchAiUsageByContextAndStatus(
			ctx,
			nemdb.FetchAiUsageByContextAndStatusParams{

				Context: req.Context.ToInt64(),

				Status: req.Status.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_ai_usage_by_context_and_status",
				Message:          "error in FetchAiUsageByContextAndStatus no order",
				EntityIdentifier: "ai_usage",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchAiUsageByContextAndStatusResponse{}, err
		}
		return types.FetchAiUsageByContextAndStatusResponse{
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
			ActionIdentifier: "fetch_ai_usage_by_context_and_status",
			Message:          "error in FetchAiUsageByContextAndStatus - order by field not found",
			EntityIdentifier: "ai_usage",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchAiUsageByContextAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_ai_usage_by_context_and_status",
				Message:          "error in FetchAiUsageByContextAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "ai_usage",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchAiUsageByContextAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchAiUsageByContextAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchAiUsageByContextAndStatusOrderedByCreatedAtASCParams{

					Context: req.Context.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_ai_usage_by_context_and_status",
					Message:          "error in FetchAiUsageByContextAndStatus - ordered asc",
					EntityIdentifier: "ai_usage",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchAiUsageByContextAndStatusResponse{}, err
			}
			return types.FetchAiUsageByContextAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchAiUsageByContextAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchAiUsageByContextAndStatusOrderedByCreatedAtDESCParams{

					Context: req.Context.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_ai_usage_by_context_and_status",
					Message:          "error in FetchAiUsageByContextAndStatus - ordered desc",
					EntityIdentifier: "ai_usage",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchAiUsageByContextAndStatusResponse{}, err
			}
			return types.FetchAiUsageByContextAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchAiUsageByContextAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchAiUsageByContextAndStatusOrderedByUpdatedAtASCParams{

					Context: req.Context.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_ai_usage_by_context_and_status",
					Message:          "error in FetchAiUsageByContextAndStatus - ordered asc",
					EntityIdentifier: "ai_usage",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchAiUsageByContextAndStatusResponse{}, err
			}
			return types.FetchAiUsageByContextAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchAiUsageByContextAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchAiUsageByContextAndStatusOrderedByUpdatedAtDESCParams{

					Context: req.Context.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_ai_usage_by_context_and_status",
					Message:          "error in FetchAiUsageByContextAndStatus - ordered desc",
					EntityIdentifier: "ai_usage",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchAiUsageByContextAndStatusResponse{}, err
			}
			return types.FetchAiUsageByContextAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_ai_usage_by_context_and_status_invalid",
		Message:          "error in FetchAiUsageByContextAndStatus - invalid request",
		EntityIdentifier: "ai_usage",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchAiUsageByContextAndStatusResponse{}, err

}
