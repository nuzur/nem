package extension

import (
	"context"

	"github.com/nuzur/nem/core/module/extension/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchExtensionByVersionAndIdentifier(
	ctx context.Context,
	req types.FetchExtensionByVersionAndIdentifierRequest,
	opts ...Option,
) (types.FetchExtensionByVersionAndIdentifierResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifier(
			ctx,
			nemdb.FetchExtensionByVersionAndIdentifierParams{

				Version: req.Version,

				Identifier: req.Identifier,

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_version_and_identifier",
				Message:          "error in FetchExtensionByVersionAndIdentifier no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByVersionAndIdentifierResponse{}, err
		}
		return types.FetchExtensionByVersionAndIdentifierResponse{
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
			ActionIdentifier: "fetch_extension_by_version_and_identifier",
			Message:          "error in FetchExtensionByVersionAndIdentifier - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByVersionAndIdentifierResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_version_and_identifier",
				Message:          "error in FetchExtensionByVersionAndIdentifier - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByVersionAndIdentifierResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierOrderedByCreatedAtASCParams{

					Version: req.Version,

					Identifier: req.Identifier,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_identifier",
					Message:          "error in FetchExtensionByVersionAndIdentifier - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierOrderedByCreatedAtDESCParams{

					Version: req.Version,

					Identifier: req.Identifier,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_identifier",
					Message:          "error in FetchExtensionByVersionAndIdentifier - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierOrderedByUpdatedAtASCParams{

					Version: req.Version,

					Identifier: req.Identifier,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_identifier",
					Message:          "error in FetchExtensionByVersionAndIdentifier - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByVersionAndIdentifierOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByVersionAndIdentifierOrderedByUpdatedAtDESCParams{

					Version: req.Version,

					Identifier: req.Identifier,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_identifier",
					Message:          "error in FetchExtensionByVersionAndIdentifier - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndIdentifierResponse{}, err
			}
			return types.FetchExtensionByVersionAndIdentifierResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_version_and_identifier_invalid",
		Message:          "error in FetchExtensionByVersionAndIdentifier - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByVersionAndIdentifierResponse{}, err

}
