package team

import (
	main_entity "nem/core/entity/team"
	nemdb "nem/core/repository/gen"

	"github.com/gofrs/uuid"
	"nem/core/entity/connection"
	"nem/core/entity/entity"
	"nem/core/entity/enviorment"
	"nem/core/entity/membership"
	"nem/core/entity/object_store"
	"nem/core/entity/review_config"
	"nem/core/entity/store"
)

func mapModelsToEntities(models []nemdb.Team) []main_entity.Team {
	result := []main_entity.Team{}
	for _, p := range models {
		result = append(result, mapModelToEntity(p))
	}
	return result
}

func mapModelToEntity(model nemdb.Team) main_entity.Team {
	return main_entity.Team{
		UUID:             uuid.FromStringOrNil(model.UUID),
		Version:          model.Version,
		Name:             model.Name,
		Enviorments:      enviorment.EnviormentSliceFromJSON(model.Enviorments),
		ReviewConfigs:    review_config.ReviewConfigSliceFromJSON(model.ReviewConfigs),
		Memberships:      membership.MembershipSliceFromJSON(model.Memberships),
		Stores:           store.StoreSliceFromJSON(model.Stores),
		Connections:      connection.ConnectionSliceFromJSON(model.Connections),
		ObjectStores:     object_store.ObjectStoreSliceFromJSON(model.ObjectStores),
		OrganizationUUID: uuid.FromStringOrNil(model.OrganizationUUID),
		DefaultEntity:    entity.EntityFromJSON(model.DefaultEntity),
		Status:           main_entity.Status(model.Status),
		CreatedAt:        model.CreatedAt,
		UpdatedAt:        model.UpdatedAt,
		CreatedByUUID:    uuid.FromStringOrNil(model.CreatedByUUID),
		UpdatedByUUID:    uuid.FromStringOrNil(model.UpdatedByUUID),
	}
}
