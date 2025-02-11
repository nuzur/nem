package types

import (
	main_entity "nem/core/entity/organization"

	"go.uber.org/zap/zapcore"
)

type FetchOrganizationByStatusRequest struct {
	Status main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchOrganizationByStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchOrganizationByStatusResponse struct {
	Results []main_entity.Organization
}
