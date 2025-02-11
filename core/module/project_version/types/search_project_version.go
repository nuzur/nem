package types

import (
	main_entity "nem/core/entity/project_version"
)

type SearchProjectVersionResponse struct {
	Results []main_entity.ProjectVersion
}
