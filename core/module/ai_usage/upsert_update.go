package ai_usage

import (
	"context"
	"errors"

	"github.com/nuzur/nem/core/module/ai_usage/types"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"

	"github.com/nuzur/nem/custom"
)

func (m *module) Update(
	ctx context.Context,
	req types.UpsertRequest,
	partial bool,
	opts ...Option,
) (types.UpsertResponse, error) {
	optConfig := applyAllOptions(opts)

	tx := optConfig.SQLTx
	createdTx := false
	if tx == nil {
		ntx, err := m.repository.DB.Begin()
		if err != nil {
			return types.UpsertResponse{}, err
		}
		tx = ntx
		defer tx.Rollback()
		createdTx = true
	}

	qtx := m.repository.Queries.WithTx(tx)
	existing, err := qtx.FetchAiUsageByUUIDForUpdate(ctx, req.AiUsage.UUID.String())
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_ai_usage",
			Message:          "error fetching existing record for UpsertAiUsage - with uuid",
			EntityIdentifier: "ai_usage",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	if len(existing) == 0 {
		err := errors.New("entity not found")
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_ai_usage",
			Message:          "not found existing record for UpsertAiUsage - with uuid",
			EntityIdentifier: "ai_usage",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	params := mapUpsertRequestToUpdateParams(req, existing[0], partial)
	err = qtx.UpdateAiUsage(
		ctx,
		params,
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_ai_usage",
			Message:          "error calling repository for UpsertAiUsage - with uuid",
			EntityIdentifier: "ai_usage",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	err = m.publishUpdateEvent(ctx, req, qtx, req.AiUsage.UUID.String(), existing)
	if err != nil {
		return types.UpsertResponse{}, err
	}

	if createdTx {
		err := tx.Commit()
		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "upsert_ai_usage",
				Message:          "error commiting for UpsertAiUsage",
				EntityIdentifier: "ai_usage",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.UpsertResponse{}, err
		}
	}

	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "upsert_ai_usage",
		Message:          "successfully handled UpsertAiUsage - with uuid",
		EntityIdentifier: "ai_usage",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return types.UpsertResponse{
		UUID: req.AiUsage.UUID,
	}, nil
}

func mapUpsertRequestToUpdateParams(req types.UpsertRequest, existing nemdb.AiUsage, partial bool) nemdb.UpdateAiUsageParams {
	if !partial {
		return nemdb.UpdateAiUsageParams{

			UUID: req.AiUsage.UUID.String(),

			UserUUID: req.AiUsage.UserUUID.String(),

			ProjectUUID: req.AiUsage.ProjectUUID.String(),

			ProjectVersionUUID: req.AiUsage.ProjectVersionUUID.String(),

			UserPrompt: req.AiUsage.UserPrompt,

			Step: req.AiUsage.Step,

			Context: req.AiUsage.Context.ToInt64(),

			Provider: req.AiUsage.Provider.ToInt64(),

			Tokens: req.AiUsage.Tokens,

			Status: req.AiUsage.Status.ToInt64(),

			CreatedAt: existing.CreatedAt,

			UpdatedAt: custom.Now(),

			CreatedByUUID: req.AiUsage.CreatedByUUID.String(),

			UpdatedByUUID: req.AiUsage.UpdatedByUUID.String(),
		}
	}

	res := nemdb.UpdateAiUsageParams{}

	res.UUID = req.AiUsage.UUID.String()

	res.UserUUID = req.AiUsage.UserUUID.String()

	res.ProjectUUID = req.AiUsage.ProjectUUID.String()

	res.ProjectVersionUUID = req.AiUsage.ProjectVersionUUID.String()

	res.UserPrompt = req.AiUsage.UserPrompt

	res.Step = req.AiUsage.Step

	res.Context = req.AiUsage.Context.ToInt64()

	res.Provider = req.AiUsage.Provider.ToInt64()

	res.Tokens = req.AiUsage.Tokens

	res.Status = req.AiUsage.Status.ToInt64()

	res.CreatedAt = existing.CreatedAt

	res.UpdatedAt = custom.Now()

	res.CreatedByUUID = req.AiUsage.CreatedByUUID.String()

	res.UpdatedByUUID = req.AiUsage.UpdatedByUUID.String()

	return res

}

func (m *module) publishUpdateEvent(ctx context.Context,
	req types.UpsertRequest,
	qtx *nemdb.Queries,
	id string,
	existing []nemdb.AiUsage) error {

	if m.events == nil {
		return nil
	}
	fetched, err := qtx.FetchAiUsageByUUID(ctx, id)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_ai_usage",
			Message:          "error fetching after UpsertAiUsage - with uuid",
			EntityIdentifier: "ai_usage",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return err
	}

	fetchedEntities := mapModelsToEntities(fetched)
	if len(fetchedEntities) != 1 {
		return errors.New("error mapping to entity")
	}
	existingEntities := mapModelsToEntities(existing)
	err = m.events.ProduceEntityUpdated(fetchedEntities[0], existingEntities[0])
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_ai_usage",
			Message:          "error producing insert event for UpsertAiUsage - with uuid",
			EntityIdentifier: "ai_usage",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return err
	}
	return nil
}
