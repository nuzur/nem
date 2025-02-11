package organization

import (
	main_entity "nem/core/entity/organization"
	nemdb "nem/core/repository/gen"

	"nem/core/entity/mapper"

	"github.com/gofrs/uuid"
	"nem/core/entity/membership"
)

func mapModelsToEntities(models []nemdb.Organization) []main_entity.Organization {
	result := []main_entity.Organization{}
	for _, p := range models {
		result = append(result, mapModelToEntity(p))
	}
	return result
}

func mapModelToEntity(model nemdb.Organization) main_entity.Organization {
	return main_entity.Organization{
		UUID:          uuid.FromStringOrNil(model.UUID),
		Version:       model.Version,
		Name:          model.Name,
		Domains:       mapper.JSONToStringSlice(model.Domains),
		AdminUUIDs:    mapper.JSONToUUIDSlice(model.AdminUUIDs),
		Memberships:   membership.MembershipSliceFromJSON(model.Memberships),
		Status:        main_entity.Status(model.Status),
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.UpdatedAt,
		CreatedByUUID: uuid.FromStringOrNil(model.CreatedByUUID),
		UpdatedByUUID: uuid.FromStringOrNil(model.UpdatedByUUID),
	}
}
