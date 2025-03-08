package extension_version

import (
	"context"
	"errors"

	"github.com/nuzur/nem/core/module/extension_version/types"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"

	main_entity "github.com/nuzur/nem/core/entity/extension_version"

	"github.com/nuzur/nem/custom"

	"time"
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
	existing, err := qtx.FetchExtensionVersionByUUIDForUpdate(ctx, req.ExtensionVersion.UUID.String())
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_extension_version",
			Message:          "error fetching existing record for UpsertExtensionVersion - with uuid",
			EntityIdentifier: "extension_version",
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
			ActionIdentifier: "upsert_extension_version",
			Message:          "not found existing record for UpsertExtensionVersion - with uuid",
			EntityIdentifier: "extension_version",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	if existing[0].Version != req.ExtensionVersion.Version {
		err := errors.New("upsert version conflict")
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_extension_version",
			Message:          "version conflict UpsertExtensionVersion",
			EntityIdentifier: "extension_version",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	req.ExtensionVersion.Version = time.Now().Unix()

	params := mapUpsertRequestToUpdateParams(req, existing[0], partial)
	err = qtx.UpdateExtensionVersion(
		ctx,
		params,
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_extension_version",
			Message:          "error calling repository for UpsertExtensionVersion - with uuid",
			EntityIdentifier: "extension_version",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	err = m.publishUpdateEvent(ctx, req, qtx, req.ExtensionVersion.UUID.String(), existing)
	if err != nil {
		return types.UpsertResponse{}, err
	}

	if createdTx {
		err := tx.Commit()
		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "upsert_extension_version",
				Message:          "error commiting for UpsertExtensionVersion",
				EntityIdentifier: "extension_version",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.UpsertResponse{}, err
		}
	}

	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "upsert_extension_version",
		Message:          "successfully handled UpsertExtensionVersion - with uuid",
		EntityIdentifier: "extension_version",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return types.UpsertResponse{
		UUID: req.ExtensionVersion.UUID,
	}, nil
}

func mapUpsertRequestToUpdateParams(req types.UpsertRequest, existing nemdb.ExtensionVersion, partial bool) nemdb.UpdateExtensionVersionParams {
	if !partial {
		return nemdb.UpdateExtensionVersionParams{

			UUID: req.ExtensionVersion.UUID.String(),

			Version: req.ExtensionVersion.Version,

			ExtensionUUID: req.ExtensionVersion.ExtensionUUID.String(),

			DisplayVersion: req.ExtensionVersion.DisplayVersion,

			Description: req.ExtensionVersion.Description,

			RepositoryTag: req.ExtensionVersion.RepositoryTag,

			ConfigurationEntity: []byte(req.ExtensionVersion.ConfigurationEntity),

			ExecutionMode: main_entity.ExecutionModeSliceToJSON(req.ExtensionVersion.ExecutionMode),

			ReviewStatus: req.ExtensionVersion.ReviewStatus.ToInt64(),

			Status: req.ExtensionVersion.Status.ToInt64(),

			CreatedAt: existing.CreatedAt,

			UpdatedAt: custom.Now(),

			CreatedByUUID: req.ExtensionVersion.CreatedByUUID.String(),

			UpdatedByUUID: req.ExtensionVersion.UpdatedByUUID.String(),
		}
	}

	res := nemdb.UpdateExtensionVersionParams{}

	res.UUID = req.ExtensionVersion.UUID.String()

	res.Version = req.ExtensionVersion.Version

	res.ExtensionUUID = req.ExtensionVersion.ExtensionUUID.String()

	res.DisplayVersion = req.ExtensionVersion.DisplayVersion

	res.Description = req.ExtensionVersion.Description

	res.RepositoryTag = req.ExtensionVersion.RepositoryTag

	res.ConfigurationEntity = []byte(req.ExtensionVersion.ConfigurationEntity)

	if string(res.ConfigurationEntity) == "" {
		res.ConfigurationEntity = []byte(`{}`)
	}

	res.ExecutionMode = main_entity.ExecutionModeSliceToJSON(req.ExtensionVersion.ExecutionMode)

	res.ReviewStatus = req.ExtensionVersion.ReviewStatus.ToInt64()

	res.Status = req.ExtensionVersion.Status.ToInt64()

	res.CreatedAt = existing.CreatedAt

	res.UpdatedAt = custom.Now()

	res.CreatedByUUID = req.ExtensionVersion.CreatedByUUID.String()

	res.UpdatedByUUID = req.ExtensionVersion.UpdatedByUUID.String()

	return res

}

func (m *module) publishUpdateEvent(ctx context.Context,
	req types.UpsertRequest,
	qtx *nemdb.Queries,
	id string,
	existing []nemdb.ExtensionVersion) error {

	if m.events == nil {
		return nil
	}
	fetched, err := qtx.FetchExtensionVersionByUUID(ctx, id)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_extension_version",
			Message:          "error fetching after UpsertExtensionVersion - with uuid",
			EntityIdentifier: "extension_version",
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
			ActionIdentifier: "upsert_extension_version",
			Message:          "error producing insert event for UpsertExtensionVersion - with uuid",
			EntityIdentifier: "extension_version",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return err
	}
	return nil
}
