package extension_execution

import (
	"context"
	"errors"

	"github.com/nuzur/nem/core/module/extension_execution/types"
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

			ProjectExtensionUUID: req.ExtensionExecution.ProjectExtensionUUID.String(),

			ProjectUUID: req.ExtensionExecution.ProjectUUID.String(),

			ProjectVersionUUID: req.ExtensionExecution.ProjectVersionUUID.String(),

			ExecutedByUUID: req.ExtensionExecution.ExecutedByUUID.String(),

			Metadata: []byte(req.ExtensionExecution.Metadata),

			Status: req.ExtensionExecution.Status.ToInt64(),

			StatusMsg: req.ExtensionExecution.StatusMsg,

			CreatedAt: existing.CreatedAt,

			UpdatedAt: custom.Now(),
		}
	}

	res := nemdb.UpdateExtensionExecutionParams{}
	emptyReq := types.UpsertRequest{}

	// regular field
	if req.ExtensionExecution.UUID == emptyReq.ExtensionExecution.UUID {

		res.UUID = existing.UUID
	} else {

		res.UUID = req.ExtensionExecution.UUID.String()

	}

	// regular field
	if req.ExtensionExecution.ExtensionUUID == emptyReq.ExtensionExecution.ExtensionUUID {

		res.ExtensionUUID = existing.ExtensionUUID
	} else {

		res.ExtensionUUID = req.ExtensionExecution.ExtensionUUID.String()

	}

	// regular field
	if req.ExtensionExecution.ExtensionVersionUUID == emptyReq.ExtensionExecution.ExtensionVersionUUID {

		res.ExtensionVersionUUID = existing.ExtensionVersionUUID
	} else {

		res.ExtensionVersionUUID = req.ExtensionExecution.ExtensionVersionUUID.String()

	}

	// regular field
	if req.ExtensionExecution.ProjectExtensionUUID == emptyReq.ExtensionExecution.ProjectExtensionUUID {

		res.ProjectExtensionUUID = existing.ProjectExtensionUUID
	} else {

		res.ProjectExtensionUUID = req.ExtensionExecution.ProjectExtensionUUID.String()

	}

	// regular field
	if req.ExtensionExecution.ProjectUUID == emptyReq.ExtensionExecution.ProjectUUID {

		res.ProjectUUID = existing.ProjectUUID
	} else {

		res.ProjectUUID = req.ExtensionExecution.ProjectUUID.String()

	}

	// regular field
	if req.ExtensionExecution.ProjectVersionUUID == emptyReq.ExtensionExecution.ProjectVersionUUID {

		res.ProjectVersionUUID = existing.ProjectVersionUUID
	} else {

		res.ProjectVersionUUID = req.ExtensionExecution.ProjectVersionUUID.String()

	}

	// regular field
	if req.ExtensionExecution.ExecutedByUUID == emptyReq.ExtensionExecution.ExecutedByUUID {

		res.ExecutedByUUID = existing.ExecutedByUUID
	} else {

		res.ExecutedByUUID = req.ExtensionExecution.ExecutedByUUID.String()

	}

	// regular field
	if req.ExtensionExecution.Metadata == emptyReq.ExtensionExecution.Metadata {

		res.Metadata = existing.Metadata
	} else {

		res.Metadata = []byte(req.ExtensionExecution.Metadata)

	}

	// regular field
	if req.ExtensionExecution.Status == emptyReq.ExtensionExecution.Status {

		res.Status = existing.Status
	} else {

		res.Status = req.ExtensionExecution.Status.ToInt64()

	}

	// regular field
	if req.ExtensionExecution.StatusMsg == emptyReq.ExtensionExecution.StatusMsg {

		res.StatusMsg = existing.StatusMsg
	} else {

		res.StatusMsg = req.ExtensionExecution.StatusMsg

	}

	// regular field
	if req.ExtensionExecution.CreatedAt == emptyReq.ExtensionExecution.CreatedAt {

		res.CreatedAt = existing.CreatedAt
	} else {

		res.CreatedAt = existing.CreatedAt

	}

	// regular field
	if req.ExtensionExecution.UpdatedAt == emptyReq.ExtensionExecution.UpdatedAt {

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
