package extension

import (
	"context"

	"github.com/nuzur/nem/core/module/extension/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchExtensionByRepositoryAndExtensionTypeAndProAndStatus(
	ctx context.Context,
	req types.FetchExtensionByRepositoryAndExtensionTypeAndProAndStatusRequest,
	opts ...Option,
) (types.FetchExtensionByRepositoryAndExtensionTypeAndProAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByRepositoryAndExtensionTypeAndProAndStatus(
			ctx,
			nemdb.FetchExtensionByRepositoryAndExtensionTypeAndProAndStatusParams{

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
				ActionIdentifier: "fetch_extension_by_repository_and_extension_type_and_pro_and_status",
				Message:          "error in FetchExtensionByRepositoryAndExtensionTypeAndProAndStatus no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByRepositoryAndExtensionTypeAndProAndStatusResponse{}, err
		}
		return types.FetchExtensionByRepositoryAndExtensionTypeAndProAndStatusResponse{
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
			ActionIdentifier: "fetch_extension_by_repository_and_extension_type_and_pro_and_status",
			Message:          "error in FetchExtensionByRepositoryAndExtensionTypeAndProAndStatus - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByRepositoryAndExtensionTypeAndProAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_repository_and_extension_type_and_pro_and_status",
				Message:          "error in FetchExtensionByRepositoryAndExtensionTypeAndProAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByRepositoryAndExtensionTypeAndProAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByRepositoryAndExtensionTypeAndProAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByRepositoryAndExtensionTypeAndProAndStatusOrderedByCreatedAtASCParams{

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
					ActionIdentifier: "fetch_extension_by_repository_and_extension_type_and_pro_and_status",
					Message:          "error in FetchExtensionByRepositoryAndExtensionTypeAndProAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByRepositoryAndExtensionTypeAndProAndStatusResponse{}, err
			}
			return types.FetchExtensionByRepositoryAndExtensionTypeAndProAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByRepositoryAndExtensionTypeAndProAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByRepositoryAndExtensionTypeAndProAndStatusOrderedByCreatedAtDESCParams{

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
					ActionIdentifier: "fetch_extension_by_repository_and_extension_type_and_pro_and_status",
					Message:          "error in FetchExtensionByRepositoryAndExtensionTypeAndProAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByRepositoryAndExtensionTypeAndProAndStatusResponse{}, err
			}
			return types.FetchExtensionByRepositoryAndExtensionTypeAndProAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByRepositoryAndExtensionTypeAndProAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByRepositoryAndExtensionTypeAndProAndStatusOrderedByUpdatedAtASCParams{

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
					ActionIdentifier: "fetch_extension_by_repository_and_extension_type_and_pro_and_status",
					Message:          "error in FetchExtensionByRepositoryAndExtensionTypeAndProAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByRepositoryAndExtensionTypeAndProAndStatusResponse{}, err
			}
			return types.FetchExtensionByRepositoryAndExtensionTypeAndProAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByRepositoryAndExtensionTypeAndProAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByRepositoryAndExtensionTypeAndProAndStatusOrderedByUpdatedAtDESCParams{

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
					ActionIdentifier: "fetch_extension_by_repository_and_extension_type_and_pro_and_status",
					Message:          "error in FetchExtensionByRepositoryAndExtensionTypeAndProAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByRepositoryAndExtensionTypeAndProAndStatusResponse{}, err
			}
			return types.FetchExtensionByRepositoryAndExtensionTypeAndProAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_repository_and_extension_type_and_pro_and_status_invalid",
		Message:          "error in FetchExtensionByRepositoryAndExtensionTypeAndProAndStatus - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByRepositoryAndExtensionTypeAndProAndStatusResponse{}, err

}
