package types

import (
	main_entity "nem/core/entity/organization"
)

type SearchOrganizationResponse struct {
	Results []main_entity.Organization
}
