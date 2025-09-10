package types

import (
	main_entity "github.com/nuzur/nem/core/entity/membership"

	"github.com/gofrs/uuid"
	"go.uber.org/zap/zapcore"
)

type FetchMembershipByUUIDRequest struct {
	UUID uuid.UUID
}

func (r FetchMembershipByUUIDRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	e.AddString("uuid", r.UUID.String())

	return nil
}

type FetchMembershipByUUIDResponse struct {
	Results []main_entity.Membership
}
