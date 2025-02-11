package types

import (
	main_entity "github.com/nuzur/nem/core/entity/project"
)

type SearchProjectResponse struct {
	Results []main_entity.Project
}
