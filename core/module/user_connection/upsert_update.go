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
	emptyReq := types.UpsertRequest{}

	// regular field
	if req.UserConnection.UUID == emptyReq.UserConnection.UUID {

		res.UUID = existing.UUID
	} else {

		res.UUID = req.UserConnection.UUID.String()

	}

	// regular field
	if req.UserConnection.UserUUID == emptyReq.UserConnection.UserUUID {

		res.UserUUID = existing.UserUUID
	} else {

		res.UserUUID = req.UserConnection.UserUUID.String()

	}

	// regular field
	if req.UserConnection.ProjectUUID == emptyReq.UserConnection.ProjectUUID {

		res.ProjectUUID = existing.ProjectUUID
	} else {

		res.ProjectUUID = req.UserConnection.ProjectUUID.String()

	}

	// regular field
	if req.UserConnection.ProjectVersionUUID == emptyReq.UserConnection.ProjectVersionUUID {

		res.ProjectVersionUUID = existing.ProjectVersionUUID
	} else {

		res.ProjectVersionUUID = req.UserConnection.ProjectVersionUUID.String()

	}

	// regular field
	if req.UserConnection.Type == emptyReq.UserConnection.Type {

		res.Type = existing.Type
	} else {

		res.Type = req.UserConnection.Type.ToInt64()

	}

	if req.UserConnection.TypeConfig.ToJSON() == nil {

		res.TypeConfig = existing.TypeConfig
	} else {

		res.TypeConfig = user_connection_type_config.UserConnectionTypeConfigToJSON(req.UserConnection.TypeConfig)

	}

	// regular field
	if req.UserConnection.DbSchema == emptyReq.UserConnection.DbSchema {

		res.DbSchema = existing.DbSchema
	} else {

		res.DbSchema = req.UserConnection.DbSchema

	}

	// json array
	if len(req.UserConnection.Executions) == 0 {

		res.Executions = existing.Executions
	} else {

		res.Executions = user_connection_execution.UserConnectionExecutionSliceToJSON(req.UserConnection.Executions)

	}

	// regular field
	if req.UserConnection.Status == emptyReq.UserConnection.Status {

		res.Status = existing.Status
	} else {

		res.Status = req.UserConnection.Status.ToInt64()

	}

	// regular field
	if req.UserConnection.CreatedAt == emptyReq.UserConnection.CreatedAt {

		res.CreatedAt = existing.CreatedAt
	} else {

		res.CreatedAt = existing.CreatedAt

	}

	// regular field
	if req.UserConnection.UpdatedAt == emptyReq.UserConnection.UpdatedAt {

		res.UpdatedAt = existing.UpdatedAt
	} else {

		res.UpdatedAt = custom.Now()

	}

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
