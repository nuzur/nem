package extension

import (
	"context"

	"github.com/nuzur/nem/core/module/extension/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchExtensionByIdentifier(
	ctx context.Context,
	req types.FetchExtensionByIdentifierRequest,
	opts ...Option,
) (types.FetchExtensionByIdentifierResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByIdentifier(
			ctx,
			nemdb.FetchExtensionByIdentifierParams{

				Identifier: req.Identifier,

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_identifier",
				Message:          "error in FetchExtensionByIdentifier no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByIdentifierResponse{}, err
		}
		return types.FetchExtensionByIdentifierResponse{
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
			ActionIdentifier: "fetch_extension_by_identifier",
			Message:          "error in FetchExtensionByIdentifier - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByIdentifierResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_identifier",
				Message:          "error in FetchExtensionByIdentifier - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByIdentifierResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByIdentifierOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByIdentifierOrderedByCreatedAtASCParams{

					Identifier: req.Identifier,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier",
					Message:          "error in FetchExtensionByIdentifier - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierResponse{}, err
			}
			return types.FetchExtensionByIdentifierResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByIdentifierOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByIdentifierOrderedByCreatedAtDESCParams{

					Identifier: req.Identifier,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier",
					Message:          "error in FetchExtensionByIdentifier - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierResponse{}, err
			}
			return types.FetchExtensionByIdentifierResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByIdentifierOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByIdentifierOrderedByUpdatedAtASCParams{

					Identifier: req.Identifier,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier",
					Message:          "error in FetchExtensionByIdentifier - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierResponse{}, err
			}
			return types.FetchExtensionByIdentifierResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByIdentifierOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByIdentifierOrderedByUpdatedAtDESCParams{

					Identifier: req.Identifier,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier",
					Message:          "error in FetchExtensionByIdentifier - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierResponse{}, err
			}
			return types.FetchExtensionByIdentifierResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_identifier_invalid",
		Message:          "error in FetchExtensionByIdentifier - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByIdentifierResponse{}, err

}
