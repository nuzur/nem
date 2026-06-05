package local_agent

import (
	"context"

	"github.com/nuzur/nem/core/module/local_agent/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchLocalAgentByStatus(
	ctx context.Context,
	req types.FetchLocalAgentByStatusRequest,
	opts ...Option,
) (types.FetchLocalAgentByStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchLocalAgentByStatus(
			ctx,
			nemdb.FetchLocalAgentByStatusParams{

				Status: req.Status.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_local_agent_by_status",
				Message:          "error in FetchLocalAgentByStatus no order",
				EntityIdentifier: "local_agent",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchLocalAgentByStatusResponse{}, err
		}
		return types.FetchLocalAgentByStatusResponse{
			Results: mapModelsToEntities(models),
		}, nil
	}

	timeFields := []string{
		"last_seen_at", "revoked_at", "created_at", "updated_at",
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
			ActionIdentifier: "fetch_local_agent_by_status",
			Message:          "error in FetchLocalAgentByStatus - order by field not found",
			EntityIdentifier: "local_agent",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchLocalAgentByStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_local_agent_by_status",
				Message:          "error in FetchLocalAgentByStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "local_agent",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchLocalAgentByStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "last_seen_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchLocalAgentByStatusOrderedByLastSeenAtASC(
				ctx,
				nemdb.FetchLocalAgentByStatusOrderedByLastSeenAtASCParams{

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_local_agent_by_status",
					Message:          "error in FetchLocalAgentByStatus - ordered asc",
					EntityIdentifier: "local_agent",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchLocalAgentByStatusResponse{}, err
			}
			return types.FetchLocalAgentByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchLocalAgentByStatusOrderedByLastSeenAtDESC(
				ctx,
				nemdb.FetchLocalAgentByStatusOrderedByLastSeenAtDESCParams{

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_local_agent_by_status",
					Message:          "error in FetchLocalAgentByStatus - ordered desc",
					EntityIdentifier: "local_agent",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchLocalAgentByStatusResponse{}, err
			}
			return types.FetchLocalAgentByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "revoked_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchLocalAgentByStatusOrderedByRevokedAtASC(
				ctx,
				nemdb.FetchLocalAgentByStatusOrderedByRevokedAtASCParams{

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_local_agent_by_status",
					Message:          "error in FetchLocalAgentByStatus - ordered asc",
					EntityIdentifier: "local_agent",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchLocalAgentByStatusResponse{}, err
			}
			return types.FetchLocalAgentByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchLocalAgentByStatusOrderedByRevokedAtDESC(
				ctx,
				nemdb.FetchLocalAgentByStatusOrderedByRevokedAtDESCParams{

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_local_agent_by_status",
					Message:          "error in FetchLocalAgentByStatus - ordered desc",
					EntityIdentifier: "local_agent",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchLocalAgentByStatusResponse{}, err
			}
			return types.FetchLocalAgentByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchLocalAgentByStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchLocalAgentByStatusOrderedByCreatedAtASCParams{

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_local_agent_by_status",
					Message:          "error in FetchLocalAgentByStatus - ordered asc",
					EntityIdentifier: "local_agent",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchLocalAgentByStatusResponse{}, err
			}
			return types.FetchLocalAgentByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchLocalAgentByStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchLocalAgentByStatusOrderedByCreatedAtDESCParams{

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_local_agent_by_status",
					Message:          "error in FetchLocalAgentByStatus - ordered desc",
					EntityIdentifier: "local_agent",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchLocalAgentByStatusResponse{}, err
			}
			return types.FetchLocalAgentByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchLocalAgentByStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchLocalAgentByStatusOrderedByUpdatedAtASCParams{

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_local_agent_by_status",
					Message:          "error in FetchLocalAgentByStatus - ordered asc",
					EntityIdentifier: "local_agent",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchLocalAgentByStatusResponse{}, err
			}
			return types.FetchLocalAgentByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchLocalAgentByStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchLocalAgentByStatusOrderedByUpdatedAtDESCParams{

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_local_agent_by_status",
					Message:          "error in FetchLocalAgentByStatus - ordered desc",
					EntityIdentifier: "local_agent",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchLocalAgentByStatusResponse{}, err
			}
			return types.FetchLocalAgentByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_local_agent_by_status_invalid",
		Message:          "error in FetchLocalAgentByStatus - invalid request",
		EntityIdentifier: "local_agent",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchLocalAgentByStatusResponse{}, err

}
