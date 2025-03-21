package extension

import (
	"context"

	"github.com/nuzur/nem/core/module/extension/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchExtensionByRepository(
	ctx context.Context,
	req types.FetchExtensionByRepositoryRequest,
	opts ...Option,
) (types.FetchExtensionByRepositoryResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchExtensionByRepository(
			ctx,
			nemdb.FetchExtensionByRepositoryParams{

				Repository: req.Repository,

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_repository",
				Message:          "error in FetchExtensionByRepository no order",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByRepositoryResponse{}, err
		}
		return types.FetchExtensionByRepositoryResponse{
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
			ActionIdentifier: "fetch_extension_by_repository",
			Message:          "error in FetchExtensionByRepository - order by field not found",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchExtensionByRepositoryResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_extension_by_repository",
				Message:          "error in FetchExtensionByRepository - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "extension",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchExtensionByRepositoryResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByRepositoryOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchExtensionByRepositoryOrderedByCreatedAtASCParams{

					Repository: req.Repository,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_repository",
					Message:          "error in FetchExtensionByRepository - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByRepositoryResponse{}, err
			}
			return types.FetchExtensionByRepositoryResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByRepositoryOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchExtensionByRepositoryOrderedByCreatedAtDESCParams{

					Repository: req.Repository,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_repository",
					Message:          "error in FetchExtensionByRepository - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByRepositoryResponse{}, err
			}
			return types.FetchExtensionByRepositoryResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchExtensionByRepositoryOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchExtensionByRepositoryOrderedByUpdatedAtASCParams{

					Repository: req.Repository,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_repository",
					Message:          "error in FetchExtensionByRepository - ordered asc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByRepositoryResponse{}, err
			}
			return types.FetchExtensionByRepositoryResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchExtensionByRepositoryOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchExtensionByRepositoryOrderedByUpdatedAtDESCParams{

					Repository: req.Repository,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_extension_by_repository",
					Message:          "error in FetchExtensionByRepository - ordered desc",
					EntityIdentifier: "extension",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchExtensionByRepositoryResponse{}, err
			}
			return types.FetchExtensionByRepositoryResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_extension_by_repository_invalid",
		Message:          "error in FetchExtensionByRepository - invalid request",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchExtensionByRepositoryResponse{}, err

}
