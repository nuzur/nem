package change_request

import (
	"context"

	"github.com/nuzur/nem/core/module/change_request/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchChangeRequestByReviewStatusAndAiGenerated(
	ctx context.Context,
	req types.FetchChangeRequestByReviewStatusAndAiGeneratedRequest,
	opts ...Option,
) (types.FetchChangeRequestByReviewStatusAndAiGeneratedResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchChangeRequestByReviewStatusAndAiGenerated(
			ctx,
			nemdb.FetchChangeRequestByReviewStatusAndAiGeneratedParams{

				ReviewStatus: req.ReviewStatus.ToInt64(),

				AiGenerated: req.AiGenerated,

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_change_request_by_review_status_and_ai_generated",
				Message:          "error in FetchChangeRequestByReviewStatusAndAiGenerated no order",
				EntityIdentifier: "change_request",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchChangeRequestByReviewStatusAndAiGeneratedResponse{}, err
		}
		return types.FetchChangeRequestByReviewStatusAndAiGeneratedResponse{
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
			ActionIdentifier: "fetch_change_request_by_review_status_and_ai_generated",
			Message:          "error in FetchChangeRequestByReviewStatusAndAiGenerated - order by field not found",
			EntityIdentifier: "change_request",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchChangeRequestByReviewStatusAndAiGeneratedResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_change_request_by_review_status_and_ai_generated",
				Message:          "error in FetchChangeRequestByReviewStatusAndAiGenerated - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "change_request",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchChangeRequestByReviewStatusAndAiGeneratedResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchChangeRequestByReviewStatusAndAiGeneratedOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchChangeRequestByReviewStatusAndAiGeneratedOrderedByCreatedAtASCParams{

					ReviewStatus: req.ReviewStatus.ToInt64(),

					AiGenerated: req.AiGenerated,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_review_status_and_ai_generated",
					Message:          "error in FetchChangeRequestByReviewStatusAndAiGenerated - ordered asc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByReviewStatusAndAiGeneratedResponse{}, err
			}
			return types.FetchChangeRequestByReviewStatusAndAiGeneratedResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchChangeRequestByReviewStatusAndAiGeneratedOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchChangeRequestByReviewStatusAndAiGeneratedOrderedByCreatedAtDESCParams{

					ReviewStatus: req.ReviewStatus.ToInt64(),

					AiGenerated: req.AiGenerated,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_review_status_and_ai_generated",
					Message:          "error in FetchChangeRequestByReviewStatusAndAiGenerated - ordered desc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByReviewStatusAndAiGeneratedResponse{}, err
			}
			return types.FetchChangeRequestByReviewStatusAndAiGeneratedResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchChangeRequestByReviewStatusAndAiGeneratedOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchChangeRequestByReviewStatusAndAiGeneratedOrderedByUpdatedAtASCParams{

					ReviewStatus: req.ReviewStatus.ToInt64(),

					AiGenerated: req.AiGenerated,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_review_status_and_ai_generated",
					Message:          "error in FetchChangeRequestByReviewStatusAndAiGenerated - ordered asc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByReviewStatusAndAiGeneratedResponse{}, err
			}
			return types.FetchChangeRequestByReviewStatusAndAiGeneratedResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchChangeRequestByReviewStatusAndAiGeneratedOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchChangeRequestByReviewStatusAndAiGeneratedOrderedByUpdatedAtDESCParams{

					ReviewStatus: req.ReviewStatus.ToInt64(),

					AiGenerated: req.AiGenerated,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_review_status_and_ai_generated",
					Message:          "error in FetchChangeRequestByReviewStatusAndAiGenerated - ordered desc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByReviewStatusAndAiGeneratedResponse{}, err
			}
			return types.FetchChangeRequestByReviewStatusAndAiGeneratedResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_change_request_by_review_status_and_ai_generated_invalid",
		Message:          "error in FetchChangeRequestByReviewStatusAndAiGenerated - invalid request",
		EntityIdentifier: "change_request",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchChangeRequestByReviewStatusAndAiGeneratedResponse{}, err

}
