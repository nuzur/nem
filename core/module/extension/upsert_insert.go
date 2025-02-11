package extension

import (
	"context"

	"errors"

	"github.com/gofrs/uuid"
	"github.com/nuzur/nem/core/module/extension/types"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"

	"github.com/nuzur/nem/custom"

	"github.com/nuzur/nem/core/entity/visibility"

	"time"

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

	params.Version = time.Now().Unix()

	_, err := qtx.InsertExtension(
		ctx,
		params,
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_extension",
			Message:          "error calling repository for UpsertExtension - no uuid",
			EntityIdentifier: "extension",
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
				ActionIdentifier: "upsert_extension",
				Message:          "error commiting for UpsertExtension (1)",
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
		Message:          "successfully handled UpsertExtension - no uuid",
		EntityIdentifier: "extension",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return types.UpsertResponse{
		UUID: uuid.FromStringOrNil(params.UUID),
	}, nil
}

func mapUpsertRequestToInsertParams(req types.UpsertRequest) nemdb.InsertExtensionParams {
	return nemdb.InsertExtensionParams{

		UUID: custom.GenerateUUID().String(),

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

		CreatedAt: custom.Now(),

		UpdatedAt: custom.Now(),

		CreatedByUUID: req.Extension.CreatedByUUID.String(),

		UpdatedByUUID: req.Extension.UpdatedByUUID.String(),
	}
}

func (m *module) publishInsertEvent(ctx context.Context,
	req types.UpsertRequest,
	qtx *nemdb.Queries,
	id string) error {

	if m.events == nil {
		return nil
	}

	fetched, err := qtx.FetchExtensionByUUID(ctx, id)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_extension",
			Message:          "error fetching after UpsertExtension - no uuid",
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
	err = m.events.ProduceEntityInserted(fetchedEntities[0])
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_extension",
			Message:          "error producing insert event for UpsertExtension - no uuid",
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
