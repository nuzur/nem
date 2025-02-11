package types

import (
	main_entity "github.com/nuzur/nem/core/entity/team"

	"go.uber.org/zap/zapcore"
)

type FetchTeamByStatusRequest struct {
	Status main_entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchTeamByStatusRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchTeamByStatusResponse struct {
	Results []main_entity.Team
}
