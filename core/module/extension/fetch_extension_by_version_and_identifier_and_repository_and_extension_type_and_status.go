package extension

import (
	"context"

	"github.com/nuzur/nem/core/module/extension/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatus(
	ctx context.Context,
	req types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatusRequest,
	opts ...Option,
) (types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatus(
			ctx,
			nemdb.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatusParams{

				Version: req.Version,

				Identifier: req.Identifier,

				Repository: req.Repository,

				ExtensionType: req.ExtensionType.ToInt64(),

				Status: req.Status.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_extension_type_and_status",
				Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatus no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatusResponse{}, err
		}
		return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatusResponse{
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
			ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_extension_type_and_status",
			Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatus - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_extension_type_and_status",
				Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatusOrderedByCreatedAtASCParams{

					Version: req.Version,

					Identifier: req.Identifier,

					Repository: req.Repository,

					ExtensionType: req.ExtensionType.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_extension_type_and_status",
					Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatusOrderedByCreatedAtDESCParams{

					Version: req.Version,

					Identifier: req.Identifier,

					Repository: req.Repository,

					ExtensionType: req.ExtensionType.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_extension_type_and_status",
					Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatusOrderedByUpdatedAtASCParams{

					Version: req.Version,

					Identifier: req.Identifier,

					Repository: req.Repository,

					ExtensionType: req.ExtensionType.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_extension_type_and_status",
					Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatusOrderedByUpdatedAtDESCParams{

					Version: req.Version,

					Identifier: req.Identifier,

					Repository: req.Repository,

					ExtensionType: req.ExtensionType.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_extension_type_and_status",
					Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_version_and_identifier_and_repository_and_extension_type_and_status_invalid",
		Message:          "error in FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatus - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatusResponse{}, err

}
