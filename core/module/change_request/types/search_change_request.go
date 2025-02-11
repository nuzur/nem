package types

import (
	main_entity "nem/core/entity/change_request"
)

type SearchChangeRequestResponse struct {
	Results []main_entity.ChangeRequest
}
