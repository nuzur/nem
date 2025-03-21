package extension

import (
	"context"

	"github.com/nuzur/nem/core/module/extension/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatus(
	ctx context.Context,
	req types.FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatusRequest,
	opts ...Option,
) (types.FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatus(
			ctx,
			nemdb.FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatusParams{

				Repository: req.Repository,

				ExtensionType: req.ExtensionType.ToInt64(),

				Public: req.Public,

				Status: req.Status.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_repository_and_extension_type_and_public_and_status",
				Message:          "error in FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatus no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatusResponse{}, err
		}
		return types.FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatusResponse{
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
			ActionIdentifier: "fetch_extension_by_repository_and_extension_type_and_public_and_status",
			Message:          "error in FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatus - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_repository_and_extension_type_and_public_and_status",
				Message:          "error in FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatusOrderedByCreatedAtASCParams{

					Repository: req.Repository,

					ExtensionType: req.ExtensionType.ToInt64(),

					Public: req.Public,

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_repository_and_extension_type_and_public_and_status",
					Message:          "error in FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatusOrderedByCreatedAtDESCParams{

					Repository: req.Repository,

					ExtensionType: req.ExtensionType.ToInt64(),

					Public: req.Public,

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_repository_and_extension_type_and_public_and_status",
					Message:          "error in FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatusOrderedByUpdatedAtASCParams{

					Repository: req.Repository,

					ExtensionType: req.ExtensionType.ToInt64(),

					Public: req.Public,

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_repository_and_extension_type_and_public_and_status",
					Message:          "error in FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatusOrderedByUpdatedAtDESCParams{

					Repository: req.Repository,

					ExtensionType: req.ExtensionType.ToInt64(),

					Public: req.Public,

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_repository_and_extension_type_and_public_and_status",
					Message:          "error in FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_repository_and_extension_type_and_public_and_status_invalid",
		Message:          "error in FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatus - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatusResponse{}, err

}
