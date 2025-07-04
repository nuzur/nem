package user_team

import (
	"context"
	"errors"

	"github.com/nuzur/nem/core/module/user_team/types"
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
	existing, err := qtx.FetchUserTeamByUUIDForUpdate(ctx, req.UserTeam.UUID.String())
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_user_team",
			Message:          "error fetching existing record for UpsertUserTeam - with uuid",
			EntityIdentifier: "user_team",
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
			ActionIdentifier: "upsert_user_team",
			Message:          "not found existing record for UpsertUserTeam - with uuid",
			EntityIdentifier: "user_team",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	params := mapUpsertRequestToUpdateParams(req, existing[0], partial)
	err = qtx.UpdateUserTeam(
		ctx,
		params,
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_user_team",
			Message:          "error calling repository for UpsertUserTeam - with uuid",
			EntityIdentifier: "user_team",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	err = m.publishUpdateEvent(ctx, req, qtx, req.UserTeam.UUID.String(), existing)
	if err != nil {
		return types.UpsertResponse{}, err
	}

	if createdTx {
		err := tx.Commit()
		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "upsert_user_team",
				Message:          "error commiting for UpsertUserTeam",
				EntityIdentifier: "user_team",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.UpsertResponse{}, err
		}
	}

	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "upsert_user_team",
		Message:          "successfully handled UpsertUserTeam - with uuid",
		EntityIdentifier: "user_team",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return types.UpsertResponse{
		UUID: req.UserTeam.UUID,
	}, nil
}

func mapUpsertRequestToUpdateParams(req types.UpsertRequest, existing nemdb.UserTeam, partial bool) nemdb.UpdateUserTeamParams {
	if !partial {
		return nemdb.UpdateUserTeamParams{

			UUID: req.UserTeam.UUID.String(),

			UserUUID: mapper.StringToSqlNullString(req.UserTeam.UserUUID.String()),

			UserEmail: mapper.StringPtrToSqlNullString(req.UserTeam.UserEmail),

			TeamUUID: req.UserTeam.TeamUUID.String(),

			Role: req.UserTeam.Role.ToInt64(),

			ReviewRequiredStructure: req.UserTeam.ReviewRequiredStructure,

			ReviewRequiredData: req.UserTeam.ReviewRequiredData,

			Status: req.UserTeam.Status.ToInt64(),

			CreatedAt: existing.CreatedAt,

			UpdatedAt: custom.Now(),

			CreatedByUUID: req.UserTeam.CreatedByUUID.String(),

			UpdatedByUUID: req.UserTeam.UpdatedByUUID.String(),
		}
	}

	res := nemdb.UpdateUserTeamParams{}

	res.UUID = req.UserTeam.UUID.String()

	res.UserUUID = mapper.StringToSqlNullString(req.UserTeam.UserUUID.String())

	res.UserEmail = mapper.StringPtrToSqlNullString(req.UserTeam.UserEmail)

	res.TeamUUID = req.UserTeam.TeamUUID.String()

	res.Role = req.UserTeam.Role.ToInt64()

	res.ReviewRequiredStructure = req.UserTeam.ReviewRequiredStructure

	res.ReviewRequiredData = req.UserTeam.ReviewRequiredData

	res.Status = req.UserTeam.Status.ToInt64()

	res.CreatedAt = existing.CreatedAt

	res.UpdatedAt = custom.Now()

	res.CreatedByUUID = req.UserTeam.CreatedByUUID.String()

	res.UpdatedByUUID = req.UserTeam.UpdatedByUUID.String()

	return res

}

func (m *module) publishUpdateEvent(ctx context.Context,
	req types.UpsertRequest,
	qtx *nemdb.Queries,
	id string,
	existing []nemdb.UserTeam) error {

	if m.events == nil {
		return nil
	}
	fetched, err := qtx.FetchUserTeamByUUID(ctx, id)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_user_team",
			Message:          "error fetching after UpsertUserTeam - with uuid",
			EntityIdentifier: "user_team",
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
			ActionIdentifier: "upsert_user_team",
			Message:          "error producing insert event for UpsertUserTeam - with uuid",
			EntityIdentifier: "user_team",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return err
	}
	return nil
}
