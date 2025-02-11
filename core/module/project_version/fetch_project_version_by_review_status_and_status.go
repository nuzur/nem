package project_version

import (
	"context"

	"github.com/nuzur/nem/core/module/project_version/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchProjectVersionByReviewStatusAndStatus(
	ctx context.Context,
	req types.FetchProjectVersionByReviewStatusAndStatusRequest,
	opts ...Option,
) (types.FetchProjectVersionByReviewStatusAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchProjectVersionByReviewStatusAndStatus(
			ctx,
			nemdb.FetchProjectVersionByReviewStatusAndStatusParams{
				ReviewStatus: req.ReviewStatus.ToInt64(),
				Status:       req.Status.ToInt64(),
				Offset:       req.Offset,
				Limit:        req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_project_version_by_review_status_and_status",
				Message:          "error in FetchProjectVersionByReviewStatusAndStatus no order",
				EntityIdentifier: "project_version",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchProjectVersionByReviewStatusAndStatusResponse{}, err
		}
		return types.FetchProjectVersionByReviewStatusAndStatusResponse{
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
			ActionIdentifier: "fetch_project_version_by_review_status_and_status",
			Message:          "error in FetchProjectVersionByReviewStatusAndStatus - order by field not found",
			EntityIdentifier: "project_version",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchProjectVersionByReviewStatusAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_project_version_by_review_status_and_status",
				Message:          "error in FetchProjectVersionByReviewStatusAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "project_version",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchProjectVersionByReviewStatusAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchProjectVersionByReviewStatusAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchProjectVersionByReviewStatusAndStatusOrderedByCreatedAtASCParams{
					ReviewStatus: req.ReviewStatus.ToInt64(),
					Status:       req.Status.ToInt64(),
					Offset:       req.Offset,
					Limit:        req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_version_by_review_status_and_status",
					Message:          "error in FetchProjectVersionByReviewStatusAndStatus - ordered asc",
					EntityIdentifier: "project_version",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectVersionByReviewStatusAndStatusResponse{}, err
			}
			return types.FetchProjectVersionByReviewStatusAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchProjectVersionByReviewStatusAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchProjectVersionByReviewStatusAndStatusOrderedByCreatedAtDESCParams{
					ReviewStatus: req.ReviewStatus.ToInt64(),
					Status:       req.Status.ToInt64(),
					Offset:       req.Offset,
					Limit:        req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_version_by_review_status_and_status",
					Message:          "error in FetchProjectVersionByReviewStatusAndStatus - ordered desc",
					EntityIdentifier: "project_version",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectVersionByReviewStatusAndStatusResponse{}, err
			}
			return types.FetchProjectVersionByReviewStatusAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchProjectVersionByReviewStatusAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchProjectVersionByReviewStatusAndStatusOrderedByUpdatedAtASCParams{
					ReviewStatus: req.ReviewStatus.ToInt64(),
					Status:       req.Status.ToInt64(),
					Offset:       req.Offset,
					Limit:        req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_version_by_review_status_and_status",
					Message:          "error in FetchProjectVersionByReviewStatusAndStatus - ordered asc",
					EntityIdentifier: "project_version",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectVersionByReviewStatusAndStatusResponse{}, err
			}
			return types.FetchProjectVersionByReviewStatusAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchProjectVersionByReviewStatusAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchProjectVersionByReviewStatusAndStatusOrderedByUpdatedAtDESCParams{
					ReviewStatus: req.ReviewStatus.ToInt64(),
					Status:       req.Status.ToInt64(),
					Offset:       req.Offset,
					Limit:        req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_version_by_review_status_and_status",
					Message:          "error in FetchProjectVersionByReviewStatusAndStatus - ordered desc",
					EntityIdentifier: "project_version",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectVersionByReviewStatusAndStatusResponse{}, err
			}
			return types.FetchProjectVersionByReviewStatusAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_project_version_by_review_status_and_status_invalid",
		Message:          "error in FetchProjectVersionByReviewStatusAndStatus - invalid request",
		EntityIdentifier: "project_version",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchProjectVersionByReviewStatusAndStatusResponse{}, err

}
