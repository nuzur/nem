package local_agent

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/nuzur/nem/core/module/local_agent/types"
	nemdb "github.com/nuzur/nem/core/repository/gen"
	"github.com/nuzur/nem/monitoring"

	"github.com/nuzur/nem/custom"

	"github.com/nuzur/nem/core/entity/local_agent_connection"

	"github.com/nuzur/nem/core/entity/mapper"
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

	_, err := qtx.InsertLocalAgent(
		ctx,
		params,
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "upsert_local_agent",
			Message:          "error calling repository for UpsertLocalAgent - no uuid",
			EntityIdentifier: "local_agent",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.UpsertResponse{}, err
	}

	if createdTx {
		err := tx.Commit()
		if err != nil {
			m.monitoring.Emit(monitoring.EmitRequest{
				ActionIdentifier: "upsert_local_agent",
				Message:          "error commiting for UpsertLocalAgent (1)",
				EntityIdentifier: "local_agent",
				Layer:            monitoring.RepositoryServiceLayer,
				Type:             monitoring.EmitTypeError,
				Data:             req,
				Error:            err,
			})
			return types.UpsertResponse{}, err
		}
	}

	m.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "upsert_local_agent",
		Message:          "successfully handled UpsertLocalAgent - no uuid",
		EntityIdentifier: "local_agent",
		Layer:            monitoring.ModuleServiceLayer,
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return types.UpsertResponse{
		UUID: uuid.FromStringOrNil(params.UUID),
	}, nil
}

func mapUpsertRequestToInsertParams(req types.UpsertRequest) nemdb.InsertLocalAgentParams {
	return nemdb.InsertLocalAgentParams{

		UUID: custom.GenerateUUID().String(),

		UserUUID: req.LocalAgent.UserUUID.String(),

		TokenHash: mapper.StringPtrToSqlNullString(req.LocalAgent.TokenHash),

		MachineName: mapper.StringPtrToSqlNullString(req.LocalAgent.MachineName),

		Os: mapper.StringPtrToSqlNullString(req.LocalAgent.Os),

		CliVersion: mapper.StringPtrToSqlNullString(req.LocalAgent.CliVersion),

		Connections: local_agent_connection.LocalAgentConnectionSliceToJSON(req.LocalAgent.Connections),

		Status: req.LocalAgent.Status.ToInt64(),

		LastSeenAt: mapper.TimePtrToSqlNullTime(req.LocalAgent.LastSeenAt),

		RevokedAt: mapper.TimePtrToSqlNullTime(req.LocalAgent.RevokedAt),

		CreatedAt: custom.Now(),

		UpdatedAt: custom.Now(),

		CreatedByUUID: req.LocalAgent.CreatedByUUID.String(),

		UpdatedByUUID: req.LocalAgent.UpdatedByUUID.String(),
	}
}
