package project_version

import (
	"context"
	"errors"

	"github.com/nuzur/nem/core/module/project_version/types"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"

	"github.com/nuzur/nem/core/entity/entity"
	"github.com/nuzur/nem/core/entity/enum"
	"github.com/nuzur/nem/core/entity/project_version_deployment"
	"github.com/nuzur/nem/core/entity/project_version_review"
	"github.com/nuzur/nem/core/entity/relationship"
	"github.com/nuzur/nem/core/entity/service"

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
	existing, err := qtx.FetchProjectVersionByUUIDForUpdate(ctx, req.ProjectVersion.UUID.String())
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_project_version",
			Message:          "error fetching existing record for UpsertProjectVersion - with uuid",
			EntityIdentifier: "project_version",
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
			ActionIdentifier: "upsert_project_version",
			Message:          "not found existing record for UpsertProjectVersion - with uuid",
			EntityIdentifier: "project_version",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	if existing[0].Version != req.ProjectVersion.Version {
		err := errors.New("upsert version conflict")
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_project_version",
			Message:          "version conflict UpsertProjectVersion",
			EntityIdentifier: "project_version",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	req.ProjectVersion.Version = time.Now().Unix()

	params := mapUpsertRequestToUpdateParams(req, existing[0], partial)
	err = qtx.UpdateProjectVersion(
		ctx,
		params,
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_project_version",
			Message:          "error calling repository for UpsertProjectVersion - with uuid",
			EntityIdentifier: "project_version",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	err = m.publishUpdateEvent(ctx, req, qtx, req.ProjectVersion.UUID.String(), existing)
	if err != nil {
		return types.UpsertResponse{}, err
	}

	if createdTx {
		err := tx.Commit()
		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "upsert_project_version",
				Message:          "error commiting for UpsertProjectVersion",
				EntityIdentifier: "project_version",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.UpsertResponse{}, err
		}
	}

	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "upsert_project_version",
		Message:          "successfully handled UpsertProjectVersion - with uuid",
		EntityIdentifier: "project_version",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return types.UpsertResponse{
		UUID: req.ProjectVersion.UUID,
	}, nil
}

func mapUpsertRequestToUpdateParams(req types.UpsertRequest, existing nemdb.ProjectVersion, partial bool) nemdb.UpdateProjectVersionParams {
	if !partial {
		return nemdb.UpdateProjectVersionParams{

			UUID: req.ProjectVersion.UUID.String(),

			Version: req.ProjectVersion.Version,

			Identifier: req.ProjectVersion.Identifier,

			Description: req.ProjectVersion.Description,

			ProjectUUID: req.ProjectVersion.ProjectUUID.String(),

			Entities: entity.EntitySliceToJSON(req.ProjectVersion.Entities),

			Relationships: relationship.RelationshipSliceToJSON(req.ProjectVersion.Relationships),

			Enums: enum.EnumSliceToJSON(req.ProjectVersion.Enums),

			Services: service.ServiceSliceToJSON(req.ProjectVersion.Services),

			BaseVersionUUID: mapper.StringToSqlNullString(req.ProjectVersion.BaseVersionUUID.String()),

			ReviewStatus: req.ProjectVersion.ReviewStatus.ToInt64(),

			Reviews: project_version_review.ProjectVersionReviewSliceToJSON(req.ProjectVersion.Reviews),

			Deployments: project_version_deployment.ProjectVersionDeploymentSliceToJSON(req.ProjectVersion.Deployments),

			Status: req.ProjectVersion.Status.ToInt64(),

			CreatedAt: existing.CreatedAt,

			UpdatedAt: custom.Now(),

			CreatedByUUID: req.ProjectVersion.CreatedByUUID.String(),

			UpdatedByUUID: req.ProjectVersion.UpdatedByUUID.String(),
		}
	}

	res := nemdb.UpdateProjectVersionParams{}

	res.UUID = req.ProjectVersion.UUID.String()

	res.Version = req.ProjectVersion.Version

	res.Identifier = req.ProjectVersion.Identifier

	res.Description = req.ProjectVersion.Description

	res.ProjectUUID = req.ProjectVersion.ProjectUUID.String()

	res.Entities = entity.EntitySliceToJSON(req.ProjectVersion.Entities)

	res.Relationships = relationship.RelationshipSliceToJSON(req.ProjectVersion.Relationships)

	res.Enums = enum.EnumSliceToJSON(req.ProjectVersion.Enums)

	res.Services = service.ServiceSliceToJSON(req.ProjectVersion.Services)

	res.BaseVersionUUID = mapper.StringToSqlNullString(req.ProjectVersion.BaseVersionUUID.String())

	res.ReviewStatus = req.ProjectVersion.ReviewStatus.ToInt64()

	res.Reviews = project_version_review.ProjectVersionReviewSliceToJSON(req.ProjectVersion.Reviews)

	res.Deployments = project_version_deployment.ProjectVersionDeploymentSliceToJSON(req.ProjectVersion.Deployments)

	res.Status = req.ProjectVersion.Status.ToInt64()

	res.CreatedAt = existing.CreatedAt

	res.UpdatedAt = custom.Now()

	res.CreatedByUUID = req.ProjectVersion.CreatedByUUID.String()

	res.UpdatedByUUID = req.ProjectVersion.UpdatedByUUID.String()

	return res

}

func (m *module) publishUpdateEvent(ctx context.Context,
	req types.UpsertRequest,
	qtx *nemdb.Queries,
	id string,
	existing []nemdb.ProjectVersion) error {

	if m.events == nil {
		return nil
	}
	fetched, err := qtx.FetchProjectVersionByUUID(ctx, id)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_project_version",
			Message:          "error fetching after UpsertProjectVersion - with uuid",
			EntityIdentifier: "project_version",
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
			ActionIdentifier: "upsert_project_version",
			Message:          "error producing insert event for UpsertProjectVersion - with uuid",
			EntityIdentifier: "project_version",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return err
	}
	return nil
}
