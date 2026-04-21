package change_request

import (
	"context"

	"github.com/nuzur/nem/core/module/change_request/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchChangeRequestByVersionAndReviewStatusAndStatusAndAiGenerated(
	ctx context.Context,
	req types.FetchChangeRequestByVersionAndReviewStatusAndStatusAndAiGeneratedRequest,
	opts ...Option,
) (types.FetchChangeRequestByVersionAndReviewStatusAndStatusAndAiGeneratedResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchChangeRequestByVersionAndReviewStatusAndStatusAndAiGenerated(
			ctx,
			nemdb.FetchChangeRequestByVersionAndReviewStatusAndStatusAndAiGeneratedParams{

				Version: req.Version,

				ReviewStatus: req.ReviewStatus.ToInt64(),

				Status: req.Status.ToInt64(),

				AiGenerated: req.AiGenerated,

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_change_request_by_version_and_review_status_and_status_and_ai_generated",
				Message:          "error in FetchChangeRequestByVersionAndReviewStatusAndStatusAndAiGenerated no order",
				EntityIdentifier: "change_request",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchChangeRequestByVersionAndReviewStatusAndStatusAndAiGeneratedResponse{}, err
		}
		return types.FetchChangeRequestByVersionAndReviewStatusAndStatusAndAiGeneratedResponse{
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
			ActionIdentifier: "fetch_change_request_by_version_and_review_status_and_status_and_ai_generated",
			Message:          "error in FetchChangeRequestByVersionAndReviewStatusAndStatusAndAiGenerated - order by field not found",
			EntityIdentifier: "change_request",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchChangeRequestByVersionAndReviewStatusAndStatusAndAiGeneratedResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_change_request_by_version_and_review_status_and_status_and_ai_generated",
				Message:          "error in FetchChangeRequestByVersionAndReviewStatusAndStatusAndAiGenerated - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "change_request",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchChangeRequestByVersionAndReviewStatusAndStatusAndAiGeneratedResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchChangeRequestByVersionAndReviewStatusAndStatusAndAiGeneratedOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchChangeRequestByVersionAndReviewStatusAndStatusAndAiGeneratedOrderedByCreatedAtASCParams{

					Version: req.Version,

					ReviewStatus: req.ReviewStatus.ToInt64(),

					Status: req.Status.ToInt64(),

					AiGenerated: req.AiGenerated,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_version_and_review_status_and_status_and_ai_generated",
					Message:          "error in FetchChangeRequestByVersionAndReviewStatusAndStatusAndAiGenerated - ordered asc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByVersionAndReviewStatusAndStatusAndAiGeneratedResponse{}, err
			}
			return types.FetchChangeRequestByVersionAndReviewStatusAndStatusAndAiGeneratedResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchChangeRequestByVersionAndReviewStatusAndStatusAndAiGeneratedOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchChangeRequestByVersionAndReviewStatusAndStatusAndAiGeneratedOrderedByCreatedAtDESCParams{

					Version: req.Version,

					ReviewStatus: req.ReviewStatus.ToInt64(),

					Status: req.Status.ToInt64(),

					AiGenerated: req.AiGenerated,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_version_and_review_status_and_status_and_ai_generated",
					Message:          "error in FetchChangeRequestByVersionAndReviewStatusAndStatusAndAiGenerated - ordered desc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByVersionAndReviewStatusAndStatusAndAiGeneratedResponse{}, err
			}
			return types.FetchChangeRequestByVersionAndReviewStatusAndStatusAndAiGeneratedResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchChangeRequestByVersionAndReviewStatusAndStatusAndAiGeneratedOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchChangeRequestByVersionAndReviewStatusAndStatusAndAiGeneratedOrderedByUpdatedAtASCParams{

					Version: req.Version,

					ReviewStatus: req.ReviewStatus.ToInt64(),

					Status: req.Status.ToInt64(),

					AiGenerated: req.AiGenerated,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_version_and_review_status_and_status_and_ai_generated",
					Message:          "error in FetchChangeRequestByVersionAndReviewStatusAndStatusAndAiGenerated - ordered asc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByVersionAndReviewStatusAndStatusAndAiGeneratedResponse{}, err
			}
			return types.FetchChangeRequestByVersionAndReviewStatusAndStatusAndAiGeneratedResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchChangeRequestByVersionAndReviewStatusAndStatusAndAiGeneratedOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchChangeRequestByVersionAndReviewStatusAndStatusAndAiGeneratedOrderedByUpdatedAtDESCParams{

					Version: req.Version,

					ReviewStatus: req.ReviewStatus.ToInt64(),

					Status: req.Status.ToInt64(),

					AiGenerated: req.AiGenerated,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_version_and_review_status_and_status_and_ai_generated",
					Message:          "error in FetchChangeRequestByVersionAndReviewStatusAndStatusAndAiGenerated - ordered desc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByVersionAndReviewStatusAndStatusAndAiGeneratedResponse{}, err
			}
			return types.FetchChangeRequestByVersionAndReviewStatusAndStatusAndAiGeneratedResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_change_request_by_version_and_review_status_and_status_and_ai_generated_invalid",
		Message:          "error in FetchChangeRequestByVersionAndReviewStatusAndStatusAndAiGenerated - invalid request",
		EntityIdentifier: "change_request",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchChangeRequestByVersionAndReviewStatusAndStatusAndAiGeneratedResponse{}, err

}
