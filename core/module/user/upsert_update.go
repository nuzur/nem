package user

import (
	"context"
	"errors"

	"nem/core/module/user/types"
	nemdb "nem/core/repository/gen"
	"nem/monitoring"

	"nem/custom"
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
	existing, err := qtx.FetchUserByUUIDForUpdate(ctx, req.User.UUID.String())
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_user",
			Message:          "error fetching existing record for UpsertUser - with uuid",
			EntityIdentifier: "user",
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
			ActionIdentifier: "upsert_user",
			Message:          "not found existing record for UpsertUser - with uuid",
			EntityIdentifier: "user",
			Layer:            monitoring.ModuleServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	params := mapUpsertRequestToUpdateParams(req, existing[0], partial)
	err = qtx.UpdateUser(
		ctx,
		params,
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_user",
			Message:          "error calling repository for UpsertUser - with uuid",
			EntityIdentifier: "user",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	err = m.publishUpdateEvent(ctx, req, qtx, req.User.UUID.String(), existing)
	if err != nil {
		return types.UpsertResponse{}, err
	}

	if createdTx {
		err := tx.Commit()
		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "upsert_user",
				Message:          "error commiting for UpsertUser",
				EntityIdentifier: "user",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.UpsertResponse{}, err
		}
	}

	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "upsert_user",
		Message:          "successfully handled UpsertUser - with uuid",
		EntityIdentifier: "user",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return types.UpsertResponse{
		UUID: req.User.UUID,
	}, nil
}

func mapUpsertRequestToUpdateParams(req types.UpsertRequest, existing nemdb.User, partial bool) nemdb.UpdateUserParams {
	if !partial {
		return nemdb.UpdateUserParams{

			UUID: req.User.UUID.String(),

			Identifier: req.User.Identifier,

			Name: req.User.Name,

			LastName: req.User.LastName,

			Email: req.User.Email,

			UserType: req.User.UserType.ToInt64(),

			CountryIos2: req.User.CountryIos2,

			Locale: req.User.Locale,

			Metadata: req.User.Metadata,

			Status: req.User.Status.ToInt64(),

			CreatedAt: existing.CreatedAt,

			UpdatedAt: custom.Now(),
		}
	}

	res := nemdb.UpdateUserParams{}
	emptyReq := types.UpsertRequest{}

	// regular field
	if req.User.UUID == emptyReq.User.UUID {

		res.UUID = existing.UUID
	} else {

		res.UUID = req.User.UUID.String()

	}

	// regular field
	if req.User.Identifier == emptyReq.User.Identifier {

		res.Identifier = existing.Identifier
	} else {

		res.Identifier = req.User.Identifier

	}

	// regular field
	if req.User.Name == emptyReq.User.Name {

		res.Name = existing.Name
	} else {

		res.Name = req.User.Name

	}

	// regular field
	if req.User.LastName == emptyReq.User.LastName {

		res.LastName = existing.LastName
	} else {

		res.LastName = req.User.LastName

	}

	// regular field
	if req.User.Email == emptyReq.User.Email {

		res.Email = existing.Email
	} else {

		res.Email = req.User.Email

	}

	// regular field
	if req.User.UserType == emptyReq.User.UserType {

		res.UserType = existing.UserType
	} else {

		res.UserType = req.User.UserType.ToInt64()

	}

	// regular field
	if req.User.CountryIos2 == emptyReq.User.CountryIos2 {

		res.CountryIos2 = existing.CountryIos2
	} else {

		res.CountryIos2 = req.User.CountryIos2

	}

	// regular field
	if req.User.Locale == emptyReq.User.Locale {

		res.Locale = existing.Locale
	} else {

		res.Locale = req.User.Locale

	}

	// regular field
	if req.User.Metadata == emptyReq.User.Metadata {

		res.Metadata = existing.Metadata
	} else {

		res.Metadata = req.User.Metadata

	}

	// regular field
	if req.User.Status == emptyReq.User.Status {

		res.Status = existing.Status
	} else {

		res.Status = req.User.Status.ToInt64()

	}

	// regular field
	if req.User.CreatedAt == emptyReq.User.CreatedAt {

		res.CreatedAt = existing.CreatedAt
	} else {

		res.CreatedAt = existing.CreatedAt

	}

	// regular field
	if req.User.UpdatedAt == emptyReq.User.UpdatedAt {

		res.UpdatedAt = existing.UpdatedAt
	} else {

		res.UpdatedAt = custom.Now()

	}

	return res

}

func (m *module) publishUpdateEvent(ctx context.Context,
	req types.UpsertRequest,
	qtx *nemdb.Queries,
	id string,
	existing []nemdb.User) error {

	if m.events == nil {
		return nil
	}
	fetched, err := qtx.FetchUserByUUID(ctx, id)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_user",
			Message:          "error fetching after UpsertUser - with uuid",
			EntityIdentifier: "user",
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
			ActionIdentifier: "upsert_user",
			Message:          "error producing insert event for UpsertUser - with uuid",
			EntityIdentifier: "user",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return err
	}
	return nil
}
