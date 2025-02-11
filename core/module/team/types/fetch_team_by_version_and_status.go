package types

import (
	main_entity "github.com/nuzur/nem/core/entity/team"

	"go.uber.org/zap/zapcore"
)

type FetchTeamByVersionAndStatusRequest struct {
	Version int64
	Status  main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchTeamByVersionAndStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchTeamByVersionAndStatusResponse struct {
	Results []main_entity.Team
}
