package types

import (
	main_entity "github.com/nuzur/nem/core/entity/membership"

	"go.uber.org/zap/zapcore"
)

type FetchMembershipByStatusRequest struct {
	Status main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchMembershipByStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchMembershipByStatusResponse struct {
	Results []main_entity.Membership
}
