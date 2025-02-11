package extension

import (
	"context"

	"github.com/nuzur/nem/core/module/extension/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchExtensionByIdentifierAndRepositoryAndStatus(
	ctx context.Context,
	req types.FetchExtensionByIdentifierAndRepositoryAndStatusRequest,
	opts ...Option,
) (types.FetchExtensionByIdentifierAndRepositoryAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByIdentifierAndRepositoryAndStatus(
			ctx,
			nemdb.FetchExtensionByIdentifierAndRepositoryAndStatusParams{
				Identifier: req.Identifier,
				Repository: req.Repository,
				Status:     req.Status.ToInt64(),
				Offset:     req.Offset,
				Limit:      req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_status",
				Message:          "error in FetchExtensionByIdentifierAndRepositoryAndStatus no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByIdentifierAndRepositoryAndStatusResponse{}, err
		}
		return types.FetchExtensionByIdentifierAndRepositoryAndStatusResponse{
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
			ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_status",
			Message:          "error in FetchExtensionByIdentifierAndRepositoryAndStatus - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByIdentifierAndRepositoryAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_status",
				Message:          "error in FetchExtensionByIdentifierAndRepositoryAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByIdentifierAndRepositoryAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndRepositoryAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndRepositoryAndStatusOrderedByCreatedAtASCParams{
					Identifier: req.Identifier,
					Repository: req.Repository,
					Status:     req.Status.ToInt64(),
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_status",
					Message:          "error in FetchExtensionByIdentifierAndRepositoryAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndRepositoryAndStatusResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndRepositoryAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndRepositoryAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndRepositoryAndStatusOrderedByCreatedAtDESCParams{
					Identifier: req.Identifier,
					Repository: req.Repository,
					Status:     req.Status.ToInt64(),
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_status",
					Message:          "error in FetchExtensionByIdentifierAndRepositoryAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndRepositoryAndStatusResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndRepositoryAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndRepositoryAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndRepositoryAndStatusOrderedByUpdatedAtASCParams{
					Identifier: req.Identifier,
					Repository: req.Repository,
					Status:     req.Status.ToInt64(),
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_status",
					Message:          "error in FetchExtensionByIdentifierAndRepositoryAndStatus - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndRepositoryAndStatusResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndRepositoryAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndRepositoryAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndRepositoryAndStatusOrderedByUpdatedAtDESCParams{
					Identifier: req.Identifier,
					Repository: req.Repository,
					Status:     req.Status.ToInt64(),
					Offset:     req.Offset,
					Limit:      req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_status",
					Message:          "error in FetchExtensionByIdentifierAndRepositoryAndStatus - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndRepositoryAndStatusResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndRepositoryAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_status_invalid",
		Message:          "error in FetchExtensionByIdentifierAndRepositoryAndStatus - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByIdentifierAndRepositoryAndStatusResponse{}, err

}
