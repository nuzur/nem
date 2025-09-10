package membership

import (
	"context"

	"github.com/nuzur/nem/core/module/membership/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchMembershipByType(
	ctx context.Context,
	req types.FetchMembershipByTypeRequest,
	opts ...Option,
) (types.FetchMembershipByTypeResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchMembershipByType(
			ctx,
			nemdb.FetchMembershipByTypeParams{

				Type: req.Type.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_membership_by_type",
				Message:          "error in FetchMembershipByType no order",
				EntityIdentifier: "membership",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchMembershipByTypeResponse{}, err
		}
		return types.FetchMembershipByTypeResponse{
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
			ActionIdentifier: "fetch_membership_by_type",
			Message:          "error in FetchMembershipByType - order by field not found",
			EntityIdentifier: "membership",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchMembershipByTypeResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_membership_by_type",
				Message:          "error in FetchMembershipByType - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "membership",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchMembershipByTypeResponse{}, err
		}
	}

	switch req.OrderBy {

	case "start_date":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchMembershipByTypeOrderedByStartDateASC(
				ctx,
				nemdb.FetchMembershipByTypeOrderedByStartDateASCParams{

					Type: req.Type.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_membership_by_type",
					Message:          "error in FetchMembershipByType - ordered asc",
					EntityIdentifier: "membership",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchMembershipByTypeResponse{}, err
			}
			return types.FetchMembershipByTypeResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchMembershipByTypeOrderedByStartDateDESC(
				ctx,
				nemdb.FetchMembershipByTypeOrderedByStartDateDESCParams{

					Type: req.Type.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_membership_by_type",
					Message:          "error in FetchMembershipByType - ordered desc",
					EntityIdentifier: "membership",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchMembershipByTypeResponse{}, err
			}
			return types.FetchMembershipByTypeResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchMembershipByTypeOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchMembershipByTypeOrderedByCreatedAtASCParams{

					Type: req.Type.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_membership_by_type",
					Message:          "error in FetchMembershipByType - ordered asc",
					EntityIdentifier: "membership",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchMembershipByTypeResponse{}, err
			}
			return types.FetchMembershipByTypeResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchMembershipByTypeOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchMembershipByTypeOrderedByCreatedAtDESCParams{

					Type: req.Type.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_membership_by_type",
					Message:          "error in FetchMembershipByType - ordered desc",
					EntityIdentifier: "membership",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchMembershipByTypeResponse{}, err
			}
			return types.FetchMembershipByTypeResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchMembershipByTypeOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchMembershipByTypeOrderedByUpdatedAtASCParams{

					Type: req.Type.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_membership_by_type",
					Message:          "error in FetchMembershipByType - ordered asc",
					EntityIdentifier: "membership",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchMembershipByTypeResponse{}, err
			}
			return types.FetchMembershipByTypeResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchMembershipByTypeOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchMembershipByTypeOrderedByUpdatedAtDESCParams{

					Type: req.Type.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_membership_by_type",
					Message:          "error in FetchMembershipByType - ordered desc",
					EntityIdentifier: "membership",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchMembershipByTypeResponse{}, err
			}
			return types.FetchMembershipByTypeResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_membership_by_type_invalid",
		Message:          "error in FetchMembershipByType - invalid request",
		EntityIdentifier: "membership",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchMembershipByTypeResponse{}, err

}
