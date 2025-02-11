package extension

import (
	"context"

	"nem/core/module/extension/types"

	"errors"
	nemdb "nem/core/repository/gen"
	"nem/monitoring"
)

func (m *module) FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatus(
	ctx context.Context,
	req types.FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatusRequest,
	opts ...Option,
) (types.FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatus(
			ctx,
			nemdb.FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatusParams{
				Identifier:    req.Identifier,
				ExtensionType: req.ExtensionType.ToInt64(),
				Public:        req.Public,
				Status:        req.Status.ToInt64(),
				Offset:        req.Offset,
				Limit:         req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_identifier_and_extension_type_and_public_and_status",
				Message:          "error in FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatus no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatusResponse{}, err
		}
		return types.FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatusResponse{
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
			ActionIdentifier: "fetch_extension_by_identifier_and_extension_type_and_public_and_status",
			Message:          "error in FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatus - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_identifier_and_extension_type_and_public_and_status",
				Message:          "error in FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatusOrderedByCreatedAtASCParams{
					Identifier:    req.Identifier,
					ExtensionType: req.ExtensionType.ToInt64(),
					Public:        req.Public,
					Status:        req.Status.ToInt64(),
					Offset:        req.Offset,
					Limit:         req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_extension_type_and_public_and_status",
					Message:          "error in FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatusOrderedByCreatedAtDESCParams{
					Identifier:    req.Identifier,
					ExtensionType: req.ExtensionType.ToInt64(),
					Public:        req.Public,
					Status:        req.Status.ToInt64(),
					Offset:        req.Offset,
					Limit:         req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_extension_type_and_public_and_status",
					Message:          "error in FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatusOrderedByUpdatedAtASCParams{
					Identifier:    req.Identifier,
					ExtensionType: req.ExtensionType.ToInt64(),
					Public:        req.Public,
					Status:        req.Status.ToInt64(),
					Offset:        req.Offset,
					Limit:         req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_extension_type_and_public_and_status",
					Message:          "error in FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatusOrderedByUpdatedAtDESCParams{
					Identifier:    req.Identifier,
					ExtensionType: req.ExtensionType.ToInt64(),
					Public:        req.Public,
					Status:        req.Status.ToInt64(),
					Offset:        req.Offset,
					Limit:         req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_extension_type_and_public_and_status",
					Message:          "error in FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_identifier_and_extension_type_and_public_and_status_invalid",
		Message:          "error in FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatus - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatusResponse{}, err

}
