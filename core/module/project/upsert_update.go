package project

import (
	"context"
	"errors"

	"github.com/nuzur/nem/core/module/project/types"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"

	"github.com/nuzur/nem/core/entity/project_extension"

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
	existing, err := qtx.FetchProjectByUUIDForUpdate(ctx, req.Project.UUID.String())
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_project",
			Message:          "error fetching existing record for UpsertProject - with uuid",
			EntityIdentifier: "project",
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
			ActionIdentifier: "upsert_project",
			Message:          "not found existing record for UpsertProject - with uuid",
			EntityIdentifier: "project",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	if existing[0].Version != req.Project.Version {
		err := errors.New("upsert version conflict")
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_project",
			Message:          "version conflict UpsertProject",
			EntityIdentifier: "project",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	req.Project.Version = time.Now().Unix()

	params := mapUpsertRequestToUpdateParams(req, existing[0], partial)
	err = qtx.UpdateProject(
		ctx,
		params,
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_project",
			Message:          "error calling repository for UpsertProject - with uuid",
			EntityIdentifier: "project",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	err = m.publishUpdateEvent(ctx, req, qtx, req.Project.UUID.String(), existing)
	if err != nil {
		return types.UpsertResponse{}, err
	}

	if createdTx {
		err := tx.Commit()
		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "upsert_project",
				Message:          "error commiting for UpsertProject",
				EntityIdentifier: "project",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.UpsertResponse{}, err
		}
	}

	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "upsert_project",
		Message:          "successfully handled UpsertProject - with uuid",
		EntityIdentifier: "project",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return types.UpsertResponse{
		UUID: req.Project.UUID,
	}, nil
}

func mapUpsertRequestToUpdateParams(req types.UpsertRequest, existing nemdb.Project, partial bool) nemdb.UpdateProjectParams {
	if !partial {
		return nemdb.UpdateProjectParams{

			UUID: req.Project.UUID.String(),

			Version: req.Project.Version,

			Name: req.Project.Name,

			Description: mapper.StringPtrToSqlNullString(req.Project.Description),

			Tags: mapper.SliceToJSON(req.Project.Tags),

			URL: mapper.StringPtrToSqlNullString(req.Project.URL),

			OwnerUUID: req.Project.OwnerUUID.String(),

			TeamUUID: req.Project.TeamUUID.String(),

			ProjectExtensions: project_extension.ProjectExtensionSliceToJSON(req.Project.ProjectExtensions),

			Status: req.Project.Status.ToInt64(),

			CreatedAt: existing.CreatedAt,

			UpdatedAt: custom.Now(),

			CreatedByUUID: req.Project.CreatedByUUID.String(),

			UpdatedByUUID: req.Project.UpdatedByUUID.String(),
		}
	}

	res := nemdb.UpdateProjectParams{}

	res.UUID = req.Project.UUID.String()

	res.Version = req.Project.Version

	res.Name = req.Project.Name

	res.Description = mapper.StringPtrToSqlNullString(req.Project.Description)

	res.Tags = mapper.SliceToJSON(req.Project.Tags)

	res.URL = mapper.StringPtrToSqlNullString(req.Project.URL)

	res.OwnerUUID = req.Project.OwnerUUID.String()

	res.TeamUUID = req.Project.TeamUUID.String()

	res.ProjectExtensions = project_extension.ProjectExtensionSliceToJSON(req.Project.ProjectExtensions)

	res.Status = req.Project.Status.ToInt64()

	res.CreatedAt = existing.CreatedAt

	res.UpdatedAt = custom.Now()

	res.CreatedByUUID = req.Project.CreatedByUUID.String()

	res.UpdatedByUUID = req.Project.UpdatedByUUID.String()

	return res

}

func (m *module) publishUpdateEvent(ctx context.Context,
	req types.UpsertRequest,
	qtx *nemdb.Queries,
	id string,
	existing []nemdb.Project) error {

	if m.events == nil {
		return nil
	}
	fetched, err := qtx.FetchProjectByUUID(ctx, id)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_project",
			Message:          "error fetching after UpsertProject - with uuid",
			EntityIdentifier: "project",
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
			ActionIdentifier: "upsert_project",
			Message:          "error producing insert event for UpsertProject - with uuid",
			EntityIdentifier: "project",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return err
	}
	return nil
}
