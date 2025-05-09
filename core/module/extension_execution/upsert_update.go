package extension_execution

import (
	"context"
	"errors"

	"github.com/nuzur/nem/core/module/extension_execution/types"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"

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
	existing, err := qtx.FetchExtensionExecutionByUUIDForUpdate(ctx, req.ExtensionExecution.UUID.String())
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_extension_execution",
			Message:          "error fetching existing record for UpsertExtensionExecution - with uuid",
			EntityIdentifier: "extension_execution",
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
			ActionIdentifier: "upsert_extension_execution",
			Message:          "not found existing record for UpsertExtensionExecution - with uuid",
			EntityIdentifier: "extension_execution",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	params := mapUpsertRequestToUpdateParams(req, existing[0], partial)
	err = qtx.UpdateExtensionExecution(
		ctx,
		params,
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_extension_execution",
			Message:          "error calling repository for UpsertExtensionExecution - with uuid",
			EntityIdentifier: "extension_execution",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	err = m.publishUpdateEvent(ctx, req, qtx, req.ExtensionExecution.UUID.String(), existing)
	if err != nil {
		return types.UpsertResponse{}, err
	}

	if createdTx {
		err := tx.Commit()
		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "upsert_extension_execution",
				Message:          "error commiting for UpsertExtensionExecution",
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
		Message:          "successfully handled UpsertExtensionExecution - with uuid",
		EntityIdentifier: "extension_execution",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return types.UpsertResponse{
		UUID: req.ExtensionExecution.UUID,
	}, nil
}

func mapUpsertRequestToUpdateParams(req types.UpsertRequest, existing nemdb.ExtensionExecution, partial bool) nemdb.UpdateExtensionExecutionParams {
	if !partial {
		return nemdb.UpdateExtensionExecutionParams{

			UUID: req.ExtensionExecution.UUID.String(),

			ExtensionUUID: req.ExtensionExecution.ExtensionUUID.String(),

			ExtensionVersionUUID: req.ExtensionExecution.ExtensionVersionUUID.String(),

			ProjectExtensionUUID: mapper.StringToSqlNullString(req.ExtensionExecution.ProjectExtensionUUID.String()),

			ProjectUUID: req.ExtensionExecution.ProjectUUID.String(),

			ProjectVersionUUID: req.ExtensionExecution.ProjectVersionUUID.String(),

			ExecutedByUUID: req.ExtensionExecution.ExecutedByUUID.String(),

			Metadata: []byte(mapper.StringPtrToString(req.ExtensionExecution.Metadata)),

			Status: req.ExtensionExecution.Status.ToInt64(),

			StatusMsg: mapper.StringPtrToSqlNullString(req.ExtensionExecution.StatusMsg),

			CreatedAt: existing.CreatedAt,

			UpdatedAt: custom.Now(),
		}
	}

	res := nemdb.UpdateExtensionExecutionParams{}

	res.UUID = req.ExtensionExecution.UUID.String()

	res.ExtensionUUID = req.ExtensionExecution.ExtensionUUID.String()

	res.ExtensionVersionUUID = req.ExtensionExecution.ExtensionVersionUUID.String()

	res.ProjectExtensionUUID = mapper.StringToSqlNullString(req.ExtensionExecution.ProjectExtensionUUID.String())

	res.ProjectUUID = req.ExtensionExecution.ProjectUUID.String()

	res.ProjectVersionUUID = req.ExtensionExecution.ProjectVersionUUID.String()

	res.ExecutedByUUID = req.ExtensionExecution.ExecutedByUUID.String()

	res.Metadata = []byte(mapper.StringPtrToString(req.ExtensionExecution.Metadata))

	if string(res.Metadata) == "" {
		res.Metadata = []byte(`{}`)
	}

	res.Status = req.ExtensionExecution.Status.ToInt64()

	res.StatusMsg = mapper.StringPtrToSqlNullString(req.ExtensionExecution.StatusMsg)

	res.CreatedAt = existing.CreatedAt

	res.UpdatedAt = custom.Now()

	return res

}

func (m *module) publishUpdateEvent(ctx context.Context,
	req types.UpsertRequest,
	qtx *nemdb.Queries,
	id string,
	existing []nemdb.ExtensionExecution) error {

	if m.events == nil {
		return nil
	}
	fetched, err := qtx.FetchExtensionExecutionByUUID(ctx, id)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_extension_execution",
			Message:          "error fetching after UpsertExtensionExecution - with uuid",
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
	existingEntities := mapModelsToEntities(existing)
	err = m.events.ProduceEntityUpdated(fetchedEntities[0], existingEntities[0])
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_extension_execution",
			Message:          "error producing insert event for UpsertExtensionExecution - with uuid",
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
