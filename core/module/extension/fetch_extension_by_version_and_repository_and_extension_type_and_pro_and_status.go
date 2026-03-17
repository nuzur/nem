package extension

import (
	"context"

	"github.com/nuzur/nem/core/module/extension/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatus(
	ctx context.Context,
	req types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatusRequest,
	opts ...Option,
) (types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatus(
			ctx,
			nemdb.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatusParams{

				Version: req.Version,

				Repository: req.Repository,

				ExtensionType: req.ExtensionType.ToInt64(),

				Pro: req.Pro,

				Status: req.Status.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type_and_pro_and_status",
				Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatus no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatusResponse{}, err
		}
		return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatusResponse{
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
			ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type_and_pro_and_status",
			Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatus - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type_and_pro_and_status",
				Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatusOrderedByCreatedAtASCParams{

					Version: req.Version,

					Repository: req.Repository,

					ExtensionType: req.ExtensionType.ToInt64(),

					Pro: req.Pro,

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type_and_pro_and_status",
					Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatusOrderedByCreatedAtDESCParams{

					Version: req.Version,

					Repository: req.Repository,

					ExtensionType: req.ExtensionType.ToInt64(),

					Pro: req.Pro,

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type_and_pro_and_status",
					Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatusOrderedByUpdatedAtASCParams{

					Version: req.Version,

					Repository: req.Repository,

					ExtensionType: req.ExtensionType.ToInt64(),

					Pro: req.Pro,

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type_and_pro_and_status",
					Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatusOrderedByUpdatedAtDESCParams{

					Version: req.Version,

					Repository: req.Repository,

					ExtensionType: req.ExtensionType.ToInt64(),

					Pro: req.Pro,

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type_and_pro_and_status",
					Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type_and_pro_and_status_invalid",
		Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatus - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatusResponse{}, err

}
