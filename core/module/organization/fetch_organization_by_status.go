package organization

import (
	"context"

	"nem/core/module/organization/types"

	"errors"
	nemdb "nem/core/repository/gen"
	"nem/monitoring"
)

func (m *module) FetchOrganizationByStatus(
	ctx context.Context,
	req types.FetchOrganizationByStatusRequest,
	opts ...Option,
) (types.FetchOrganizationByStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchOrganizationByStatus(
			ctx,
			nemdb.FetchOrganizationByStatusParams{
				Status: req.Status.ToInt64(),
				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_organization_by_status",
				Message:          "error in FetchOrganizationByStatus no order",
				EntityIdentifier: "organization",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchOrganizationByStatusResponse{}, err
		}
		return types.FetchOrganizationByStatusResponse{
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
			ActionIdentifier: "fetch_organization_by_status",
			Message:          "error in FetchOrganizationByStatus - order by field not found",
			EntityIdentifier: "organization",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchOrganizationByStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_organization_by_status",
				Message:          "error in FetchOrganizationByStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "organization",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchOrganizationByStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchOrganizationByStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchOrganizationByStatusOrderedByCreatedAtASCParams{
					Status: req.Status.ToInt64(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_organization_by_status",
					Message:          "error in FetchOrganizationByStatus - ordered asc",
					EntityIdentifier: "organization",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchOrganizationByStatusResponse{}, err
			}
			return types.FetchOrganizationByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchOrganizationByStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchOrganizationByStatusOrderedByCreatedAtDESCParams{
					Status: req.Status.ToInt64(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_organization_by_status",
					Message:          "error in FetchOrganizationByStatus - ordered desc",
					EntityIdentifier: "organization",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchOrganizationByStatusResponse{}, err
			}
			return types.FetchOrganizationByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchOrganizationByStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchOrganizationByStatusOrderedByUpdatedAtASCParams{
					Status: req.Status.ToInt64(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_organization_by_status",
					Message:          "error in FetchOrganizationByStatus - ordered asc",
					EntityIdentifier: "organization",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchOrganizationByStatusResponse{}, err
			}
			return types.FetchOrganizationByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchOrganizationByStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchOrganizationByStatusOrderedByUpdatedAtDESCParams{
					Status: req.Status.ToInt64(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_organization_by_status",
					Message:          "error in FetchOrganizationByStatus - ordered desc",
					EntityIdentifier: "organization",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchOrganizationByStatusResponse{}, err
			}
			return types.FetchOrganizationByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_organization_by_status_invalid",
		Message:          "error in FetchOrganizationByStatus - invalid request",
		EntityIdentifier: "organization",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchOrganizationByStatusResponse{}, err

}
