package change_request

import (
	"context"

	"github.com/nuzur/nem/core/module/change_request/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchChangeRequestByVersionAndReviewStatusAndStatus(
	ctx context.Context,
	req types.FetchChangeRequestByVersionAndReviewStatusAndStatusRequest,
	opts ...Option,
) (types.FetchChangeRequestByVersionAndReviewStatusAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchChangeRequestByVersionAndReviewStatusAndStatus(
			ctx,
			nemdb.FetchChangeRequestByVersionAndReviewStatusAndStatusParams{

				Version: req.Version,

				ReviewStatus: req.ReviewStatus.ToInt64(),

				Status: req.Status.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_change_request_by_version_and_review_status_and_status",
				Message:          "error in FetchChangeRequestByVersionAndReviewStatusAndStatus no order",
				EntityIdentifier: "change_request",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchChangeRequestByVersionAndReviewStatusAndStatusResponse{}, err
		}
		return types.FetchChangeRequestByVersionAndReviewStatusAndStatusResponse{
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
			ActionIdentifier: "fetch_change_request_by_version_and_review_status_and_status",
			Message:          "error in FetchChangeRequestByVersionAndReviewStatusAndStatus - order by field not found",
			EntityIdentifier: "change_request",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchChangeRequestByVersionAndReviewStatusAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_change_request_by_version_and_review_status_and_status",
				Message:          "error in FetchChangeRequestByVersionAndReviewStatusAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "change_request",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchChangeRequestByVersionAndReviewStatusAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchChangeRequestByVersionAndReviewStatusAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchChangeRequestByVersionAndReviewStatusAndStatusOrderedByCreatedAtASCParams{

					Version: req.Version,

					ReviewStatus: req.ReviewStatus.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_version_and_review_status_and_status",
					Message:          "error in FetchChangeRequestByVersionAndReviewStatusAndStatus - ordered asc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByVersionAndReviewStatusAndStatusResponse{}, err
			}
			return types.FetchChangeRequestByVersionAndReviewStatusAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchChangeRequestByVersionAndReviewStatusAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchChangeRequestByVersionAndReviewStatusAndStatusOrderedByCreatedAtDESCParams{

					Version: req.Version,

					ReviewStatus: req.ReviewStatus.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_version_and_review_status_and_status",
					Message:          "error in FetchChangeRequestByVersionAndReviewStatusAndStatus - ordered desc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByVersionAndReviewStatusAndStatusResponse{}, err
			}
			return types.FetchChangeRequestByVersionAndReviewStatusAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchChangeRequestByVersionAndReviewStatusAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchChangeRequestByVersionAndReviewStatusAndStatusOrderedByUpdatedAtASCParams{

					Version: req.Version,

					ReviewStatus: req.ReviewStatus.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_version_and_review_status_and_status",
					Message:          "error in FetchChangeRequestByVersionAndReviewStatusAndStatus - ordered asc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByVersionAndReviewStatusAndStatusResponse{}, err
			}
			return types.FetchChangeRequestByVersionAndReviewStatusAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchChangeRequestByVersionAndReviewStatusAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchChangeRequestByVersionAndReviewStatusAndStatusOrderedByUpdatedAtDESCParams{

					Version: req.Version,

					ReviewStatus: req.ReviewStatus.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_version_and_review_status_and_status",
					Message:          "error in FetchChangeRequestByVersionAndReviewStatusAndStatus - ordered desc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByVersionAndReviewStatusAndStatusResponse{}, err
			}
			return types.FetchChangeRequestByVersionAndReviewStatusAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_change_request_by_version_and_review_status_and_status_invalid",
		Message:          "error in FetchChangeRequestByVersionAndReviewStatusAndStatus - invalid request",
		EntityIdentifier: "change_request",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchChangeRequestByVersionAndReviewStatusAndStatusResponse{}, err

}
