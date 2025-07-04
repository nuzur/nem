package user_project

import (
	"context"
	"errors"

	"github.com/nuzur/nem/core/module/user_project/types"
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
	existing, err := qtx.FetchUserProjectByUUIDForUpdate(ctx, req.UserProject.UUID.String())
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_user_project",
			Message:          "error fetching existing record for UpsertUserProject - with uuid",
			EntityIdentifier: "user_project",
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
			ActionIdentifier: "upsert_user_project",
			Message:          "not found existing record for UpsertUserProject - with uuid",
			EntityIdentifier: "user_project",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	params := mapUpsertRequestToUpdateParams(req, existing[0], partial)
	err = qtx.UpdateUserProject(
		ctx,
		params,
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_user_project",
			Message:          "error calling repository for UpsertUserProject - with uuid",
			EntityIdentifier: "user_project",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	err = m.publishUpdateEvent(ctx, req, qtx, req.UserProject.UUID.String(), existing)
	if err != nil {
		return types.UpsertResponse{}, err
	}

	if createdTx {
		err := tx.Commit()
		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "upsert_user_project",
				Message:          "error commiting for UpsertUserProject",
				EntityIdentifier: "user_project",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.UpsertResponse{}, err
		}
	}

	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "upsert_user_project",
		Message:          "successfully handled UpsertUserProject - with uuid",
		EntityIdentifier: "user_project",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return types.UpsertResponse{
		UUID: req.UserProject.UUID,
	}, nil
}

func mapUpsertRequestToUpdateParams(req types.UpsertRequest, existing nemdb.UserProject, partial bool) nemdb.UpdateUserProjectParams {
	if !partial {
		return nemdb.UpdateUserProjectParams{

			UUID: req.UserProject.UUID.String(),

			UserUUID: mapper.StringToSqlNullString(req.UserProject.UserUUID.String()),

			UserEmail: mapper.StringPtrToSqlNullString(req.UserProject.UserEmail),

			ProjectUUID: req.UserProject.ProjectUUID.String(),

			Role: req.UserProject.Role.ToInt64(),

			ReviewRequiredStructure: req.UserProject.ReviewRequiredStructure,

			ReviewRequiredData: req.UserProject.ReviewRequiredData,

			Status: req.UserProject.Status.ToInt64(),

			CreatedAt: existing.CreatedAt,

			UpdatedAt: custom.Now(),

			CreatedByUUID: req.UserProject.CreatedByUUID.String(),

			UpdatedByUUID: req.UserProject.UpdatedByUUID.String(),
		}
	}

	res := nemdb.UpdateUserProjectParams{}

	res.UUID = req.UserProject.UUID.String()

	res.UserUUID = mapper.StringToSqlNullString(req.UserProject.UserUUID.String())

	res.UserEmail = mapper.StringPtrToSqlNullString(req.UserProject.UserEmail)

	res.ProjectUUID = req.UserProject.ProjectUUID.String()

	res.Role = req.UserProject.Role.ToInt64()

	res.ReviewRequiredStructure = req.UserProject.ReviewRequiredStructure

	res.ReviewRequiredData = req.UserProject.ReviewRequiredData

	res.Status = req.UserProject.Status.ToInt64()

	res.CreatedAt = existing.CreatedAt

	res.UpdatedAt = custom.Now()

	res.CreatedByUUID = req.UserProject.CreatedByUUID.String()

	res.UpdatedByUUID = req.UserProject.UpdatedByUUID.String()

	return res

}

func (m *module) publishUpdateEvent(ctx context.Context,
	req types.UpsertRequest,
	qtx *nemdb.Queries,
	id string,
	existing []nemdb.UserProject) error {

	if m.events == nil {
		return nil
	}
	fetched, err := qtx.FetchUserProjectByUUID(ctx, id)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_user_project",
			Message:          "error fetching after UpsertUserProject - with uuid",
			EntityIdentifier: "user_project",
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
			ActionIdentifier: "upsert_user_project",
			Message:          "error producing insert event for UpsertUserProject - with uuid",
			EntityIdentifier: "user_project",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return err
	}
	return nil
}
