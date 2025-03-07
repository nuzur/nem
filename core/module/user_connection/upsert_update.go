package user_connection

import (
	"context"
	"errors"

	"github.com/nuzur/nem/core/module/user_connection/types"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"

	"github.com/nuzur/nem/core/entity/user_connection_execution"
	"github.com/nuzur/nem/core/entity/user_connection_type_config"

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
	existing, err := qtx.FetchUserConnectionByUUIDForUpdate(ctx, req.UserConnection.UUID.String())
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_user_connection",
			Message:          "error fetching existing record for UpsertUserConnection - with uuid",
			EntityIdentifier: "user_connection",
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
			ActionIdentifier: "upsert_user_connection",
			Message:          "not found existing record for UpsertUserConnection - with uuid",
			EntityIdentifier: "user_connection",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	params := mapUpsertRequestToUpdateParams(req, existing[0], partial)
	err = qtx.UpdateUserConnection(
		ctx,
		params,
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_user_connection",
			Message:          "error calling repository for UpsertUserConnection - with uuid",
			EntityIdentifier: "user_connection",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	err = m.publishUpdateEvent(ctx, req, qtx, req.UserConnection.UUID.String(), existing)
	if err != nil {
		return types.UpsertResponse{}, err
	}

	if createdTx {
		err := tx.Commit()
		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "upsert_user_connection",
				Message:          "error commiting for UpsertUserConnection",
				EntityIdentifier: "user_connection",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.UpsertResponse{}, err
		}
	}

	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "upsert_user_connection",
		Message:          "successfully handled UpsertUserConnection - with uuid",
		EntityIdentifier: "user_connection",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return types.UpsertResponse{
		UUID: req.UserConnection.UUID,
	}, nil
}

func mapUpsertRequestToUpdateParams(req types.UpsertRequest, existing nemdb.UserConnection, partial bool) nemdb.UpdateUserConnectionParams {
	if !partial {
		return nemdb.UpdateUserConnectionParams{

			UUID: req.UserConnection.UUID.String(),

			UserUUID: req.UserConnection.UserUUID.String(),

			ProjectUUID: req.UserConnection.ProjectUUID.String(),

			ProjectVersionUUID: req.UserConnection.ProjectVersionUUID.String(),

			Type: req.UserConnection.Type.ToInt64(),

			TypeConfig: user_connection_type_config.UserConnectionTypeConfigToJSON(req.UserConnection.TypeConfig),

			DbSchema: req.UserConnection.DbSchema,

			Executions: user_connection_execution.UserConnectionExecutionSliceToJSON(req.UserConnection.Executions),

			Status: req.UserConnection.Status.ToInt64(),

			CreatedAt: existing.CreatedAt,

			UpdatedAt: custom.Now(),
		}
	}

	res := nemdb.UpdateUserConnectionParams{}

	res.UUID = req.UserConnection.UUID.String()

	res.UserUUID = req.UserConnection.UserUUID.String()

	res.ProjectUUID = req.UserConnection.ProjectUUID.String()

	res.ProjectVersionUUID = req.UserConnection.ProjectVersionUUID.String()

	res.Type = req.UserConnection.Type.ToInt64()

	res.TypeConfig = user_connection_type_config.UserConnectionTypeConfigToJSON(req.UserConnection.TypeConfig)

	res.DbSchema = req.UserConnection.DbSchema

	res.Executions = user_connection_execution.UserConnectionExecutionSliceToJSON(req.UserConnection.Executions)

	res.Status = req.UserConnection.Status.ToInt64()

	res.CreatedAt = existing.CreatedAt

	res.UpdatedAt = custom.Now()

	return res

}

func (m *module) publishUpdateEvent(ctx context.Context,
	req types.UpsertRequest,
	qtx *nemdb.Queries,
	id string,
	existing []nemdb.UserConnection) error {

	if m.events == nil {
		return nil
	}
	fetched, err := qtx.FetchUserConnectionByUUID(ctx, id)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_user_connection",
			Message:          "error fetching after UpsertUserConnection - with uuid",
			EntityIdentifier: "user_connection",
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
			ActionIdentifier: "upsert_user_connection",
			Message:          "error producing insert event for UpsertUserConnection - with uuid",
			EntityIdentifier: "user_connection",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return err
	}
	return nil
}
