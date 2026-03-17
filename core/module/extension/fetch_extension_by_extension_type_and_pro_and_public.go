package extension

import (
	"context"

	"github.com/nuzur/nem/core/module/extension/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchExtensionByExtensionTypeAndProAndPublic(
	ctx context.Context,
	req types.FetchExtensionByExtensionTypeAndProAndPublicRequest,
	opts ...Option,
) (types.FetchExtensionByExtensionTypeAndProAndPublicResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByExtensionTypeAndProAndPublic(
			ctx,
			nemdb.FetchExtensionByExtensionTypeAndProAndPublicParams{

				ExtensionType: req.ExtensionType.ToInt64(),

				Pro: req.Pro,

				Public: req.Public,

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_extension_type_and_pro_and_public",
				Message:          "error in FetchExtensionByExtensionTypeAndProAndPublic no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByExtensionTypeAndProAndPublicResponse{}, err
		}
		return types.FetchExtensionByExtensionTypeAndProAndPublicResponse{
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
			ActionIdentifier: "fetch_extension_by_extension_type_and_pro_and_public",
			Message:          "error in FetchExtensionByExtensionTypeAndProAndPublic - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByExtensionTypeAndProAndPublicResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_extension_type_and_pro_and_public",
				Message:          "error in FetchExtensionByExtensionTypeAndProAndPublic - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByExtensionTypeAndProAndPublicResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByExtensionTypeAndProAndPublicOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByExtensionTypeAndProAndPublicOrderedByCreatedAtASCParams{

					ExtensionType: req.ExtensionType.ToInt64(),

					Pro: req.Pro,

					Public: req.Public,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_extension_type_and_pro_and_public",
					Message:          "error in FetchExtensionByExtensionTypeAndProAndPublic - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByExtensionTypeAndProAndPublicResponse{}, err
			}
			return types.FetchExtensionByExtensionTypeAndProAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByExtensionTypeAndProAndPublicOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByExtensionTypeAndProAndPublicOrderedByCreatedAtDESCParams{

					ExtensionType: req.ExtensionType.ToInt64(),

					Pro: req.Pro,

					Public: req.Public,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_extension_type_and_pro_and_public",
					Message:          "error in FetchExtensionByExtensionTypeAndProAndPublic - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByExtensionTypeAndProAndPublicResponse{}, err
			}
			return types.FetchExtensionByExtensionTypeAndProAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByExtensionTypeAndProAndPublicOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByExtensionTypeAndProAndPublicOrderedByUpdatedAtASCParams{

					ExtensionType: req.ExtensionType.ToInt64(),

					Pro: req.Pro,

					Public: req.Public,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_extension_type_and_pro_and_public",
					Message:          "error in FetchExtensionByExtensionTypeAndProAndPublic - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByExtensionTypeAndProAndPublicResponse{}, err
			}
			return types.FetchExtensionByExtensionTypeAndProAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByExtensionTypeAndProAndPublicOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByExtensionTypeAndProAndPublicOrderedByUpdatedAtDESCParams{

					ExtensionType: req.ExtensionType.ToInt64(),

					Pro: req.Pro,

					Public: req.Public,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_extension_type_and_pro_and_public",
					Message:          "error in FetchExtensionByExtensionTypeAndProAndPublic - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByExtensionTypeAndProAndPublicResponse{}, err
			}
			return types.FetchExtensionByExtensionTypeAndProAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_extension_type_and_pro_and_public_invalid",
		Message:          "error in FetchExtensionByExtensionTypeAndProAndPublic - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByExtensionTypeAndProAndPublicResponse{}, err

}
