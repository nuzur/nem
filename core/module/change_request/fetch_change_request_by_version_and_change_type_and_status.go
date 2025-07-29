package change_request

import (
	"context"

	"github.com/nuzur/nem/core/module/change_request/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchChangeRequestByVersionAndChangeTypeAndStatus(
	ctx context.Context,
	req types.FetchChangeRequestByVersionAndChangeTypeAndStatusRequest,
	opts ...Option,
) (types.FetchChangeRequestByVersionAndChangeTypeAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchChangeRequestByVersionAndChangeTypeAndStatus(
			ctx,
			nemdb.FetchChangeRequestByVersionAndChangeTypeAndStatusParams{

				Version: req.Version,

				ChangeType: req.ChangeType.ToInt64(),

				Status: req.Status.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_change_request_by_version_and_change_type_and_status",
				Message:          "error in FetchChangeRequestByVersionAndChangeTypeAndStatus no order",
				EntityIdentifier: "change_request",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchChangeRequestByVersionAndChangeTypeAndStatusResponse{}, err
		}
		return types.FetchChangeRequestByVersionAndChangeTypeAndStatusResponse{
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
			ActionIdentifier: "fetch_change_request_by_version_and_change_type_and_status",
			Message:          "error in FetchChangeRequestByVersionAndChangeTypeAndStatus - order by field not found",
			EntityIdentifier: "change_request",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchChangeRequestByVersionAndChangeTypeAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_change_request_by_version_and_change_type_and_status",
				Message:          "error in FetchChangeRequestByVersionAndChangeTypeAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "change_request",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchChangeRequestByVersionAndChangeTypeAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchChangeRequestByVersionAndChangeTypeAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchChangeRequestByVersionAndChangeTypeAndStatusOrderedByCreatedAtASCParams{

					Version: req.Version,

					ChangeType: req.ChangeType.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_version_and_change_type_and_status",
					Message:          "error in FetchChangeRequestByVersionAndChangeTypeAndStatus - ordered asc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByVersionAndChangeTypeAndStatusResponse{}, err
			}
			return types.FetchChangeRequestByVersionAndChangeTypeAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchChangeRequestByVersionAndChangeTypeAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchChangeRequestByVersionAndChangeTypeAndStatusOrderedByCreatedAtDESCParams{

					Version: req.Version,

					ChangeType: req.ChangeType.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_version_and_change_type_and_status",
					Message:          "error in FetchChangeRequestByVersionAndChangeTypeAndStatus - ordered desc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByVersionAndChangeTypeAndStatusResponse{}, err
			}
			return types.FetchChangeRequestByVersionAndChangeTypeAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchChangeRequestByVersionAndChangeTypeAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchChangeRequestByVersionAndChangeTypeAndStatusOrderedByUpdatedAtASCParams{

					Version: req.Version,

					ChangeType: req.ChangeType.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_version_and_change_type_and_status",
					Message:          "error in FetchChangeRequestByVersionAndChangeTypeAndStatus - ordered asc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByVersionAndChangeTypeAndStatusResponse{}, err
			}
			return types.FetchChangeRequestByVersionAndChangeTypeAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchChangeRequestByVersionAndChangeTypeAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchChangeRequestByVersionAndChangeTypeAndStatusOrderedByUpdatedAtDESCParams{

					Version: req.Version,

					ChangeType: req.ChangeType.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_version_and_change_type_and_status",
					Message:          "error in FetchChangeRequestByVersionAndChangeTypeAndStatus - ordered desc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByVersionAndChangeTypeAndStatusResponse{}, err
			}
			return types.FetchChangeRequestByVersionAndChangeTypeAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_change_request_by_version_and_change_type_and_status_invalid",
		Message:          "error in FetchChangeRequestByVersionAndChangeTypeAndStatus - invalid request",
		EntityIdentifier: "change_request",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchChangeRequestByVersionAndChangeTypeAndStatusResponse{}, err

}
