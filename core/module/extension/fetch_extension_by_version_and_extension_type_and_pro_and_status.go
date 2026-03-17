package extension

import (
	"context"

	"github.com/nuzur/nem/core/module/extension/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchExtensionByVersionAndExtensionTypeAndProAndStatus(
	ctx context.Context,
	req types.FetchExtensionByVersionAndExtensionTypeAndProAndStatusRequest,
	opts ...Option,
) (types.FetchExtensionByVersionAndExtensionTypeAndProAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByVersionAndExtensionTypeAndProAndStatus(
			ctx,
			nemdb.FetchExtensionByVersionAndExtensionTypeAndProAndStatusParams{

				Version: req.Version,

				ExtensionType: req.ExtensionType.ToInt64(),

				Pro: req.Pro,

				Status: req.Status.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_version_and_extension_type_and_pro_and_status",
				Message:          "error in FetchExtensionByVersionAndExtensionTypeAndProAndStatus no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByVersionAndExtensionTypeAndProAndStatusResponse{}, err
		}
		return types.FetchExtensionByVersionAndExtensionTypeAndProAndStatusResponse{
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
			ActionIdentifier: "fetch_extension_by_version_and_extension_type_and_pro_and_status",
			Message:          "error in FetchExtensionByVersionAndExtensionTypeAndProAndStatus - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByVersionAndExtensionTypeAndProAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_version_and_extension_type_and_pro_and_status",
				Message:          "error in FetchExtensionByVersionAndExtensionTypeAndProAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByVersionAndExtensionTypeAndProAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByVersionAndExtensionTypeAndProAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByVersionAndExtensionTypeAndProAndStatusOrderedByCreatedAtASCParams{

					Version: req.Version,

					ExtensionType: req.ExtensionType.ToInt64(),

					Pro: req.Pro,

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_extension_type_and_pro_and_status",
					Message:          "error in FetchExtensionByVersionAndExtensionTypeAndProAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndExtensionTypeAndProAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndExtensionTypeAndProAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByVersionAndExtensionTypeAndProAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByVersionAndExtensionTypeAndProAndStatusOrderedByCreatedAtDESCParams{

					Version: req.Version,

					ExtensionType: req.ExtensionType.ToInt64(),

					Pro: req.Pro,

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_extension_type_and_pro_and_status",
					Message:          "error in FetchExtensionByVersionAndExtensionTypeAndProAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndExtensionTypeAndProAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndExtensionTypeAndProAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByVersionAndExtensionTypeAndProAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByVersionAndExtensionTypeAndProAndStatusOrderedByUpdatedAtASCParams{

					Version: req.Version,

					ExtensionType: req.ExtensionType.ToInt64(),

					Pro: req.Pro,

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_extension_type_and_pro_and_status",
					Message:          "error in FetchExtensionByVersionAndExtensionTypeAndProAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndExtensionTypeAndProAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndExtensionTypeAndProAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByVersionAndExtensionTypeAndProAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByVersionAndExtensionTypeAndProAndStatusOrderedByUpdatedAtDESCParams{

					Version: req.Version,

					ExtensionType: req.ExtensionType.ToInt64(),

					Pro: req.Pro,

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_extension_type_and_pro_and_status",
					Message:          "error in FetchExtensionByVersionAndExtensionTypeAndProAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndExtensionTypeAndProAndStatusResponse{}, err
			}
			return types.FetchExtensionByVersionAndExtensionTypeAndProAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_version_and_extension_type_and_pro_and_status_invalid",
		Message:          "error in FetchExtensionByVersionAndExtensionTypeAndProAndStatus - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByVersionAndExtensionTypeAndProAndStatusResponse{}, err

}
