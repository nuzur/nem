package extension_version

import (
	"context"

	"github.com/nuzur/nem/core/module/extension_version/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchExtensionVersionByVersionAndStatus(
	ctx context.Context,
	req types.FetchExtensionVersionByVersionAndStatusRequest,
	opts ...Option,
) (types.FetchExtensionVersionByVersionAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionVersionByVersionAndStatus(
			ctx,
			nemdb.FetchExtensionVersionByVersionAndStatusParams{

				Version: req.Version,

				Status: req.Status.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_version_by_version_and_status",
				Message:          "error in FetchExtensionVersionByVersionAndStatus no order",
				EntityIdentifier: "extension_version",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionVersionByVersionAndStatusResponse{}, err
		}
		return types.FetchExtensionVersionByVersionAndStatusResponse{
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
			ActionIdentifier: "fetch_extension_version_by_version_and_status",
			Message:          "error in FetchExtensionVersionByVersionAndStatus - order by field not found",
			EntityIdentifier: "extension_version",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionVersionByVersionAndStatusResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_version_by_version_and_status",
				Message:          "error in FetchExtensionVersionByVersionAndStatus - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension_version",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionVersionByVersionAndStatusResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionVersionByVersionAndStatusOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionVersionByVersionAndStatusOrderedByCreatedAtASCParams{

					Version: req.Version,

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_version_by_version_and_status",
					Message:          "error in FetchExtensionVersionByVersionAndStatus - ordered asc",
					EntityIdentifier: "extension_version",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionVersionByVersionAndStatusResponse{}, err
			}
			return types.FetchExtensionVersionByVersionAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionVersionByVersionAndStatusOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionVersionByVersionAndStatusOrderedByCreatedAtDESCParams{

					Version: req.Version,

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_version_by_version_and_status",
					Message:          "error in FetchExtensionVersionByVersionAndStatus - ordered desc",
					EntityIdentifier: "extension_version",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionVersionByVersionAndStatusResponse{}, err
			}
			return types.FetchExtensionVersionByVersionAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionVersionByVersionAndStatusOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionVersionByVersionAndStatusOrderedByUpdatedAtASCParams{

					Version: req.Version,

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_version_by_version_and_status",
					Message:          "error in FetchExtensionVersionByVersionAndStatus - ordered asc",
					EntityIdentifier: "extension_version",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionVersionByVersionAndStatusResponse{}, err
			}
			return types.FetchExtensionVersionByVersionAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionVersionByVersionAndStatusOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionVersionByVersionAndStatusOrderedByUpdatedAtDESCParams{

					Version: req.Version,

					Status: req.Status.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_version_by_version_and_status",
					Message:          "error in FetchExtensionVersionByVersionAndStatus - ordered desc",
					EntityIdentifier: "extension_version",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionVersionByVersionAndStatusResponse{}, err
			}
			return types.FetchExtensionVersionByVersionAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_version_by_version_and_status_invalid",
		Message:          "error in FetchExtensionVersionByVersionAndStatus - invalid request",
		EntityIdentifier: "extension_version",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionVersionByVersionAndStatusResponse{}, err

}
