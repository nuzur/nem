package project_version

import (
	"context"

	"github.com/nuzur/nem/core/module/project_version/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchProjectVersionByReviewStatus(
	ctx context.Context,
	req types.FetchProjectVersionByReviewStatusRequest,
	opts ...Option,
) (types.FetchProjectVersionByReviewStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchProjectVersionByReviewStatus(
			ctx,
			nemdb.FetchProjectVersionByReviewStatusParams{

				ReviewStatus: req.ReviewStatus.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_project_version_by_review_status",
				Message:          "error in FetchProjectVersionByReviewStatus no order",
				EntityIdentifier: "project_version",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchProjectVersionByReviewStatusResponse{}, err
		}
		return types.FetchProjectVersionByReviewStatusResponse{
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
			ActionIdentifier: "fetch_project_version_by_review_status",
			Message:          "error in FetchProjectVersionByReviewStatus - order by field not found",
			EntityIdentifier: "project_version",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchProjectVersionByReviewStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_project_version_by_review_status",
				Message:          "error in FetchProjectVersionByReviewStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "project_version",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchProjectVersionByReviewStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchProjectVersionByReviewStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchProjectVersionByReviewStatusOrderedByCreatedAtASCParams{

					ReviewStatus: req.ReviewStatus.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_version_by_review_status",
					Message:          "error in FetchProjectVersionByReviewStatus - ordered asc",
					EntityIdentifier: "project_version",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectVersionByReviewStatusResponse{}, err
			}
			return types.FetchProjectVersionByReviewStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchProjectVersionByReviewStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchProjectVersionByReviewStatusOrderedByCreatedAtDESCParams{

					ReviewStatus: req.ReviewStatus.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_version_by_review_status",
					Message:          "error in FetchProjectVersionByReviewStatus - ordered desc",
					EntityIdentifier: "project_version",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectVersionByReviewStatusResponse{}, err
			}
			return types.FetchProjectVersionByReviewStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchProjectVersionByReviewStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchProjectVersionByReviewStatusOrderedByUpdatedAtASCParams{

					ReviewStatus: req.ReviewStatus.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_version_by_review_status",
					Message:          "error in FetchProjectVersionByReviewStatus - ordered asc",
					EntityIdentifier: "project_version",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectVersionByReviewStatusResponse{}, err
			}
			return types.FetchProjectVersionByReviewStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchProjectVersionByReviewStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchProjectVersionByReviewStatusOrderedByUpdatedAtDESCParams{

					ReviewStatus: req.ReviewStatus.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_version_by_review_status",
					Message:          "error in FetchProjectVersionByReviewStatus - ordered desc",
					EntityIdentifier: "project_version",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectVersionByReviewStatusResponse{}, err
			}
			return types.FetchProjectVersionByReviewStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_project_version_by_review_status_invalid",
		Message:          "error in FetchProjectVersionByReviewStatus - invalid request",
		EntityIdentifier: "project_version",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchProjectVersionByReviewStatusResponse{}, err

}
