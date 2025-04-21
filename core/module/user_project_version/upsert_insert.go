package user_project_version

import (
	"context"

	"errors"

	"github.com/gofrs/uuid"
	"github.com/nuzur/nem/core/module/user_project_version/types"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"

	"github.com/nuzur/nem/custom"

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

	_, err := qtx.InsertUserProjectVersion(
		ctx,
		params,
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_user_project_version",
			Message:          "error calling repository for UpsertUserProjectVersion - no uuid",
			EntityIdentifier: "user_project_version",
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
				ActionIdentifier: "upsert_user_project_version",
				Message:          "error commiting for UpsertUserProjectVersion (1)",
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
		Message:          "successfully handled UpsertUserProjectVersion - no uuid",
		EntityIdentifier: "user_project_version",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return types.UpsertResponse{
		UUID: uuid.FromStringOrNil(params.UUID),
	}, nil
}

func mapUpsertRequestToInsertParams(req types.UpsertRequest) nemdb.InsertUserProjectVersionParams {
	return nemdb.InsertUserProjectVersionParams{

		UUID: custom.GenerateUUID().String(),

		Version: req.UserProjectVersion.Version,

		ProjectVersionUUID: req.UserProjectVersion.ProjectVersionUUID.String(),

		UserUUID: req.UserProjectVersion.UserUUID.String(),

		Data: []byte(mapper.StringPtrToString(req.UserProjectVersion.Data)),

		Status: req.UserProjectVersion.Status.ToInt64(),

		CreatedAt: custom.Now(),

		UpdatedAt: custom.Now(),

		CreatedByUUID: req.UserProjectVersion.CreatedByUUID.String(),

		UpdatedByUUID: req.UserProjectVersion.UpdatedByUUID.String(),
	}
}

func (m *module) publishInsertEvent(ctx context.Context,
	req types.UpsertRequest,
	qtx *nemdb.Queries,
	id string) error {

	if m.events == nil {
		return nil
	}

	fetched, err := qtx.FetchUserProjectVersionByUUID(ctx, id)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_user_project_version",
			Message:          "error fetching after UpsertUserProjectVersion - no uuid",
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
	err = m.events.ProduceEntityInserted(fetchedEntities[0])
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_user_project_version",
			Message:          "error producing insert event for UpsertUserProjectVersion - no uuid",
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
