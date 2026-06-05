package local_agent

import (
	"context"
	"errors"

	"github.com/nuzur/nem/core/module/local_agent/types"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"

	"github.com/nuzur/nem/core/entity/local_agent_connection"

	"github.com/nuzur/nem/custom"

	"github.com/nuzur/nem/core/entity/mapper"
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
	existing, err := qtx.FetchLocalAgentByUUIDForUpdate(ctx, req.LocalAgent.UUID.String())
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_local_agent",
			Message:          "error fetching existing record for UpsertLocalAgent - with uuid",
			EntityIdentifier: "local_agent",
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
			ActionIdentifier: "upsert_local_agent",
			Message:          "not found existing record for UpsertLocalAgent - with uuid",
			EntityIdentifier: "local_agent",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	params := mapUpsertRequestToUpdateParams(req, existing[0], partial)
	err = qtx.UpdateLocalAgent(
		ctx,
		params,
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_local_agent",
			Message:          "error calling repository for UpsertLocalAgent - with uuid",
			EntityIdentifier: "local_agent",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	err = m.publishUpdateEvent(ctx, req, qtx, req.LocalAgent.UUID.String(), existing)
	if err != nil {
		return types.UpsertResponse{}, err
	}

	if createdTx {
		err := tx.Commit()
		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "upsert_local_agent",
				Message:          "error commiting for UpsertLocalAgent",
				EntityIdentifier: "local_agent",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.UpsertResponse{}, err
		}
	}

	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "upsert_local_agent",
		Message:          "successfully handled UpsertLocalAgent - with uuid",
		EntityIdentifier: "local_agent",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return types.UpsertResponse{
		UUID: req.LocalAgent.UUID,
	}, nil
}

func mapUpsertRequestToUpdateParams(req types.UpsertRequest, existing nemdb.LocalAgent, partial bool) nemdb.UpdateLocalAgentParams {
	if !partial {
		return nemdb.UpdateLocalAgentParams{

			UUID: req.LocalAgent.UUID.String(),

			UserUUID: req.LocalAgent.UserUUID.String(),

			TokenHash: mapper.StringPtrToSqlNullString(req.LocalAgent.TokenHash),

			MachineName: mapper.StringPtrToSqlNullString(req.LocalAgent.MachineName),

			Os: mapper.StringPtrToSqlNullString(req.LocalAgent.Os),

			CliVersion: mapper.StringPtrToSqlNullString(req.LocalAgent.CliVersion),

			Connections: local_agent_connection.LocalAgentConnectionSliceToJSON(req.LocalAgent.Connections),

			Status: req.LocalAgent.Status.ToInt64(),

			LastSeenAt: mapper.TimePtrToSqlNullTime(req.LocalAgent.LastSeenAt),

			RevokedAt: mapper.TimePtrToSqlNullTime(req.LocalAgent.RevokedAt),

			CreatedAt: existing.CreatedAt,

			UpdatedAt: custom.Now(),

			CreatedByUUID: req.LocalAgent.CreatedByUUID.String(),

			UpdatedByUUID: req.LocalAgent.UpdatedByUUID.String(),
		}
	}

	res := nemdb.UpdateLocalAgentParams{}

	res.UUID = req.LocalAgent.UUID.String()

	res.UserUUID = req.LocalAgent.UserUUID.String()

	res.TokenHash = mapper.StringPtrToSqlNullString(req.LocalAgent.TokenHash)

	res.MachineName = mapper.StringPtrToSqlNullString(req.LocalAgent.MachineName)

	res.Os = mapper.StringPtrToSqlNullString(req.LocalAgent.Os)

	res.CliVersion = mapper.StringPtrToSqlNullString(req.LocalAgent.CliVersion)

	res.Connections = local_agent_connection.LocalAgentConnectionSliceToJSON(req.LocalAgent.Connections)

	res.Status = req.LocalAgent.Status.ToInt64()

	res.LastSeenAt = mapper.TimePtrToSqlNullTime(req.LocalAgent.LastSeenAt)

	res.RevokedAt = mapper.TimePtrToSqlNullTime(req.LocalAgent.RevokedAt)

	res.CreatedAt = existing.CreatedAt

	res.UpdatedAt = custom.Now()

	res.CreatedByUUID = req.LocalAgent.CreatedByUUID.String()

	res.UpdatedByUUID = req.LocalAgent.UpdatedByUUID.String()

	return res

}

func (m *module) publishUpdateEvent(ctx context.Context,
	req types.UpsertRequest,
	qtx *nemdb.Queries,
	id string,
	existing []nemdb.LocalAgent) error {

	if m.events == nil {
		return nil
	}
	fetched, err := qtx.FetchLocalAgentByUUID(ctx, id)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_local_agent",
			Message:          "error fetching after UpsertLocalAgent - with uuid",
			EntityIdentifier: "local_agent",
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
			ActionIdentifier: "upsert_local_agent",
			Message:          "error producing insert event for UpsertLocalAgent - with uuid",
			EntityIdentifier: "local_agent",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return err
	}
	return nil
}
