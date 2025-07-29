package change_request

import (
	"context"

	"github.com/nuzur/nem/core/module/change_request/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchChangeRequestByChangeTypeAndReviewStatus(
	ctx context.Context,
	req types.FetchChangeRequestByChangeTypeAndReviewStatusRequest,
	opts ...Option,
) (types.FetchChangeRequestByChangeTypeAndReviewStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchChangeRequestByChangeTypeAndReviewStatus(
			ctx,
			nemdb.FetchChangeRequestByChangeTypeAndReviewStatusParams{

				ChangeType: req.ChangeType.ToInt64(),

				ReviewStatus: req.ReviewStatus.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_change_request_by_change_type_and_review_status",
				Message:          "error in FetchChangeRequestByChangeTypeAndReviewStatus no order",
				EntityIdentifier: "change_request",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchChangeRequestByChangeTypeAndReviewStatusResponse{}, err
		}
		return types.FetchChangeRequestByChangeTypeAndReviewStatusResponse{
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
			ActionIdentifier: "fetch_change_request_by_change_type_and_review_status",
			Message:          "error in FetchChangeRequestByChangeTypeAndReviewStatus - order by field not found",
			EntityIdentifier: "change_request",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchChangeRequestByChangeTypeAndReviewStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_change_request_by_change_type_and_review_status",
				Message:          "error in FetchChangeRequestByChangeTypeAndReviewStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "change_request",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchChangeRequestByChangeTypeAndReviewStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchChangeRequestByChangeTypeAndReviewStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchChangeRequestByChangeTypeAndReviewStatusOrderedByCreatedAtASCParams{

					ChangeType: req.ChangeType.ToInt64(),

					ReviewStatus: req.ReviewStatus.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_change_type_and_review_status",
					Message:          "error in FetchChangeRequestByChangeTypeAndReviewStatus - ordered asc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByChangeTypeAndReviewStatusResponse{}, err
			}
			return types.FetchChangeRequestByChangeTypeAndReviewStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchChangeRequestByChangeTypeAndReviewStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchChangeRequestByChangeTypeAndReviewStatusOrderedByCreatedAtDESCParams{

					ChangeType: req.ChangeType.ToInt64(),

					ReviewStatus: req.ReviewStatus.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_change_type_and_review_status",
					Message:          "error in FetchChangeRequestByChangeTypeAndReviewStatus - ordered desc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByChangeTypeAndReviewStatusResponse{}, err
			}
			return types.FetchChangeRequestByChangeTypeAndReviewStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchChangeRequestByChangeTypeAndReviewStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchChangeRequestByChangeTypeAndReviewStatusOrderedByUpdatedAtASCParams{

					ChangeType: req.ChangeType.ToInt64(),

					ReviewStatus: req.ReviewStatus.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_change_type_and_review_status",
					Message:          "error in FetchChangeRequestByChangeTypeAndReviewStatus - ordered asc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByChangeTypeAndReviewStatusResponse{}, err
			}
			return types.FetchChangeRequestByChangeTypeAndReviewStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchChangeRequestByChangeTypeAndReviewStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchChangeRequestByChangeTypeAndReviewStatusOrderedByUpdatedAtDESCParams{

					ChangeType: req.ChangeType.ToInt64(),

					ReviewStatus: req.ReviewStatus.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_change_type_and_review_status",
					Message:          "error in FetchChangeRequestByChangeTypeAndReviewStatus - ordered desc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByChangeTypeAndReviewStatusResponse{}, err
			}
			return types.FetchChangeRequestByChangeTypeAndReviewStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_change_request_by_change_type_and_review_status_invalid",
		Message:          "error in FetchChangeRequestByChangeTypeAndReviewStatus - invalid request",
		EntityIdentifier: "change_request",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchChangeRequestByChangeTypeAndReviewStatusResponse{}, err

}
