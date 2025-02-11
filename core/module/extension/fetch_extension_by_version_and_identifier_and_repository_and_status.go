package extension

import (
	"context"

	"nem/core/module/extension/types"

	"errors"
	nemdb "nem/core/repository/gen"
	"nem/monitoring"
)

func (m *module) FetchExtensionByVersionAndIdentifierAndRepositoryAndStatus(
	ctx context.Context,
	req types.FetchExtensionByVersionAndIdentifierAndRepositoryAndStatusRequest,
	opts ...Option,
) (types.FetchExtensionByVersionAndIdentifierAndRepositoryAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndRepositoryAndStatus(
			ctx,
			nemdb.FetchExtensionByVersionAndIdentifierAndRepositoryAndStatusParams{
				Version:    req.Version,
				Identifier: req.Identifier,
				Repository: req.Repository,
				Status:     req.Status.ToInt64(),
				Offset:     req.Offset,
				Limit:      req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_status",
				Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndStatus no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndStatusResponse{}, err
		}
		return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndStatusResponse{
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
			ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_status",
			Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndStatus - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_status",
				Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndRepositoryAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierAndRepositoryAndStatusOrderedByCreatedAtASCParams{
					Version:    req.Version,
					Identifier: req.Identifier,
					Repository: req.Repository,
					Status:     req.Status.ToInt64(),
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_status",
					Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndRepositoryAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierAndRepositoryAndStatusOrderedByCreatedAtDESCParams{
					Version:    req.Version,
					Identifier: req.Identifier,
					Repository: req.Repository,
					Status:     req.Status.ToInt64(),
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_status",
					Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndRepositoryAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierAndRepositoryAndStatusOrderedByUpdatedAtASCParams{
					Version:    req.Version,
					Identifier: req.Identifier,
					Repository: req.Repository,
					Status:     req.Status.ToInt64(),
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_status",
					Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndRepositoryAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierAndRepositoryAndStatusOrderedByUpdatedAtDESCParams{
					Version:    req.Version,
					Identifier: req.Identifier,
					Repository: req.Repository,
					Status:     req.Status.ToInt64(),
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_status",
					Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_status_invalid",
		Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndStatus - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndStatusResponse{}, err

}
