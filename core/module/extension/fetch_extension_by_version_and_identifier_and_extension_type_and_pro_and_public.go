package extension

import (
	"context"

	"github.com/nuzur/nem/core/module/extension/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublic(
	ctx context.Context,
	req types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublicRequest,
	opts ...Option,
) (types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublicResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublic(
			ctx,
			nemdb.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublicParams{

				Version: req.Version,

				Identifier: req.Identifier,

				ExtensionType: req.ExtensionType.ToInt64(),

				Pro: req.Pro,

				Public: req.Public,

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_version_and_identifier_and_extension_type_and_pro_and_public",
				Message:          "error in FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublic no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublicResponse{}, err
		}
		return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublicResponse{
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
			ActionIdentifier: "fetch_extension_by_version_and_identifier_and_extension_type_and_pro_and_public",
			Message:          "error in FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublic - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublicResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_version_and_identifier_and_extension_type_and_pro_and_public",
				Message:          "error in FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublic - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublicResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublicOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublicOrderedByCreatedAtASCParams{

					Version: req.Version,

					Identifier: req.Identifier,

					ExtensionType: req.ExtensionType.ToInt64(),

					Pro: req.Pro,

					Public: req.Public,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_identifier_and_extension_type_and_pro_and_public",
					Message:          "error in FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublic - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublicResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublicOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublicOrderedByCreatedAtDESCParams{

					Version: req.Version,

					Identifier: req.Identifier,

					ExtensionType: req.ExtensionType.ToInt64(),

					Pro: req.Pro,

					Public: req.Public,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_identifier_and_extension_type_and_pro_and_public",
					Message:          "error in FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublic - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublicResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublicOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublicOrderedByUpdatedAtASCParams{

					Version: req.Version,

					Identifier: req.Identifier,

					ExtensionType: req.ExtensionType.ToInt64(),

					Pro: req.Pro,

					Public: req.Public,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_identifier_and_extension_type_and_pro_and_public",
					Message:          "error in FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublic - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublicResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublicOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublicOrderedByUpdatedAtDESCParams{

					Version: req.Version,

					Identifier: req.Identifier,

					ExtensionType: req.ExtensionType.ToInt64(),

					Pro: req.Pro,

					Public: req.Public,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_identifier_and_extension_type_and_pro_and_public",
					Message:          "error in FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublic - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublicResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_version_and_identifier_and_extension_type_and_pro_and_public_invalid",
		Message:          "error in FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublic - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublicResponse{}, err

}
