package local_agent

import (
	main_entity "github.com/nuzur/nem/core/entity/local_agent"
	nemdb "github.com/nuzur/nem/core/repository/gen"

	"github.com/nuzur/nem/core/entity/mapper"

	"github.com/gofrs/uuid"
	"github.com/nuzur/nem/core/entity/local_agent_connection"
)

func mapModelsToEntities(models []nemdb.LocalAgent) []main_entity.LocalAgent {
	result := []main_entity.LocalAgent{}
	for _, p := range models {
		result = append(result, mapModelToEntity(p))
	}
	return result
}

func mapModelToEntity(model nemdb.LocalAgent) main_entity.LocalAgent {
	return main_entity.LocalAgent{
		UUID:          uuid.FromStringOrNil(model.UUID),
		UserUUID:      uuid.FromStringOrNil(model.UserUUID),
		MachineName:   mapper.SqlNullStringToStringPtr(model.MachineName),
		Os:            mapper.SqlNullStringToStringPtr(model.Os),
		CliVersion:    mapper.SqlNullStringToStringPtr(model.CliVersion),
		Connections:   local_agent_connection.LocalAgentConnectionSliceFromJSON(model.Connections),
		Status:        main_entity.Status(model.Status),
		LastSeenAt:    mapper.SqlNullTimeToTimePtr(model.LastSeenAt),
		RevokedAt:     mapper.SqlNullTimeToTimePtr(model.RevokedAt),
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.UpdatedAt,
		CreatedByUUID: uuid.FromStringOrNil(model.CreatedByUUID),
		UpdatedByUUID: uuid.FromStringOrNil(model.UpdatedByUUID),
	}
}
