package organization

import (
	"context"

	"github.com/nuzur/nem/core/module/organization/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchOrganizationByVersion(
	ctx context.Context,
	req types.FetchOrganizationByVersionRequest,
	opts ...Option,
) (types.FetchOrganizationByVersionResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchOrganizationByVersion(
			ctx,
			nemdb.FetchOrganizationByVersionParams{
				Version: req.Version,
				Offset:  req.Offset,
				Limit:   req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_organization_by_version",
				Message:          "error in FetchOrganizationByVersion no order",
				EntityIdentifier: "organization",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchOrganizationByVersionResponse{}, err
		}
		return types.FetchOrganizationByVersionResponse{
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
			ActionIdentifier: "fetch_organization_by_version",
			Message:          "error in FetchOrganizationByVersion - order by field not found",
			EntityIdentifier: "organization",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchOrganizationByVersionResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_organization_by_version",
				Message:          "error in FetchOrganizationByVersion - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "organization",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchOrganizationByVersionResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchOrganizationByVersionOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchOrganizationByVersionOrderedByCreatedAtASCParams{
					Version: req.Version,
					Offset:  req.Offset,
					Limit:   req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_organization_by_version",
					Message:          "error in FetchOrganizationByVersion - ordered asc",
					EntityIdentifier: "organization",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchOrganizationByVersionResponse{}, err
			}
			return types.FetchOrganizationByVersionResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchOrganizationByVersionOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchOrganizationByVersionOrderedByCreatedAtDESCParams{
					Version: req.Version,
					Offset:  req.Offset,
					Limit:   req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_organization_by_version",
					Message:          "error in FetchOrganizationByVersion - ordered desc",
					EntityIdentifier: "organization",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchOrganizationByVersionResponse{}, err
			}
			return types.FetchOrganizationByVersionResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchOrganizationByVersionOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchOrganizationByVersionOrderedByUpdatedAtASCParams{
					Version: req.Version,
					Offset:  req.Offset,
					Limit:   req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_organization_by_version",
					Message:          "error in FetchOrganizationByVersion - ordered asc",
					EntityIdentifier: "organization",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchOrganizationByVersionResponse{}, err
			}
			return types.FetchOrganizationByVersionResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchOrganizationByVersionOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchOrganizationByVersionOrderedByUpdatedAtDESCParams{
					Version: req.Version,
					Offset:  req.Offset,
					Limit:   req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_organization_by_version",
					Message:          "error in FetchOrganizationByVersion - ordered desc",
					EntityIdentifier: "organization",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchOrganizationByVersionResponse{}, err
			}
			return types.FetchOrganizationByVersionResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_organization_by_version_invalid",
		Message:          "error in FetchOrganizationByVersion - invalid request",
		EntityIdentifier: "organization",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchOrganizationByVersionResponse{}, err

}
