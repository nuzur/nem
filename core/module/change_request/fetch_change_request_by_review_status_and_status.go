package change_request

import (
	"context"

	"github.com/nuzur/nem/core/module/change_request/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchChangeRequestByReviewStatusAndStatus(
	ctx context.Context,
	req types.FetchChangeRequestByReviewStatusAndStatusRequest,
	opts ...Option,
) (types.FetchChangeRequestByReviewStatusAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchChangeRequestByReviewStatusAndStatus(
			ctx,
			nemdb.FetchChangeRequestByReviewStatusAndStatusParams{

				ReviewStatus: req.ReviewStatus.ToInt64(),

				Status: req.Status.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_change_request_by_review_status_and_status",
				Message:          "error in FetchChangeRequestByReviewStatusAndStatus no order",
				EntityIdentifier: "change_request",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchChangeRequestByReviewStatusAndStatusResponse{}, err
		}
		return types.FetchChangeRequestByReviewStatusAndStatusResponse{
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
			ActionIdentifier: "fetch_change_request_by_review_status_and_status",
			Message:          "error in FetchChangeRequestByReviewStatusAndStatus - order by field not found",
			EntityIdentifier: "change_request",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchChangeRequestByReviewStatusAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_change_request_by_review_status_and_status",
				Message:          "error in FetchChangeRequestByReviewStatusAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "change_request",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchChangeRequestByReviewStatusAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchChangeRequestByReviewStatusAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchChangeRequestByReviewStatusAndStatusOrderedByCreatedAtASCParams{

					ReviewStatus: req.ReviewStatus.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_review_status_and_status",
					Message:          "error in FetchChangeRequestByReviewStatusAndStatus - ordered asc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByReviewStatusAndStatusResponse{}, err
			}
			return types.FetchChangeRequestByReviewStatusAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchChangeRequestByReviewStatusAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchChangeRequestByReviewStatusAndStatusOrderedByCreatedAtDESCParams{

					ReviewStatus: req.ReviewStatus.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_review_status_and_status",
					Message:          "error in FetchChangeRequestByReviewStatusAndStatus - ordered desc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByReviewStatusAndStatusResponse{}, err
			}
			return types.FetchChangeRequestByReviewStatusAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchChangeRequestByReviewStatusAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchChangeRequestByReviewStatusAndStatusOrderedByUpdatedAtASCParams{

					ReviewStatus: req.ReviewStatus.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_review_status_and_status",
					Message:          "error in FetchChangeRequestByReviewStatusAndStatus - ordered asc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByReviewStatusAndStatusResponse{}, err
			}
			return types.FetchChangeRequestByReviewStatusAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchChangeRequestByReviewStatusAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchChangeRequestByReviewStatusAndStatusOrderedByUpdatedAtDESCParams{

					ReviewStatus: req.ReviewStatus.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_review_status_and_status",
					Message:          "error in FetchChangeRequestByReviewStatusAndStatus - ordered desc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByReviewStatusAndStatusResponse{}, err
			}
			return types.FetchChangeRequestByReviewStatusAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_change_request_by_review_status_and_status_invalid",
		Message:          "error in FetchChangeRequestByReviewStatusAndStatus - invalid request",
		EntityIdentifier: "change_request",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchChangeRequestByReviewStatusAndStatusResponse{}, err

}
