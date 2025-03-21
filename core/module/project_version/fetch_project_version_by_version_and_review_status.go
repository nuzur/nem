package project_version

import (
	"context"

	"github.com/nuzur/nem/core/module/project_version/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchProjectVersionByVersionAndReviewStatus(
	ctx context.Context,
	req types.FetchProjectVersionByVersionAndReviewStatusRequest,
	opts ...Option,
) (types.FetchProjectVersionByVersionAndReviewStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchProjectVersionByVersionAndReviewStatus(
			ctx,
			nemdb.FetchProjectVersionByVersionAndReviewStatusParams{

				Version: req.Version,

				ReviewStatus: req.ReviewStatus.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_project_version_by_version_and_review_status",
				Message:          "error in FetchProjectVersionByVersionAndReviewStatus no order",
				EntityIdentifier: "project_version",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchProjectVersionByVersionAndReviewStatusResponse{}, err
		}
		return types.FetchProjectVersionByVersionAndReviewStatusResponse{
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
			ActionIdentifier: "fetch_project_version_by_version_and_review_status",
			Message:          "error in FetchProjectVersionByVersionAndReviewStatus - order by field not found",
			EntityIdentifier: "project_version",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchProjectVersionByVersionAndReviewStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_project_version_by_version_and_review_status",
				Message:          "error in FetchProjectVersionByVersionAndReviewStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "project_version",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchProjectVersionByVersionAndReviewStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchProjectVersionByVersionAndReviewStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchProjectVersionByVersionAndReviewStatusOrderedByCreatedAtASCParams{

					Version: req.Version,

					ReviewStatus: req.ReviewStatus.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_version_by_version_and_review_status",
					Message:          "error in FetchProjectVersionByVersionAndReviewStatus - ordered asc",
					EntityIdentifier: "project_version",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectVersionByVersionAndReviewStatusResponse{}, err
			}
			return types.FetchProjectVersionByVersionAndReviewStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchProjectVersionByVersionAndReviewStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchProjectVersionByVersionAndReviewStatusOrderedByCreatedAtDESCParams{

					Version: req.Version,

					ReviewStatus: req.ReviewStatus.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_version_by_version_and_review_status",
					Message:          "error in FetchProjectVersionByVersionAndReviewStatus - ordered desc",
					EntityIdentifier: "project_version",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectVersionByVersionAndReviewStatusResponse{}, err
			}
			return types.FetchProjectVersionByVersionAndReviewStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchProjectVersionByVersionAndReviewStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchProjectVersionByVersionAndReviewStatusOrderedByUpdatedAtASCParams{

					Version: req.Version,

					ReviewStatus: req.ReviewStatus.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_version_by_version_and_review_status",
					Message:          "error in FetchProjectVersionByVersionAndReviewStatus - ordered asc",
					EntityIdentifier: "project_version",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectVersionByVersionAndReviewStatusResponse{}, err
			}
			return types.FetchProjectVersionByVersionAndReviewStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchProjectVersionByVersionAndReviewStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchProjectVersionByVersionAndReviewStatusOrderedByUpdatedAtDESCParams{

					Version: req.Version,

					ReviewStatus: req.ReviewStatus.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_project_version_by_version_and_review_status",
					Message:          "error in FetchProjectVersionByVersionAndReviewStatus - ordered desc",
					EntityIdentifier: "project_version",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchProjectVersionByVersionAndReviewStatusResponse{}, err
			}
			return types.FetchProjectVersionByVersionAndReviewStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_project_version_by_version_and_review_status_invalid",
		Message:          "error in FetchProjectVersionByVersionAndReviewStatus - invalid request",
		EntityIdentifier: "project_version",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchProjectVersionByVersionAndReviewStatusResponse{}, err

}
