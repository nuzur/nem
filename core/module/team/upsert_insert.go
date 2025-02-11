package team

import (
	"context"

	"errors"

	"github.com/gofrs/uuid"
	"nem/core/module/team/types"
	nemdb "nem/core/repository/gen"
	"nem/monitoring"

	"nem/custom"

	"nem/core/entity/connection"
	"nem/core/entity/entity"
	"nem/core/entity/enviorment"
	"nem/core/entity/membership"
	"nem/core/entity/object_store"
	"nem/core/entity/review_config"
	"nem/core/entity/store"

	"time"
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

	_, err := qtx.InsertTeam(
		ctx,
		params,
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_team",
			Message:          "error calling repository for UpsertTeam - no uuid",
			EntityIdentifier: "team",
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
				ActionIdentifier: "upsert_team",
				Message:          "error commiting for UpsertTeam (1)",
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
		Message:          "successfully handled UpsertTeam - no uuid",
		EntityIdentifier: "team",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return types.UpsertResponse{
		UUID: uuid.FromStringOrNil(params.UUID),
	}, nil
}

func mapUpsertRequestToInsertParams(req types.UpsertRequest) nemdb.InsertTeamParams {
	return nemdb.InsertTeamParams{

		UUID: custom.GenerateUUID().String(),

		Version: req.Team.Version,

		Name: req.Team.Name,

		Enviorments: enviorment.EnviormentSliceToJSON(req.Team.Enviorments),

		ReviewConfigs: review_config.ReviewConfigSliceToJSON(req.Team.ReviewConfigs),

		Memberships: membership.MembershipSliceToJSON(req.Team.Memberships),

		Stores: store.StoreSliceToJSON(req.Team.Stores),

		Connections: connection.ConnectionSliceToJSON(req.Team.Connections),

		ObjectStores: object_store.ObjectStoreSliceToJSON(req.Team.ObjectStores),

		OrganizationUUID: custom.GenerateUUID().String(),

		DefaultEntity: entity.EntityToJSON(req.Team.DefaultEntity),

		Status: req.Team.Status.ToInt64(),

		CreatedAt: custom.Now(),

		UpdatedAt: custom.Now(),

		CreatedByUUID: req.Team.CreatedByUUID.String(),

		UpdatedByUUID: req.Team.UpdatedByUUID.String(),
	}
}

func (m *module) publishInsertEvent(ctx context.Context,
	req types.UpsertRequest,
	qtx *nemdb.Queries,
	id string) error {

	if m.events == nil {
		return nil
	}

	fetched, err := qtx.FetchTeamByUUID(ctx, id)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_team",
			Message:          "error fetching after UpsertTeam - no uuid",
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
	err = m.events.ProduceEntityInserted(fetchedEntities[0])
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_team",
			Message:          "error producing insert event for UpsertTeam - no uuid",
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
