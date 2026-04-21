package change_request

import (
	"context"

	"github.com/nuzur/nem/core/module/change_request/types"

	"errors"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchChangeRequestByVersionAndChangeTypeAndStatusAndAiGenerated(
	ctx context.Context,
	req types.FetchChangeRequestByVersionAndChangeTypeAndStatusAndAiGeneratedRequest,
	opts ...Option,
) (types.FetchChangeRequestByVersionAndChangeTypeAndStatusAndAiGeneratedResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.Queries.FetchChangeRequestByVersionAndChangeTypeAndStatusAndAiGenerated(
			ctx,
			nemdb.FetchChangeRequestByVersionAndChangeTypeAndStatusAndAiGeneratedParams{

				Version: req.Version,

				ChangeType: req.ChangeType.ToInt64(),

				Status: req.Status.ToInt64(),

				AiGenerated: req.AiGenerated,

				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)

		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_change_request_by_version_and_change_type_and_status_and_ai_generated",
				Message:          "error in FetchChangeRequestByVersionAndChangeTypeAndStatusAndAiGenerated no order",
				EntityIdentifier: "change_request",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchChangeRequestByVersionAndChangeTypeAndStatusAndAiGeneratedResponse{}, err
		}
		return types.FetchChangeRequestByVersionAndChangeTypeAndStatusAndAiGeneratedResponse{
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
			ActionIdentifier: "fetch_change_request_by_version_and_change_type_and_status_and_ai_generated",
			Message:          "error in FetchChangeRequestByVersionAndChangeTypeAndStatusAndAiGenerated - order by field not found",
			EntityIdentifier: "change_request",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchChangeRequestByVersionAndChangeTypeAndStatusAndAiGeneratedResponse{}, err
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			err := errors.New("invalid sort value only ASC or DESC are valid")
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "fetch_change_request_by_version_and_change_type_and_status_and_ai_generated",
				Message:          "error in FetchChangeRequestByVersionAndChangeTypeAndStatusAndAiGenerated - invalid sort value only ASC or DESC are valid",
				EntityIdentifier: "change_request",
				Layer:            monitoring.ModuleServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.FetchChangeRequestByVersionAndChangeTypeAndStatusAndAiGeneratedResponse{}, err
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchChangeRequestByVersionAndChangeTypeAndStatusAndAiGeneratedOrderedByCreatedAtASC(
				ctx,
				nemdb.FetchChangeRequestByVersionAndChangeTypeAndStatusAndAiGeneratedOrderedByCreatedAtASCParams{

					Version: req.Version,

					ChangeType: req.ChangeType.ToInt64(),

					Status: req.Status.ToInt64(),

					AiGenerated: req.AiGenerated,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_version_and_change_type_and_status_and_ai_generated",
					Message:          "error in FetchChangeRequestByVersionAndChangeTypeAndStatusAndAiGenerated - ordered asc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByVersionAndChangeTypeAndStatusAndAiGeneratedResponse{}, err
			}
			return types.FetchChangeRequestByVersionAndChangeTypeAndStatusAndAiGeneratedResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchChangeRequestByVersionAndChangeTypeAndStatusAndAiGeneratedOrderedByCreatedAtDESC(
				ctx,
				nemdb.FetchChangeRequestByVersionAndChangeTypeAndStatusAndAiGeneratedOrderedByCreatedAtDESCParams{

					Version: req.Version,

					ChangeType: req.ChangeType.ToInt64(),

					Status: req.Status.ToInt64(),

					AiGenerated: req.AiGenerated,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_version_and_change_type_and_status_and_ai_generated",
					Message:          "error in FetchChangeRequestByVersionAndChangeTypeAndStatusAndAiGenerated - ordered desc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByVersionAndChangeTypeAndStatusAndAiGeneratedResponse{}, err
			}
			return types.FetchChangeRequestByVersionAndChangeTypeAndStatusAndAiGeneratedResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.Queries.FetchChangeRequestByVersionAndChangeTypeAndStatusAndAiGeneratedOrderedByUpdatedAtASC(
				ctx,
				nemdb.FetchChangeRequestByVersionAndChangeTypeAndStatusAndAiGeneratedOrderedByUpdatedAtASCParams{

					Version: req.Version,

					ChangeType: req.ChangeType.ToInt64(),

					Status: req.Status.ToInt64(),

					AiGenerated: req.AiGenerated,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_version_and_change_type_and_status_and_ai_generated",
					Message:          "error in FetchChangeRequestByVersionAndChangeTypeAndStatusAndAiGenerated - ordered asc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByVersionAndChangeTypeAndStatusAndAiGeneratedResponse{}, err
			}
			return types.FetchChangeRequestByVersionAndChangeTypeAndStatusAndAiGeneratedResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.Queries.FetchChangeRequestByVersionAndChangeTypeAndStatusAndAiGeneratedOrderedByUpdatedAtDESC(
				ctx,
				nemdb.FetchChangeRequestByVersionAndChangeTypeAndStatusAndAiGeneratedOrderedByUpdatedAtDESCParams{

					Version: req.Version,

					ChangeType: req.ChangeType.ToInt64(),

					Status: req.Status.ToInt64(),

					AiGenerated: req.AiGenerated,

					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				m.monitoring.Emit(monitoring.EmitRequest{
					ActionIdentifier: "fetch_change_request_by_version_and_change_type_and_status_and_ai_generated",
					Message:          "error in FetchChangeRequestByVersionAndChangeTypeAndStatusAndAiGenerated - ordered desc",
					EntityIdentifier: "change_request",
					Layer:            monitoring.RepositoryServiceLayer,
					Type:             monitoring.EmitTypeError,
					Data:             req,
					Error:            err,
				})
				return types.FetchChangeRequestByVersionAndChangeTypeAndStatusAndAiGeneratedResponse{}, err
			}
			return types.FetchChangeRequestByVersionAndChangeTypeAndStatusAndAiGeneratedResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	err := errors.New("could not process request")
	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "fetch_change_request_by_version_and_change_type_and_status_and_ai_generated_invalid",
		Message:          "error in FetchChangeRequestByVersionAndChangeTypeAndStatusAndAiGenerated - invalid request",
		EntityIdentifier: "change_request",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeError,
		Data:             req,
		Error:            err,
	})
	return types.FetchChangeRequestByVersionAndChangeTypeAndStatusAndAiGeneratedResponse{}, err

}
