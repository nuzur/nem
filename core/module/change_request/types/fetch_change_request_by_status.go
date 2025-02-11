package types

import (
	main_entity "nem/core/entity/change_request"

	"go.uber.org/zap/zapcore"
)

type FetchChangeRequestByStatusRequest struct {
	Status main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchChangeRequestByStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchChangeRequestByStatusResponse struct {
	Results []main_entity.ChangeRequest
}
