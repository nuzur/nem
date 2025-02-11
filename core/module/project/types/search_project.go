package types

import (
	main_entity "nem/core/entity/project"
)

type SearchProjectResponse struct {
	Results []main_entity.Project
}
