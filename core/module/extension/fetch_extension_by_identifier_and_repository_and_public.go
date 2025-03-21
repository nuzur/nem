package extension

import (
	"context"

	"github.com/nuzur/nem/core/module/extension/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchExtensionByIdentifierAndRepositoryAndPublic(
	ctx context.Context,
	req types.FetchExtensionByIdentifierAndRepositoryAndPublicRequest,
	opts ...Option,
) (types.FetchExtensionByIdentifierAndRepositoryAndPublicResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByIdentifierAndRepositoryAndPublic(
			ctx,
			nemdb.FetchExtensionByIdentifierAndRepositoryAndPublicParams{

				Identifier: req.Identifier,

				Repository: req.Repository,

				Public: req.Public,

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_public",
				Message:          "error in FetchExtensionByIdentifierAndRepositoryAndPublic no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByIdentifierAndRepositoryAndPublicResponse{}, err
		}
		return types.FetchExtensionByIdentifierAndRepositoryAndPublicResponse{
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
			ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_public",
			Message:          "error in FetchExtensionByIdentifierAndRepositoryAndPublic - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByIdentifierAndRepositoryAndPublicResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_public",
				Message:          "error in FetchExtensionByIdentifierAndRepositoryAndPublic - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByIdentifierAndRepositoryAndPublicResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndRepositoryAndPublicOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndRepositoryAndPublicOrderedByCreatedAtASCParams{

					Identifier: req.Identifier,

					Repository: req.Repository,

					Public: req.Public,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_public",
					Message:          "error in FetchExtensionByIdentifierAndRepositoryAndPublic - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndRepositoryAndPublicResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndRepositoryAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndRepositoryAndPublicOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndRepositoryAndPublicOrderedByCreatedAtDESCParams{

					Identifier: req.Identifier,

					Repository: req.Repository,

					Public: req.Public,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_public",
					Message:          "error in FetchExtensionByIdentifierAndRepositoryAndPublic - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndRepositoryAndPublicResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndRepositoryAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndRepositoryAndPublicOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndRepositoryAndPublicOrderedByUpdatedAtASCParams{

					Identifier: req.Identifier,

					Repository: req.Repository,

					Public: req.Public,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_public",
					Message:          "error in FetchExtensionByIdentifierAndRepositoryAndPublic - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndRepositoryAndPublicResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndRepositoryAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndRepositoryAndPublicOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndRepositoryAndPublicOrderedByUpdatedAtDESCParams{

					Identifier: req.Identifier,

					Repository: req.Repository,

					Public: req.Public,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_public",
					Message:          "error in FetchExtensionByIdentifierAndRepositoryAndPublic - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndRepositoryAndPublicResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndRepositoryAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_public_invalid",
		Message:          "error in FetchExtensionByIdentifierAndRepositoryAndPublic - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByIdentifierAndRepositoryAndPublicResponse{}, err

}
