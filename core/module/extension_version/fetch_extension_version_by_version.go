package extension_version

import (
	"context"

	"nem/core/module/extension_version/types"

	"errors"
	nemdb "nem/core/repository/gen"
	"nem/monitoring"
)

func (m *module) FetchExtensionVersionByVersion(
	ctx context.Context,
	req types.FetchExtensionVersionByVersionRequest,
	opts ...Option,
) (types.FetchExtensionVersionByVersionResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionVersionByVersion(
			ctx,
			nemdb.FetchExtensionVersionByVersionParams{
				Version: req.Version,
				Offset:  req.Offset,
				Limit:   req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_version_by_version",
				Message:          "error in FetchExtensionVersionByVersion no order",
				EntityIdentifier: "extension_version",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionVersionByVersionResponse{}, err
		}
		return types.FetchExtensionVersionByVersionResponse{
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
			ActionIdentifier: "fetch_extension_version_by_version",
			Message:          "error in FetchExtensionVersionByVersion - order by field not found",
			EntityIdentifier: "extension_version",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionVersionByVersionResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_version_by_version",
				Message:          "error in FetchExtensionVersionByVersion - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension_version",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionVersionByVersionResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionVersionByVersionOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionVersionByVersionOrderedByCreatedAtASCParams{
					Version: req.Version,
					Offset:  req.Offset,
					Limit:   req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_version_by_version",
					Message:          "error in FetchExtensionVersionByVersion - ordered asc",
					EntityIdentifier: "extension_version",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionVersionByVersionResponse{}, err
			}
			return types.FetchExtensionVersionByVersionResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionVersionByVersionOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionVersionByVersionOrderedByCreatedAtDESCParams{
					Version: req.Version,
					Offset:  req.Offset,
					Limit:   req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_version_by_version",
					Message:          "error in FetchExtensionVersionByVersion - ordered desc",
					EntityIdentifier: "extension_version",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionVersionByVersionResponse{}, err
			}
			return types.FetchExtensionVersionByVersionResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionVersionByVersionOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionVersionByVersionOrderedByUpdatedAtASCParams{
					Version: req.Version,
					Offset:  req.Offset,
					Limit:   req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_version_by_version",
					Message:          "error in FetchExtensionVersionByVersion - ordered asc",
					EntityIdentifier: "extension_version",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionVersionByVersionResponse{}, err
			}
			return types.FetchExtensionVersionByVersionResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionVersionByVersionOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionVersionByVersionOrderedByUpdatedAtDESCParams{
					Version: req.Version,
					Offset:  req.Offset,
					Limit:   req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_version_by_version",
					Message:          "error in FetchExtensionVersionByVersion - ordered desc",
					EntityIdentifier: "extension_version",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionVersionByVersionResponse{}, err
			}
			return types.FetchExtensionVersionByVersionResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_version_by_version_invalid",
		Message:          "error in FetchExtensionVersionByVersion - invalid request",
		EntityIdentifier: "extension_version",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionVersionByVersionResponse{}, err

}
