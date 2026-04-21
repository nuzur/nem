package change_request

import (
	"context"

	"github.com/nuzur/nem/core/module/change_request/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchChangeRequestByReviewStatusAndStatusAndAiGenerated(
	ctx context.Context,
	req types.FetchChangeRequestByReviewStatusAndStatusAndAiGeneratedRequest,
	opts ...Option,
) (types.FetchChangeRequestByReviewStatusAndStatusAndAiGeneratedResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchChangeRequestByReviewStatusAndStatusAndAiGenerated(
			ctx,
			nemdb.FetchChangeRequestByReviewStatusAndStatusAndAiGeneratedParams{

				ReviewStatus: req.ReviewStatus.ToInt64(),

				Status: req.Status.ToInt64(),

				AiGenerated: req.AiGenerated,

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_change_request_by_review_status_and_status_and_ai_generated",
				Message:          "error in FetchChangeRequestByReviewStatusAndStatusAndAiGenerated no order",
				EntityIdentifier: "change_request",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchChangeRequestByReviewStatusAndStatusAndAiGeneratedResponse{}, err
		}
		return types.FetchChangeRequestByReviewStatusAndStatusAndAiGeneratedResponse{
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
			ActionIdentifier: "fetch_change_request_by_review_status_and_status_and_ai_generated",
			Message:          "error in FetchChangeRequestByReviewStatusAndStatusAndAiGenerated - order by field not found",
			EntityIdentifier: "change_request",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchChangeRequestByReviewStatusAndStatusAndAiGeneratedResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_change_request_by_review_status_and_status_and_ai_generated",
				Message:          "error in FetchChangeRequestByReviewStatusAndStatusAndAiGenerated - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "change_request",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchChangeRequestByReviewStatusAndStatusAndAiGeneratedResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchChangeRequestByReviewStatusAndStatusAndAiGeneratedOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchChangeRequestByReviewStatusAndStatusAndAiGeneratedOrderedByCreatedAtASCParams{

					ReviewStatus: req.ReviewStatus.ToInt64(),

					Status: req.Status.ToInt64(),

					AiGenerated: req.AiGenerated,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_review_status_and_status_and_ai_generated",
					Message:          "error in FetchChangeRequestByReviewStatusAndStatusAndAiGenerated - ordered asc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByReviewStatusAndStatusAndAiGeneratedResponse{}, err
			}
			return types.FetchChangeRequestByReviewStatusAndStatusAndAiGeneratedResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchChangeRequestByReviewStatusAndStatusAndAiGeneratedOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchChangeRequestByReviewStatusAndStatusAndAiGeneratedOrderedByCreatedAtDESCParams{

					ReviewStatus: req.ReviewStatus.ToInt64(),

					Status: req.Status.ToInt64(),

					AiGenerated: req.AiGenerated,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_review_status_and_status_and_ai_generated",
					Message:          "error in FetchChangeRequestByReviewStatusAndStatusAndAiGenerated - ordered desc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByReviewStatusAndStatusAndAiGeneratedResponse{}, err
			}
			return types.FetchChangeRequestByReviewStatusAndStatusAndAiGeneratedResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchChangeRequestByReviewStatusAndStatusAndAiGeneratedOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchChangeRequestByReviewStatusAndStatusAndAiGeneratedOrderedByUpdatedAtASCParams{

					ReviewStatus: req.ReviewStatus.ToInt64(),

					Status: req.Status.ToInt64(),

					AiGenerated: req.AiGenerated,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_review_status_and_status_and_ai_generated",
					Message:          "error in FetchChangeRequestByReviewStatusAndStatusAndAiGenerated - ordered asc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByReviewStatusAndStatusAndAiGeneratedResponse{}, err
			}
			return types.FetchChangeRequestByReviewStatusAndStatusAndAiGeneratedResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchChangeRequestByReviewStatusAndStatusAndAiGeneratedOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchChangeRequestByReviewStatusAndStatusAndAiGeneratedOrderedByUpdatedAtDESCParams{

					ReviewStatus: req.ReviewStatus.ToInt64(),

					Status: req.Status.ToInt64(),

					AiGenerated: req.AiGenerated,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_review_status_and_status_and_ai_generated",
					Message:          "error in FetchChangeRequestByReviewStatusAndStatusAndAiGenerated - ordered desc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByReviewStatusAndStatusAndAiGeneratedResponse{}, err
			}
			return types.FetchChangeRequestByReviewStatusAndStatusAndAiGeneratedResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_change_request_by_review_status_and_status_and_ai_generated_invalid",
		Message:          "error in FetchChangeRequestByReviewStatusAndStatusAndAiGenerated - invalid request",
		EntityIdentifier: "change_request",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchChangeRequestByReviewStatusAndStatusAndAiGeneratedResponse{}, err

}
