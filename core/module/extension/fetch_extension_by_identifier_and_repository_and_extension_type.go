package extension

import (
	"context"

	"github.com/nuzur/nem/core/module/extension/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchExtensionByIdentifierAndRepositoryAndExtensionType(
	ctx context.Context,
	req types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeRequest,
	opts ...Option,
) (types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByIdentifierAndRepositoryAndExtensionType(
			ctx,
			nemdb.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeParams{

				Identifier: req.Identifier,

				Repository: req.Repository,

				ExtensionType: req.ExtensionType.ToInt64(),

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_extension_type",
				Message:          "error in FetchExtensionByIdentifierAndRepositoryAndExtensionType no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeResponse{}, err
		}
		return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeResponse{
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
			ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_extension_type",
			Message:          "error in FetchExtensionByIdentifierAndRepositoryAndExtensionType - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_extension_type",
				Message:          "error in FetchExtensionByIdentifierAndRepositoryAndExtensionType - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeOrderedByCreatedAtASCParams{

					Identifier: req.Identifier,

					Repository: req.Repository,

					ExtensionType: req.ExtensionType.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_extension_type",
					Message:          "error in FetchExtensionByIdentifierAndRepositoryAndExtensionType - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeOrderedByCreatedAtDESCParams{

					Identifier: req.Identifier,

					Repository: req.Repository,

					ExtensionType: req.ExtensionType.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_extension_type",
					Message:          "error in FetchExtensionByIdentifierAndRepositoryAndExtensionType - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeOrderedByUpdatedAtASCParams{

					Identifier: req.Identifier,

					Repository: req.Repository,

					ExtensionType: req.ExtensionType.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_extension_type",
					Message:          "error in FetchExtensionByIdentifierAndRepositoryAndExtensionType - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeOrderedByUpdatedAtDESCParams{

					Identifier: req.Identifier,

					Repository: req.Repository,

					ExtensionType: req.ExtensionType.ToInt64(),

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_extension_type",
					Message:          "error in FetchExtensionByIdentifierAndRepositoryAndExtensionType - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeResponse{}, err
			}
			return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_identifier_and_repository_and_extension_type_invalid",
		Message:          "error in FetchExtensionByIdentifierAndRepositoryAndExtensionType - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeResponse{}, err

}
