package extension

import (
	"context"

	"github.com/nuzur/nem/core/module/extension/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublic(
	ctx context.Context,
	req types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicRequest,
	opts ...Option,
) (types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublic(
			ctx,
			nemdb.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicParams{

				Version: req.Version,

				Identifier: req.Identifier,

				ExtensionType: req.ExtensionType.ToInt64(),

				Public: req.Public,

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_version_and_identifier_and_extension_type_and_public",
				Message:          "error in FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublic no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicResponse{}, err
		}
		return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicResponse{
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
			ActionIdentifier: "fetch_extension_by_version_and_identifier_and_extension_type_and_public",
			Message:          "error in FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublic - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_version_and_identifier_and_extension_type_and_public",
				Message:          "error in FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublic - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicOrderedByCreatedAtASCParams{

					Version: req.Version,

					Identifier: req.Identifier,

					ExtensionType: req.ExtensionType.ToInt64(),

					Public: req.Public,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_identifier_and_extension_type_and_public",
					Message:          "error in FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublic - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicOrderedByCreatedAtDESCParams{

					Version: req.Version,

					Identifier: req.Identifier,

					ExtensionType: req.ExtensionType.ToInt64(),

					Public: req.Public,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_identifier_and_extension_type_and_public",
					Message:          "error in FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublic - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicOrderedByUpdatedAtASCParams{

					Version: req.Version,

					Identifier: req.Identifier,

					ExtensionType: req.ExtensionType.ToInt64(),

					Public: req.Public,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_identifier_and_extension_type_and_public",
					Message:          "error in FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublic - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicOrderedByUpdatedAtDESCParams{

					Version: req.Version,

					Identifier: req.Identifier,

					ExtensionType: req.ExtensionType.ToInt64(),

					Public: req.Public,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_identifier_and_extension_type_and_public",
					Message:          "error in FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublic - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_version_and_identifier_and_extension_type_and_public_invalid",
		Message:          "error in FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublic - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicResponse{}, err

}
