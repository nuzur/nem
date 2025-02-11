package extension

import (
	"context"

	"nem/core/module/extension/types"

	"errors"
	nemdb "nem/core/repository/gen"
	"nem/monitoring"
)

func (m *module) FetchExtensionByVersionAndIdentifierAndExtensionType(
	ctx context.Context,
	req types.FetchExtensionByVersionAndIdentifierAndExtensionTypeRequest,
	opts ...Option,
) (types.FetchExtensionByVersionAndIdentifierAndExtensionTypeResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndExtensionType(
			ctx,
			nemdb.FetchExtensionByVersionAndIdentifierAndExtensionTypeParams{
				Version:       req.Version,
				Identifier:    req.Identifier,
				ExtensionType: req.ExtensionType.ToInt64(),
				Offset:        req.Offset,
				Limit:         req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_version_and_identifier_and_extension_type",
				Message:          "error in FetchExtensionByVersionAndIdentifierAndExtensionType no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeResponse{}, err
		}
		return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeResponse{
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
			ActionIdentifier: "fetch_extension_by_version_and_identifier_and_extension_type",
			Message:          "error in FetchExtensionByVersionAndIdentifierAndExtensionType - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_version_and_identifier_and_extension_type",
				Message:          "error in FetchExtensionByVersionAndIdentifierAndExtensionType - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndExtensionTypeOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierAndExtensionTypeOrderedByCreatedAtASCParams{
					Version:       req.Version,
					Identifier:    req.Identifier,
					ExtensionType: req.ExtensionType.ToInt64(),
					Offset:        req.Offset,
					Limit:         req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_identifier_and_extension_type",
					Message:          "error in FetchExtensionByVersionAndIdentifierAndExtensionType - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndExtensionTypeOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierAndExtensionTypeOrderedByCreatedAtDESCParams{
					Version:       req.Version,
					Identifier:    req.Identifier,
					ExtensionType: req.ExtensionType.ToInt64(),
					Offset:        req.Offset,
					Limit:         req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_identifier_and_extension_type",
					Message:          "error in FetchExtensionByVersionAndIdentifierAndExtensionType - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndExtensionTypeOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierAndExtensionTypeOrderedByUpdatedAtASCParams{
					Version:       req.Version,
					Identifier:    req.Identifier,
					ExtensionType: req.ExtensionType.ToInt64(),
					Offset:        req.Offset,
					Limit:         req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_identifier_and_extension_type",
					Message:          "error in FetchExtensionByVersionAndIdentifierAndExtensionType - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndExtensionTypeOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierAndExtensionTypeOrderedByUpdatedAtDESCParams{
					Version:       req.Version,
					Identifier:    req.Identifier,
					ExtensionType: req.ExtensionType.ToInt64(),
					Offset:        req.Offset,
					Limit:         req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_identifier_and_extension_type",
					Message:          "error in FetchExtensionByVersionAndIdentifierAndExtensionType - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_version_and_identifier_and_extension_type_invalid",
		Message:          "error in FetchExtensionByVersionAndIdentifierAndExtensionType - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeResponse{}, err

}
