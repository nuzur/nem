package extension

import (
	"context"

	"nem/core/module/extension/types"

	"errors"
	nemdb "nem/core/repository/gen"
	"nem/monitoring"
)

func (m *module) FetchExtensionByRepositoryAndExtensionTypeAndPublic(
	ctx context.Context,
	req types.FetchExtensionByRepositoryAndExtensionTypeAndPublicRequest,
	opts ...Option,
) (types.FetchExtensionByRepositoryAndExtensionTypeAndPublicResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByRepositoryAndExtensionTypeAndPublic(
			ctx,
			nemdb.FetchExtensionByRepositoryAndExtensionTypeAndPublicParams{
				Repository:    req.Repository,
				ExtensionType: req.ExtensionType.ToInt64(),
				Public:        req.Public,
				Offset:        req.Offset,
				Limit:         req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_repository_and_extension_type_and_public",
				Message:          "error in FetchExtensionByRepositoryAndExtensionTypeAndPublic no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByRepositoryAndExtensionTypeAndPublicResponse{}, err
		}
		return types.FetchExtensionByRepositoryAndExtensionTypeAndPublicResponse{
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
			ActionIdentifier: "fetch_extension_by_repository_and_extension_type_and_public",
			Message:          "error in FetchExtensionByRepositoryAndExtensionTypeAndPublic - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByRepositoryAndExtensionTypeAndPublicResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_repository_and_extension_type_and_public",
				Message:          "error in FetchExtensionByRepositoryAndExtensionTypeAndPublic - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByRepositoryAndExtensionTypeAndPublicResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByRepositoryAndExtensionTypeAndPublicOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByRepositoryAndExtensionTypeAndPublicOrderedByCreatedAtASCParams{
					Repository:    req.Repository,
					ExtensionType: req.ExtensionType.ToInt64(),
					Public:        req.Public,
					Offset:        req.Offset,
					Limit:         req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_repository_and_extension_type_and_public",
					Message:          "error in FetchExtensionByRepositoryAndExtensionTypeAndPublic - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByRepositoryAndExtensionTypeAndPublicResponse{}, err
			}
			return types.FetchExtensionByRepositoryAndExtensionTypeAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByRepositoryAndExtensionTypeAndPublicOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByRepositoryAndExtensionTypeAndPublicOrderedByCreatedAtDESCParams{
					Repository:    req.Repository,
					ExtensionType: req.ExtensionType.ToInt64(),
					Public:        req.Public,
					Offset:        req.Offset,
					Limit:         req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_repository_and_extension_type_and_public",
					Message:          "error in FetchExtensionByRepositoryAndExtensionTypeAndPublic - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByRepositoryAndExtensionTypeAndPublicResponse{}, err
			}
			return types.FetchExtensionByRepositoryAndExtensionTypeAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByRepositoryAndExtensionTypeAndPublicOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByRepositoryAndExtensionTypeAndPublicOrderedByUpdatedAtASCParams{
					Repository:    req.Repository,
					ExtensionType: req.ExtensionType.ToInt64(),
					Public:        req.Public,
					Offset:        req.Offset,
					Limit:         req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_repository_and_extension_type_and_public",
					Message:          "error in FetchExtensionByRepositoryAndExtensionTypeAndPublic - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByRepositoryAndExtensionTypeAndPublicResponse{}, err
			}
			return types.FetchExtensionByRepositoryAndExtensionTypeAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByRepositoryAndExtensionTypeAndPublicOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByRepositoryAndExtensionTypeAndPublicOrderedByUpdatedAtDESCParams{
					Repository:    req.Repository,
					ExtensionType: req.ExtensionType.ToInt64(),
					Public:        req.Public,
					Offset:        req.Offset,
					Limit:         req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_repository_and_extension_type_and_public",
					Message:          "error in FetchExtensionByRepositoryAndExtensionTypeAndPublic - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByRepositoryAndExtensionTypeAndPublicResponse{}, err
			}
			return types.FetchExtensionByRepositoryAndExtensionTypeAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_repository_and_extension_type_and_public_invalid",
		Message:          "error in FetchExtensionByRepositoryAndExtensionTypeAndPublic - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByRepositoryAndExtensionTypeAndPublicResponse{}, err

}
