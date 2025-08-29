package change_request

import (
	"context"
	"errors"

	"github.com/nuzur/nem/core/module/change_request/types"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"

	"github.com/nuzur/nem/core/entity/change_request_data_change"
	"github.com/nuzur/nem/core/entity/change_request_metadata"
	"github.com/nuzur/nem/core/entity/change_request_review"

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
	existing, err := qtx.FetchChangeRequestByUUIDForUpdate(ctx, req.ChangeRequest.UUID.String())
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_change_request",
			Message:          "error fetching existing record for UpsertChangeRequest - with uuid",
			EntityIdentifier: "change_request",
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
			ActionIdentifier: "upsert_change_request",
			Message:          "not found existing record for UpsertChangeRequest - with uuid",
			EntityIdentifier: "change_request",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	if existing[0].Version != req.ChangeRequest.Version {
		err := errors.New("upsert version conflict")
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_change_request",
			Message:          "version conflict UpsertChangeRequest",
			EntityIdentifier: "change_request",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	req.ChangeRequest.Version = time.Now().Unix()

	params := mapUpsertRequestToUpdateParams(req, existing[0], partial)
	err = qtx.UpdateChangeRequest(
		ctx,
		params,
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_change_request",
			Message:          "error calling repository for UpsertChangeRequest - with uuid",
			EntityIdentifier: "change_request",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	err = m.publishUpdateEvent(ctx, req, qtx, req.ChangeRequest.UUID.String(), existing)
	if err != nil {
		return types.UpsertResponse{}, err
	}

	if createdTx {
		err := tx.Commit()
		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "upsert_change_request",
				Message:          "error commiting for UpsertChangeRequest",
				EntityIdentifier: "change_request",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.UpsertResponse{}, err
		}
	}

	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "upsert_change_request",
		Message:          "successfully handled UpsertChangeRequest - with uuid",
		EntityIdentifier: "change_request",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return types.UpsertResponse{
		UUID: req.ChangeRequest.UUID,
	}, nil
}

func mapUpsertRequestToUpdateParams(req types.UpsertRequest, existing nemdb.ChangeRequest, partial bool) nemdb.UpdateChangeRequestParams {
	if !partial {
		return nemdb.UpdateChangeRequestParams{

			UUID: req.ChangeRequest.UUID.String(),

			Version: req.ChangeRequest.Version,

			Title: req.ChangeRequest.Title,

			Description: mapper.StringPtrToSqlNullString(req.ChangeRequest.Description),

			ProjectUUID: req.ChangeRequest.ProjectUUID.String(),

			ProjectVersionUUID: req.ChangeRequest.ProjectVersionUUID.String(),

			ChangeType: req.ChangeRequest.ChangeType.ToInt64(),

			DataChanges: change_request_data_change.ChangeRequestDataChangeSliceToJSON(req.ChangeRequest.DataChanges),

			Metadata: change_request_metadata.ChangeRequestMetadataToJSON(req.ChangeRequest.Metadata),

			Reviews: change_request_review.ChangeRequestReviewSliceToJSON(req.ChangeRequest.Reviews),

			ReviewStatus: req.ChangeRequest.ReviewStatus.ToInt64(),

			OwnerUUID: req.ChangeRequest.OwnerUUID.String(),

			Status: req.ChangeRequest.Status.ToInt64(),

			CreatedAt: existing.CreatedAt,

			UpdatedAt: custom.Now(),

			CreatedByUUID: req.ChangeRequest.CreatedByUUID.String(),

			UpdatedByUUID: req.ChangeRequest.UpdatedByUUID.String(),
		}
	}

	res := nemdb.UpdateChangeRequestParams{}

	res.UUID = req.ChangeRequest.UUID.String()

	res.Version = req.ChangeRequest.Version

	res.Title = req.ChangeRequest.Title

	res.Description = mapper.StringPtrToSqlNullString(req.ChangeRequest.Description)

	res.ProjectUUID = req.ChangeRequest.ProjectUUID.String()

	res.ProjectVersionUUID = req.ChangeRequest.ProjectVersionUUID.String()

	res.ChangeType = req.ChangeRequest.ChangeType.ToInt64()

	res.DataChanges = change_request_data_change.ChangeRequestDataChangeSliceToJSON(req.ChangeRequest.DataChanges)

	res.Metadata = change_request_metadata.ChangeRequestMetadataToJSON(req.ChangeRequest.Metadata)

	res.Reviews = change_request_review.ChangeRequestReviewSliceToJSON(req.ChangeRequest.Reviews)

	res.ReviewStatus = req.ChangeRequest.ReviewStatus.ToInt64()

	res.OwnerUUID = req.ChangeRequest.OwnerUUID.String()

	res.Status = req.ChangeRequest.Status.ToInt64()

	res.CreatedAt = existing.CreatedAt

	res.UpdatedAt = custom.Now()

	res.CreatedByUUID = req.ChangeRequest.CreatedByUUID.String()

	res.UpdatedByUUID = req.ChangeRequest.UpdatedByUUID.String()

	return res

}

func (m *module) publishUpdateEvent(ctx context.Context,
	req types.UpsertRequest,
	qtx *nemdb.Queries,
	id string,
	existing []nemdb.ChangeRequest) error {

	if m.events == nil {
		return nil
	}
	fetched, err := qtx.FetchChangeRequestByUUID(ctx, id)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_change_request",
			Message:          "error fetching after UpsertChangeRequest - with uuid",
			EntityIdentifier: "change_request",
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
			ActionIdentifier: "upsert_change_request",
			Message:          "error producing insert event for UpsertChangeRequest - with uuid",
			EntityIdentifier: "change_request",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return err
	}
	return nil
}
