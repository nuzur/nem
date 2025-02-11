package extension

import (
	"context"

	"github.com/nuzur/nem/core/module/extension/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchExtensionByIdentifierAndStatus(
	ctx context.Context,
	req types.FetchExtensionByIdentifierAndStatusRequest,
	opts ...Option,
) (types.FetchExtensionByIdentifierAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByIdentifierAndStatus(
			ctx,
			nemdb.FetchExtensionByIdentifierAndStatusParams{
				Identifier: req.Identifier,
				Status:     req.Status.ToInt64(),
				Offset:     req.Offset,
				Limit:      req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_identifier_and_status",
				Message:          "error in FetchExtensionByIdentifierAndStatus no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByIdentifierAndStatusResponse{}, err
		}
		return types.FetchExtensionByIdentifierAndStatusResponse{
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
			ActionIdentifier: "fetch_extension_by_identifier_and_status",
			Message:          "error in FetchExtensionByIdentifierAndStatus - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByIdentifierAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_identifier_and_status",
				Message:          "error in FetchExtensionByIdentifierAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByIdentifierAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndStatusOrderedByCreatedAtASCParams{
					Identifier: req.Identifier,
					Status:     req.Status.ToInt64(),
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_status",
					Message:          "error in FetchExtensionByIdentifierAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndStatusResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndStatusOrderedByCreatedAtDESCParams{
					Identifier: req.Identifier,
					Status:     req.Status.ToInt64(),
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_status",
					Message:          "error in FetchExtensionByIdentifierAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndStatusResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndStatusOrderedByUpdatedAtASCParams{
					Identifier: req.Identifier,
					Status:     req.Status.ToInt64(),
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_status",
					Message:          "error in FetchExtensionByIdentifierAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndStatusResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndStatusOrderedByUpdatedAtDESCParams{
					Identifier: req.Identifier,
					Status:     req.Status.ToInt64(),
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_status",
					Message:          "error in FetchExtensionByIdentifierAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndStatusResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_identifier_and_status_invalid",
		Message:          "error in FetchExtensionByIdentifierAndStatus - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByIdentifierAndStatusResponse{}, err

}
