package extension

import (
	"context"

	"github.com/nuzur/nem/core/module/extension/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatus(
	ctx context.Context,
	req types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatusRequest,
	opts ...Option,
) (types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatus(
			ctx,
			nemdb.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatusParams{

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
				ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_extension_type_and_status",
				Message:          "error in FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatus no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatusResponse{}, err
		}
		return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatusResponse{
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
			ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_extension_type_and_status",
			Message:          "error in FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatus - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_extension_type_and_status",
				Message:          "error in FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatusOrderedByCreatedAtASCParams{

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
					ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_extension_type_and_status",
					Message:          "error in FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatusResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatusOrderedByCreatedAtDESCParams{

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
					ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_extension_type_and_status",
					Message:          "error in FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatusResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatusOrderedByUpdatedAtASCParams{

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
					ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_extension_type_and_status",
					Message:          "error in FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatusResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatusOrderedByUpdatedAtDESCParams{

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
					ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_extension_type_and_status",
					Message:          "error in FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatusResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_extension_type_and_status_invalid",
		Message:          "error in FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatus - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatusResponse{}, err

}
