package extension

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/nuzur/nem/core/module/extension/types"
)

func (m *module) Upsert(
	ctx context.Context,
	req types.UpsertRequest,
	partial bool,
	opts ...Option,
) (types.UpsertResponse, error) {
	if req.Extension.UUID == uuid.Nil {
		return m.Insert(ctx, req, opts...)
	}

	return m.Update(ctx, req, partial, opts...)
}
