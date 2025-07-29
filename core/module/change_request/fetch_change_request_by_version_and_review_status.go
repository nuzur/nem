package change_request

import (
	"context"

	"github.com/nuzur/nem/core/module/change_request/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchChangeRequestByVersionAndReviewStatus(
	ctx context.Context,
	req types.FetchChangeRequestByVersionAndReviewStatusRequest,
	opts ...Option,
) (types.FetchChangeRequestByVersionAndReviewStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchChangeRequestByVersionAndReviewStatus(
			ctx,
			nemdb.FetchChangeRequestByVersionAndReviewStatusParams{

				Version: req.Version,

				ReviewStatus: req.ReviewStatus.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_change_request_by_version_and_review_status",
				Message:          "error in FetchChangeRequestByVersionAndReviewStatus no order",
				EntityIdentifier: "change_request",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchChangeRequestByVersionAndReviewStatusResponse{}, err
		}
		return types.FetchChangeRequestByVersionAndReviewStatusResponse{
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
			ActionIdentifier: "fetch_change_request_by_version_and_review_status",
			Message:          "error in FetchChangeRequestByVersionAndReviewStatus - order by field not found",
			EntityIdentifier: "change_request",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchChangeRequestByVersionAndReviewStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_change_request_by_version_and_review_status",
				Message:          "error in FetchChangeRequestByVersionAndReviewStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "change_request",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchChangeRequestByVersionAndReviewStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchChangeRequestByVersionAndReviewStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchChangeRequestByVersionAndReviewStatusOrderedByCreatedAtASCParams{

					Version: req.Version,

					ReviewStatus: req.ReviewStatus.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_version_and_review_status",
					Message:          "error in FetchChangeRequestByVersionAndReviewStatus - ordered asc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByVersionAndReviewStatusResponse{}, err
			}
			return types.FetchChangeRequestByVersionAndReviewStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchChangeRequestByVersionAndReviewStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchChangeRequestByVersionAndReviewStatusOrderedByCreatedAtDESCParams{

					Version: req.Version,

					ReviewStatus: req.ReviewStatus.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_version_and_review_status",
					Message:          "error in FetchChangeRequestByVersionAndReviewStatus - ordered desc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByVersionAndReviewStatusResponse{}, err
			}
			return types.FetchChangeRequestByVersionAndReviewStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchChangeRequestByVersionAndReviewStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchChangeRequestByVersionAndReviewStatusOrderedByUpdatedAtASCParams{

					Version: req.Version,

					ReviewStatus: req.ReviewStatus.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_version_and_review_status",
					Message:          "error in FetchChangeRequestByVersionAndReviewStatus - ordered asc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByVersionAndReviewStatusResponse{}, err
			}
			return types.FetchChangeRequestByVersionAndReviewStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchChangeRequestByVersionAndReviewStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchChangeRequestByVersionAndReviewStatusOrderedByUpdatedAtDESCParams{

					Version: req.Version,

					ReviewStatus: req.ReviewStatus.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_version_and_review_status",
					Message:          "error in FetchChangeRequestByVersionAndReviewStatus - ordered desc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByVersionAndReviewStatusResponse{}, err
			}
			return types.FetchChangeRequestByVersionAndReviewStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_change_request_by_version_and_review_status_invalid",
		Message:          "error in FetchChangeRequestByVersionAndReviewStatus - invalid request",
		EntityIdentifier: "change_request",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchChangeRequestByVersionAndReviewStatusResponse{}, err

}
