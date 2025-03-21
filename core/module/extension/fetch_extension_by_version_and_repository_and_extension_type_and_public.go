package extension

import (
	"context"

	"github.com/nuzur/nem/core/module/extension/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublic(
	ctx context.Context,
	req types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicRequest,
	opts ...Option,
) (types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublic(
			ctx,
			nemdb.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicParams{

				Version: req.Version,

				Repository: req.Repository,

				ExtensionType: req.ExtensionType.ToInt64(),

				Public: req.Public,

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type_and_public",
				Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublic no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicResponse{}, err
		}
		return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicResponse{
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
			ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type_and_public",
			Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublic - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type_and_public",
				Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublic - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicOrderedByCreatedAtASCParams{

					Version: req.Version,

					Repository: req.Repository,

					ExtensionType: req.ExtensionType.ToInt64(),

					Public: req.Public,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type_and_public",
					Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublic - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicResponse{}, err
			}
			return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicOrderedByCreatedAtDESCParams{

					Version: req.Version,

					Repository: req.Repository,

					ExtensionType: req.ExtensionType.ToInt64(),

					Public: req.Public,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type_and_public",
					Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublic - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicResponse{}, err
			}
			return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicOrderedByUpdatedAtASCParams{

					Version: req.Version,

					Repository: req.Repository,

					ExtensionType: req.ExtensionType.ToInt64(),

					Public: req.Public,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type_and_public",
					Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublic - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicResponse{}, err
			}
			return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicOrderedByUpdatedAtDESCParams{

					Version: req.Version,

					Repository: req.Repository,

					ExtensionType: req.ExtensionType.ToInt64(),

					Public: req.Public,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type_and_public",
					Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublic - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicResponse{}, err
			}
			return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type_and_public_invalid",
		Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublic - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicResponse{}, err

}
