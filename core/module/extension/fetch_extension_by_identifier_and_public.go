package extension

import (
	"context"

	"github.com/nuzur/nem/core/module/extension/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchExtensionByIdentifierAndPublic(
	ctx context.Context,
	req types.FetchExtensionByIdentifierAndPublicRequest,
	opts ...Option,
) (types.FetchExtensionByIdentifierAndPublicResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByIdentifierAndPublic(
			ctx,
			nemdb.FetchExtensionByIdentifierAndPublicParams{
				Identifier: req.Identifier,
				Public:     req.Public,
				Offset:     req.Offset,
				Limit:      req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_identifier_and_public",
				Message:          "error in FetchExtensionByIdentifierAndPublic no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByIdentifierAndPublicResponse{}, err
		}
		return types.FetchExtensionByIdentifierAndPublicResponse{
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
			ActionIdentifier: "fetch_extension_by_identifier_and_public",
			Message:          "error in FetchExtensionByIdentifierAndPublic - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByIdentifierAndPublicResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_identifier_and_public",
				Message:          "error in FetchExtensionByIdentifierAndPublic - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByIdentifierAndPublicResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndPublicOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndPublicOrderedByCreatedAtASCParams{
					Identifier: req.Identifier,
					Public:     req.Public,
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_public",
					Message:          "error in FetchExtensionByIdentifierAndPublic - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndPublicResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndPublicOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndPublicOrderedByCreatedAtDESCParams{
					Identifier: req.Identifier,
					Public:     req.Public,
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_public",
					Message:          "error in FetchExtensionByIdentifierAndPublic - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndPublicResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndPublicOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndPublicOrderedByUpdatedAtASCParams{
					Identifier: req.Identifier,
					Public:     req.Public,
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_public",
					Message:          "error in FetchExtensionByIdentifierAndPublic - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndPublicResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndPublicOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndPublicOrderedByUpdatedAtDESCParams{
					Identifier: req.Identifier,
					Public:     req.Public,
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_public",
					Message:          "error in FetchExtensionByIdentifierAndPublic - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndPublicResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndPublicResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_identifier_and_public_invalid",
		Message:          "error in FetchExtensionByIdentifierAndPublic - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByIdentifierAndPublicResponse{}, err

}
