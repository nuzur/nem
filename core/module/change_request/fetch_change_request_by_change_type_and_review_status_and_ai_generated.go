package change_request

import (
	"context"

	"github.com/nuzur/nem/core/module/change_request/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchChangeRequestByChangeTypeAndReviewStatusAndAiGenerated(
	ctx context.Context,
	req types.FetchChangeRequestByChangeTypeAndReviewStatusAndAiGeneratedRequest,
	opts ...Option,
) (types.FetchChangeRequestByChangeTypeAndReviewStatusAndAiGeneratedResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchChangeRequestByChangeTypeAndReviewStatusAndAiGenerated(
			ctx,
			nemdb.FetchChangeRequestByChangeTypeAndReviewStatusAndAiGeneratedParams{

				ChangeType: req.ChangeType.ToInt64(),

				ReviewStatus: req.ReviewStatus.ToInt64(),

				AiGenerated: req.AiGenerated,

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_change_request_by_change_type_and_review_status_and_ai_generated",
				Message:          "error in FetchChangeRequestByChangeTypeAndReviewStatusAndAiGenerated no order",
				EntityIdentifier: "change_request",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchChangeRequestByChangeTypeAndReviewStatusAndAiGeneratedResponse{}, err
		}
		return types.FetchChangeRequestByChangeTypeAndReviewStatusAndAiGeneratedResponse{
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
			ActionIdentifier: "fetch_change_request_by_change_type_and_review_status_and_ai_generated",
			Message:          "error in FetchChangeRequestByChangeTypeAndReviewStatusAndAiGenerated - order by field not found",
			EntityIdentifier: "change_request",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchChangeRequestByChangeTypeAndReviewStatusAndAiGeneratedResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_change_request_by_change_type_and_review_status_and_ai_generated",
				Message:          "error in FetchChangeRequestByChangeTypeAndReviewStatusAndAiGenerated - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "change_request",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchChangeRequestByChangeTypeAndReviewStatusAndAiGeneratedResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchChangeRequestByChangeTypeAndReviewStatusAndAiGeneratedOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchChangeRequestByChangeTypeAndReviewStatusAndAiGeneratedOrderedByCreatedAtASCParams{

					ChangeType: req.ChangeType.ToInt64(),

					ReviewStatus: req.ReviewStatus.ToInt64(),

					AiGenerated: req.AiGenerated,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_change_type_and_review_status_and_ai_generated",
					Message:          "error in FetchChangeRequestByChangeTypeAndReviewStatusAndAiGenerated - ordered asc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByChangeTypeAndReviewStatusAndAiGeneratedResponse{}, err
			}
			return types.FetchChangeRequestByChangeTypeAndReviewStatusAndAiGeneratedResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchChangeRequestByChangeTypeAndReviewStatusAndAiGeneratedOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchChangeRequestByChangeTypeAndReviewStatusAndAiGeneratedOrderedByCreatedAtDESCParams{

					ChangeType: req.ChangeType.ToInt64(),

					ReviewStatus: req.ReviewStatus.ToInt64(),

					AiGenerated: req.AiGenerated,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_change_type_and_review_status_and_ai_generated",
					Message:          "error in FetchChangeRequestByChangeTypeAndReviewStatusAndAiGenerated - ordered desc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByChangeTypeAndReviewStatusAndAiGeneratedResponse{}, err
			}
			return types.FetchChangeRequestByChangeTypeAndReviewStatusAndAiGeneratedResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchChangeRequestByChangeTypeAndReviewStatusAndAiGeneratedOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchChangeRequestByChangeTypeAndReviewStatusAndAiGeneratedOrderedByUpdatedAtASCParams{

					ChangeType: req.ChangeType.ToInt64(),

					ReviewStatus: req.ReviewStatus.ToInt64(),

					AiGenerated: req.AiGenerated,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_change_type_and_review_status_and_ai_generated",
					Message:          "error in FetchChangeRequestByChangeTypeAndReviewStatusAndAiGenerated - ordered asc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByChangeTypeAndReviewStatusAndAiGeneratedResponse{}, err
			}
			return types.FetchChangeRequestByChangeTypeAndReviewStatusAndAiGeneratedResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchChangeRequestByChangeTypeAndReviewStatusAndAiGeneratedOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchChangeRequestByChangeTypeAndReviewStatusAndAiGeneratedOrderedByUpdatedAtDESCParams{

					ChangeType: req.ChangeType.ToInt64(),

					ReviewStatus: req.ReviewStatus.ToInt64(),

					AiGenerated: req.AiGenerated,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_change_type_and_review_status_and_ai_generated",
					Message:          "error in FetchChangeRequestByChangeTypeAndReviewStatusAndAiGenerated - ordered desc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByChangeTypeAndReviewStatusAndAiGeneratedResponse{}, err
			}
			return types.FetchChangeRequestByChangeTypeAndReviewStatusAndAiGeneratedResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_change_request_by_change_type_and_review_status_and_ai_generated_invalid",
		Message:          "error in FetchChangeRequestByChangeTypeAndReviewStatusAndAiGenerated - invalid request",
		EntityIdentifier: "change_request",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchChangeRequestByChangeTypeAndReviewStatusAndAiGeneratedResponse{}, err

}
