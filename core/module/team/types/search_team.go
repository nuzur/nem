package types

import (
	main_entity "github.com/nuzur/nem/core/entity/team"
)

type SearchTeamResponse struct {
	Results []main_entity.Team
}
