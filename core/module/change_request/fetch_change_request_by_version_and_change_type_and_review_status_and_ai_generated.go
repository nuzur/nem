package change_request

import (
	"context"

	"github.com/nuzur/nem/core/module/change_request/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndAiGenerated(
	ctx context.Context,
	req types.FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndAiGeneratedRequest,
	opts ...Option,
) (types.FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndAiGeneratedResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndAiGenerated(
			ctx,
			nemdb.FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndAiGeneratedParams{

				Version: req.Version,

				ChangeType: req.ChangeType.ToInt64(),

				ReviewStatus: req.ReviewStatus.ToInt64(),

				AiGenerated: req.AiGenerated,

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_change_request_by_version_and_change_type_and_review_status_and_ai_generated",
				Message:          "error in FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndAiGenerated no order",
				EntityIdentifier: "change_request",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndAiGeneratedResponse{}, err
		}
		return types.FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndAiGeneratedResponse{
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
			ActionIdentifier: "fetch_change_request_by_version_and_change_type_and_review_status_and_ai_generated",
			Message:          "error in FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndAiGenerated - order by field not found",
			EntityIdentifier: "change_request",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndAiGeneratedResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_change_request_by_version_and_change_type_and_review_status_and_ai_generated",
				Message:          "error in FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndAiGenerated - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "change_request",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndAiGeneratedResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndAiGeneratedOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndAiGeneratedOrderedByCreatedAtASCParams{

					Version: req.Version,

					ChangeType: req.ChangeType.ToInt64(),

					ReviewStatus: req.ReviewStatus.ToInt64(),

					AiGenerated: req.AiGenerated,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_version_and_change_type_and_review_status_and_ai_generated",
					Message:          "error in FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndAiGenerated - ordered asc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndAiGeneratedResponse{}, err
			}
			return types.FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndAiGeneratedResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndAiGeneratedOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndAiGeneratedOrderedByCreatedAtDESCParams{

					Version: req.Version,

					ChangeType: req.ChangeType.ToInt64(),

					ReviewStatus: req.ReviewStatus.ToInt64(),

					AiGenerated: req.AiGenerated,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_version_and_change_type_and_review_status_and_ai_generated",
					Message:          "error in FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndAiGenerated - ordered desc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndAiGeneratedResponse{}, err
			}
			return types.FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndAiGeneratedResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndAiGeneratedOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndAiGeneratedOrderedByUpdatedAtASCParams{

					Version: req.Version,

					ChangeType: req.ChangeType.ToInt64(),

					ReviewStatus: req.ReviewStatus.ToInt64(),

					AiGenerated: req.AiGenerated,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_version_and_change_type_and_review_status_and_ai_generated",
					Message:          "error in FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndAiGenerated - ordered asc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndAiGeneratedResponse{}, err
			}
			return types.FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndAiGeneratedResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndAiGeneratedOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndAiGeneratedOrderedByUpdatedAtDESCParams{

					Version: req.Version,

					ChangeType: req.ChangeType.ToInt64(),

					ReviewStatus: req.ReviewStatus.ToInt64(),

					AiGenerated: req.AiGenerated,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_version_and_change_type_and_review_status_and_ai_generated",
					Message:          "error in FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndAiGenerated - ordered desc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndAiGeneratedResponse{}, err
			}
			return types.FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndAiGeneratedResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_change_request_by_version_and_change_type_and_review_status_and_ai_generated_invalid",
		Message:          "error in FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndAiGenerated - invalid request",
		EntityIdentifier: "change_request",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndAiGeneratedResponse{}, err

}
