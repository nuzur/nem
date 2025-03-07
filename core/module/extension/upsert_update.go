package extension

import (
	"context"
	"errors"

	"github.com/nuzur/nem/core/module/extension/types"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"

	"github.com/nuzur/nem/core/entity/visibility"

	"github.com/nuzur/nem/custom"

	"time"

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
	existing, err := qtx.FetchExtensionByUUIDForUpdate(ctx, req.Extension.UUID.String())
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_extension",
			Message:          "error fetching existing record for UpsertExtension - with uuid",
			EntityIdentifier: "extension",
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
			ActionIdentifier: "upsert_extension",
			Message:          "not found existing record for UpsertExtension - with uuid",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	if existing[0].Version != req.Extension.Version {
		err := errors.New("upsert version conflict")
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_extension",
			Message:          "version conflict UpsertExtension",
			EntityIdentifier: "extension",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	req.Extension.Version = time.Now().Unix()

	params := mapUpsertRequestToUpdateParams(req, existing[0], partial)
	err = qtx.UpdateExtension(
		ctx,
		params,
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_extension",
			Message:          "error calling repository for UpsertExtension - with uuid",
			EntityIdentifier: "extension",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	err = m.publishUpdateEvent(ctx, req, qtx, req.Extension.UUID.String(), existing)
	if err != nil {
		return types.UpsertResponse{}, err
	}

	if createdTx {
		err := tx.Commit()
		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "upsert_extension",
				Message:          "error commiting for UpsertExtension",
				EntityIdentifier: "extension",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.UpsertResponse{}, err
		}
	}

	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "upsert_extension",
		Message:          "successfully handled UpsertExtension - with uuid",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return types.UpsertResponse{
		UUID: req.Extension.UUID,
	}, nil
}

func mapUpsertRequestToUpdateParams(req types.UpsertRequest, existing nemdb.Extension, partial bool) nemdb.UpdateExtensionParams {
	if !partial {
		return nemdb.UpdateExtensionParams{

			UUID: req.Extension.UUID.String(),

			Version: req.Extension.Version,

			Identifier: req.Extension.Identifier,

			DisplayName: req.Extension.DisplayName,

			DisplayAuthorName: req.Extension.DisplayAuthorName,

			Description: req.Extension.Description,

			URL: req.Extension.URL,

			Verfied: req.Extension.Verfied,

			Repository: req.Extension.Repository,

			ExtensionType: req.Extension.ExtensionType.ToInt64(),

			Tags: mapper.SliceToJSON(req.Extension.Tags),

			Public: req.Extension.Public,

			Visibility: visibility.VisibilitySliceToJSON(req.Extension.Visibility),

			Status: req.Extension.Status.ToInt64(),

			OwnerUUID: req.Extension.OwnerUUID.String(),

			CreatedAt: existing.CreatedAt,

			UpdatedAt: custom.Now(),

			CreatedByUUID: req.Extension.CreatedByUUID.String(),

			UpdatedByUUID: req.Extension.UpdatedByUUID.String(),
		}
	}

	res := nemdb.UpdateExtensionParams{}

	res.UUID = req.Extension.UUID.String()

	res.Version = req.Extension.Version

	res.Identifier = req.Extension.Identifier

	res.DisplayName = req.Extension.DisplayName

	res.DisplayAuthorName = req.Extension.DisplayAuthorName

	res.Description = req.Extension.Description

	res.URL = req.Extension.URL

	res.Verfied = req.Extension.Verfied

	res.Repository = req.Extension.Repository

	res.ExtensionType = req.Extension.ExtensionType.ToInt64()

	res.Tags = mapper.SliceToJSON(req.Extension.Tags)

	res.Public = req.Extension.Public

	res.Visibility = visibility.VisibilitySliceToJSON(req.Extension.Visibility)

	res.Status = req.Extension.Status.ToInt64()

	res.OwnerUUID = req.Extension.OwnerUUID.String()

	res.CreatedAt = existing.CreatedAt

	res.UpdatedAt = custom.Now()

	res.CreatedByUUID = req.Extension.CreatedByUUID.String()

	res.UpdatedByUUID = req.Extension.UpdatedByUUID.String()

	return res

}

func (m *module) publishUpdateEvent(ctx context.Context,
	req types.UpsertRequest,
	qtx *nemdb.Queries,
	id string,
	existing []nemdb.Extension) error {

	if m.events == nil {
		return nil
	}
	fetched, err := qtx.FetchExtensionByUUID(ctx, id)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_extension",
			Message:          "error fetching after UpsertExtension - with uuid",
			EntityIdentifier: "extension",
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
			ActionIdentifier: "upsert_extension",
			Message:          "error producing insert event for UpsertExtension - with uuid",
			EntityIdentifier: "extension",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return err
	}
	return nil
}
