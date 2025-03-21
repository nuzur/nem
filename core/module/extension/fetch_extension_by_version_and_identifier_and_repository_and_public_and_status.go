package extension

import (
	"context"

	"github.com/nuzur/nem/core/module/extension/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatus(
	ctx context.Context,
	req types.FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatusRequest,
	opts ...Option,
) (types.FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatus(
			ctx,
			nemdb.FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatusParams{

				Version: req.Version,

				Identifier: req.Identifier,

				Repository: req.Repository,

				Public: req.Public,

				Status: req.Status.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_public_and_status",
				Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatus no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatusResponse{}, err
		}
		return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatusResponse{
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
			ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_public_and_status",
			Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatus - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_public_and_status",
				Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatusOrderedByCreatedAtASCParams{

					Version: req.Version,

					Identifier: req.Identifier,

					Repository: req.Repository,

					Public: req.Public,

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_public_and_status",
					Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatusOrderedByCreatedAtDESCParams{

					Version: req.Version,

					Identifier: req.Identifier,

					Repository: req.Repository,

					Public: req.Public,

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_public_and_status",
					Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatusOrderedByUpdatedAtASCParams{

					Version: req.Version,

					Identifier: req.Identifier,

					Repository: req.Repository,

					Public: req.Public,

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_public_and_status",
					Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatusOrderedByUpdatedAtDESCParams{

					Version: req.Version,

					Identifier: req.Identifier,

					Repository: req.Repository,

					Public: req.Public,

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_public_and_status",
					Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_public_and_status_invalid",
		Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatus - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatusResponse{}, err

}
