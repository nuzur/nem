package user_project_version

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/nuzur/nem/core/module/user_project_version/types"
)

func (m *module) Upsert(
	ctx context.Context,
	req types.UpsertRequest,
	partial bool,
	opts ...Option,
) (types.UpsertResponse, error) {
	if req.UserProjectVersion.UUID == uuid.Nil {
		return m.Insert(ctx, req, opts...)
	}

	return m.Update(ctx, req, partial, opts...)
}
