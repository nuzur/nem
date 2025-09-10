package team

import (
	"context"
	"errors"

	"github.com/nuzur/nem/core/module/team/types"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"

	"github.com/nuzur/nem/core/entity/connection"
	"github.com/nuzur/nem/core/entity/entity"
	"github.com/nuzur/nem/core/entity/enviorment"
	"github.com/nuzur/nem/core/entity/object_store"
	"github.com/nuzur/nem/core/entity/review_config"
	"github.com/nuzur/nem/core/entity/store"

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
	existing, err := qtx.FetchTeamByUUIDForUpdate(ctx, req.Team.UUID.String())
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_team",
			Message:          "error fetching existing record for UpsertTeam - with uuid",
			EntityIdentifier: "team",
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
			ActionIdentifier: "upsert_team",
			Message:          "not found existing record for UpsertTeam - with uuid",
			EntityIdentifier: "team",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	if existing[0].Version != req.Team.Version {
		err := errors.New("upsert version conflict")
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_team",
			Message:          "version conflict UpsertTeam",
			EntityIdentifier: "team",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	req.Team.Version = time.Now().Unix()

	params := mapUpsertRequestToUpdateParams(req, existing[0], partial)
	err = qtx.UpdateTeam(
		ctx,
		params,
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_team",
			Message:          "error calling repository for UpsertTeam - with uuid",
			EntityIdentifier: "team",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	err = m.publishUpdateEvent(ctx, req, qtx, req.Team.UUID.String(), existing)
	if err != nil {
		return types.UpsertResponse{}, err
	}

	if createdTx {
		err := tx.Commit()
		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "upsert_team",
				Message:          "error commiting for UpsertTeam",
				EntityIdentifier: "team",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.UpsertResponse{}, err
		}
	}

	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "upsert_team",
		Message:          "successfully handled UpsertTeam - with uuid",
		EntityIdentifier: "team",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return types.UpsertResponse{
		UUID: req.Team.UUID,
	}, nil
}

func mapUpsertRequestToUpdateParams(req types.UpsertRequest, existing nemdb.Team, partial bool) nemdb.UpdateTeamParams {
	if !partial {
		return nemdb.UpdateTeamParams{

			UUID: req.Team.UUID.String(),

			Version: req.Team.Version,

			Name: req.Team.Name,

			Enviorments: enviorment.EnviormentSliceToJSON(req.Team.Enviorments),

			ReviewConfigs: review_config.ReviewConfigSliceToJSON(req.Team.ReviewConfigs),

			Stores: store.StoreSliceToJSON(req.Team.Stores),

			Connections: connection.ConnectionSliceToJSON(req.Team.Connections),

			ObjectStores: object_store.ObjectStoreSliceToJSON(req.Team.ObjectStores),

			OrganizationUUID: mapper.StringToSqlNullString(req.Team.OrganizationUUID.String()),

			DefaultEntity: entity.EntityToJSON(req.Team.DefaultEntity),

			Status: req.Team.Status.ToInt64(),

			CreatedAt: existing.CreatedAt,

			UpdatedAt: custom.Now(),

			CreatedByUUID: req.Team.CreatedByUUID.String(),

			UpdatedByUUID: req.Team.UpdatedByUUID.String(),
		}
	}

	res := nemdb.UpdateTeamParams{}

	res.UUID = req.Team.UUID.String()

	res.Version = req.Team.Version

	res.Name = req.Team.Name

	res.Enviorments = enviorment.EnviormentSliceToJSON(req.Team.Enviorments)

	res.ReviewConfigs = review_config.ReviewConfigSliceToJSON(req.Team.ReviewConfigs)

	res.Stores = store.StoreSliceToJSON(req.Team.Stores)

	res.Connections = connection.ConnectionSliceToJSON(req.Team.Connections)

	res.ObjectStores = object_store.ObjectStoreSliceToJSON(req.Team.ObjectStores)

	res.OrganizationUUID = mapper.StringToSqlNullString(req.Team.OrganizationUUID.String())

	res.DefaultEntity = entity.EntityToJSON(req.Team.DefaultEntity)

	res.Status = req.Team.Status.ToInt64()

	res.CreatedAt = existing.CreatedAt

	res.UpdatedAt = custom.Now()

	res.CreatedByUUID = req.Team.CreatedByUUID.String()

	res.UpdatedByUUID = req.Team.UpdatedByUUID.String()

	return res

}

func (m *module) publishUpdateEvent(ctx context.Context,
	req types.UpsertRequest,
	qtx *nemdb.Queries,
	id string,
	existing []nemdb.Team) error {

	if m.events == nil {
		return nil
	}
	fetched, err := qtx.FetchTeamByUUID(ctx, id)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_team",
			Message:          "error fetching after UpsertTeam - with uuid",
			EntityIdentifier: "team",
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
			ActionIdentifier: "upsert_team",
			Message:          "error producing insert event for UpsertTeam - with uuid",
			EntityIdentifier: "team",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return err
	}
	return nil
}
