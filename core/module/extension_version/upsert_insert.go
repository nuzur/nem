package extension_version

import (
	"context"

	"errors"

	"github.com/gofrs/uuid"
	"github.com/nuzur/nem/core/module/extension_version/types"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"

	"github.com/nuzur/nem/custom"

	main_entity "github.com/nuzur/nem/core/entity/extension_version"

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

	_, err := qtx.InsertExtensionVersion(
		ctx,
		params,
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_extension_version",
			Message:          "error calling repository for UpsertExtensionVersion - no uuid",
			EntityIdentifier: "extension_version",
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
				ActionIdentifier: "upsert_extension_version",
				Message:          "error commiting for UpsertExtensionVersion (1)",
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
		Message:          "successfully handled UpsertExtensionVersion - no uuid",
		EntityIdentifier: "extension_version",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return types.UpsertResponse{
		UUID: uuid.FromStringOrNil(params.UUID),
	}, nil
}

func mapUpsertRequestToInsertParams(req types.UpsertRequest) nemdb.InsertExtensionVersionParams {
	return nemdb.InsertExtensionVersionParams{

		UUID: custom.GenerateUUID().String(),

		Version: req.ExtensionVersion.Version,

		ExtensionUUID: req.ExtensionVersion.ExtensionUUID.String(),

		DisplayVersion: mapper.StringPtrToSqlNullString(req.ExtensionVersion.DisplayVersion),

		Description: mapper.StringPtrToSqlNullString(req.ExtensionVersion.Description),

		RepositoryTag: req.ExtensionVersion.RepositoryTag,

		ConfigurationEntity: []byte(mapper.StringPtrToString(req.ExtensionVersion.ConfigurationEntity)),

		ExecutionMode: main_entity.ExecutionModeSliceToJSON(req.ExtensionVersion.ExecutionMode),

		ReviewStatus: req.ExtensionVersion.ReviewStatus.ToInt64(),

		Status: req.ExtensionVersion.Status.ToInt64(),

		CreatedAt: custom.Now(),

		UpdatedAt: custom.Now(),

		CreatedByUUID: req.ExtensionVersion.CreatedByUUID.String(),

		UpdatedByUUID: req.ExtensionVersion.UpdatedByUUID.String(),
	}
}

func (m *module) publishInsertEvent(ctx context.Context,
	req types.UpsertRequest,
	qtx *nemdb.Queries,
	id string) error {

	if m.events == nil {
		return nil
	}

	fetched, err := qtx.FetchExtensionVersionByUUID(ctx, id)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_extension_version",
			Message:          "error fetching after UpsertExtensionVersion - no uuid",
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
	err = m.events.ProduceEntityInserted(fetchedEntities[0])
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_extension_version",
			Message:          "error producing insert event for UpsertExtensionVersion - no uuid",
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
