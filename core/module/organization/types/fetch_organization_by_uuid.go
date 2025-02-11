package types

import (
	main_entity "nem/core/entity/organization"

	"github.com/gofrs/uuid"
	"go.uber.org/zap/zapcore"
)

type FetchOrganizationByUUIDRequest struct {
	UUID uuid.UUID
}

func (r FetchOrganizationByUUIDRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	e.AddString("uuid", r.UUID.String())

	return nil
}

type FetchOrganizationByUUIDResponse struct {
	Results []main_entity.Organization
}
