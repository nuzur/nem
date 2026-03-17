package extension

import (
	"context"

	"github.com/nuzur/nem/core/module/extension/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchExtensionByVersionAndExtensionTypeAndProAndPublicAndStatus(
	ctx context.Context,
	req types.FetchExtensionByVersionAndExtensionTypeAndProAndPublicAndStatusRequest,
	opts ...Option,
) (types.FetchExtensionByVersionAndExtensionTypeAndProAndPublicAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByVersionAndExtensionTypeAndProAndPublicAndStatus(
			ctx,
			nemdb.FetchExtensionByVersionAndExtensionTypeAndProAndPublicAndStatusParams{

				Version: req.Version,

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
				ActionIdentifier: "fetch_extension_by_version_and_extension_type_and_pro_and_public_and_status",
				Message:          "error in FetchExtensionByVersionAndExtensionTypeAndProAndPublicAndStatus no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByVersionAndExtensionTypeAndProAndPublicAndStatusResponse{}, err
		}
		return types.FetchExtensionByVersionAndExtensionTypeAndProAndPublicAndStatusResponse{
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
			ActionIdentifier: "fetch_extension_by_version_and_extension_type_and_pro_and_public_and_status",
			Message:          "error in FetchExtensionByVersionAndExtensionTypeAndProAndPublicAndStatus - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByVersionAndExtensionTypeAndProAndPublicAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_version_and_extension_type_and_pro_and_public_and_status",
				Message:          "error in FetchExtensionByVersionAndExtensionTypeAndProAndPublicAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByVersionAndExtensionTypeAndProAndPublicAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByVersionAndExtensionTypeAndProAndPublicAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByVersionAndExtensionTypeAndProAndPublicAndStatusOrderedByCreatedAtASCParams{

					Version: req.Version,

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
					ActionIdentifier: "fetch_extension_by_version_and_extension_type_and_pro_and_public_and_status",
					Message:          "error in FetchExtensionByVersionAndExtensionTypeAndProAndPublicAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndExtensionTypeAndProAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndExtensionTypeAndProAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByVersionAndExtensionTypeAndProAndPublicAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByVersionAndExtensionTypeAndProAndPublicAndStatusOrderedByCreatedAtDESCParams{

					Version: req.Version,

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
					ActionIdentifier: "fetch_extension_by_version_and_extension_type_and_pro_and_public_and_status",
					Message:          "error in FetchExtensionByVersionAndExtensionTypeAndProAndPublicAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndExtensionTypeAndProAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndExtensionTypeAndProAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByVersionAndExtensionTypeAndProAndPublicAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByVersionAndExtensionTypeAndProAndPublicAndStatusOrderedByUpdatedAtASCParams{

					Version: req.Version,

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
					ActionIdentifier: "fetch_extension_by_version_and_extension_type_and_pro_and_public_and_status",
					Message:          "error in FetchExtensionByVersionAndExtensionTypeAndProAndPublicAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndExtensionTypeAndProAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndExtensionTypeAndProAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByVersionAndExtensionTypeAndProAndPublicAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByVersionAndExtensionTypeAndProAndPublicAndStatusOrderedByUpdatedAtDESCParams{

					Version: req.Version,

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
					ActionIdentifier: "fetch_extension_by_version_and_extension_type_and_pro_and_public_and_status",
					Message:          "error in FetchExtensionByVersionAndExtensionTypeAndProAndPublicAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndExtensionTypeAndProAndPublicAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndExtensionTypeAndProAndPublicAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_version_and_extension_type_and_pro_and_public_and_status_invalid",
		Message:          "error in FetchExtensionByVersionAndExtensionTypeAndProAndPublicAndStatus - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByVersionAndExtensionTypeAndProAndPublicAndStatusResponse{}, err

}
