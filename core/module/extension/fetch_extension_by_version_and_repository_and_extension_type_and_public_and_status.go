package extension

import (
	"context"

	"github.com/nuzur/nem/core/module/extension/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatus(
	ctx context.Context,
	req types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatusRequest,
	opts ...Option,
) (types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatus(
			ctx,
			nemdb.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatusParams{

				Version: req.Version,

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
				ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type_and_public_and_status",
				Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatus no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatusResponse{}, err
		}
		return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatusResponse{
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
			ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type_and_public_and_status",
			Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatus - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type_and_public_and_status",
				Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatusOrderedByCreatedAtASCParams{

					Version: req.Version,

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
					ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type_and_public_and_status",
					Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatusOrderedByCreatedAtDESCParams{

					Version: req.Version,

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
					ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type_and_public_and_status",
					Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatusOrderedByUpdatedAtASCParams{

					Version: req.Version,

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
					ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type_and_public_and_status",
					Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatusOrderedByUpdatedAtDESCParams{

					Version: req.Version,

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
					ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type_and_public_and_status",
					Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type_and_public_and_status_invalid",
		Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatus - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatusResponse{}, err

}
