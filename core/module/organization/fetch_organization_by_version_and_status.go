package organization

import (
	"context"

	"github.com/nuzur/nem/core/module/organization/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchOrganizationByVersionAndStatus(
	ctx context.Context,
	req types.FetchOrganizationByVersionAndStatusRequest,
	opts ...Option,
) (types.FetchOrganizationByVersionAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchOrganizationByVersionAndStatus(
			ctx,
			nemdb.FetchOrganizationByVersionAndStatusParams{

				Version: req.Version,

				Status: req.Status.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_organization_by_version_and_status",
				Message:          "error in FetchOrganizationByVersionAndStatus no order",
				EntityIdentifier: "organization",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchOrganizationByVersionAndStatusResponse{}, err
		}
		return types.FetchOrganizationByVersionAndStatusResponse{
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
			ActionIdentifier: "fetch_organization_by_version_and_status",
			Message:          "error in FetchOrganizationByVersionAndStatus - order by field not found",
			EntityIdentifier: "organization",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchOrganizationByVersionAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_organization_by_version_and_status",
				Message:          "error in FetchOrganizationByVersionAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "organization",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchOrganizationByVersionAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchOrganizationByVersionAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchOrganizationByVersionAndStatusOrderedByCreatedAtASCParams{

					Version: req.Version,

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_organization_by_version_and_status",
					Message:          "error in FetchOrganizationByVersionAndStatus - ordered asc",
					EntityIdentifier: "organization",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchOrganizationByVersionAndStatusResponse{}, err
			}
			return types.FetchOrganizationByVersionAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchOrganizationByVersionAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchOrganizationByVersionAndStatusOrderedByCreatedAtDESCParams{

					Version: req.Version,

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_organization_by_version_and_status",
					Message:          "error in FetchOrganizationByVersionAndStatus - ordered desc",
					EntityIdentifier: "organization",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchOrganizationByVersionAndStatusResponse{}, err
			}
			return types.FetchOrganizationByVersionAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchOrganizationByVersionAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchOrganizationByVersionAndStatusOrderedByUpdatedAtASCParams{

					Version: req.Version,

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_organization_by_version_and_status",
					Message:          "error in FetchOrganizationByVersionAndStatus - ordered asc",
					EntityIdentifier: "organization",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchOrganizationByVersionAndStatusResponse{}, err
			}
			return types.FetchOrganizationByVersionAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchOrganizationByVersionAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchOrganizationByVersionAndStatusOrderedByUpdatedAtDESCParams{

					Version: req.Version,

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_organization_by_version_and_status",
					Message:          "error in FetchOrganizationByVersionAndStatus - ordered desc",
					EntityIdentifier: "organization",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchOrganizationByVersionAndStatusResponse{}, err
			}
			return types.FetchOrganizationByVersionAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_organization_by_version_and_status_invalid",
		Message:          "error in FetchOrganizationByVersionAndStatus - invalid request",
		EntityIdentifier: "organization",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchOrganizationByVersionAndStatusResponse{}, err

}
