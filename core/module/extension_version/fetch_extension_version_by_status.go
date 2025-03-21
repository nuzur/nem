package extension_version

import (
	"context"

	"github.com/nuzur/nem/core/module/extension_version/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchExtensionVersionByStatus(
	ctx context.Context,
	req types.FetchExtensionVersionByStatusRequest,
	opts ...Option,
) (types.FetchExtensionVersionByStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionVersionByStatus(
			ctx,
			nemdb.FetchExtensionVersionByStatusParams{

				Status: req.Status.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_version_by_status",
				Message:          "error in FetchExtensionVersionByStatus no order",
				EntityIdentifier: "extension_version",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionVersionByStatusResponse{}, err
		}
		return types.FetchExtensionVersionByStatusResponse{
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
			ActionIdentifier: "fetch_extension_version_by_status",
			Message:          "error in FetchExtensionVersionByStatus - order by field not found",
			EntityIdentifier: "extension_version",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionVersionByStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_version_by_status",
				Message:          "error in FetchExtensionVersionByStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension_version",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionVersionByStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionVersionByStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionVersionByStatusOrderedByCreatedAtASCParams{

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_version_by_status",
					Message:          "error in FetchExtensionVersionByStatus - ordered asc",
					EntityIdentifier: "extension_version",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionVersionByStatusResponse{}, err
			}
			return types.FetchExtensionVersionByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionVersionByStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionVersionByStatusOrderedByCreatedAtDESCParams{

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_version_by_status",
					Message:          "error in FetchExtensionVersionByStatus - ordered desc",
					EntityIdentifier: "extension_version",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionVersionByStatusResponse{}, err
			}
			return types.FetchExtensionVersionByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionVersionByStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionVersionByStatusOrderedByUpdatedAtASCParams{

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_version_by_status",
					Message:          "error in FetchExtensionVersionByStatus - ordered asc",
					EntityIdentifier: "extension_version",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionVersionByStatusResponse{}, err
			}
			return types.FetchExtensionVersionByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionVersionByStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionVersionByStatusOrderedByUpdatedAtDESCParams{

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_version_by_status",
					Message:          "error in FetchExtensionVersionByStatus - ordered desc",
					EntityIdentifier: "extension_version",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionVersionByStatusResponse{}, err
			}
			return types.FetchExtensionVersionByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_version_by_status_invalid",
		Message:          "error in FetchExtensionVersionByStatus - invalid request",
		EntityIdentifier: "extension_version",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionVersionByStatusResponse{}, err

}
