package types

import (
	main_entity "github.com/nuzur/nem/core/entity/project_version"
)

type SearchProjectVersionResponse struct {
	Results []main_entity.ProjectVersion
}
