package extension

import (
	"context"

	"github.com/nuzur/nem/core/module/extension/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchExtensionByVersionAndRepositoryAndExtensionType(
	ctx context.Context,
	req types.FetchExtensionByVersionAndRepositoryAndExtensionTypeRequest,
	opts ...Option,
) (types.FetchExtensionByVersionAndRepositoryAndExtensionTypeResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByVersionAndRepositoryAndExtensionType(
			ctx,
			nemdb.FetchExtensionByVersionAndRepositoryAndExtensionTypeParams{

				Version: req.Version,

				Repository: req.Repository,

				ExtensionType: req.ExtensionType.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type",
				Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionType no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeResponse{}, err
		}
		return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeResponse{
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
			ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type",
			Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionType - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type",
				Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionType - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByVersionAndRepositoryAndExtensionTypeOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByVersionAndRepositoryAndExtensionTypeOrderedByCreatedAtASCParams{

					Version: req.Version,

					Repository: req.Repository,

					ExtensionType: req.ExtensionType.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type",
					Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionType - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeResponse{}, err
			}
			return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByVersionAndRepositoryAndExtensionTypeOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByVersionAndRepositoryAndExtensionTypeOrderedByCreatedAtDESCParams{

					Version: req.Version,

					Repository: req.Repository,

					ExtensionType: req.ExtensionType.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type",
					Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionType - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeResponse{}, err
			}
			return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByVersionAndRepositoryAndExtensionTypeOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByVersionAndRepositoryAndExtensionTypeOrderedByUpdatedAtASCParams{

					Version: req.Version,

					Repository: req.Repository,

					ExtensionType: req.ExtensionType.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type",
					Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionType - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeResponse{}, err
			}
			return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByVersionAndRepositoryAndExtensionTypeOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByVersionAndRepositoryAndExtensionTypeOrderedByUpdatedAtDESCParams{

					Version: req.Version,

					Repository: req.Repository,

					ExtensionType: req.ExtensionType.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type",
					Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionType - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeResponse{}, err
			}
			return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type_invalid",
		Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionType - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeResponse{}, err

}
