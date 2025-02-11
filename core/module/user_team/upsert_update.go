package user_team

import (
	"context"
	"errors"

	"nem/core/module/user_team/types"
	nemdb "nem/core/repository/gen"
	"nem/monitoring"

	main_entity "nem/core/entity/user_team"

	"nem/custom"
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

			UserUUID: req.UserTeam.UserUUID.String(),

			UserEmail: req.UserTeam.UserEmail,

			TeamUUID: req.UserTeam.TeamUUID.String(),

			Roles: main_entity.RoleSliceToJSON(req.UserTeam.Roles),

			Status: req.UserTeam.Status.ToInt64(),

			CreatedAt: existing.CreatedAt,

			UpdatedAt: custom.Now(),

			CreatedByUUID: req.UserTeam.CreatedByUUID.String(),

			UpdatedByUUID: req.UserTeam.UpdatedByUUID.String(),
		}
	}

	res := nemdb.UpdateUserTeamParams{}
	emptyReq := types.UpsertRequest{}

	// regular field
	if req.UserTeam.UUID == emptyReq.UserTeam.UUID {

		res.UUID = existing.UUID
	} else {

		res.UUID = req.UserTeam.UUID.String()

	}

	// regular field
	if req.UserTeam.UserUUID == emptyReq.UserTeam.UserUUID {

		res.UserUUID = existing.UserUUID
	} else {

		res.UserUUID = req.UserTeam.UserUUID.String()

	}

	// regular field
	if req.UserTeam.UserEmail == emptyReq.UserTeam.UserEmail {

		res.UserEmail = existing.UserEmail
	} else {

		res.UserEmail = req.UserTeam.UserEmail

	}

	// regular field
	if req.UserTeam.TeamUUID == emptyReq.UserTeam.TeamUUID {

		res.TeamUUID = existing.TeamUUID
	} else {

		res.TeamUUID = req.UserTeam.TeamUUID.String()

	}

	// raw json is a pointer
	if req.UserTeam.Roles == nil {

		res.Roles = existing.Roles
	} else {

		res.Roles = main_entity.RoleSliceToJSON(req.UserTeam.Roles)

	}

	// regular field
	if req.UserTeam.Status == emptyReq.UserTeam.Status {

		res.Status = existing.Status
	} else {

		res.Status = req.UserTeam.Status.ToInt64()

	}

	// regular field
	if req.UserTeam.CreatedAt == emptyReq.UserTeam.CreatedAt {

		res.CreatedAt = existing.CreatedAt
	} else {

		res.CreatedAt = existing.CreatedAt

	}

	// regular field
	if req.UserTeam.UpdatedAt == emptyReq.UserTeam.UpdatedAt {

		res.UpdatedAt = existing.UpdatedAt
	} else {

		res.UpdatedAt = custom.Now()

	}

	// regular field
	if req.UserTeam.CreatedByUUID == emptyReq.UserTeam.CreatedByUUID {

		res.CreatedByUUID = existing.CreatedByUUID
	} else {

		res.CreatedByUUID = req.UserTeam.CreatedByUUID.String()

	}

	// regular field
	if req.UserTeam.UpdatedByUUID == emptyReq.UserTeam.UpdatedByUUID {

		res.UpdatedByUUID = existing.UpdatedByUUID
	} else {

		res.UpdatedByUUID = req.UserTeam.UpdatedByUUID.String()

	}

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
