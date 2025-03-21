package extension

import (
	"context"

	"github.com/nuzur/nem/core/module/extension/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchExtensionByVersionAndRepositoryAndPublic(
	ctx context.Context,
	req types.FetchExtensionByVersionAndRepositoryAndPublicRequest,
	opts ...Option,
) (types.FetchExtensionByVersionAndRepositoryAndPublicResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByVersionAndRepositoryAndPublic(
			ctx,
			nemdb.FetchExtensionByVersionAndRepositoryAndPublicParams{

				Version: req.Version,

				Repository: req.Repository,

				Public: req.Public,

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_version_and_repository_and_public",
				Message:          "error in FetchExtensionByVersionAndRepositoryAndPublic no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByVersionAndRepositoryAndPublicResponse{}, err
		}
		return types.FetchExtensionByVersionAndRepositoryAndPublicResponse{
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
			ActionIdentifier: "fetch_extension_by_version_and_repository_and_public",
			Message:          "error in FetchExtensionByVersionAndRepositoryAndPublic - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByVersionAndRepositoryAndPublicResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_version_and_repository_and_public",
				Message:          "error in FetchExtensionByVersionAndRepositoryAndPublic - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByVersionAndRepositoryAndPublicResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByVersionAndRepositoryAndPublicOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByVersionAndRepositoryAndPublicOrderedByCreatedAtASCParams{

					Version: req.Version,

					Repository: req.Repository,

					Public: req.Public,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_repository_and_public",
					Message:          "error in FetchExtensionByVersionAndRepositoryAndPublic - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndRepositoryAndPublicResponse{}, err
			}
			return types.FetchExtensionByVersionAndRepositoryAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByVersionAndRepositoryAndPublicOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByVersionAndRepositoryAndPublicOrderedByCreatedAtDESCParams{

					Version: req.Version,

					Repository: req.Repository,

					Public: req.Public,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_repository_and_public",
					Message:          "error in FetchExtensionByVersionAndRepositoryAndPublic - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndRepositoryAndPublicResponse{}, err
			}
			return types.FetchExtensionByVersionAndRepositoryAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByVersionAndRepositoryAndPublicOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByVersionAndRepositoryAndPublicOrderedByUpdatedAtASCParams{

					Version: req.Version,

					Repository: req.Repository,

					Public: req.Public,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_repository_and_public",
					Message:          "error in FetchExtensionByVersionAndRepositoryAndPublic - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndRepositoryAndPublicResponse{}, err
			}
			return types.FetchExtensionByVersionAndRepositoryAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByVersionAndRepositoryAndPublicOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByVersionAndRepositoryAndPublicOrderedByUpdatedAtDESCParams{

					Version: req.Version,

					Repository: req.Repository,

					Public: req.Public,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_repository_and_public",
					Message:          "error in FetchExtensionByVersionAndRepositoryAndPublic - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndRepositoryAndPublicResponse{}, err
			}
			return types.FetchExtensionByVersionAndRepositoryAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_version_and_repository_and_public_invalid",
		Message:          "error in FetchExtensionByVersionAndRepositoryAndPublic - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByVersionAndRepositoryAndPublicResponse{}, err

}
