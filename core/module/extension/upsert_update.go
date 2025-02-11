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
	emptyReq := types.UpsertRequest{}

	// regular field
	if req.Extension.UUID == emptyReq.Extension.UUID {

		res.UUID = existing.UUID
	} else {

		res.UUID = req.Extension.UUID.String()

	}

	// regular field
	if req.Extension.Version == emptyReq.Extension.Version {

		res.Version = existing.Version
	} else {

		res.Version = req.Extension.Version

	}

	// regular field
	if req.Extension.Identifier == emptyReq.Extension.Identifier {

		res.Identifier = existing.Identifier
	} else {

		res.Identifier = req.Extension.Identifier

	}

	// regular field
	if req.Extension.DisplayName == emptyReq.Extension.DisplayName {

		res.DisplayName = existing.DisplayName
	} else {

		res.DisplayName = req.Extension.DisplayName

	}

	// regular field
	if req.Extension.DisplayAuthorName == emptyReq.Extension.DisplayAuthorName {

		res.DisplayAuthorName = existing.DisplayAuthorName
	} else {

		res.DisplayAuthorName = req.Extension.DisplayAuthorName

	}

	// regular field
	if req.Extension.Description == emptyReq.Extension.Description {

		res.Description = existing.Description
	} else {

		res.Description = req.Extension.Description

	}

	// regular field
	if req.Extension.URL == emptyReq.Extension.URL {

		res.URL = existing.URL
	} else {

		res.URL = req.Extension.URL

	}

	// regular field
	if req.Extension.Verfied == emptyReq.Extension.Verfied {

		res.Verfied = existing.Verfied
	} else {

		res.Verfied = req.Extension.Verfied

	}

	// regular field
	if req.Extension.Repository == emptyReq.Extension.Repository {

		res.Repository = existing.Repository
	} else {

		res.Repository = req.Extension.Repository

	}

	// regular field
	if req.Extension.ExtensionType == emptyReq.Extension.ExtensionType {

		res.ExtensionType = existing.ExtensionType
	} else {

		res.ExtensionType = req.Extension.ExtensionType.ToInt64()

	}

	// raw json is a pointer
	if req.Extension.Tags == nil {

		res.Tags = existing.Tags
	} else {

		res.Tags = mapper.SliceToJSON(req.Extension.Tags)

	}

	// regular field
	if req.Extension.Public == emptyReq.Extension.Public {

		res.Public = existing.Public
	} else {

		res.Public = req.Extension.Public

	}

	// json array
	if len(req.Extension.Visibility) == 0 {

		res.Visibility = existing.Visibility
	} else {

		res.Visibility = visibility.VisibilitySliceToJSON(req.Extension.Visibility)

	}

	// regular field
	if req.Extension.Status == emptyReq.Extension.Status {

		res.Status = existing.Status
	} else {

		res.Status = req.Extension.Status.ToInt64()

	}

	// regular field
	if req.Extension.OwnerUUID == emptyReq.Extension.OwnerUUID {

		res.OwnerUUID = existing.OwnerUUID
	} else {

		res.OwnerUUID = req.Extension.OwnerUUID.String()

	}

	// regular field
	if req.Extension.CreatedAt == emptyReq.Extension.CreatedAt {

		res.CreatedAt = existing.CreatedAt
	} else {

		res.CreatedAt = existing.CreatedAt

	}

	// regular field
	if req.Extension.UpdatedAt == emptyReq.Extension.UpdatedAt {

		res.UpdatedAt = existing.UpdatedAt
	} else {

		res.UpdatedAt = custom.Now()

	}

	// regular field
	if req.Extension.CreatedByUUID == emptyReq.Extension.CreatedByUUID {

		res.CreatedByUUID = existing.CreatedByUUID
	} else {

		res.CreatedByUUID = req.Extension.CreatedByUUID.String()

	}

	// regular field
	if req.Extension.UpdatedByUUID == emptyReq.Extension.UpdatedByUUID {

		res.UpdatedByUUID = existing.UpdatedByUUID
	} else {

		res.UpdatedByUUID = req.Extension.UpdatedByUUID.String()

	}

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
