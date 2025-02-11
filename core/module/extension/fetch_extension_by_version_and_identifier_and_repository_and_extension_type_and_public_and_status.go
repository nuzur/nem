package extension

import (
	"context"

	"github.com/nuzur/nem/core/module/extension/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatus(
	ctx context.Context,
	req types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusRequest,
	opts ...Option,
) (types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatus(
			ctx,
			nemdb.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusParams{
				Version:       req.Version,
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
				ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_extension_type_and_public_and_status",
				Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatus no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusResponse{}, err
		}
		return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusResponse{
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
			ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_extension_type_and_public_and_status",
			Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatus - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_extension_type_and_public_and_status",
				Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusOrderedByCreatedAtASCParams{
					Version:       req.Version,
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
					ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_extension_type_and_public_and_status",
					Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusOrderedByCreatedAtDESCParams{
					Version:       req.Version,
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
					ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_extension_type_and_public_and_status",
					Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusOrderedByUpdatedAtASCParams{
					Version:       req.Version,
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
					ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_extension_type_and_public_and_status",
					Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusOrderedByUpdatedAtDESCParams{
					Version:       req.Version,
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
					ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_extension_type_and_public_and_status",
					Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_extension_type_and_public_and_status_invalid",
		Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatus - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusResponse{}, err

}
