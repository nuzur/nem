package types

import (
	main_entity "nem/core/entity/organization"

	"go.uber.org/zap/zapcore"
)

type FetchOrganizationByVersionRequest struct {
	Version int64

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchOrganizationByVersionRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchOrganizationByVersionResponse struct {
	Results []main_entity.Organization
}
