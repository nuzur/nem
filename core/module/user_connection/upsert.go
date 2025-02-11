package user_connection

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/nuzur/nem/core/module/user_connection/types"
)

func (m *module) Upsert(
	ctx context.Context,
	req types.UpsertRequest,
	partial bool,
	opts ...Option,
) (types.UpsertResponse, error) {
	if req.UserConnection.UUID == uuid.Nil {
		return m.Insert(ctx, req, opts...)
	}

	return m.Update(ctx, req, partial, opts...)
}
