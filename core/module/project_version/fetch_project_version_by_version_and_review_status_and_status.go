package project_version

import (
	"context"

	"github.com/nuzur/nem/core/module/project_version/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchProjectVersionByVersionAndReviewStatusAndStatus(
	ctx context.Context,
	req types.FetchProjectVersionByVersionAndReviewStatusAndStatusRequest,
	opts ...Option,
) (types.FetchProjectVersionByVersionAndReviewStatusAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchProjectVersionByVersionAndReviewStatusAndStatus(
			ctx,
			nemdb.FetchProjectVersionByVersionAndReviewStatusAndStatusParams{
				Version:      req.Version,
				ReviewStatus: req.ReviewStatus.ToInt64(),
				Status:       req.Status.ToInt64(),
				Offset:       req.Offset,
				Limit:        req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_project_version_by_version_and_review_status_and_status",
				Message:          "error in FetchProjectVersionByVersionAndReviewStatusAndStatus no order",
				EntityIdentifier: "project_version",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchProjectVersionByVersionAndReviewStatusAndStatusResponse{}, err
		}
		return types.FetchProjectVersionByVersionAndReviewStatusAndStatusResponse{
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
			ActionIdentifier: "fetch_project_version_by_version_and_review_status_and_status",
			Message:          "error in FetchProjectVersionByVersionAndReviewStatusAndStatus - order by field not found",
			EntityIdentifier: "project_version",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchProjectVersionByVersionAndReviewStatusAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_project_version_by_version_and_review_status_and_status",
				Message:          "error in FetchProjectVersionByVersionAndReviewStatusAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "project_version",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchProjectVersionByVersionAndReviewStatusAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchProjectVersionByVersionAndReviewStatusAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchProjectVersionByVersionAndReviewStatusAndStatusOrderedByCreatedAtASCParams{
					Version:      req.Version,
					ReviewStatus: req.ReviewStatus.ToInt64(),
					Status:       req.Status.ToInt64(),
					Offset:       req.Offset,
					Limit:        req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_version_by_version_and_review_status_and_status",
					Message:          "error in FetchProjectVersionByVersionAndReviewStatusAndStatus - ordered asc",
					EntityIdentifier: "project_version",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectVersionByVersionAndReviewStatusAndStatusResponse{}, err
			}
			return types.FetchProjectVersionByVersionAndReviewStatusAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchProjectVersionByVersionAndReviewStatusAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchProjectVersionByVersionAndReviewStatusAndStatusOrderedByCreatedAtDESCParams{
					Version:      req.Version,
					ReviewStatus: req.ReviewStatus.ToInt64(),
					Status:       req.Status.ToInt64(),
					Offset:       req.Offset,
					Limit:        req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_version_by_version_and_review_status_and_status",
					Message:          "error in FetchProjectVersionByVersionAndReviewStatusAndStatus - ordered desc",
					EntityIdentifier: "project_version",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectVersionByVersionAndReviewStatusAndStatusResponse{}, err
			}
			return types.FetchProjectVersionByVersionAndReviewStatusAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchProjectVersionByVersionAndReviewStatusAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchProjectVersionByVersionAndReviewStatusAndStatusOrderedByUpdatedAtASCParams{
					Version:      req.Version,
					ReviewStatus: req.ReviewStatus.ToInt64(),
					Status:       req.Status.ToInt64(),
					Offset:       req.Offset,
					Limit:        req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_version_by_version_and_review_status_and_status",
					Message:          "error in FetchProjectVersionByVersionAndReviewStatusAndStatus - ordered asc",
					EntityIdentifier: "project_version",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectVersionByVersionAndReviewStatusAndStatusResponse{}, err
			}
			return types.FetchProjectVersionByVersionAndReviewStatusAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchProjectVersionByVersionAndReviewStatusAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchProjectVersionByVersionAndReviewStatusAndStatusOrderedByUpdatedAtDESCParams{
					Version:      req.Version,
					ReviewStatus: req.ReviewStatus.ToInt64(),
					Status:       req.Status.ToInt64(),
					Offset:       req.Offset,
					Limit:        req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_version_by_version_and_review_status_and_status",
					Message:          "error in FetchProjectVersionByVersionAndReviewStatusAndStatus - ordered desc",
					EntityIdentifier: "project_version",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectVersionByVersionAndReviewStatusAndStatusResponse{}, err
			}
			return types.FetchProjectVersionByVersionAndReviewStatusAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_project_version_by_version_and_review_status_and_status_invalid",
		Message:          "error in FetchProjectVersionByVersionAndReviewStatusAndStatus - invalid request",
		EntityIdentifier: "project_version",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchProjectVersionByVersionAndReviewStatusAndStatusResponse{}, err

}
