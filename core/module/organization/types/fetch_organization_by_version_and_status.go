package types

import (
	main_entity "github.com/nuzur/nem/core/entity/organization"

	"go.uber.org/zap/zapcore"
)

type FetchOrganizationByVersionAndStatusRequest struct {
	Version int64
	Status  main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchOrganizationByVersionAndStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchOrganizationByVersionAndStatusResponse struct {
	Results []main_entity.Organization
}
