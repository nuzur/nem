package change_request

import (
	"context"

	"github.com/nuzur/nem/core/module/change_request/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchChangeRequestByStatusAndAiGenerated(
	ctx context.Context,
	req types.FetchChangeRequestByStatusAndAiGeneratedRequest,
	opts ...Option,
) (types.FetchChangeRequestByStatusAndAiGeneratedResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchChangeRequestByStatusAndAiGenerated(
			ctx,
			nemdb.FetchChangeRequestByStatusAndAiGeneratedParams{

				Status: req.Status.ToInt64(),

				AiGenerated: req.AiGenerated,

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_change_request_by_status_and_ai_generated",
				Message:          "error in FetchChangeRequestByStatusAndAiGenerated no order",
				EntityIdentifier: "change_request",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchChangeRequestByStatusAndAiGeneratedResponse{}, err
		}
		return types.FetchChangeRequestByStatusAndAiGeneratedResponse{
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
			ActionIdentifier: "fetch_change_request_by_status_and_ai_generated",
			Message:          "error in FetchChangeRequestByStatusAndAiGenerated - order by field not found",
			EntityIdentifier: "change_request",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchChangeRequestByStatusAndAiGeneratedResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_change_request_by_status_and_ai_generated",
				Message:          "error in FetchChangeRequestByStatusAndAiGenerated - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "change_request",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchChangeRequestByStatusAndAiGeneratedResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchChangeRequestByStatusAndAiGeneratedOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchChangeRequestByStatusAndAiGeneratedOrderedByCreatedAtASCParams{

					Status: req.Status.ToInt64(),

					AiGenerated: req.AiGenerated,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_status_and_ai_generated",
					Message:          "error in FetchChangeRequestByStatusAndAiGenerated - ordered asc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByStatusAndAiGeneratedResponse{}, err
			}
			return types.FetchChangeRequestByStatusAndAiGeneratedResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchChangeRequestByStatusAndAiGeneratedOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchChangeRequestByStatusAndAiGeneratedOrderedByCreatedAtDESCParams{

					Status: req.Status.ToInt64(),

					AiGenerated: req.AiGenerated,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_status_and_ai_generated",
					Message:          "error in FetchChangeRequestByStatusAndAiGenerated - ordered desc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByStatusAndAiGeneratedResponse{}, err
			}
			return types.FetchChangeRequestByStatusAndAiGeneratedResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchChangeRequestByStatusAndAiGeneratedOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchChangeRequestByStatusAndAiGeneratedOrderedByUpdatedAtASCParams{

					Status: req.Status.ToInt64(),

					AiGenerated: req.AiGenerated,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_status_and_ai_generated",
					Message:          "error in FetchChangeRequestByStatusAndAiGenerated - ordered asc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByStatusAndAiGeneratedResponse{}, err
			}
			return types.FetchChangeRequestByStatusAndAiGeneratedResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchChangeRequestByStatusAndAiGeneratedOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchChangeRequestByStatusAndAiGeneratedOrderedByUpdatedAtDESCParams{

					Status: req.Status.ToInt64(),

					AiGenerated: req.AiGenerated,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_status_and_ai_generated",
					Message:          "error in FetchChangeRequestByStatusAndAiGenerated - ordered desc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByStatusAndAiGeneratedResponse{}, err
			}
			return types.FetchChangeRequestByStatusAndAiGeneratedResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_change_request_by_status_and_ai_generated_invalid",
		Message:          "error in FetchChangeRequestByStatusAndAiGenerated - invalid request",
		EntityIdentifier: "change_request",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchChangeRequestByStatusAndAiGeneratedResponse{}, err

}
