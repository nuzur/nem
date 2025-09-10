package types

import (
	main_entity "github.com/nuzur/nem/core/entity/membership"

	"go.uber.org/zap/zapcore"
)

type FetchMembershipByTypeRequest struct {
	Type main_entity.Type

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchMembershipByTypeRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchMembershipByTypeResponse struct {
	Results []main_entity.Membership
}
