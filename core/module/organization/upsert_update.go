package organization

import (
	"context"
	"errors"

	"github.com/nuzur/nem/core/module/organization/types"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"

	"github.com/nuzur/nem/core/entity/membership"

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
	existing, err := qtx.FetchOrganizationByUUIDForUpdate(ctx, req.Organization.UUID.String())
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_organization",
			Message:          "error fetching existing record for UpsertOrganization - with uuid",
			EntityIdentifier: "organization",
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
			ActionIdentifier: "upsert_organization",
			Message:          "not found existing record for UpsertOrganization - with uuid",
			EntityIdentifier: "organization",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	if existing[0].Version != req.Organization.Version {
		err := errors.New("upsert version conflict")
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_organization",
			Message:          "version conflict UpsertOrganization",
			EntityIdentifier: "organization",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	req.Organization.Version = time.Now().Unix()

	params := mapUpsertRequestToUpdateParams(req, existing[0], partial)
	err = qtx.UpdateOrganization(
		ctx,
		params,
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_organization",
			Message:          "error calling repository for UpsertOrganization - with uuid",
			EntityIdentifier: "organization",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	err = m.publishUpdateEvent(ctx, req, qtx, req.Organization.UUID.String(), existing)
	if err != nil {
		return types.UpsertResponse{}, err
	}

	if createdTx {
		err := tx.Commit()
		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "upsert_organization",
				Message:          "error commiting for UpsertOrganization",
				EntityIdentifier: "organization",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.UpsertResponse{}, err
		}
	}

	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "upsert_organization",
		Message:          "successfully handled UpsertOrganization - with uuid",
		EntityIdentifier: "organization",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return types.UpsertResponse{
		UUID: req.Organization.UUID,
	}, nil
}

func mapUpsertRequestToUpdateParams(req types.UpsertRequest, existing nemdb.Organization, partial bool) nemdb.UpdateOrganizationParams {
	if !partial {
		return nemdb.UpdateOrganizationParams{

			UUID: req.Organization.UUID.String(),

			Version: req.Organization.Version,

			Name: req.Organization.Name,

			Domains: mapper.SliceToJSON(req.Organization.Domains),

			AdminUUIDs: mapper.SliceToJSON(req.Organization.AdminUUIDs),

			Memberships: membership.MembershipSliceToJSON(req.Organization.Memberships),

			Status: req.Organization.Status.ToInt64(),

			CreatedAt: existing.CreatedAt,

			UpdatedAt: custom.Now(),

			CreatedByUUID: req.Organization.CreatedByUUID.String(),

			UpdatedByUUID: req.Organization.UpdatedByUUID.String(),
		}
	}

	res := nemdb.UpdateOrganizationParams{}
	emptyReq := types.UpsertRequest{}

	// regular field
	if req.Organization.UUID == emptyReq.Organization.UUID {

		res.UUID = existing.UUID
	} else {

		res.UUID = req.Organization.UUID.String()

	}

	// regular field
	if req.Organization.Version == emptyReq.Organization.Version {

		res.Version = existing.Version
	} else {

		res.Version = req.Organization.Version

	}

	// regular field
	if req.Organization.Name == emptyReq.Organization.Name {

		res.Name = existing.Name
	} else {

		res.Name = req.Organization.Name

	}

	// raw json is a pointer
	if req.Organization.Domains == nil {

		res.Domains = existing.Domains
	} else {

		res.Domains = mapper.SliceToJSON(req.Organization.Domains)

	}

	// raw json is a pointer
	if req.Organization.AdminUUIDs == nil {

		res.AdminUUIDs = existing.AdminUUIDs
	} else {

		res.AdminUUIDs = mapper.SliceToJSON(req.Organization.AdminUUIDs)

	}

	// json array
	if len(req.Organization.Memberships) == 0 {

		res.Memberships = existing.Memberships
	} else {

		res.Memberships = membership.MembershipSliceToJSON(req.Organization.Memberships)

	}

	// regular field
	if req.Organization.Status == emptyReq.Organization.Status {

		res.Status = existing.Status
	} else {

		res.Status = req.Organization.Status.ToInt64()

	}

	// regular field
	if req.Organization.CreatedAt == emptyReq.Organization.CreatedAt {

		res.CreatedAt = existing.CreatedAt
	} else {

		res.CreatedAt = existing.CreatedAt

	}

	// regular field
	if req.Organization.UpdatedAt == emptyReq.Organization.UpdatedAt {

		res.UpdatedAt = existing.UpdatedAt
	} else {

		res.UpdatedAt = custom.Now()

	}

	// regular field
	if req.Organization.CreatedByUUID == emptyReq.Organization.CreatedByUUID {

		res.CreatedByUUID = existing.CreatedByUUID
	} else {

		res.CreatedByUUID = req.Organization.CreatedByUUID.String()

	}

	// regular field
	if req.Organization.UpdatedByUUID == emptyReq.Organization.UpdatedByUUID {

		res.UpdatedByUUID = existing.UpdatedByUUID
	} else {

		res.UpdatedByUUID = req.Organization.UpdatedByUUID.String()

	}

	return res

}

func (m *module) publishUpdateEvent(ctx context.Context,
	req types.UpsertRequest,
	qtx *nemdb.Queries,
	id string,
	existing []nemdb.Organization) error {

	if m.events == nil {
		return nil
	}
	fetched, err := qtx.FetchOrganizationByUUID(ctx, id)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_organization",
			Message:          "error fetching after UpsertOrganization - with uuid",
			EntityIdentifier: "organization",
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
			ActionIdentifier: "upsert_organization",
			Message:          "error producing insert event for UpsertOrganization - with uuid",
			EntityIdentifier: "organization",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return err
	}
	return nil
}
