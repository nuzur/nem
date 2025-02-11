package extension

import (
	"context"

	"nem/core/module/extension/types"

	"errors"
	nemdb "nem/core/repository/gen"
	"nem/monitoring"
)

func (m *module) FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatus(
	ctx context.Context,
	req types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusRequest,
	opts ...Option,
) (types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatus(
			ctx,
			nemdb.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusParams{
				Identifier:    req.Identifier,
				Repository:    req.Repository,
				ExtensionType: req.ExtensionType.ToInt64(),
				Public:        req.Public,
				Status:        req.Status.ToInt64(),
				Offset:        req.Offset,
				Limit:         req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_extension_type_and_public_and_status",
				Message:          "error in FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatus no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusResponse{}, err
		}
		return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusResponse{
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
			ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_extension_type_and_public_and_status",
			Message:          "error in FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatus - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_extension_type_and_public_and_status",
				Message:          "error in FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusOrderedByCreatedAtASCParams{
					Identifier:    req.Identifier,
					Repository:    req.Repository,
					ExtensionType: req.ExtensionType.ToInt64(),
					Public:        req.Public,
					Status:        req.Status.ToInt64(),
					Offset:        req.Offset,
					Limit:         req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_extension_type_and_public_and_status",
					Message:          "error in FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusOrderedByCreatedAtDESCParams{
					Identifier:    req.Identifier,
					Repository:    req.Repository,
					ExtensionType: req.ExtensionType.ToInt64(),
					Public:        req.Public,
					Status:        req.Status.ToInt64(),
					Offset:        req.Offset,
					Limit:         req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_extension_type_and_public_and_status",
					Message:          "error in FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusOrderedByUpdatedAtASCParams{
					Identifier:    req.Identifier,
					Repository:    req.Repository,
					ExtensionType: req.ExtensionType.ToInt64(),
					Public:        req.Public,
					Status:        req.Status.ToInt64(),
					Offset:        req.Offset,
					Limit:         req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_extension_type_and_public_and_status",
					Message:          "error in FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusOrderedByUpdatedAtDESCParams{
					Identifier:    req.Identifier,
					Repository:    req.Repository,
					ExtensionType: req.ExtensionType.ToInt64(),
					Public:        req.Public,
					Status:        req.Status.ToInt64(),
					Offset:        req.Offset,
					Limit:         req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_extension_type_and_public_and_status",
					Message:          "error in FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_extension_type_and_public_and_status_invalid",
		Message:          "error in FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatus - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusResponse{}, err

}
