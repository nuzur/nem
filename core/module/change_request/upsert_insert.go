package change_request

import (
	"context"

	"errors"

	"github.com/gofrs/uuid"
	"github.com/nuzur/nem/core/module/change_request/types"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"

	"github.com/nuzur/nem/custom"

	"github.com/nuzur/nem/core/entity/change_request_review"

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

	_, err := qtx.InsertChangeRequest(
		ctx,
		params,
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_change_request",
			Message:          "error calling repository for UpsertChangeRequest - no uuid",
			EntityIdentifier: "change_request",
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
				ActionIdentifier: "upsert_change_request",
				Message:          "error commiting for UpsertChangeRequest (1)",
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
		Message:          "successfully handled UpsertChangeRequest - no uuid",
		EntityIdentifier: "change_request",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return types.UpsertResponse{
		UUID: uuid.FromStringOrNil(params.UUID),
	}, nil
}

func mapUpsertRequestToInsertParams(req types.UpsertRequest) nemdb.InsertChangeRequestParams {
	return nemdb.InsertChangeRequestParams{

		UUID: custom.GenerateUUID().String(),

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

		CreatedAt: custom.Now(),

		UpdatedAt: custom.Now(),

		CreatedByUUID: req.ChangeRequest.CreatedByUUID.String(),

		UpdatedByUUID: req.ChangeRequest.UpdatedByUUID.String(),
	}
}

func (m *module) publishInsertEvent(ctx context.Context,
	req types.UpsertRequest,
	qtx *nemdb.Queries,
	id string) error {

	if m.events == nil {
		return nil
	}

	fetched, err := qtx.FetchChangeRequestByUUID(ctx, id)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_change_request",
			Message:          "error fetching after UpsertChangeRequest - no uuid",
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
	err = m.events.ProduceEntityInserted(fetchedEntities[0])
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_change_request",
			Message:          "error producing insert event for UpsertChangeRequest - no uuid",
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
