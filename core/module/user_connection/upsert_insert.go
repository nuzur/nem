package user_connection

import (
	"context"

	"errors"

	"github.com/gofrs/uuid"
	"github.com/nuzur/nem/core/module/user_connection/types"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"

	"github.com/nuzur/nem/custom"

	"github.com/nuzur/nem/core/entity/user_connection_execution"
	"github.com/nuzur/nem/core/entity/user_connection_type_config"
)

func (m *module) Insert(
	ctx context.Context,
	req types.UpsertRequest,
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
	params := mapUpsertRequestToInsertParams(req)

	_, err := qtx.InsertUserConnection(
		ctx,
		params,
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_user_connection",
			Message:          "error calling repository for UpsertUserConnection - no uuid",
			EntityIdentifier: "user_connection",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	err = m.publishInsertEvent(ctx, req, qtx, params.UUID)
	if err != nil {
		return types.UpsertResponse{}, err
	}

	if createdTx {
		err := tx.Commit()
		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "upsert_user_connection",
				Message:          "error commiting for UpsertUserConnection (1)",
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
		Message:          "successfully handled UpsertUserConnection - no uuid",
		EntityIdentifier: "user_connection",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return types.UpsertResponse{
		UUID: uuid.FromStringOrNil(params.UUID),
	}, nil
}

func mapUpsertRequestToInsertParams(req types.UpsertRequest) nemdb.InsertUserConnectionParams {
	return nemdb.InsertUserConnectionParams{

		UUID: custom.GenerateUUID().String(),

		UserUUID: req.UserConnection.UserUUID.String(),

		ProjectUUID: req.UserConnection.ProjectUUID.String(),

		ProjectVersionUUID: req.UserConnection.ProjectVersionUUID.String(),

		Type: req.UserConnection.Type.ToInt64(),

		TypeConfig: user_connection_type_config.UserConnectionTypeConfigToJSON(req.UserConnection.TypeConfig),

		DbSchema: req.UserConnection.DbSchema,

		Executions: user_connection_execution.UserConnectionExecutionSliceToJSON(req.UserConnection.Executions),

		Status: req.UserConnection.Status.ToInt64(),

		CreatedAt: custom.Now(),

		UpdatedAt: custom.Now(),
	}
}

func (m *module) publishInsertEvent(ctx context.Context,
	req types.UpsertRequest,
	qtx *nemdb.Queries,
	id string) error {

	if m.events == nil {
		return nil
	}

	fetched, err := qtx.FetchUserConnectionByUUID(ctx, id)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_user_connection",
			Message:          "error fetching after UpsertUserConnection - no uuid",
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
	err = m.events.ProduceEntityInserted(fetchedEntities[0])
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_user_connection",
			Message:          "error producing insert event for UpsertUserConnection - no uuid",
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
