package change_request

import (
	"context"

	"github.com/nuzur/nem/core/module/change_request/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchChangeRequestByVersionAndStatusAndAiGenerated(
	ctx context.Context,
	req types.FetchChangeRequestByVersionAndStatusAndAiGeneratedRequest,
	opts ...Option,
) (types.FetchChangeRequestByVersionAndStatusAndAiGeneratedResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchChangeRequestByVersionAndStatusAndAiGenerated(
			ctx,
			nemdb.FetchChangeRequestByVersionAndStatusAndAiGeneratedParams{

				Version: req.Version,

				Status: req.Status.ToInt64(),

				AiGenerated: req.AiGenerated,

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_change_request_by_version_and_status_and_ai_generated",
				Message:          "error in FetchChangeRequestByVersionAndStatusAndAiGenerated no order",
				EntityIdentifier: "change_request",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchChangeRequestByVersionAndStatusAndAiGeneratedResponse{}, err
		}
		return types.FetchChangeRequestByVersionAndStatusAndAiGeneratedResponse{
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
			ActionIdentifier: "fetch_change_request_by_version_and_status_and_ai_generated",
			Message:          "error in FetchChangeRequestByVersionAndStatusAndAiGenerated - order by field not found",
			EntityIdentifier: "change_request",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchChangeRequestByVersionAndStatusAndAiGeneratedResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_change_request_by_version_and_status_and_ai_generated",
				Message:          "error in FetchChangeRequestByVersionAndStatusAndAiGenerated - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "change_request",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchChangeRequestByVersionAndStatusAndAiGeneratedResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchChangeRequestByVersionAndStatusAndAiGeneratedOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchChangeRequestByVersionAndStatusAndAiGeneratedOrderedByCreatedAtASCParams{

					Version: req.Version,

					Status: req.Status.ToInt64(),

					AiGenerated: req.AiGenerated,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_version_and_status_and_ai_generated",
					Message:          "error in FetchChangeRequestByVersionAndStatusAndAiGenerated - ordered asc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByVersionAndStatusAndAiGeneratedResponse{}, err
			}
			return types.FetchChangeRequestByVersionAndStatusAndAiGeneratedResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchChangeRequestByVersionAndStatusAndAiGeneratedOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchChangeRequestByVersionAndStatusAndAiGeneratedOrderedByCreatedAtDESCParams{

					Version: req.Version,

					Status: req.Status.ToInt64(),

					AiGenerated: req.AiGenerated,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_version_and_status_and_ai_generated",
					Message:          "error in FetchChangeRequestByVersionAndStatusAndAiGenerated - ordered desc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByVersionAndStatusAndAiGeneratedResponse{}, err
			}
			return types.FetchChangeRequestByVersionAndStatusAndAiGeneratedResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchChangeRequestByVersionAndStatusAndAiGeneratedOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchChangeRequestByVersionAndStatusAndAiGeneratedOrderedByUpdatedAtASCParams{

					Version: req.Version,

					Status: req.Status.ToInt64(),

					AiGenerated: req.AiGenerated,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_version_and_status_and_ai_generated",
					Message:          "error in FetchChangeRequestByVersionAndStatusAndAiGenerated - ordered asc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByVersionAndStatusAndAiGeneratedResponse{}, err
			}
			return types.FetchChangeRequestByVersionAndStatusAndAiGeneratedResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchChangeRequestByVersionAndStatusAndAiGeneratedOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchChangeRequestByVersionAndStatusAndAiGeneratedOrderedByUpdatedAtDESCParams{

					Version: req.Version,

					Status: req.Status.ToInt64(),

					AiGenerated: req.AiGenerated,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_version_and_status_and_ai_generated",
					Message:          "error in FetchChangeRequestByVersionAndStatusAndAiGenerated - ordered desc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByVersionAndStatusAndAiGeneratedResponse{}, err
			}
			return types.FetchChangeRequestByVersionAndStatusAndAiGeneratedResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_change_request_by_version_and_status_and_ai_generated_invalid",
		Message:          "error in FetchChangeRequestByVersionAndStatusAndAiGenerated - invalid request",
		EntityIdentifier: "change_request",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchChangeRequestByVersionAndStatusAndAiGeneratedResponse{}, err

}
