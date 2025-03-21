package extension

import (
	"context"

	"github.com/nuzur/nem/core/module/extension/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublic(
	ctx context.Context,
	req types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicRequest,
	opts ...Option,
) (types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublic(
			ctx,
			nemdb.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicParams{

				Version: req.Version,

				Identifier: req.Identifier,

				Repository: req.Repository,

				ExtensionType: req.ExtensionType.ToInt64(),

				Public: req.Public,

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_extension_type_and_public",
				Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublic no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicResponse{}, err
		}
		return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicResponse{
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
			ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_extension_type_and_public",
			Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublic - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_extension_type_and_public",
				Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublic - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicOrderedByCreatedAtASCParams{

					Version: req.Version,

					Identifier: req.Identifier,

					Repository: req.Repository,

					ExtensionType: req.ExtensionType.ToInt64(),

					Public: req.Public,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_extension_type_and_public",
					Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublic - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicOrderedByCreatedAtDESCParams{

					Version: req.Version,

					Identifier: req.Identifier,

					Repository: req.Repository,

					ExtensionType: req.ExtensionType.ToInt64(),

					Public: req.Public,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_extension_type_and_public",
					Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublic - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicOrderedByUpdatedAtASCParams{

					Version: req.Version,

					Identifier: req.Identifier,

					Repository: req.Repository,

					ExtensionType: req.ExtensionType.ToInt64(),

					Public: req.Public,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_extension_type_and_public",
					Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublic - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicOrderedByUpdatedAtDESCParams{

					Version: req.Version,

					Identifier: req.Identifier,

					Repository: req.Repository,

					ExtensionType: req.ExtensionType.ToInt64(),

					Public: req.Public,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_extension_type_and_public",
					Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublic - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_extension_type_and_public_invalid",
		Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublic - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicResponse{}, err

}
