package types

import (
	"github.com/gofrs/uuid"
	main_entity "github.com/nuzur/nem/core/entity/ai_usage"
)

type UpsertRequest struct {
	AiUsage main_entity.AiUsage
}

type UpsertResponse struct {
	UUID uuid.UUID
}
