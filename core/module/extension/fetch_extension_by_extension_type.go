package extension

import (
	"context"

	"github.com/nuzur/nem/core/module/extension/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchExtensionByExtensionType(
	ctx context.Context,
	req types.FetchExtensionByExtensionTypeRequest,
	opts ...Option,
) (types.FetchExtensionByExtensionTypeResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByExtensionType(
			ctx,
			nemdb.FetchExtensionByExtensionTypeParams{
				ExtensionType: req.ExtensionType.ToInt64(),
				Offset:        req.Offset,
				Limit:         req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_extension_type",
				Message:          "error in FetchExtensionByExtensionType no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByExtensionTypeResponse{}, err
		}
		return types.FetchExtensionByExtensionTypeResponse{
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
			ActionIdentifier: "fetch_extension_by_extension_type",
			Message:          "error in FetchExtensionByExtensionType - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByExtensionTypeResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_extension_type",
				Message:          "error in FetchExtensionByExtensionType - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByExtensionTypeResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByExtensionTypeOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByExtensionTypeOrderedByCreatedAtASCParams{
					ExtensionType: req.ExtensionType.ToInt64(),
					Offset:        req.Offset,
					Limit:         req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_extension_type",
					Message:          "error in FetchExtensionByExtensionType - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByExtensionTypeResponse{}, err
			}
			return types.FetchExtensionByExtensionTypeResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByExtensionTypeOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByExtensionTypeOrderedByCreatedAtDESCParams{
					ExtensionType: req.ExtensionType.ToInt64(),
					Offset:        req.Offset,
					Limit:         req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_extension_type",
					Message:          "error in FetchExtensionByExtensionType - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByExtensionTypeResponse{}, err
			}
			return types.FetchExtensionByExtensionTypeResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByExtensionTypeOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByExtensionTypeOrderedByUpdatedAtASCParams{
					ExtensionType: req.ExtensionType.ToInt64(),
					Offset:        req.Offset,
					Limit:         req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_extension_type",
					Message:          "error in FetchExtensionByExtensionType - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByExtensionTypeResponse{}, err
			}
			return types.FetchExtensionByExtensionTypeResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByExtensionTypeOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByExtensionTypeOrderedByUpdatedAtDESCParams{
					ExtensionType: req.ExtensionType.ToInt64(),
					Offset:        req.Offset,
					Limit:         req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_extension_type",
					Message:          "error in FetchExtensionByExtensionType - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByExtensionTypeResponse{}, err
			}
			return types.FetchExtensionByExtensionTypeResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_extension_type_invalid",
		Message:          "error in FetchExtensionByExtensionType - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByExtensionTypeResponse{}, err

}
