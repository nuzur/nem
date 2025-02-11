package change_request

import (
	"context"
	"errors"

	"github.com/nuzur/nem/core/module/change_request/types"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"

	"github.com/nuzur/nem/core/entity/change_request_review"

	"github.com/nuzur/nem/custom"

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

			Description: req.ChangeRequest.Description,

			ReviewType: req.ChangeRequest.ReviewType.ToInt64(),

			RefUUID: req.ChangeRequest.RefUUID.String(),

			OldData: []byte(req.ChangeRequest.OldData),

			OldDataRef: req.ChangeRequest.OldDataRef,

			NewData: []byte(req.ChangeRequest.NewData),

			NewDataRef: req.ChangeRequest.NewDataRef,

			Reviews: change_request_review.ChangeRequestReviewSliceToJSON(req.ChangeRequest.Reviews),

			OwnerUUID: req.ChangeRequest.OwnerUUID.String(),

			Status: req.ChangeRequest.Status.ToInt64(),

			CreatedAt: existing.CreatedAt,

			UpdatedAt: custom.Now(),

			CreatedByUUID: req.ChangeRequest.CreatedByUUID.String(),

			UpdatedByUUID: req.ChangeRequest.UpdatedByUUID.String(),
		}
	}

	res := nemdb.UpdateChangeRequestParams{}
	emptyReq := types.UpsertRequest{}

	// regular field
	if req.ChangeRequest.UUID == emptyReq.ChangeRequest.UUID {

		res.UUID = existing.UUID
	} else {

		res.UUID = req.ChangeRequest.UUID.String()

	}

	// regular field
	if req.ChangeRequest.Version == emptyReq.ChangeRequest.Version {

		res.Version = existing.Version
	} else {

		res.Version = req.ChangeRequest.Version

	}

	// regular field
	if req.ChangeRequest.Title == emptyReq.ChangeRequest.Title {

		res.Title = existing.Title
	} else {

		res.Title = req.ChangeRequest.Title

	}

	// regular field
	if req.ChangeRequest.Description == emptyReq.ChangeRequest.Description {

		res.Description = existing.Description
	} else {

		res.Description = req.ChangeRequest.Description

	}

	// regular field
	if req.ChangeRequest.ReviewType == emptyReq.ChangeRequest.ReviewType {

		res.ReviewType = existing.ReviewType
	} else {

		res.ReviewType = req.ChangeRequest.ReviewType.ToInt64()

	}

	// regular field
	if req.ChangeRequest.RefUUID == emptyReq.ChangeRequest.RefUUID {

		res.RefUUID = existing.RefUUID
	} else {

		res.RefUUID = req.ChangeRequest.RefUUID.String()

	}

	// regular field
	if req.ChangeRequest.OldData == emptyReq.ChangeRequest.OldData {

		res.OldData = existing.OldData
	} else {

		res.OldData = []byte(req.ChangeRequest.OldData)

	}

	// regular field
	if req.ChangeRequest.OldDataRef == emptyReq.ChangeRequest.OldDataRef {

		res.OldDataRef = existing.OldDataRef
	} else {

		res.OldDataRef = req.ChangeRequest.OldDataRef

	}

	// regular field
	if req.ChangeRequest.NewData == emptyReq.ChangeRequest.NewData {

		res.NewData = existing.NewData
	} else {

		res.NewData = []byte(req.ChangeRequest.NewData)

	}

	// regular field
	if req.ChangeRequest.NewDataRef == emptyReq.ChangeRequest.NewDataRef {

		res.NewDataRef = existing.NewDataRef
	} else {

		res.NewDataRef = req.ChangeRequest.NewDataRef

	}

	// json array
	if len(req.ChangeRequest.Reviews) == 0 {

		res.Reviews = existing.Reviews
	} else {

		res.Reviews = change_request_review.ChangeRequestReviewSliceToJSON(req.ChangeRequest.Reviews)

	}

	// regular field
	if req.ChangeRequest.OwnerUUID == emptyReq.ChangeRequest.OwnerUUID {

		res.OwnerUUID = existing.OwnerUUID
	} else {

		res.OwnerUUID = req.ChangeRequest.OwnerUUID.String()

	}

	// regular field
	if req.ChangeRequest.Status == emptyReq.ChangeRequest.Status {

		res.Status = existing.Status
	} else {

		res.Status = req.ChangeRequest.Status.ToInt64()

	}

	// regular field
	if req.ChangeRequest.CreatedAt == emptyReq.ChangeRequest.CreatedAt {

		res.CreatedAt = existing.CreatedAt
	} else {

		res.CreatedAt = existing.CreatedAt

	}

	// regular field
	if req.ChangeRequest.UpdatedAt == emptyReq.ChangeRequest.UpdatedAt {

		res.UpdatedAt = existing.UpdatedAt
	} else {

		res.UpdatedAt = custom.Now()

	}

	// regular field
	if req.ChangeRequest.CreatedByUUID == emptyReq.ChangeRequest.CreatedByUUID {

		res.CreatedByUUID = existing.CreatedByUUID
	} else {

		res.CreatedByUUID = req.ChangeRequest.CreatedByUUID.String()

	}

	// regular field
	if req.ChangeRequest.UpdatedByUUID == emptyReq.ChangeRequest.UpdatedByUUID {

		res.UpdatedByUUID = existing.UpdatedByUUID
	} else {

		res.UpdatedByUUID = req.ChangeRequest.UpdatedByUUID.String()

	}

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
