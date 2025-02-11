package extension_execution

import (
	"context"

	"github.com/nuzur/nem/core/module/extension_execution/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchExtensionExecutionByStatus(
	ctx context.Context,
	req types.FetchExtensionExecutionByStatusRequest,
	opts ...Option,
) (types.FetchExtensionExecutionByStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionExecutionByStatus(
			ctx,
			nemdb.FetchExtensionExecutionByStatusParams{
				Status: req.Status.ToInt64(),
				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_execution_by_status",
				Message:          "error in FetchExtensionExecutionByStatus no order",
				EntityIdentifier: "extension_execution",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionExecutionByStatusResponse{}, err
		}
		return types.FetchExtensionExecutionByStatusResponse{
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
			ActionIdentifier: "fetch_extension_execution_by_status",
			Message:          "error in FetchExtensionExecutionByStatus - order by field not found",
			EntityIdentifier: "extension_execution",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionExecutionByStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_execution_by_status",
				Message:          "error in FetchExtensionExecutionByStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension_execution",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionExecutionByStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionExecutionByStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionExecutionByStatusOrderedByCreatedAtASCParams{
					Status: req.Status.ToInt64(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_execution_by_status",
					Message:          "error in FetchExtensionExecutionByStatus - ordered asc",
					EntityIdentifier: "extension_execution",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionExecutionByStatusResponse{}, err
			}
			return types.FetchExtensionExecutionByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionExecutionByStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionExecutionByStatusOrderedByCreatedAtDESCParams{
					Status: req.Status.ToInt64(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_execution_by_status",
					Message:          "error in FetchExtensionExecutionByStatus - ordered desc",
					EntityIdentifier: "extension_execution",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionExecutionByStatusResponse{}, err
			}
			return types.FetchExtensionExecutionByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionExecutionByStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionExecutionByStatusOrderedByUpdatedAtASCParams{
					Status: req.Status.ToInt64(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_execution_by_status",
					Message:          "error in FetchExtensionExecutionByStatus - ordered asc",
					EntityIdentifier: "extension_execution",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionExecutionByStatusResponse{}, err
			}
			return types.FetchExtensionExecutionByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionExecutionByStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionExecutionByStatusOrderedByUpdatedAtDESCParams{
					Status: req.Status.ToInt64(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_execution_by_status",
					Message:          "error in FetchExtensionExecutionByStatus - ordered desc",
					EntityIdentifier: "extension_execution",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionExecutionByStatusResponse{}, err
			}
			return types.FetchExtensionExecutionByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_execution_by_status_invalid",
		Message:          "error in FetchExtensionExecutionByStatus - invalid request",
		EntityIdentifier: "extension_execution",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionExecutionByStatusResponse{}, err

}
