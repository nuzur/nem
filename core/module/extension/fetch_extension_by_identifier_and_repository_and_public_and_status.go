package extension

import (
	"context"

	"nem/core/module/extension/types"

	"errors"
	nemdb "nem/core/repository/gen"
	"nem/monitoring"
)

func (m *module) FetchExtensionByIdentifierAndRepositoryAndPublicAndStatus(
	ctx context.Context,
	req types.FetchExtensionByIdentifierAndRepositoryAndPublicAndStatusRequest,
	opts ...Option,
) (types.FetchExtensionByIdentifierAndRepositoryAndPublicAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByIdentifierAndRepositoryAndPublicAndStatus(
			ctx,
			nemdb.FetchExtensionByIdentifierAndRepositoryAndPublicAndStatusParams{
				Identifier: req.Identifier,
				Repository: req.Repository,
				Public:     req.Public,
				Status:     req.Status.ToInt64(),
				Offset:     req.Offset,
				Limit:      req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_public_and_status",
				Message:          "error in FetchExtensionByIdentifierAndRepositoryAndPublicAndStatus no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByIdentifierAndRepositoryAndPublicAndStatusResponse{}, err
		}
		return types.FetchExtensionByIdentifierAndRepositoryAndPublicAndStatusResponse{
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
			ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_public_and_status",
			Message:          "error in FetchExtensionByIdentifierAndRepositoryAndPublicAndStatus - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByIdentifierAndRepositoryAndPublicAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_public_and_status",
				Message:          "error in FetchExtensionByIdentifierAndRepositoryAndPublicAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByIdentifierAndRepositoryAndPublicAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndRepositoryAndPublicAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndRepositoryAndPublicAndStatusOrderedByCreatedAtASCParams{
					Identifier: req.Identifier,
					Repository: req.Repository,
					Public:     req.Public,
					Status:     req.Status.ToInt64(),
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_public_and_status",
					Message:          "error in FetchExtensionByIdentifierAndRepositoryAndPublicAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndRepositoryAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndRepositoryAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndRepositoryAndPublicAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndRepositoryAndPublicAndStatusOrderedByCreatedAtDESCParams{
					Identifier: req.Identifier,
					Repository: req.Repository,
					Public:     req.Public,
					Status:     req.Status.ToInt64(),
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_public_and_status",
					Message:          "error in FetchExtensionByIdentifierAndRepositoryAndPublicAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndRepositoryAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndRepositoryAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndRepositoryAndPublicAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndRepositoryAndPublicAndStatusOrderedByUpdatedAtASCParams{
					Identifier: req.Identifier,
					Repository: req.Repository,
					Public:     req.Public,
					Status:     req.Status.ToInt64(),
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_public_and_status",
					Message:          "error in FetchExtensionByIdentifierAndRepositoryAndPublicAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndRepositoryAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndRepositoryAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndRepositoryAndPublicAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndRepositoryAndPublicAndStatusOrderedByUpdatedAtDESCParams{
					Identifier: req.Identifier,
					Repository: req.Repository,
					Public:     req.Public,
					Status:     req.Status.ToInt64(),
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_public_and_status",
					Message:          "error in FetchExtensionByIdentifierAndRepositoryAndPublicAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndRepositoryAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndRepositoryAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_public_and_status_invalid",
		Message:          "error in FetchExtensionByIdentifierAndRepositoryAndPublicAndStatus - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByIdentifierAndRepositoryAndPublicAndStatusResponse{}, err

}
