package user_connection

import (
	main_entity "github.com/nuzur/nem/core/entity/user_connection"
	nemdb "github.com/nuzur/nem/core/repository/gen"

	"github.com/nuzur/nem/core/entity/mapper"

	"github.com/gofrs/uuid"
	"github.com/nuzur/nem/core/entity/user_connection_execution"
	"github.com/nuzur/nem/core/entity/user_connection_type_config"
)

func mapModelsToEntities(models []nemdb.UserConnection) []main_entity.UserConnection {
	result := []main_entity.UserConnection{}
	for _, p := range models {
		result = append(result, mapModelToEntity(p))
	}
	return result
}

func mapModelToEntity(model nemdb.UserConnection) main_entity.UserConnection {
	return main_entity.UserConnection{
		UUID:               uuid.FromStringOrNil(model.UUID),
		UserUUID:           uuid.FromStringOrNil(model.UserUUID),
		ProjectUUID:        uuid.FromStringOrNil(model.ProjectUUID),
		ProjectVersionUUID: uuid.FromStringOrNil(mapper.SqlNullStringToString(model.ProjectVersionUUID)),
		Type:               main_entity.Type(model.Type),
		TypeConfig:         user_connection_type_config.UserConnectionTypeConfigFromJSON(model.TypeConfig),
		DbSchema:           model.DbSchema,
		Executions:         user_connection_execution.UserConnectionExecutionSliceFromJSON(model.Executions),
		Status:             main_entity.Status(model.Status),
		CreatedAt:          model.CreatedAt,
		UpdatedAt:          model.UpdatedAt,
	}
}
