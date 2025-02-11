package types

import (
	main_entity "github.com/nuzur/nem/core/entity/extension"
)

type SearchExtensionResponse struct {
	Results []main_entity.Extension
}
