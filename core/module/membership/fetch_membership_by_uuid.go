package membership

import (
	"context"

	"github.com/nuzur/nem/core/module/membership/types"

	"github.com/nuzur/nem/monitoring"
)

func (m *module) FetchMembershipByUUID(
	ctx context.Context,
	req types.FetchMembershipByUUIDRequest,
	opts ...Option,
) (types.FetchMembershipByUUIDResponse, error) {

	models, err := m.repository.Queries.FetchMembershipByUUID(
		ctx,
		req.UUID.String(),
	)
	if err != nil {
		m.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "fetch_membership_by_uuid",
			Message:          "error in FetchMembershipByUUID",
			EntityIdentifier: "membership",
			Layer:            monitoring.RepositoryServiceLayer,
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return types.FetchMembershipByUUIDResponse{}, err
	}
	return types.FetchMembershipByUUIDResponse{
		Results: mapModelsToEntities(models),
	}, nil

}
