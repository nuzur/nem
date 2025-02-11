package types

import (
	main_entity "github.com/nuzur/nem/core/entity/organization"
)

type SearchOrganizationResponse struct {
	Results []main_entity.Organization
}
