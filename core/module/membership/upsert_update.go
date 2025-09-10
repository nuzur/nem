package membership

import (
	"context"
	"errors"

	"github.com/nuzur/nem/core/module/membership/types"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"

	"github.com/nuzur/nem/custom"
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
	existing, err := qtx.FetchMembershipByUUIDForUpdate(ctx, req.Membership.UUID.String())
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_membership",
			Message:          "error fetching existing record for UpsertMembership - with uuid",
			EntityIdentifier: "membership",
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
			ActionIdentifier: "upsert_membership",
			Message:          "not found existing record for UpsertMembership - with uuid",
			EntityIdentifier: "membership",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	params := mapUpsertRequestToUpdateParams(req, existing[0], partial)
	err = qtx.UpdateMembership(
		ctx,
		params,
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_membership",
			Message:          "error calling repository for UpsertMembership - with uuid",
			EntityIdentifier: "membership",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	err = m.publishUpdateEvent(ctx, req, qtx, req.Membership.UUID.String(), existing)
	if err != nil {
		return types.UpsertResponse{}, err
	}

	if createdTx {
		err := tx.Commit()
		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "upsert_membership",
				Message:          "error commiting for UpsertMembership",
				EntityIdentifier: "membership",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.UpsertResponse{}, err
		}
	}

	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "upsert_membership",
		Message:          "successfully handled UpsertMembership - with uuid",
		EntityIdentifier: "membership",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return types.UpsertResponse{
		UUID: req.Membership.UUID,
	}, nil
}

func mapUpsertRequestToUpdateParams(req types.UpsertRequest, existing nemdb.Membership, partial bool) nemdb.UpdateMembershipParams {
	if !partial {
		return nemdb.UpdateMembershipParams{

			UUID: req.Membership.UUID.String(),

			OwnerUUID: req.Membership.OwnerUUID.String(),

			Type: req.Membership.Type.ToInt64(),

			StartDate: req.Membership.StartDate,

			BillingMetadata: []byte(req.Membership.BillingMetadata),

			Status: req.Membership.Status.ToInt64(),

			CreatedAt: existing.CreatedAt,

			UpdatedAt: custom.Now(),

			CreatedByUUID: req.Membership.CreatedByUUID.String(),

			UpdatedByUUID: req.Membership.UpdatedByUUID.String(),
		}
	}

	res := nemdb.UpdateMembershipParams{}

	res.UUID = req.Membership.UUID.String()

	res.OwnerUUID = req.Membership.OwnerUUID.String()

	res.Type = req.Membership.Type.ToInt64()

	res.StartDate = req.Membership.StartDate

	res.BillingMetadata = []byte(req.Membership.BillingMetadata)

	if string(res.BillingMetadata) == "" {
		res.BillingMetadata = []byte(`{}`)
	}

	res.Status = req.Membership.Status.ToInt64()

	res.CreatedAt = existing.CreatedAt

	res.UpdatedAt = custom.Now()

	res.CreatedByUUID = req.Membership.CreatedByUUID.String()

	res.UpdatedByUUID = req.Membership.UpdatedByUUID.String()

	return res

}

func (m *module) publishUpdateEvent(ctx context.Context,
	req types.UpsertRequest,
	qtx *nemdb.Queries,
	id string,
	existing []nemdb.Membership) error {

	if m.events == nil {
		return nil
	}
	fetched, err := qtx.FetchMembershipByUUID(ctx, id)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_membership",
			Message:          "error fetching after UpsertMembership - with uuid",
			EntityIdentifier: "membership",
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
			ActionIdentifier: "upsert_membership",
			Message:          "error producing insert event for UpsertMembership - with uuid",
			EntityIdentifier: "membership",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return err
	}
	return nil
}
