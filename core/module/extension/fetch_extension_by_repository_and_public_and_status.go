package extension

import (
	"context"

	"github.com/nuzur/nem/core/module/extension/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchExtensionByRepositoryAndPublicAndStatus(
	ctx context.Context,
	req types.FetchExtensionByRepositoryAndPublicAndStatusRequest,
	opts ...Option,
) (types.FetchExtensionByRepositoryAndPublicAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByRepositoryAndPublicAndStatus(
			ctx,
			nemdb.FetchExtensionByRepositoryAndPublicAndStatusParams{

				Repository: req.Repository,

				Public: req.Public,

				Status: req.Status.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_repository_and_public_and_status",
				Message:          "error in FetchExtensionByRepositoryAndPublicAndStatus no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByRepositoryAndPublicAndStatusResponse{}, err
		}
		return types.FetchExtensionByRepositoryAndPublicAndStatusResponse{
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
			ActionIdentifier: "fetch_extension_by_repository_and_public_and_status",
			Message:          "error in FetchExtensionByRepositoryAndPublicAndStatus - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByRepositoryAndPublicAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_repository_and_public_and_status",
				Message:          "error in FetchExtensionByRepositoryAndPublicAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByRepositoryAndPublicAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByRepositoryAndPublicAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByRepositoryAndPublicAndStatusOrderedByCreatedAtASCParams{

					Repository: req.Repository,

					Public: req.Public,

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_repository_and_public_and_status",
					Message:          "error in FetchExtensionByRepositoryAndPublicAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByRepositoryAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByRepositoryAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByRepositoryAndPublicAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByRepositoryAndPublicAndStatusOrderedByCreatedAtDESCParams{

					Repository: req.Repository,

					Public: req.Public,

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_repository_and_public_and_status",
					Message:          "error in FetchExtensionByRepositoryAndPublicAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByRepositoryAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByRepositoryAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByRepositoryAndPublicAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByRepositoryAndPublicAndStatusOrderedByUpdatedAtASCParams{

					Repository: req.Repository,

					Public: req.Public,

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_repository_and_public_and_status",
					Message:          "error in FetchExtensionByRepositoryAndPublicAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByRepositoryAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByRepositoryAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByRepositoryAndPublicAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByRepositoryAndPublicAndStatusOrderedByUpdatedAtDESCParams{

					Repository: req.Repository,

					Public: req.Public,

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_repository_and_public_and_status",
					Message:          "error in FetchExtensionByRepositoryAndPublicAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByRepositoryAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByRepositoryAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_repository_and_public_and_status_invalid",
		Message:          "error in FetchExtensionByRepositoryAndPublicAndStatus - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByRepositoryAndPublicAndStatusResponse{}, err

}
