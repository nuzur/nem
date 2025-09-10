package membership

import (
	main_entity "github.com/nuzur/nem/core/entity/membership"
	nemdb "github.com/nuzur/nem/core/repository/gen"

	"github.com/gofrs/uuid"
)

func mapModelsToEntities(models []nemdb.Membership) []main_entity.Membership {
	result := []main_entity.Membership{}
	for _, p := range models {
		result = append(result, mapModelToEntity(p))
	}
	return result
}

func mapModelToEntity(model nemdb.Membership) main_entity.Membership {
	return main_entity.Membership{
		UUID:            uuid.FromStringOrNil(model.UUID),
		OwnerUUID:       uuid.FromStringOrNil(model.OwnerUUID),
		Type:            main_entity.Type(model.Type),
		StartDate:       model.StartDate,
		BillingMetadata: string(model.BillingMetadata),
		Status:          main_entity.Status(model.Status),
		CreatedAt:       model.CreatedAt,
		UpdatedAt:       model.UpdatedAt,
		CreatedByUUID:   uuid.FromStringOrNil(model.CreatedByUUID),
		UpdatedByUUID:   uuid.FromStringOrNil(model.UpdatedByUUID),
	}
}
