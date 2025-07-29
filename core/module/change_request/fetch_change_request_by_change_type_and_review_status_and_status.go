package change_request

import (
	"context"

	"github.com/nuzur/nem/core/module/change_request/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchChangeRequestByChangeTypeAndReviewStatusAndStatus(
	ctx context.Context,
	req types.FetchChangeRequestByChangeTypeAndReviewStatusAndStatusRequest,
	opts ...Option,
) (types.FetchChangeRequestByChangeTypeAndReviewStatusAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchChangeRequestByChangeTypeAndReviewStatusAndStatus(
			ctx,
			nemdb.FetchChangeRequestByChangeTypeAndReviewStatusAndStatusParams{

				ChangeType: req.ChangeType.ToInt64(),

				ReviewStatus: req.ReviewStatus.ToInt64(),

				Status: req.Status.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_change_request_by_change_type_and_review_status_and_status",
				Message:          "error in FetchChangeRequestByChangeTypeAndReviewStatusAndStatus no order",
				EntityIdentifier: "change_request",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchChangeRequestByChangeTypeAndReviewStatusAndStatusResponse{}, err
		}
		return types.FetchChangeRequestByChangeTypeAndReviewStatusAndStatusResponse{
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
			ActionIdentifier: "fetch_change_request_by_change_type_and_review_status_and_status",
			Message:          "error in FetchChangeRequestByChangeTypeAndReviewStatusAndStatus - order by field not found",
			EntityIdentifier: "change_request",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchChangeRequestByChangeTypeAndReviewStatusAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_change_request_by_change_type_and_review_status_and_status",
				Message:          "error in FetchChangeRequestByChangeTypeAndReviewStatusAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "change_request",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchChangeRequestByChangeTypeAndReviewStatusAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchChangeRequestByChangeTypeAndReviewStatusAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchChangeRequestByChangeTypeAndReviewStatusAndStatusOrderedByCreatedAtASCParams{

					ChangeType: req.ChangeType.ToInt64(),

					ReviewStatus: req.ReviewStatus.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_change_type_and_review_status_and_status",
					Message:          "error in FetchChangeRequestByChangeTypeAndReviewStatusAndStatus - ordered asc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByChangeTypeAndReviewStatusAndStatusResponse{}, err
			}
			return types.FetchChangeRequestByChangeTypeAndReviewStatusAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchChangeRequestByChangeTypeAndReviewStatusAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchChangeRequestByChangeTypeAndReviewStatusAndStatusOrderedByCreatedAtDESCParams{

					ChangeType: req.ChangeType.ToInt64(),

					ReviewStatus: req.ReviewStatus.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_change_type_and_review_status_and_status",
					Message:          "error in FetchChangeRequestByChangeTypeAndReviewStatusAndStatus - ordered desc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByChangeTypeAndReviewStatusAndStatusResponse{}, err
			}
			return types.FetchChangeRequestByChangeTypeAndReviewStatusAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchChangeRequestByChangeTypeAndReviewStatusAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchChangeRequestByChangeTypeAndReviewStatusAndStatusOrderedByUpdatedAtASCParams{

					ChangeType: req.ChangeType.ToInt64(),

					ReviewStatus: req.ReviewStatus.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_change_type_and_review_status_and_status",
					Message:          "error in FetchChangeRequestByChangeTypeAndReviewStatusAndStatus - ordered asc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByChangeTypeAndReviewStatusAndStatusResponse{}, err
			}
			return types.FetchChangeRequestByChangeTypeAndReviewStatusAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchChangeRequestByChangeTypeAndReviewStatusAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchChangeRequestByChangeTypeAndReviewStatusAndStatusOrderedByUpdatedAtDESCParams{

					ChangeType: req.ChangeType.ToInt64(),

					ReviewStatus: req.ReviewStatus.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_change_type_and_review_status_and_status",
					Message:          "error in FetchChangeRequestByChangeTypeAndReviewStatusAndStatus - ordered desc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByChangeTypeAndReviewStatusAndStatusResponse{}, err
			}
			return types.FetchChangeRequestByChangeTypeAndReviewStatusAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_change_request_by_change_type_and_review_status_and_status_invalid",
		Message:          "error in FetchChangeRequestByChangeTypeAndReviewStatusAndStatus - invalid request",
		EntityIdentifier: "change_request",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchChangeRequestByChangeTypeAndReviewStatusAndStatusResponse{}, err

}
