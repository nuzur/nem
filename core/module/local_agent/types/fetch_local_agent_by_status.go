package types

import (
	main_entity "github.com/nuzur/nem/core/entity/local_agent"

	"go.uber.org/zap/zapcore"
)

type FetchLocalAgentByStatusRequest struct {
	Status main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchLocalAgentByStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchLocalAgentByStatusResponse struct {
	Results []main_entity.LocalAgent
}
