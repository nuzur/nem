package extension

import (
	"context"

	"github.com/nuzur/nem/core/module/extension/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchExtensionByIdentifierAndExtensionType(
	ctx context.Context,
	req types.FetchExtensionByIdentifierAndExtensionTypeRequest,
	opts ...Option,
) (types.FetchExtensionByIdentifierAndExtensionTypeResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByIdentifierAndExtensionType(
			ctx,
			nemdb.FetchExtensionByIdentifierAndExtensionTypeParams{
				Identifier:    req.Identifier,
				ExtensionType: req.ExtensionType.ToInt64(),
				Offset:        req.Offset,
				Limit:         req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_identifier_and_extension_type",
				Message:          "error in FetchExtensionByIdentifierAndExtensionType no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByIdentifierAndExtensionTypeResponse{}, err
		}
		return types.FetchExtensionByIdentifierAndExtensionTypeResponse{
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
			ActionIdentifier: "fetch_extension_by_identifier_and_extension_type",
			Message:          "error in FetchExtensionByIdentifierAndExtensionType - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByIdentifierAndExtensionTypeResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_identifier_and_extension_type",
				Message:          "error in FetchExtensionByIdentifierAndExtensionType - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByIdentifierAndExtensionTypeResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndExtensionTypeOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndExtensionTypeOrderedByCreatedAtASCParams{
					Identifier:    req.Identifier,
					ExtensionType: req.ExtensionType.ToInt64(),
					Offset:        req.Offset,
					Limit:         req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_extension_type",
					Message:          "error in FetchExtensionByIdentifierAndExtensionType - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndExtensionTypeResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndExtensionTypeResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndExtensionTypeOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndExtensionTypeOrderedByCreatedAtDESCParams{
					Identifier:    req.Identifier,
					ExtensionType: req.ExtensionType.ToInt64(),
					Offset:        req.Offset,
					Limit:         req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_extension_type",
					Message:          "error in FetchExtensionByIdentifierAndExtensionType - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndExtensionTypeResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndExtensionTypeResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndExtensionTypeOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndExtensionTypeOrderedByUpdatedAtASCParams{
					Identifier:    req.Identifier,
					ExtensionType: req.ExtensionType.ToInt64(),
					Offset:        req.Offset,
					Limit:         req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_extension_type",
					Message:          "error in FetchExtensionByIdentifierAndExtensionType - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndExtensionTypeResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndExtensionTypeResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndExtensionTypeOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndExtensionTypeOrderedByUpdatedAtDESCParams{
					Identifier:    req.Identifier,
					ExtensionType: req.ExtensionType.ToInt64(),
					Offset:        req.Offset,
					Limit:         req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_extension_type",
					Message:          "error in FetchExtensionByIdentifierAndExtensionType - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndExtensionTypeResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndExtensionTypeResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_identifier_and_extension_type_invalid",
		Message:          "error in FetchExtensionByIdentifierAndExtensionType - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByIdentifierAndExtensionTypeResponse{}, err

}
