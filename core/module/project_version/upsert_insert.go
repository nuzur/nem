package project_version

import (
	"context"

	"errors"

	"github.com/gofrs/uuid"
	"nem/core/module/project_version/types"
	nemdb "nem/core/repository/gen"
	"nem/monitoring"

	"nem/custom"

	"nem/core/entity/entity"
	"nem/core/entity/enum"
	"nem/core/entity/project_version_deployment"
	"nem/core/entity/relationship"
	"nem/core/entity/service"

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

	_, err := qtx.InsertProjectVersion(
		ctx,
		params,
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_project_version",
			Message:          "error calling repository for UpsertProjectVersion - no uuid",
			EntityIdentifier: "project_version",
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
				ActionIdentifier: "upsert_project_version",
				Message:          "error commiting for UpsertProjectVersion (1)",
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
		Message:          "successfully handled UpsertProjectVersion - no uuid",
		EntityIdentifier: "project_version",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return types.UpsertResponse{
		UUID: uuid.FromStringOrNil(params.UUID),
	}, nil
}

func mapUpsertRequestToInsertParams(req types.UpsertRequest) nemdb.InsertProjectVersionParams {
	return nemdb.InsertProjectVersionParams{

		UUID: custom.GenerateUUID().String(),

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

		CreatedAt: custom.Now(),

		UpdatedAt: custom.Now(),

		CreatedByUUID: req.ProjectVersion.CreatedByUUID.String(),

		UpdatedByUUID: req.ProjectVersion.UpdatedByUUID.String(),
	}
}

func (m *module) publishInsertEvent(ctx context.Context,
	req types.UpsertRequest,
	qtx *nemdb.Queries,
	id string) error {

	if m.events == nil {
		return nil
	}

	fetched, err := qtx.FetchProjectVersionByUUID(ctx, id)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_project_version",
			Message:          "error fetching after UpsertProjectVersion - no uuid",
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
	err = m.events.ProduceEntityInserted(fetchedEntities[0])
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_project_version",
			Message:          "error producing insert event for UpsertProjectVersion - no uuid",
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
