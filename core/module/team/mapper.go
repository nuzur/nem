package team

import (
	main_entity "github.com/nuzur/nem/core/entity/team"
	nemdb "github.com/nuzur/nem/core/repository/gen"

	"github.com/gofrs/uuid"
	"github.com/nuzur/nem/core/entity/connection"
	"github.com/nuzur/nem/core/entity/entity"
	"github.com/nuzur/nem/core/entity/enviorment"
	"github.com/nuzur/nem/core/entity/object_store"
	"github.com/nuzur/nem/core/entity/review_config"
	"github.com/nuzur/nem/core/entity/store"
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
		UUID:          uuid.FromStringOrNil(model.UUID),
		Version:       model.Version,
		Name:          model.Name,
		Enviorments:   enviorment.EnviormentSliceFromJSON(model.Enviorments),
		ReviewConfigs: review_config.ReviewConfigSliceFromJSON(model.ReviewConfigs),
		Stores:        store.StoreSliceFromJSON(model.Stores),
		Connections:   connection.ConnectionSliceFromJSON(model.Connections),
		ObjectStores:  object_store.ObjectStoreSliceFromJSON(model.ObjectStores),
		DefaultEntity: entity.EntityFromJSON(model.DefaultEntity),
		Status:        main_entity.Status(model.Status),
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.UpdatedAt,
		CreatedByUUID: uuid.FromStringOrNil(model.CreatedByUUID),
		UpdatedByUUID: uuid.FromStringOrNil(model.UpdatedByUUID),
	}
}
