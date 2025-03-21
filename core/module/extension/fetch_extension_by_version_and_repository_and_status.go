package extension

import (
	"context"

	"github.com/nuzur/nem/core/module/extension/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchExtensionByVersionAndRepositoryAndStatus(
	ctx context.Context,
	req types.FetchExtensionByVersionAndRepositoryAndStatusRequest,
	opts ...Option,
) (types.FetchExtensionByVersionAndRepositoryAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByVersionAndRepositoryAndStatus(
			ctx,
			nemdb.FetchExtensionByVersionAndRepositoryAndStatusParams{

				Version: req.Version,

				Repository: req.Repository,

				Status: req.Status.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_version_and_repository_and_status",
				Message:          "error in FetchExtensionByVersionAndRepositoryAndStatus no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByVersionAndRepositoryAndStatusResponse{}, err
		}
		return types.FetchExtensionByVersionAndRepositoryAndStatusResponse{
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
			ActionIdentifier: "fetch_extension_by_version_and_repository_and_status",
			Message:          "error in FetchExtensionByVersionAndRepositoryAndStatus - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByVersionAndRepositoryAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_version_and_repository_and_status",
				Message:          "error in FetchExtensionByVersionAndRepositoryAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByVersionAndRepositoryAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByVersionAndRepositoryAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByVersionAndRepositoryAndStatusOrderedByCreatedAtASCParams{

					Version: req.Version,

					Repository: req.Repository,

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_repository_and_status",
					Message:          "error in FetchExtensionByVersionAndRepositoryAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndRepositoryAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndRepositoryAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByVersionAndRepositoryAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByVersionAndRepositoryAndStatusOrderedByCreatedAtDESCParams{

					Version: req.Version,

					Repository: req.Repository,

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_repository_and_status",
					Message:          "error in FetchExtensionByVersionAndRepositoryAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndRepositoryAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndRepositoryAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByVersionAndRepositoryAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByVersionAndRepositoryAndStatusOrderedByUpdatedAtASCParams{

					Version: req.Version,

					Repository: req.Repository,

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_repository_and_status",
					Message:          "error in FetchExtensionByVersionAndRepositoryAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndRepositoryAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndRepositoryAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByVersionAndRepositoryAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByVersionAndRepositoryAndStatusOrderedByUpdatedAtDESCParams{

					Version: req.Version,

					Repository: req.Repository,

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_repository_and_status",
					Message:          "error in FetchExtensionByVersionAndRepositoryAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndRepositoryAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndRepositoryAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_version_and_repository_and_status_invalid",
		Message:          "error in FetchExtensionByVersionAndRepositoryAndStatus - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByVersionAndRepositoryAndStatusResponse{}, err

}
