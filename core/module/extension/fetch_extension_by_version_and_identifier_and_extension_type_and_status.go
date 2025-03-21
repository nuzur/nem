package extension

import (
	"context"

	"github.com/nuzur/nem/core/module/extension/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatus(
	ctx context.Context,
	req types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatusRequest,
	opts ...Option,
) (types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatus(
			ctx,
			nemdb.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatusParams{

				Version: req.Version,

				Identifier: req.Identifier,

				ExtensionType: req.ExtensionType.ToInt64(),

				Status: req.Status.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_version_and_identifier_and_extension_type_and_status",
				Message:          "error in FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatus no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatusResponse{}, err
		}
		return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatusResponse{
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
			ActionIdentifier: "fetch_extension_by_version_and_identifier_and_extension_type_and_status",
			Message:          "error in FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatus - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_version_and_identifier_and_extension_type_and_status",
				Message:          "error in FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatusOrderedByCreatedAtASCParams{

					Version: req.Version,

					Identifier: req.Identifier,

					ExtensionType: req.ExtensionType.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_identifier_and_extension_type_and_status",
					Message:          "error in FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatusOrderedByCreatedAtDESCParams{

					Version: req.Version,

					Identifier: req.Identifier,

					ExtensionType: req.ExtensionType.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_identifier_and_extension_type_and_status",
					Message:          "error in FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatusOrderedByUpdatedAtASCParams{

					Version: req.Version,

					Identifier: req.Identifier,

					ExtensionType: req.ExtensionType.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_identifier_and_extension_type_and_status",
					Message:          "error in FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatusOrderedByUpdatedAtDESCParams{

					Version: req.Version,

					Identifier: req.Identifier,

					ExtensionType: req.ExtensionType.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_identifier_and_extension_type_and_status",
					Message:          "error in FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_version_and_identifier_and_extension_type_and_status_invalid",
		Message:          "error in FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatus - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatusResponse{}, err

}
