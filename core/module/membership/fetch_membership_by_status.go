package membership

import (
	"context"

	"github.com/nuzur/nem/core/module/membership/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchMembershipByStatus(
	ctx context.Context,
	req types.FetchMembershipByStatusRequest,
	opts ...Option,
) (types.FetchMembershipByStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchMembershipByStatus(
			ctx,
			nemdb.FetchMembershipByStatusParams{

				Status: req.Status.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_membership_by_status",
				Message:          "error in FetchMembershipByStatus no order",
				EntityIdentifier: "membership",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchMembershipByStatusResponse{}, err
		}
		return types.FetchMembershipByStatusResponse{
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
			ActionIdentifier: "fetch_membership_by_status",
			Message:          "error in FetchMembershipByStatus - order by field not found",
			EntityIdentifier: "membership",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchMembershipByStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_membership_by_status",
				Message:          "error in FetchMembershipByStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "membership",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchMembershipByStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "start_date":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchMembershipByStatusOrderedByStartDateASC(
				ctx,
				nemdb.FetchMembershipByStatusOrderedByStartDateASCParams{

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_membership_by_status",
					Message:          "error in FetchMembershipByStatus - ordered asc",
					EntityIdentifier: "membership",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchMembershipByStatusResponse{}, err
			}
			return types.FetchMembershipByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchMembershipByStatusOrderedByStartDateDESC(
				ctx,
				nemdb.FetchMembershipByStatusOrderedByStartDateDESCParams{

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_membership_by_status",
					Message:          "error in FetchMembershipByStatus - ordered desc",
					EntityIdentifier: "membership",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchMembershipByStatusResponse{}, err
			}
			return types.FetchMembershipByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchMembershipByStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchMembershipByStatusOrderedByCreatedAtASCParams{

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_membership_by_status",
					Message:          "error in FetchMembershipByStatus - ordered asc",
					EntityIdentifier: "membership",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchMembershipByStatusResponse{}, err
			}
			return types.FetchMembershipByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchMembershipByStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchMembershipByStatusOrderedByCreatedAtDESCParams{

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_membership_by_status",
					Message:          "error in FetchMembershipByStatus - ordered desc",
					EntityIdentifier: "membership",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchMembershipByStatusResponse{}, err
			}
			return types.FetchMembershipByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchMembershipByStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchMembershipByStatusOrderedByUpdatedAtASCParams{

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_membership_by_status",
					Message:          "error in FetchMembershipByStatus - ordered asc",
					EntityIdentifier: "membership",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchMembershipByStatusResponse{}, err
			}
			return types.FetchMembershipByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchMembershipByStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchMembershipByStatusOrderedByUpdatedAtDESCParams{

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_membership_by_status",
					Message:          "error in FetchMembershipByStatus - ordered desc",
					EntityIdentifier: "membership",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchMembershipByStatusResponse{}, err
			}
			return types.FetchMembershipByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_membership_by_status_invalid",
		Message:          "error in FetchMembershipByStatus - invalid request",
		EntityIdentifier: "membership",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchMembershipByStatusResponse{}, err

}
