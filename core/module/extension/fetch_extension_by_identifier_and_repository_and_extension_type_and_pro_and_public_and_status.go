package extension

import (
	"context"

	"github.com/nuzur/nem/core/module/extension/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatus(
	ctx context.Context,
	req types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatusRequest,
	opts ...Option,
) (types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatus(
			ctx,
			nemdb.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatusParams{

				Identifier: req.Identifier,

				Repository: req.Repository,

				ExtensionType: req.ExtensionType.ToInt64(),

				Pro: req.Pro,

				Public: req.Public,

				Status: req.Status.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_extension_type_and_pro_and_public_and_status",
				Message:          "error in FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatus no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatusResponse{}, err
		}
		return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatusResponse{
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
			ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_extension_type_and_pro_and_public_and_status",
			Message:          "error in FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatus - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_extension_type_and_pro_and_public_and_status",
				Message:          "error in FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatusOrderedByCreatedAtASCParams{

					Identifier: req.Identifier,

					Repository: req.Repository,

					ExtensionType: req.ExtensionType.ToInt64(),

					Pro: req.Pro,

					Public: req.Public,

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_extension_type_and_pro_and_public_and_status",
					Message:          "error in FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatusOrderedByCreatedAtDESCParams{

					Identifier: req.Identifier,

					Repository: req.Repository,

					ExtensionType: req.ExtensionType.ToInt64(),

					Pro: req.Pro,

					Public: req.Public,

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_extension_type_and_pro_and_public_and_status",
					Message:          "error in FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatusOrderedByUpdatedAtASCParams{

					Identifier: req.Identifier,

					Repository: req.Repository,

					ExtensionType: req.ExtensionType.ToInt64(),

					Pro: req.Pro,

					Public: req.Public,

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_extension_type_and_pro_and_public_and_status",
					Message:          "error in FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatusOrderedByUpdatedAtDESCParams{

					Identifier: req.Identifier,

					Repository: req.Repository,

					ExtensionType: req.ExtensionType.ToInt64(),

					Pro: req.Pro,

					Public: req.Public,

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_extension_type_and_pro_and_public_and_status",
					Message:          "error in FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_extension_type_and_pro_and_public_and_status_invalid",
		Message:          "error in FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatus - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatusResponse{}, err

}
