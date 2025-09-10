package types

import (
	main_entity "github.com/nuzur/nem/core/entity/membership"

	"go.uber.org/zap/zapcore"
)

type FetchMembershipByTypeAndStatusRequest struct {
	Type   main_entity.Type
	Status main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchMembershipByTypeAndStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchMembershipByTypeAndStatusResponse struct {
	Results []main_entity.Membership
}
