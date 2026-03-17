package extension

import (
	"context"

	"github.com/nuzur/nem/core/module/extension/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPro(
	ctx context.Context,
	req types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProRequest,
	opts ...Option,
) (types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPro(
			ctx,
			nemdb.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProParams{

				Version: req.Version,

				Repository: req.Repository,

				ExtensionType: req.ExtensionType.ToInt64(),

				Pro: req.Pro,

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type_and_pro",
				Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPro no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProResponse{}, err
		}
		return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProResponse{
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
			ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type_and_pro",
			Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPro - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type_and_pro",
				Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPro - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProOrderedByCreatedAtASCParams{

					Version: req.Version,

					Repository: req.Repository,

					ExtensionType: req.ExtensionType.ToInt64(),

					Pro: req.Pro,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type_and_pro",
					Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPro - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProResponse{}, err
			}
			return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProOrderedByCreatedAtDESCParams{

					Version: req.Version,

					Repository: req.Repository,

					ExtensionType: req.ExtensionType.ToInt64(),

					Pro: req.Pro,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type_and_pro",
					Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPro - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProResponse{}, err
			}
			return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProOrderedByUpdatedAtASCParams{

					Version: req.Version,

					Repository: req.Repository,

					ExtensionType: req.ExtensionType.ToInt64(),

					Pro: req.Pro,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type_and_pro",
					Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPro - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProResponse{}, err
			}
			return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProOrderedByUpdatedAtDESCParams{

					Version: req.Version,

					Repository: req.Repository,

					ExtensionType: req.ExtensionType.ToInt64(),

					Pro: req.Pro,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type_and_pro",
					Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPro - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProResponse{}, err
			}
			return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_version_and_repository_and_extension_type_and_pro_invalid",
		Message:          "error in FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPro - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProResponse{}, err

}
