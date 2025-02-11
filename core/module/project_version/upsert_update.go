package project_version

import (
	"context"
	"errors"

	"nem/core/module/project_version/types"
	nemdb "nem/core/repository/gen"
	"nem/monitoring"

	"nem/core/entity/entity"
	"nem/core/entity/enum"
	"nem/core/entity/project_version_deployment"
	"nem/core/entity/relationship"
	"nem/core/entity/service"

	"nem/custom"

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

			BaseVersionUUID: req.ProjectVersion.BaseVersionUUID.String(),

			ReviewStatus: req.ProjectVersion.ReviewStatus.ToInt64(),

			Deployments: project_version_deployment.ProjectVersionDeploymentSliceToJSON(req.ProjectVersion.Deployments),

			Status: req.ProjectVersion.Status.ToInt64(),

			CreatedAt: existing.CreatedAt,

			UpdatedAt: custom.Now(),

			CreatedByUUID: req.ProjectVersion.CreatedByUUID.String(),

			UpdatedByUUID: req.ProjectVersion.UpdatedByUUID.String(),
		}
	}

	res := nemdb.UpdateProjectVersionParams{}
	emptyReq := types.UpsertRequest{}

	// regular field
	if req.ProjectVersion.UUID == emptyReq.ProjectVersion.UUID {

		res.UUID = existing.UUID
	} else {

		res.UUID = req.ProjectVersion.UUID.String()

	}

	// regular field
	if req.ProjectVersion.Version == emptyReq.ProjectVersion.Version {

		res.Version = existing.Version
	} else {

		res.Version = req.ProjectVersion.Version

	}

	// regular field
	if req.ProjectVersion.Identifier == emptyReq.ProjectVersion.Identifier {

		res.Identifier = existing.Identifier
	} else {

		res.Identifier = req.ProjectVersion.Identifier

	}

	// regular field
	if req.ProjectVersion.Description == emptyReq.ProjectVersion.Description {

		res.Description = existing.Description
	} else {

		res.Description = req.ProjectVersion.Description

	}

	// regular field
	if req.ProjectVersion.ProjectUUID == emptyReq.ProjectVersion.ProjectUUID {

		res.ProjectUUID = existing.ProjectUUID
	} else {

		res.ProjectUUID = req.ProjectVersion.ProjectUUID.String()

	}

	// json array
	if len(req.ProjectVersion.Entities) == 0 {

		res.Entities = existing.Entities
	} else {

		res.Entities = entity.EntitySliceToJSON(req.ProjectVersion.Entities)

	}

	// json array
	if len(req.ProjectVersion.Relationships) == 0 {

		res.Relationships = existing.Relationships
	} else {

		res.Relationships = relationship.RelationshipSliceToJSON(req.ProjectVersion.Relationships)

	}

	// json array
	if len(req.ProjectVersion.Enums) == 0 {

		res.Enums = existing.Enums
	} else {

		res.Enums = enum.EnumSliceToJSON(req.ProjectVersion.Enums)

	}

	// json array
	if len(req.ProjectVersion.Services) == 0 {

		res.Services = existing.Services
	} else {

		res.Services = service.ServiceSliceToJSON(req.ProjectVersion.Services)

	}

	// regular field
	if req.ProjectVersion.BaseVersionUUID == emptyReq.ProjectVersion.BaseVersionUUID {

		res.BaseVersionUUID = existing.BaseVersionUUID
	} else {

		res.BaseVersionUUID = req.ProjectVersion.BaseVersionUUID.String()

	}

	// regular field
	if req.ProjectVersion.ReviewStatus == emptyReq.ProjectVersion.ReviewStatus {

		res.ReviewStatus = existing.ReviewStatus
	} else {

		res.ReviewStatus = req.ProjectVersion.ReviewStatus.ToInt64()

	}

	// json array
	if len(req.ProjectVersion.Deployments) == 0 {

		res.Deployments = existing.Deployments
	} else {

		res.Deployments = project_version_deployment.ProjectVersionDeploymentSliceToJSON(req.ProjectVersion.Deployments)

	}

	// regular field
	if req.ProjectVersion.Status == emptyReq.ProjectVersion.Status {

		res.Status = existing.Status
	} else {

		res.Status = req.ProjectVersion.Status.ToInt64()

	}

	// regular field
	if req.ProjectVersion.CreatedAt == emptyReq.ProjectVersion.CreatedAt {

		res.CreatedAt = existing.CreatedAt
	} else {

		res.CreatedAt = existing.CreatedAt

	}

	// regular field
	if req.ProjectVersion.UpdatedAt == emptyReq.ProjectVersion.UpdatedAt {

		res.UpdatedAt = existing.UpdatedAt
	} else {

		res.UpdatedAt = custom.Now()

	}

	// regular field
	if req.ProjectVersion.CreatedByUUID == emptyReq.ProjectVersion.CreatedByUUID {

		res.CreatedByUUID = existing.CreatedByUUID
	} else {

		res.CreatedByUUID = req.ProjectVersion.CreatedByUUID.String()

	}

	// regular field
	if req.ProjectVersion.UpdatedByUUID == emptyReq.ProjectVersion.UpdatedByUUID {

		res.UpdatedByUUID = existing.UpdatedByUUID
	} else {

		res.UpdatedByUUID = req.ProjectVersion.UpdatedByUUID.String()

	}

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
