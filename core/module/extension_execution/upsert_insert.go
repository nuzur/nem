package extension_execution

import (
	"context"

	"errors"

	"github.com/gofrs/uuid"
	"github.com/nuzur/nem/core/module/extension_execution/types"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"

	"github.com/nuzur/nem/custom"

	"github.com/nuzur/nem/core/entity/mapper"
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

	_, err := qtx.InsertExtensionExecution(
		ctx,
		params,
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_extension_execution",
			Message:          "error calling repository for UpsertExtensionExecution - no uuid",
			EntityIdentifier: "extension_execution",
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
				ActionIdentifier: "upsert_extension_execution",
				Message:          "error commiting for UpsertExtensionExecution (1)",
				EntityIdentifier: "extension_execution",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.UpsertResponse{}, err
		}
	}

	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "upsert_extension_execution",
		Message:          "successfully handled UpsertExtensionExecution - no uuid",
		EntityIdentifier: "extension_execution",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return types.UpsertResponse{
		UUID: uuid.FromStringOrNil(params.UUID),
	}, nil
}

func mapUpsertRequestToInsertParams(req types.UpsertRequest) nemdb.InsertExtensionExecutionParams {
	return nemdb.InsertExtensionExecutionParams{

		UUID: custom.GenerateUUID().String(),

		ExtensionUUID: req.ExtensionExecution.ExtensionUUID.String(),

		ExtensionVersionUUID: req.ExtensionExecution.ExtensionVersionUUID.String(),

		ProjectExtensionUUID: mapper.StringToSqlNullString(req.ExtensionExecution.ProjectExtensionUUID.String()),

		ProjectUUID: req.ExtensionExecution.ProjectUUID.String(),

		ProjectVersionUUID: req.ExtensionExecution.ProjectVersionUUID.String(),

		ExecutedByUUID: req.ExtensionExecution.ExecutedByUUID.String(),

		Metadata: []byte(mapper.StringPtrToString(req.ExtensionExecution.Metadata)),

		Status: req.ExtensionExecution.Status.ToInt64(),

		StatusMsg: mapper.StringPtrToSqlNullString(req.ExtensionExecution.StatusMsg),

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

	fetched, err := qtx.FetchExtensionExecutionByUUID(ctx, id)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_extension_execution",
			Message:          "error fetching after UpsertExtensionExecution - no uuid",
			EntityIdentifier: "extension_execution",
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
			ActionIdentifier: "upsert_extension_execution",
			Message:          "error producing insert event for UpsertExtensionExecution - no uuid",
			EntityIdentifier: "extension_execution",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return err
	}
	return nil
}
