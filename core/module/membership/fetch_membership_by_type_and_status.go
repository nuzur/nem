package membership

import (
	"context"

	"github.com/nuzur/nem/core/module/membership/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchMembershipByTypeAndStatus(
	ctx context.Context,
	req types.FetchMembershipByTypeAndStatusRequest,
	opts ...Option,
) (types.FetchMembershipByTypeAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchMembershipByTypeAndStatus(
			ctx,
			nemdb.FetchMembershipByTypeAndStatusParams{

				Type: req.Type.ToInt64(),

				Status: req.Status.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_membership_by_type_and_status",
				Message:          "error in FetchMembershipByTypeAndStatus no order",
				EntityIdentifier: "membership",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchMembershipByTypeAndStatusResponse{}, err
		}
		return types.FetchMembershipByTypeAndStatusResponse{
			Results: mapModelsToEntities(models),
		}, nil
	}

	timeFields := []string{
		"start_date", "created_at", "updated_at",
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
			ActionIdentifier: "fetch_membership_by_type_and_status",
			Message:          "error in FetchMembershipByTypeAndStatus - order by field not found",
			EntityIdentifier: "membership",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchMembershipByTypeAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_membership_by_type_and_status",
				Message:          "error in FetchMembershipByTypeAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "membership",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchMembershipByTypeAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "start_date":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchMembershipByTypeAndStatusOrderedByStartDateASC(
				ctx,
				nemdb.FetchMembershipByTypeAndStatusOrderedByStartDateASCParams{

					Type: req.Type.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_membership_by_type_and_status",
					Message:          "error in FetchMembershipByTypeAndStatus - ordered asc",
					EntityIdentifier: "membership",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchMembershipByTypeAndStatusResponse{}, err
			}
			return types.FetchMembershipByTypeAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchMembershipByTypeAndStatusOrderedByStartDateDESC(
				ctx,
				nemdb.FetchMembershipByTypeAndStatusOrderedByStartDateDESCParams{

					Type: req.Type.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_membership_by_type_and_status",
					Message:          "error in FetchMembershipByTypeAndStatus - ordered desc",
					EntityIdentifier: "membership",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchMembershipByTypeAndStatusResponse{}, err
			}
			return types.FetchMembershipByTypeAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchMembershipByTypeAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchMembershipByTypeAndStatusOrderedByCreatedAtASCParams{

					Type: req.Type.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_membership_by_type_and_status",
					Message:          "error in FetchMembershipByTypeAndStatus - ordered asc",
					EntityIdentifier: "membership",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchMembershipByTypeAndStatusResponse{}, err
			}
			return types.FetchMembershipByTypeAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchMembershipByTypeAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchMembershipByTypeAndStatusOrderedByCreatedAtDESCParams{

					Type: req.Type.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_membership_by_type_and_status",
					Message:          "error in FetchMembershipByTypeAndStatus - ordered desc",
					EntityIdentifier: "membership",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchMembershipByTypeAndStatusResponse{}, err
			}
			return types.FetchMembershipByTypeAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchMembershipByTypeAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchMembershipByTypeAndStatusOrderedByUpdatedAtASCParams{

					Type: req.Type.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_membership_by_type_and_status",
					Message:          "error in FetchMembershipByTypeAndStatus - ordered asc",
					EntityIdentifier: "membership",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchMembershipByTypeAndStatusResponse{}, err
			}
			return types.FetchMembershipByTypeAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchMembershipByTypeAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchMembershipByTypeAndStatusOrderedByUpdatedAtDESCParams{

					Type: req.Type.ToInt64(),

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_membership_by_type_and_status",
					Message:          "error in FetchMembershipByTypeAndStatus - ordered desc",
					EntityIdentifier: "membership",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchMembershipByTypeAndStatusResponse{}, err
			}
			return types.FetchMembershipByTypeAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_membership_by_type_and_status_invalid",
		Message:          "error in FetchMembershipByTypeAndStatus - invalid request",
		EntityIdentifier: "membership",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchMembershipByTypeAndStatusResponse{}, err

}
