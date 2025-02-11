package types

import (
	main_entity "github.com/nuzur/nem/core/entity/team"

	"go.uber.org/zap/zapcore"
)

type FetchTeamByVersionRequest struct {
	Version int64

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

func (r FetchTeamByVersionRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchTeamByVersionResponse struct {
	Results []main_entity.Team
}
