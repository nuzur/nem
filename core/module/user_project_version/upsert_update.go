package user_project_version

import (
	"context"
	"errors"

	"github.com/nuzur/nem/core/module/user_project_version/types"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"

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
	existing, err := qtx.FetchUserProjectVersionByUUIDForUpdate(ctx, req.UserProjectVersion.UUID.String())
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_user_project_version",
			Message:          "error fetching existing record for UpsertUserProjectVersion - with uuid",
			EntityIdentifier: "user_project_version",
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
			ActionIdentifier: "upsert_user_project_version",
			Message:          "not found existing record for UpsertUserProjectVersion - with uuid",
			EntityIdentifier: "user_project_version",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	if existing[0].Version != req.UserProjectVersion.Version {
		err := errors.New("upsert version conflict")
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_user_project_version",
			Message:          "version conflict UpsertUserProjectVersion",
			EntityIdentifier: "user_project_version",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	req.UserProjectVersion.Version = time.Now().Unix()

	params := mapUpsertRequestToUpdateParams(req, existing[0], partial)
	err = qtx.UpdateUserProjectVersion(
		ctx,
		params,
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_user_project_version",
			Message:          "error calling repository for UpsertUserProjectVersion - with uuid",
			EntityIdentifier: "user_project_version",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	err = m.publishUpdateEvent(ctx, req, qtx, req.UserProjectVersion.UUID.String(), existing)
	if err != nil {
		return types.UpsertResponse{}, err
	}

	if createdTx {
		err := tx.Commit()
		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "upsert_user_project_version",
				Message:          "error commiting for UpsertUserProjectVersion",
				EntityIdentifier: "user_project_version",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.UpsertResponse{}, err
		}
	}

	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "upsert_user_project_version",
		Message:          "successfully handled UpsertUserProjectVersion - with uuid",
		EntityIdentifier: "user_project_version",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return types.UpsertResponse{
		UUID: req.UserProjectVersion.UUID,
	}, nil
}

func mapUpsertRequestToUpdateParams(req types.UpsertRequest, existing nemdb.UserProjectVersion, partial bool) nemdb.UpdateUserProjectVersionParams {
	if !partial {
		return nemdb.UpdateUserProjectVersionParams{

			UUID: req.UserProjectVersion.UUID.String(),

			Version: req.UserProjectVersion.Version,

			ProjectVersionUUID: req.UserProjectVersion.ProjectVersionUUID.String(),

			UserUUID: req.UserProjectVersion.UserUUID.String(),

			Data: []byte(mapper.StringPtrToString(req.UserProjectVersion.Data)),

			Status: req.UserProjectVersion.Status.ToInt64(),

			CreatedAt: existing.CreatedAt,

			UpdatedAt: custom.Now(),

			CreatedByUUID: req.UserProjectVersion.CreatedByUUID.String(),

			UpdatedByUUID: req.UserProjectVersion.UpdatedByUUID.String(),
		}
	}

	res := nemdb.UpdateUserProjectVersionParams{}

	res.UUID = req.UserProjectVersion.UUID.String()

	res.Version = req.UserProjectVersion.Version

	res.ProjectVersionUUID = req.UserProjectVersion.ProjectVersionUUID.String()

	res.UserUUID = req.UserProjectVersion.UserUUID.String()

	res.Data = []byte(mapper.StringPtrToString(req.UserProjectVersion.Data))

	if string(res.Data) == "" {
		res.Data = []byte(`{}`)
	}

	res.Status = req.UserProjectVersion.Status.ToInt64()

	res.CreatedAt = existing.CreatedAt

	res.UpdatedAt = custom.Now()

	res.CreatedByUUID = req.UserProjectVersion.CreatedByUUID.String()

	res.UpdatedByUUID = req.UserProjectVersion.UpdatedByUUID.String()

	return res

}

func (m *module) publishUpdateEvent(ctx context.Context,
	req types.UpsertRequest,
	qtx *nemdb.Queries,
	id string,
	existing []nemdb.UserProjectVersion) error {

	if m.events == nil {
		return nil
	}
	fetched, err := qtx.FetchUserProjectVersionByUUID(ctx, id)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_user_project_version",
			Message:          "error fetching after UpsertUserProjectVersion - with uuid",
			EntityIdentifier: "user_project_version",
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
			ActionIdentifier: "upsert_user_project_version",
			Message:          "error producing insert event for UpsertUserProjectVersion - with uuid",
			EntityIdentifier: "user_project_version",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return err
	}
	return nil
}
