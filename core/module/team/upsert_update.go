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
	"github.com/nuzur/nem/core/entity/membership"
	"github.com/nuzur/nem/core/entity/object_store"
	"github.com/nuzur/nem/core/entity/review_config"
	"github.com/nuzur/nem/core/entity/store"

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

			Memberships: membership.MembershipSliceToJSON(req.Team.Memberships),

			Stores: store.StoreSliceToJSON(req.Team.Stores),

			Connections: connection.ConnectionSliceToJSON(req.Team.Connections),

			ObjectStores: object_store.ObjectStoreSliceToJSON(req.Team.ObjectStores),

			OrganizationUUID: req.Team.OrganizationUUID.String(),

			DefaultEntity: entity.EntityToJSON(req.Team.DefaultEntity),

			Status: req.Team.Status.ToInt64(),

			CreatedAt: existing.CreatedAt,

			UpdatedAt: custom.Now(),

			CreatedByUUID: req.Team.CreatedByUUID.String(),

			UpdatedByUUID: req.Team.UpdatedByUUID.String(),
		}
	}

	res := nemdb.UpdateTeamParams{}
	emptyReq := types.UpsertRequest{}

	// regular field
	if req.Team.UUID == emptyReq.Team.UUID {

		res.UUID = existing.UUID
	} else {

		res.UUID = req.Team.UUID.String()

	}

	// regular field
	if req.Team.Version == emptyReq.Team.Version {

		res.Version = existing.Version
	} else {

		res.Version = req.Team.Version

	}

	// regular field
	if req.Team.Name == emptyReq.Team.Name {

		res.Name = existing.Name
	} else {

		res.Name = req.Team.Name

	}

	// json array
	if len(req.Team.Enviorments) == 0 {

		res.Enviorments = existing.Enviorments
	} else {

		res.Enviorments = enviorment.EnviormentSliceToJSON(req.Team.Enviorments)

	}

	// json array
	if len(req.Team.ReviewConfigs) == 0 {

		res.ReviewConfigs = existing.ReviewConfigs
	} else {

		res.ReviewConfigs = review_config.ReviewConfigSliceToJSON(req.Team.ReviewConfigs)

	}

	// json array
	if len(req.Team.Memberships) == 0 {

		res.Memberships = existing.Memberships
	} else {

		res.Memberships = membership.MembershipSliceToJSON(req.Team.Memberships)

	}

	// json array
	if len(req.Team.Stores) == 0 {

		res.Stores = existing.Stores
	} else {

		res.Stores = store.StoreSliceToJSON(req.Team.Stores)

	}

	// json array
	if len(req.Team.Connections) == 0 {

		res.Connections = existing.Connections
	} else {

		res.Connections = connection.ConnectionSliceToJSON(req.Team.Connections)

	}

	// json array
	if len(req.Team.ObjectStores) == 0 {

		res.ObjectStores = existing.ObjectStores
	} else {

		res.ObjectStores = object_store.ObjectStoreSliceToJSON(req.Team.ObjectStores)

	}

	// regular field
	if req.Team.OrganizationUUID == emptyReq.Team.OrganizationUUID {

		res.OrganizationUUID = existing.OrganizationUUID
	} else {

		res.OrganizationUUID = req.Team.OrganizationUUID.String()

	}

	if req.Team.DefaultEntity.ToJSON() == nil {

		res.DefaultEntity = existing.DefaultEntity
	} else {

		res.DefaultEntity = entity.EntityToJSON(req.Team.DefaultEntity)

	}

	// regular field
	if req.Team.Status == emptyReq.Team.Status {

		res.Status = existing.Status
	} else {

		res.Status = req.Team.Status.ToInt64()

	}

	// regular field
	if req.Team.CreatedAt == emptyReq.Team.CreatedAt {

		res.CreatedAt = existing.CreatedAt
	} else {

		res.CreatedAt = existing.CreatedAt

	}

	// regular field
	if req.Team.UpdatedAt == emptyReq.Team.UpdatedAt {

		res.UpdatedAt = existing.UpdatedAt
	} else {

		res.UpdatedAt = custom.Now()

	}

	// regular field
	if req.Team.CreatedByUUID == emptyReq.Team.CreatedByUUID {

		res.CreatedByUUID = existing.CreatedByUUID
	} else {

		res.CreatedByUUID = req.Team.CreatedByUUID.String()

	}

	// regular field
	if req.Team.UpdatedByUUID == emptyReq.Team.UpdatedByUUID {

		res.UpdatedByUUID = existing.UpdatedByUUID
	} else {

		res.UpdatedByUUID = req.Team.UpdatedByUUID.String()

	}

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
