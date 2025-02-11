package types

import (
	main_entity "nem/core/entity/team"
)

type SearchTeamResponse struct {
	Results []main_entity.Team
}
