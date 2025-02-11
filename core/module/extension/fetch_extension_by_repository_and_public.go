package extension

import (
	"context"

	"nem/core/module/extension/types"

	"errors"
	nemdb "nem/core/repository/gen"
	"nem/monitoring"
)

func (m *module) FetchExtensionByRepositoryAndPublic(
	ctx context.Context,
	req types.FetchExtensionByRepositoryAndPublicRequest,
	opts ...Option,
) (types.FetchExtensionByRepositoryAndPublicResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByRepositoryAndPublic(
			ctx,
			nemdb.FetchExtensionByRepositoryAndPublicParams{
				Repository: req.Repository,
				Public:     req.Public,
				Offset:     req.Offset,
				Limit:      req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_repository_and_public",
				Message:          "error in FetchExtensionByRepositoryAndPublic no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByRepositoryAndPublicResponse{}, err
		}
		return types.FetchExtensionByRepositoryAndPublicResponse{
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
			ActionIdentifier: "fetch_extension_by_repository_and_public",
			Message:          "error in FetchExtensionByRepositoryAndPublic - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByRepositoryAndPublicResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_repository_and_public",
				Message:          "error in FetchExtensionByRepositoryAndPublic - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByRepositoryAndPublicResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByRepositoryAndPublicOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByRepositoryAndPublicOrderedByCreatedAtASCParams{
					Repository: req.Repository,
					Public:     req.Public,
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_repository_and_public",
					Message:          "error in FetchExtensionByRepositoryAndPublic - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByRepositoryAndPublicResponse{}, err
			}
			return types.FetchExtensionByRepositoryAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByRepositoryAndPublicOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByRepositoryAndPublicOrderedByCreatedAtDESCParams{
					Repository: req.Repository,
					Public:     req.Public,
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_repository_and_public",
					Message:          "error in FetchExtensionByRepositoryAndPublic - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByRepositoryAndPublicResponse{}, err
			}
			return types.FetchExtensionByRepositoryAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByRepositoryAndPublicOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByRepositoryAndPublicOrderedByUpdatedAtASCParams{
					Repository: req.Repository,
					Public:     req.Public,
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_repository_and_public",
					Message:          "error in FetchExtensionByRepositoryAndPublic - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByRepositoryAndPublicResponse{}, err
			}
			return types.FetchExtensionByRepositoryAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByRepositoryAndPublicOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByRepositoryAndPublicOrderedByUpdatedAtDESCParams{
					Repository: req.Repository,
					Public:     req.Public,
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_repository_and_public",
					Message:          "error in FetchExtensionByRepositoryAndPublic - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByRepositoryAndPublicResponse{}, err
			}
			return types.FetchExtensionByRepositoryAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_repository_and_public_invalid",
		Message:          "error in FetchExtensionByRepositoryAndPublic - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByRepositoryAndPublicResponse{}, err

}
