package extension

import (
	"context"

	"nem/core/module/extension/types"

	"errors"
	nemdb "nem/core/repository/gen"
	"nem/monitoring"
)

func (m *module) FetchExtensionByIdentifierAndExtensionTypeAndPublic(
	ctx context.Context,
	req types.FetchExtensionByIdentifierAndExtensionTypeAndPublicRequest,
	opts ...Option,
) (types.FetchExtensionByIdentifierAndExtensionTypeAndPublicResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByIdentifierAndExtensionTypeAndPublic(
			ctx,
			nemdb.FetchExtensionByIdentifierAndExtensionTypeAndPublicParams{
				Identifier:    req.Identifier,
				ExtensionType: req.ExtensionType.ToInt64(),
				Public:        req.Public,
				Offset:        req.Offset,
				Limit:         req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_identifier_and_extension_type_and_public",
				Message:          "error in FetchExtensionByIdentifierAndExtensionTypeAndPublic no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByIdentifierAndExtensionTypeAndPublicResponse{}, err
		}
		return types.FetchExtensionByIdentifierAndExtensionTypeAndPublicResponse{
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
			ActionIdentifier: "fetch_extension_by_identifier_and_extension_type_and_public",
			Message:          "error in FetchExtensionByIdentifierAndExtensionTypeAndPublic - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByIdentifierAndExtensionTypeAndPublicResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_identifier_and_extension_type_and_public",
				Message:          "error in FetchExtensionByIdentifierAndExtensionTypeAndPublic - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByIdentifierAndExtensionTypeAndPublicResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndExtensionTypeAndPublicOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndExtensionTypeAndPublicOrderedByCreatedAtASCParams{
					Identifier:    req.Identifier,
					ExtensionType: req.ExtensionType.ToInt64(),
					Public:        req.Public,
					Offset:        req.Offset,
					Limit:         req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_extension_type_and_public",
					Message:          "error in FetchExtensionByIdentifierAndExtensionTypeAndPublic - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndExtensionTypeAndPublicResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndExtensionTypeAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndExtensionTypeAndPublicOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndExtensionTypeAndPublicOrderedByCreatedAtDESCParams{
					Identifier:    req.Identifier,
					ExtensionType: req.ExtensionType.ToInt64(),
					Public:        req.Public,
					Offset:        req.Offset,
					Limit:         req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_extension_type_and_public",
					Message:          "error in FetchExtensionByIdentifierAndExtensionTypeAndPublic - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndExtensionTypeAndPublicResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndExtensionTypeAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndExtensionTypeAndPublicOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndExtensionTypeAndPublicOrderedByUpdatedAtASCParams{
					Identifier:    req.Identifier,
					ExtensionType: req.ExtensionType.ToInt64(),
					Public:        req.Public,
					Offset:        req.Offset,
					Limit:         req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_extension_type_and_public",
					Message:          "error in FetchExtensionByIdentifierAndExtensionTypeAndPublic - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndExtensionTypeAndPublicResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndExtensionTypeAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndExtensionTypeAndPublicOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndExtensionTypeAndPublicOrderedByUpdatedAtDESCParams{
					Identifier:    req.Identifier,
					ExtensionType: req.ExtensionType.ToInt64(),
					Public:        req.Public,
					Offset:        req.Offset,
					Limit:         req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_extension_type_and_public",
					Message:          "error in FetchExtensionByIdentifierAndExtensionTypeAndPublic - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndExtensionTypeAndPublicResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndExtensionTypeAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_identifier_and_extension_type_and_public_invalid",
		Message:          "error in FetchExtensionByIdentifierAndExtensionTypeAndPublic - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByIdentifierAndExtensionTypeAndPublicResponse{}, err

}
