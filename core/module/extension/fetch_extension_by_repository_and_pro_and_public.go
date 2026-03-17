package extension

import (
	"context"

	"github.com/nuzur/nem/core/module/extension/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchExtensionByRepositoryAndProAndPublic(
	ctx context.Context,
	req types.FetchExtensionByRepositoryAndProAndPublicRequest,
	opts ...Option,
) (types.FetchExtensionByRepositoryAndProAndPublicResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByRepositoryAndProAndPublic(
			ctx,
			nemdb.FetchExtensionByRepositoryAndProAndPublicParams{

				Repository: req.Repository,

				Pro: req.Pro,

				Public: req.Public,

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_repository_and_pro_and_public",
				Message:          "error in FetchExtensionByRepositoryAndProAndPublic no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByRepositoryAndProAndPublicResponse{}, err
		}
		return types.FetchExtensionByRepositoryAndProAndPublicResponse{
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
			ActionIdentifier: "fetch_extension_by_repository_and_pro_and_public",
			Message:          "error in FetchExtensionByRepositoryAndProAndPublic - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByRepositoryAndProAndPublicResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_repository_and_pro_and_public",
				Message:          "error in FetchExtensionByRepositoryAndProAndPublic - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByRepositoryAndProAndPublicResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByRepositoryAndProAndPublicOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByRepositoryAndProAndPublicOrderedByCreatedAtASCParams{

					Repository: req.Repository,

					Pro: req.Pro,

					Public: req.Public,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_repository_and_pro_and_public",
					Message:          "error in FetchExtensionByRepositoryAndProAndPublic - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByRepositoryAndProAndPublicResponse{}, err
			}
			return types.FetchExtensionByRepositoryAndProAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByRepositoryAndProAndPublicOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByRepositoryAndProAndPublicOrderedByCreatedAtDESCParams{

					Repository: req.Repository,

					Pro: req.Pro,

					Public: req.Public,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_repository_and_pro_and_public",
					Message:          "error in FetchExtensionByRepositoryAndProAndPublic - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByRepositoryAndProAndPublicResponse{}, err
			}
			return types.FetchExtensionByRepositoryAndProAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByRepositoryAndProAndPublicOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByRepositoryAndProAndPublicOrderedByUpdatedAtASCParams{

					Repository: req.Repository,

					Pro: req.Pro,

					Public: req.Public,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_repository_and_pro_and_public",
					Message:          "error in FetchExtensionByRepositoryAndProAndPublic - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByRepositoryAndProAndPublicResponse{}, err
			}
			return types.FetchExtensionByRepositoryAndProAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByRepositoryAndProAndPublicOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByRepositoryAndProAndPublicOrderedByUpdatedAtDESCParams{

					Repository: req.Repository,

					Pro: req.Pro,

					Public: req.Public,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_repository_and_pro_and_public",
					Message:          "error in FetchExtensionByRepositoryAndProAndPublic - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByRepositoryAndProAndPublicResponse{}, err
			}
			return types.FetchExtensionByRepositoryAndProAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_repository_and_pro_and_public_invalid",
		Message:          "error in FetchExtensionByRepositoryAndProAndPublic - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByRepositoryAndProAndPublicResponse{}, err

}
